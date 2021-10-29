package es

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/olivere/elastic/v7"
)

var (
	EsClient *Elastic
)

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

func init() {
	EsClient = InitES("10.4.7.71", 9200, "elastic", "hZksYkpkcweABXu68qh0", "micrograme-test")
}

// 新建索引
// postman查看: http://10.4.7.71:9200/why_index; 记得authorization-> basic auth-> 用户名密码登录
// Kibana 查看: GET /why_index
func Test_CreateIndex(t *testing.T) {
	index := "go-micro-frame-test"

	searchResult := EsClient.CreateIndex(index)
	fmt.Println(searchResult)
}

// 删除索引
func Test_DelIndex(t *testing.T) {
	index := "why_index"

	searchResult := EsClient.DelIndex(index)
	fmt.Println(searchResult)
}

// 增加数据
// 查看：/go-micro-frame/employee/1
func Test_Put(t *testing.T) {
	//使用结构体
	{
		index := "go-micro-frame"
		typeName := "employee"
		id := "1"
		e1 := Employee{"Jane", "Smith", 32, "I like go-micro", []string{"music"}}

		searchResult := EsClient.PutAny(index, typeName, id, e1)
		fmt.Println(searchResult)
	}

	// 使用json字符串
	{
		index := "go-micro-frame"
		typeName := "employee"
		id2 := "2"
		e2 := `{"first_name":"jett","last_name":"jia","age":25,"about":"I love to go","interests":["sports","music"]}`

		searchResult := EsClient.PutAny(index, typeName, id2, e2)
		fmt.Println(searchResult)
	}
}

// 修改
func Test_Update(t *testing.T) {
	index := "go-micro-frame"
	typeName := "employee"
	id := "1"
	updateMap := map[string]interface{}{
		"age":        88,
		"first_name": "lili",
	}
	EsClient.Update(index, typeName, id, updateMap)
}

// 删除
func Test_Del(t *testing.T) {
	index := "go-micro-frame"
	typeName := "employee"
	id := "2"
	EsClient.Del(index, typeName, id)
}

// 查收数据
func Test_QueryString(t *testing.T) {
	index := "go-micro-frame"
	typeName := "employee"
	//query := "first_name:jett OR age:25"
	query := "first_name:Jane"
	rsp := EsClient.QueryString(index, typeName, query, 10)
	printEmployee(rsp, nil)
}

// 查询数据, 有缺陷
func Test_QueryStringMap(t *testing.T) {
	index := "go-micro-frame"
	typeName := "employee"
	//query := "first_name:jett OR age:25"
	query := "first_name:jett"
	stringMap := EsClient.QueryStringMap(index, typeName, query, 1, 10)
	fmt.Println(stringMap)
}

//打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ Employee
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Employee)
		fmt.Printf("%#v\n", t)
	}
}
