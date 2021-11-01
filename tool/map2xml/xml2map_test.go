package map2xml

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Xml2Map(t *testing.T) {
	data := `<container uid="FA6666D9-EC9F-4DA3-9C3D-4B2460A4E1F6" lifetime="2019-10-10T18:00:11">
				<cats>
					<cat>
						<id>CDA035B6-D453-4A17-B090-84295AE2DEC5</id>
						<name>moritz</name>
						<age>7</age> 	
						<items>
							<n>1293</n>
							<n>1255</n>
							<n>1257</n>
						</items>
					</cat>
					<cat>
						<id>1634C644-975F-4302-8336-1EF1366EC6A4</id>
						<name>oliver</name>
						<age>12</age>
					</cat>
					<dog color="gray">hello</dog>
				</cats>
				<color>white</color>
				<city>NY</city>
			</container>`

	decoder := NewDecoder(strings.NewReader(data))
	result, err := decoder.Decode()

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v\n", result)
	}

	v := result["container"].
	(map[string]interface{})["cats"].
	(map[string]interface{})["cat"].
	([]map[string]interface{})[0]["items"].
	(map[string]interface{})["n"].([]string)[1]

	fmt.Printf("n[1]: %v\n", v)
}
