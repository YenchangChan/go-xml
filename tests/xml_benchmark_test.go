package tests

import (
	stdxml "encoding/xml"
	"fmt"
	"github.com/YenchangChan/xml"
	"testing"
)

type Address struct {
	City, State string
}
type Person struct {
	XMLName stdxml.Name `xml:"person"`
	Id int `xml:"id,attr"`
	FirstName string `xml:"name>first"`
	LastName string `xml:"name>last"`
	Age int `xml:"age"`
	Height float32 `xml:"height,omitempty"`
	Married bool
	Address
	Comment string `xml:",comment"`
}


func BenchmarkEncodingXml(b *testing.B){
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
		v.Comment = "Need more details."
		v.Address = Address{"Hanga Roa", "Easter Island"}

		_, err := stdxml.MarshalIndent(v, "    ", "    ")
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}
}

func BenchmarkGoXml(b *testing.B){
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f := xml.NewXmlFile("bench.xml")
		f.BeginwithAttr("person", []xml.XMLAttr{{Key:"id", Value:13}})
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
		f.GetContext()
	}
}
