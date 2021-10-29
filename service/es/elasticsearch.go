package es

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/olivere/elastic/v7"

	"github.com/jettjia/go-micro-frame/service/es/util"
	"github.com/jettjia/go-micro-frame/service/logger"
)

type Elastic struct {
	Client *elastic.Client
	host   string
}

// es 缓存池
var (
	ElasticPool = make(map[string]*Elastic)
)

func InitES(addr string, port int, user, password, name string) *Elastic {
	if es, ok := ElasticPool[name]; ok {
		return es
	}
	host := fmt.Sprintf("http://%s:%d", addr, port)

	logger := log.New(os.Stdout, "microframe", log.LstdFlags)

	var err error
	client, err := elastic.NewClient(
		elastic.SetURL(host), elastic.SetSniff(false),
		elastic.SetBasicAuth(user, password),
		elastic.SetTraceLog(logger),
	)
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esVersion, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esVersion)
	es := &Elastic{
		Client: client,
		host:   host,
	}
	ElasticPool[name] = es
	return es
}

// 创建 index
func (self *Elastic) CreateIndex(index string) bool {
	// 判断索引是否存在
	exists, err := self.Client.IndexExists(index).Do(context.Background())
	if err != nil {
		logger.Error(err, logger.String("<CreateIndex> some error occurred when check exists, index", index))
		return false
	}
	if exists {
		logger.Error(err, logger.String("<CreateIndex> index:{%s} is already exists", index))
		return true
	}
	createIndex, err := self.Client.CreateIndex(index).Do(context.Background())
	if err != nil {
		logger.Error(err, logger.String("<CreateIndex> some error occurred when create. index: %s", index))
		return false
	}
	if !createIndex.Acknowledged {
		// Not acknowledged
		logger.Error(err, logger.String("<CreateIndex> Not acknowledged, index: %s", index))
		return false
	}
	return true
}

/*
新建索引
*/
func (self *Elastic) CreateIndexAndMapping(index, mapping string) bool {
	// 判断索引是否存在
	exists, err := self.Client.IndexExists(index).Do(context.Background())
	if err != nil {
		logger.Error(err, logger.String("<CreateIndex> some error occurred when check exists, index", index))
		return false
	}
	if exists {
		logger.Error(err, logger.String("<CreateIndex> index:{%s} is already exists", index))
		return true
	}
	createIndex, err := self.Client.CreateIndex(index).Body(mapping).Do(context.Background())
	if err != nil {
		logger.Error(err, logger.String("<CreateIndex> some error occurred when create. index: %s", index))
		return false
	}
	if !createIndex.Acknowledged {
		// Not acknowledged
		logger.Error(err, logger.String("<CreateIndex> Not acknowledged, index: %s", index))
		return false
	}
	return true
}

/*
删除索引
*/
func (self *Elastic) DelIndex(index string) bool {
	// Delete an index.
	deleteIndex, err := self.Client.DeleteIndex(index).Do(context.Background())
	if err != nil {
		// Handle error
		logger.Error(err, logger.String("<DelIndex> some error occurred when delete. index: %s", index))
		return false
	}
	if !deleteIndex.Acknowledged {
		// Not acknowledged
		logger.Error(err, logger.String("<DelIndex> acknowledged. index: %s", index))
		return false
	}
	return true
}

// 存储-string类型
func (self *Elastic) Put(index, typ, id, bodyJson string) bool {
	_, err := self.Client.Index().
		Index(index).
		Type(typ).
		Id(id).
		BodyJson(bodyJson).
		Do(context.Background())
	if err != nil {
		// Handle error
		logger.Error(err, logger.String("<Put> some error occurred when put.  err:%s", index))
		return false
	}
	//logger.Info("<Put> success", logger.String(" id: %s", put.Id))
	return true
}

// 存储-任何数据
func (self *Elastic) PutAny(index, typ, id string, body interface{}) bool {
	_, err := self.Client.Index().
		Index(index).
		Type(typ).
		Id(id).
		BodyJson(body).
		Do(context.Background())
	if err != nil {
		// Handle error
		logger.Error(err, logger.String("<Put> some error occurred when put.  err:%s", index))
		return false
	}
	//logger.Info("<Put> success", logger.String(" id: %s", put.Id))
	return true
}

