package map2xml

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func Test_Map2Xml(t *testing.T) {
	inputMap := map[string]interface{}{
		"first_name": "No",
		"last_name":  "Name",
		"age":        42,
		"got_a_job":  true,
		"address": map[string]interface{}{
			"street":   "124 Oxford Street",
			"city":     "Somewhere",
			"zip_code": 32784,
			"state":    "Deep state",
		},
		"fileList": []map[string]interface{}{
			{
				"username": "jett1",
				"age":      20,
			},
			{
				"username": "jett2",
				"age":      22,
			},
		},
	}

	config := New(inputMap)
	config.WithIndent("", "  ")
	config.WithRoot("person", map[string]string{"mood": "happy"})
	xmlBytes, err := config.MarshalToString()
	if err != nil {
		t.Error(err)
	}

	// 写入xml
	f, err := os.Create("index.xml")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(xmlBytes)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}
