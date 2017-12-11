// Crappiest program to ping multiple domains - for learning purpose
package main

import (
	"fmt"
	"net"
	"os"
	"time"

	tablewriter "github.com/olekukonko/tablewriter"
	fastping "github.com/tatsushid/go-fastping"
)

func getRTT(ip string) time.Duration {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", ip)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p.AddIPAddr(ra)
	var rrtt time.Duration
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		rrtt = rtt
	}
	p.OnIdle = func() {
		// fmt.Println("finish")
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)

	}
	return rrtt
}

func main() {
	args := os.Args[1:]
	data := [][]string{}
	var data_item []string
	data_item = make([]string, 2)
	//data := make([][]string,3)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"IP", "RTT"})

	for _, arg := range args {
		ip_rtt := getRTT(arg)
		data_item[0] = string(arg)
		data_item[1] = ip_rtt.String()
		data = data[:0]
		data = append(data, data_item)
		for _, v := range data {
			table.Append(v)
		}

	}
	table.Render()

}
