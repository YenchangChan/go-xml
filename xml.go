package xml

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"reflect"
	"strings"
)

type XMLFile struct {
	name     string
	builder  strings.Builder
	indent   int
}

type XMLAttr struct {
	Key   string
	Value interface{}
}

func NewXmlFile(name string) *XMLFile {
	return &XMLFile{
		name:   name,
		indent: 0,
	}
}

func escape(context string) string {
	context = strings.ReplaceAll(context, "<", "&lt;")
	context = strings.ReplaceAll(context, ">", "&gt;")
	context = strings.ReplaceAll(context, "&", "&amp;")
	context = strings.ReplaceAll(context, "'", "&apos;")
	context = strings.ReplaceAll(context, "\"", "&quot;")
	return context
}

func finalValue(value interface{}) interface{} {
	if _, ok := value.(string); ok {
		value = escape(value.(string))
	}
	rv := reflect.ValueOf(value)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	return rv.Interface()
}

func (xml *XMLFile) GetContext() string {
	return xml.builder.String()
}

func (xml *XMLFile) SetContext(context string) {
	xml.builder.Reset()
	xml.builder.WriteString(context)
}

func (xml *XMLFile) GetIndent() int {
	return xml.indent
}

func (xml *XMLFile) SetIndent(indent int) {
	xml.indent = indent
}

func (xml *XMLFile) Write(tag string, value interface{}) {
	value = finalValue(value)
	if value == nil {
		return
	}
	xml.builder.WriteString(fmt.Sprintf("%s<%s>%v</%s>\n", strings.Repeat(" ", xml.indent*4), tag, value, tag))
}

func (xml *XMLFile) WritewithAttr(tag string, value interface{}, attrs []XMLAttr) {
	value = finalValue(value)
	if value == nil {
		return
	}

	xml.builder.WriteString(fmt.Sprintf("%s<%s", strings.Repeat(" ", xml.indent*4), tag))
	for idx, attr := range attrs {
		if idx < len(attrs) {
			xml.builder.WriteString(" ")
		}
		xml.builder.WriteString(fmt.Sprintf("%s=\"%v\"", attr.Key, finalValue(attr.Value)))
	}
	xml.builder.WriteString(fmt.Sprintf(">%v</%s>\n", value, tag))
}

func (xml *XMLFile) Begin(tag string) {
	xml.builder.WriteString(fmt.Sprintf("%s<%s>\n", strings.Repeat(" ", xml.indent*4), tag))
	xml.indent++
}

func (xml *XMLFile) BeginwithAttr(tag string, attrs []XMLAttr) {
	xml.builder.WriteString(fmt.Sprintf("%s<%s", strings.Repeat(" ", xml.indent*4), tag))
	for idx, attr := range attrs {
		if idx < len(attrs) {
			xml.builder.WriteString(" ")
		}
		xml.builder.WriteString(fmt.Sprintf("%s=\"%v\"", attr.Key, finalValue(attr.Value)))
	}
	xml.builder.WriteString(">\n")
	xml.indent++
}

func (xml *XMLFile) End(tag string) {
	xml.indent--
	xml.builder.WriteString(fmt.Sprintf("%s</%s>\n", strings.Repeat(" ", xml.indent*4), tag))
}

func (xml *XMLFile) Comment(comment string) {
	xml.builder.WriteString(fmt.Sprintf("%s<!-- %s -->\n", strings.Repeat(" ", xml.indent*4), comment))
}

func (xml *XMLFile) Append(context string) {
	xml.builder.WriteString(context)
}

func (xml *XMLFile) Dump() error {
	if xml.name == "" {
		return errors.Errorf("xml name is not exist")
	}
	context := xml.builder.String()
	if context == "" {
		return errors.Errorf("xml %s context is empty", xml.name)
	}
	fi, err := os.OpenFile(xml.name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fi.Close()
	nbytes, err := fi.WriteString(context)
	if err != nil {
		return err
	}
	if nbytes != len(context) {
		return errors.Errorf("write xml file %s failed.", xml.name)
	}
	return nil
}
