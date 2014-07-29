package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type ElasticSearch struct {
	client     *http.Client
	baseUrl    string
	indiceName string
}

func NewElasticSearch(port int) (ret *ElasticSearch) {
	ret = new(ElasticSearch)
	tr := &http.Transport{}

	portStr := strconv.Itoa(port)

	ret.baseUrl = "http://localhost:" + portStr + "/"
	ret.client = &http.Client{Transport: tr}
	ret.indiceName = "images_similar"
	return
}

func (e *ElasticSearch) getInfo() (ret string) {
	resp, _ := e.client.Get(e.baseUrl)
	body, _ := ioutil.ReadAll(resp.Body)
	ret = string(body)
	return
}

func (e *ElasticSearch) saveData(id string, fields []int64, line string) {

	strFields := make([]string, len(fields))
	for i := 0; i < len(fields); i++ {
		strFields[i] = strconv.Itoa(int(fields[i]))
	}

	data := `{"f" : "` + strings.Join(strFields, " ") + `", "line" : "` + line + `"}`
	req, err := http.NewRequest("PUT", e.baseUrl+e.indiceName+"/img/"+id, strings.NewReader(data))
	req.ContentLength = int64(len(data))
	resp, err := e.client.Do(req)
	if err != nil {
		fmt.Println(err, resp)
	}

	defer resp.Body.Close()
}

func (e *ElasticSearch) putMapping() {

	data := `
	{
		"mappings": {
			"img": {
				"properties": {
					"f": {
						"type": "string"
					}
				}
			}
		}
	}`
	req, err := http.NewRequest("PUT", e.baseUrl+e.indiceName+"/", strings.NewReader(data))
	req.ContentLength = int64(len(data))
	resp, err := e.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err, resp)
	}

}

func (e *ElasticSearch) deleteIndex() {

	req, err := http.NewRequest("DELETE", e.baseUrl+e.indiceName+"/", nil)
	resp, err := e.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err, resp)
	}
}

func (e *ElasticSearch) getImage(id string, limit int) (string, string) {

	resp, _ := e.client.Get(e.baseUrl + e.indiceName + "/img/" + id)
	body, _ := ioutil.ReadAll(resp.Body)

	var f interface{}

	err := json.Unmarshal(body, &f)

	if err != nil {
		fmt.Println(err)
	}

	line := f.(map[string]interface{})["_source"].(map[string]interface{})["line"].(string)

	tagString := f.(map[string]interface{})["_source"].(map[string]interface{})["f"].(string)

	fields := strings.Split(tagString, " ")

	if len(fields) >= limit {
		fields = fields[0:limit]
	}

	ret := strings.Join(fields, " ")

	return ret, line
}

func (e *ElasticSearch) getSimilar(tags, line string, maxResults int) PairList {

	vec, _ := ParseVectors(line)
	data := `{"size": ` + strconv.Itoa(maxResults) + `, "query" : {"match" : { "f" : "` + tags + `" }}}`

	req, err := http.NewRequest("GET", e.baseUrl+e.indiceName+"/img/_search", strings.NewReader(data))
	req.ContentLength = int64(len(data))
	resp, err := e.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err, resp)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	var iface interface{}

	err = json.Unmarshal(body, &iface)

	hits := iface.(map[string]interface{})["hits"].(map[string]interface{})["hits"].([]interface{})

	results := make(PairList, len(hits))

	for i := 0; i < len(hits); i++ {
		id := hits[i].(map[string]interface{})["_id"].(string)
		l2 := hits[i].(map[string]interface{})["_source"].(map[string]interface{})["line"].(string)
		vec2, _ := ParseVectors(l2)

		distance := VectorDistance(vec, vec2, 100000)

		results[i] = Pair{id, distance}
	}

	if err != nil {
		fmt.Println(err)
	}

	sort.Sort(results)

	return results
}
