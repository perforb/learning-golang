package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"` // attr 追加で <plant id="27"> のようになる
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}
	out, _ := xml.MarshalIndent(coffee, " ", " ")
	// <plant id="27">
	//  <name>Coffee</name>
	//  <origin>Ethiopia</origin>
	//  <origin>Brazil</origin>
	// </plant>
	fmt.Println(string(out))
	//<?xml version="1.0" encoding="UTF-8"?>
	// <plant id="27">
	//  <name>Coffee</name>
	//  <origin>Ethiopia</origin>
	//  <origin>Brazil</origin>
	// </plant>
	fmt.Println(xml.Header + string(out))

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	// Plant id=27, name=Coffee, origin=[Ethiopia Brazil]
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}
	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}
	out, _ = xml.MarshalIndent(nesting, " ", " ")
	// <nesting>
	//  <parent>
	//   <child>
	//    <plant id="27">
	//     <name>Coffee</name>
	//     <origin>Ethiopia</origin>
	//     <origin>Brazil</origin>
	//    </plant>
	//    <plant id="81">
	//     <name>Tomato</name>
	//     <origin>Mexico</origin>
	//     <origin>California</origin>
	//    </plant>
	//   </child>
	//  </parent>
	// </nesting>
	fmt.Println(string(out))
}
