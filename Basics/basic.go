package main

import (
	"fmt"
)

//Struct

type CMember struct {
	name      string
	age       int
	address   string
	rank      string
	clearance int
}

func main() {
	cm := CMember{"Arun", 23, "Station Mars", "Sergeant", 5}

	cm1 := CMember{
		name:      "Kevin",
		clearance: 5,
		age:       32,
		rank:      "Sergeant",
		address:   "Station Mars",
	}

	var cm2 CMember
	cm.name = "Sam"
	cm.age = 23
	cm.address = "Earth"
	cm.rank = "Sergeant"
	cm.clearance = 5

	cmp := &cm

	fmt.Println(string(cmp.name))
	// Slices
	var crew []CMember
	crew = append(crew, cm, cm1, cm2)
	for i, v := range crew {
		fmt.Println(i, v.name)
	}
	//maps with structs
	var m map[string]CMember
	m = make(map[string]CMember)
	m["keyring"] = cm

	m1 := map[string]CMember{
		"kevin": CMember{name: "Kevin", address: "Earth2"},
		"Jo":    CMember{name: "Jo", address: "Earth2"},
	}
	m["cis"] = CMember{name: "Cisco", address: "Mars"}

	//retrieve
	elem := m1["Jo"]
	fmt.Println(elem)

	//Check if the value exists
	v, ok := m1["Jo"]
	fmt.Println(ok, v)
	delete(m, "Jo")

	for k, v := range m1 {
		fmt.Println("Key:", k, "Value", v)
	}
	cm.PrintSecurityClearance()
}

//Methods - A function with receiver
func (cm CMember) PrintSecurityClearance() {
	fmt.Println(cm.clearance)
}
