# xml
go-xml is a library for wirte data to xml file written by Go.
# Quick Start
## Installation
```bash
go get github.com/YenchangChan/gxml
```
## Example
```go
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
```
You can get a xml file like this:
```xml
<person id="13">
    <name>
        <first>John</first>
        <last>Doe</last>
    </name>
    <age>42</age>
    <Married>false</Married>
    <City>Hanga Roa</City>
    <State>Easter Island</State>
    <!-- Need more details. -->
</person>
```

# Why not encoding/xml?
You must define a struct with `xml` tag first, and then use reflect to marshal struct to xml.
In this way, you can't generate xml file dynamiclly.

# BenchMark
```bash
goos: darwin
goarch: amd64
pkg: github.com/YenchangChan/xml/tests
cpu: Intel(R) Core(TM) i5-8257U CPU @ 1.40GHz
BenchmarkEncodingXml-8            295854              4077 ns/op            5184 B/op         14 allocs/op
BenchmarkGoXml-8                  292113              3841 ns/op            1336 B/op         58 allocs/op

```