/*
数据删除
*/
func (self *Elastic) Del(index, typ, id string) bool {
	_, err := self.Client.Delete().
		Index(index).
		Type(typ).
		Id(id).
		Do(context.Background())
	if err != nil {
		// Handle error
		logger.Error(err, logger.String("<Del> some error occurred when put", index))
		return false
	}
	//logger.Info("<Del> success", logger.String(" id: %s", del.Id))
	return true
}

/*
更新数据
*/
func (self *Elastic) Update(index, typ, id string, updateMap map[string]interface{}) bool {
	res, err := self.Client.Update().
		Index(index).Type(typ).Id(id).
		Doc(updateMap).
		FetchSource(true).
		Do(context.Background())
	if err != nil {
		logger.Error(err, logger.String("<Update> some error occurred when update. index", index))
		return false
	}
	if res == nil {
		logger.Error(err, logger.String("<Update> res == nil when update. index", index))
		return false
	}
	if res.GetResult == nil {
		logger.Error(err, logger.String("<Update> expected GetResult != nil when update. index", index))
		return false
	}
	//data, _ := json.Marshal(res.GetResult.Source)
	//logger.Info("<Update> success", logger.String(" data: %s", string(data)))
	return true
}

//func (self *Elastic) TermQueryMap(index, typ string, term *elastic.TermQuery, start, end int) map[string]interface{} {
//	//elastic.NewTermQuery()
//	searchResult, err := self.Client.Search().
//		Index(index).
//		Type(typ).
//		Query(term). // specify the query
//		From(start).Size(end). // take documents start-end
//		Pretty(true). // pretty print request and response JSON
//		DoString(context.Background()) // execute
//	if err != nil {
//		fmt.Println(err.Error())
//		// Handle error
//		logger.Error(err, logger.String("<TermQuery> some error occurred when search. index:%s", index))
//		return make(map[string]interface{})
//	}
//	return util.JsonToMap(searchResult)
//}

func (self *Elastic) TermQuery(index, typ string, term *elastic.TermQuery, start, end int) *elastic.SearchResult {
	//elastic.NewTermQuery()
	searchResult, err := self.Client.Search().
		Index(index).
		Type(typ).
		Query(term). // specify the query
		From(start).Size(end). // take documents start-end
		Pretty(true). // pretty print request and response JSON
		Do(context.Background()) // execute
	if err != nil {
		fmt.Println(err.Error())
		// Handle error
		logger.Error(err, logger.String("<TermQuery> some error occurred when search. index:%s", index))
		return nil
	}
	return searchResult
}

func (self *Elastic) QueryStringMap(index, typ, query string, start, end int) map[string]interface{} {
	q := elastic.NewQueryStringQuery(query)
	// Match all should return all documents
	searchResult, err := self.Client.Search().
		Index(index).
		Type(typ). // type of Index
		Query(q).
		Pretty(true).
		From(start).Size(end).
		Do(context.Background())
	if err != nil {
		// Handle error
		logger.Error(err, logger.String("<QueryString> some error occurred when search. index:%s", index))
		return nil
	}

	b, err := json.Marshal(&searchResult)
	if err != nil {
		logger.Error(err, logger.String("<QueryString> error, when json.Marshal the data. index:%s", index))
	}

	return util.JsonToMap(string(b))
}

// https://www.elastic.co/guide/en/elasticsearch/reference/6.8/query-dsl-query-string-query.html
func (self *Elastic) QueryString(index, typ, query string, size int) *elastic.SearchResult {

	q := elastic.NewQueryStringQuery(query)
	// Match all should return all documents
	searchResult, err := self.Client.Search().
		Index(index).
		Type(typ). // type of Index
		Query(q).
		Pretty(true).
		From(0).Size(size).
		Do(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		// Handle error
		logger.Error(err, logger.String("<QueryString> some error occurred when search. index:%s", index))
		return nil
	}
	return searchResult
}

