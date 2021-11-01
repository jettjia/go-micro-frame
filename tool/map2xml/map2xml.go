// map2xml
// from https://github.com/yoda-of-soda/map2xml
package map2xml

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"reflect"
)

type Indentation struct {
	Prefix string
	Indent string
}

type Root struct {
	Name          *xml.Name
	XMLAttributes *[]xml.Attr
	Attributes    map[string]string
}
type StructMap struct {
	CData  bool
	Map    map[string]interface{}
	Indent *Indentation
	Root   *Root
}

type xmlMapEntry struct {
	XMLName    xml.Name
	Value      interface{} `xml:",innerxml"`
	CDataValue interface{} `xml:",cdata"`
}

//Initializes the builder. Required to do anything with this library
func New(input map[string]interface{}) *StructMap {
	return &StructMap{Map: input}
}

//Add indentation to your xml
func (smap *StructMap) WithIndent(prefix string, indent string) *StructMap {
	smap.Indent = &Indentation{Prefix: prefix, Indent: indent}
	return smap
}

//Add a root node to your xml
func (smap *StructMap) WithRoot(name string, attributes map[string]string) *StructMap {
	attr := []xml.Attr{}
	for k, v := range attributes {
		attr = append(attr, xml.Attr{Name: xml.Name{Local: k}, Value: v})
	}
	smap.Root = &Root{Name: &xml.Name{Local: name}, XMLAttributes: &attr, Attributes: attributes}
	return smap
}

//Add CDATA tags to all nodes
func (smap *StructMap) AsCData() *StructMap {
	smap.CData = true
	return smap
}

//Prints your configuration in json format
func (smap *StructMap) Print() *StructMap {
	var indent interface{} = smap.Indent
	var root interface{}
	if smap.Indent != nil {
		indent = map[string]int{"indent_spaces": len(*&smap.Indent.Indent), "prefix_spaces": len(*&smap.Indent.Prefix)}
	}
	if root = smap.Root; root != nil {
		root = map[string]interface{}{"name": *&smap.Root.Name.Local, "attributes": smap.Root.Attributes}
	}
	b, _ := json.MarshalIndent(map[string]interface{}{"root": root, "cdata": smap.CData, "indent": indent}, " ", "  ")
	fmt.Println(string(b))
	return smap
}

//Builds XML as bytes
func (smap *StructMap) Marshal() ([]byte, error) {
	if smap.Indent == nil {
		return xml.Marshal(smap)
	} else {
		return xml.MarshalIndent(smap, smap.Indent.Prefix, smap.Indent.Indent)
	}
}

//Builds XML as string
func (smap *StructMap) MarshalToString() (string, error) {
	xmlBytes, err := smap.Marshal()
	return string(xmlBytes), err
}

func (smap StructMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(smap.Map) == 0 {
		return nil
	}
	if smap.Root != nil {
		start = xml.StartElement{Name: *smap.Root.Name, Attr: *smap.Root.XMLAttributes}
		if err := e.EncodeToken(start); err != nil {
			return err
		}
	}

	for k, v := range smap.Map {
		if err := handleChildren(e, k, v, smap.CData); err != nil {
			return err
		}
	}

	if smap.Root != nil {
		return e.EncodeToken(start.End())
	}
	return nil
}

func handleChildren(e *xml.Encoder, fieldName string, v interface{}, cdata bool) error {
	if reflect.TypeOf(v) == nil {
		return e.Encode(xmlMapEntry{XMLName: xml.Name{Local: fieldName}, Value: ""})
	} else if reflect.TypeOf(v).Kind() == reflect.Map {
		e.EncodeToken(xml.StartElement{Name: xml.Name{Local: fieldName}})
		for key, val := range v.(map[string]interface{}) {
			handleChildren(e, key, val, cdata)
		}
		return e.EncodeToken(xml.EndElement{Name: xml.Name{Local: fieldName}})
	} else if reflect.TypeOf(v).Kind() == reflect.Slice {
		e.EncodeToken(xml.StartElement{Name: xml.Name{Local: fieldName}})
		childName := fieldName + "_child"
		if _, hasChildName := v.([]map[string]interface{})[0]["xml_child_name"]; hasChildName {
			childName = v.([]map[string]interface{})[0]["xml_child_name"].(string)
		}
		for _, elem := range v.([]map[string]interface{}) {
			handleChildren(e, childName, elem, cdata)
		}
		return e.EncodeToken(xml.EndElement{Name: xml.Name{Local: fieldName}})
	}
	if cdata {
		return e.Encode(xmlMapEntry{XMLName: xml.Name{Local: fieldName}, CDataValue: v})
	} else {
		return e.Encode(xmlMapEntry{XMLName: xml.Name{Local: fieldName}, Value: fmt.Sprintf("%v", v)})
	}
}
