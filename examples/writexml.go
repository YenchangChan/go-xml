package main

import (
	"fmt"
	"github.com/YenchangChan/xml"
)

func main() {
	f := xml.NewXmlFile("example.xml")
	f.BeginwithAttr("person", []xml.XMLAttr{{Key: "id", Value: 13}})
	f.Begin("name")
	f.Write("first", "John")
	f.Write("last", "Doe")
	f.End("name")
	f.Write("age", 42)
	f.Write("Married", false)
	f.Write("City", "Hanga Roa")
	f.Write("State", "Easter Island")
	f.Comment("Need more details.")
	f.End("person")
	if err := f.Dump(); err != nil {
		fmt.Println(err)
	}
}