/*
多条件参考：https://stackoverflow.com/questions/49942373/golang-elasticsearch-multiple-query-parameters
*/
func (self *Elastic) MultiMatchQueryBestFields(index, typ, text string, start, end int, fields ...string) *elastic.SearchResult {
	q := elastic.NewMultiMatchQuery(text, fields...)
	searchResult, err := self.Client.Search().
		Index(index). // name of Index
		Type(typ). // type of Index
		Query(q).
		Sort("_score", false).
		From(start).Size(end).
		Do(context.Background())
	if err != nil {
		// Handle error
		logger.Error(err, logger.String("<MultiMatchQueryBestFields> some error occurred when search. index:%s", index))
		return nil
	}
	return searchResult
}

//func (self *Elastic) MultiMatchQueryBestFieldsMap(index, typ, text string, start, end int, fields ...string) map[string]interface{} {
//	q := elastic.NewMultiMatchQuery(text, fields...)
//	searchResult, err := self.Client.Search().
//		Index(index). // name of Index
//		Type(typ). // type of Index
//		Query(q).
//		Sort("_score", false).
//		From(start).Size(end).
//		DoString(context.Background())
//	if err != nil {
//		// Handle error
//		log.Errorf(logHeader, "<MultiMatchQueryBestFields> some error occurred when search. index:%s, text:%v,  err:%s", index, text, err.Error())
//		return nil
//	}
//	return util.JsonToMap(searchResult)
//}

func (self *Elastic) QueryStringRandomSearch(client *elastic.Client, index, typ, query string, size int) *elastic.SearchResult {
	q := elastic.NewFunctionScoreQuery()
	queryString := elastic.NewQueryStringQuery(query)
	q = q.Query(queryString)
	q = q.AddScoreFunc(elastic.NewRandomFunction())
	searchResult, err := client.Search().
		Index(index).
		Type(typ).
		Query(q).
		Size(size).
		Do(context.Background())
	if err != nil {
		// Handle error
		logger.Error(err, logger.String("<QueryStringRandomSearch> some error occurred when search. index:%s", index))
		return nil
	}
	return searchResult
}

func (self *Elastic) RangeQueryLoginDate(index string, typ string, start, end int) *elastic.SearchResult {
	q := elastic.NewRangeQuery("latest_time").
		Gte("now-30d/d")
	searchResult, err := self.Client.Search().
		Index(index).
		Type(typ).
		Query(q).
		Sort("latest_time", false).
		From(start).Size(end).
		Do(context.Background())
	if err != nil {
		logger.Error(err, logger.String("<RangeQueryLoginDate> some error occurred when search. index:%s", index))
		return nil
	}
	return searchResult
}

func (self *Elastic) JsonMap(index, typ, query string, fields []string, from, size int,
	terms map[string]interface{}, mustNot, filter, sort []map[string]interface{}) map[string]interface{} {
	data := make(map[string]interface{})

	var must []map[string]interface{}
	if len(terms) != 0 {
		must = append(must, map[string]interface{}{
			"terms": terms,
		})
	}
	if query != "" {
		multiMatch := make(map[string]interface{})
		multiMatch["query"] = "小龙虾批发"
		if len(fields) != 0 {
			multiMatch["fields"] = fields
		}
		must = append(must, map[string]interface{}{
			"multi_match": multiMatch,
		})
	}
	data["from"] = from
	data["size"] = size
	data["sort"] = sort
	data["query"] = map[string]interface{}{
		"bool": map[string]interface{}{
			"must":     must,
			"must_not": mustNot,
			"filter":   filter,
		},
	}

	byteDates, err := json.Marshal(data)
	util.Must(err)
	reader := bytes.NewReader(byteDates)

	client := &http.Client{}
	url := self.host + "/" + index + "/" + typ + "/_search"
	req, err := http.NewRequest("POST", url, reader)
	util.Must(err)

	req.Header.Add("content-type", "application/json")

	resp, err := client.Do(req)
	util.Must(err)
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	util.Must(err)

	ret := util.JsonToMap(string(b))
	if ret["error"] != nil {
		fmt.Println(fmt.Sprintf("%s", string(b)))
		logger.Error(err, logger.String("<JsonMap> error",  string(b)))
		panic("es error")
		return nil
	}

	return ret
}
