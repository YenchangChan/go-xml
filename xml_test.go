package xml

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXMLFile_Dump(t *testing.T) {
	f := NewXmlFile("test.xml")
	f.Comment("This is a test xml file")
	f.Begin("top")
	f.Write("name", "zhangsan")
	f.Write("age", 20)
	f.BeginwithAttr("subTag", []XMLAttr{{Key: "id", Value:1}})
	f.Write("escape", "abc&?32>")
	f.WritewithAttr("attr", "value", []XMLAttr{{Key:"value", Value:"abc"}})
	f.End("subTag")
	f.End("top")
	except := `<!-- This is a test xml file -->
<top>
    <name>zhangsan</name>
    <age>20</age>
    <subTag id="1">
        <escape>abc&amp;?32&amp;gt;</escape>
        <attr value="abc">value</attr>
    </subTag>
</top>
`
	assert.Equal(t, f.GetContext(), except)
}

func TestEscape(t *testing.T){
	context := "abc<sdsf>sf&zd'sd\\ade"
	assert.Equal(t, escape(context), "abc&amp;lt;sdsf&amp;gt;sf&amp;zd&apos;sd\\ade")
}
