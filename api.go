package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

func similarHTMLHandler(w http.ResponseWriter, r *http.Request, es *ElasticSearch, l, k int) {

	var content bytes.Buffer
	id := r.URL.Query().Get("id")

	tags, line := es.getImage(id, k)

	content.WriteString("<img src=\"http://mufin.fi.muni.cz/profimedia/bigImages/" + id + "\"><br><hr>")

	similar := es.getSimilar(tags, line, l)

	for i := 0; i < len(similar); i++ {

		score := strconv.FormatFloat(similar[i].Value, 'f', 2, 32)

		content.WriteString(score + "<img src=\"http://mufin.fi.muni.cz/profimedia/bigImages/" + similar[i].Key.(string) + "\"><br>")
	}

	fmt.Fprintf(w, "<html>"+content.String()+"</html>")
}

func similarJSONHandler(w http.ResponseWriter, r *http.Request, es *ElasticSearch, l, k int) {
	var content bytes.Buffer
	id := r.URL.Query().Get("id")

	tags, line := es.getImage(id, k)

	content.WriteString("[")

	similar := es.getSimilar(tags, line, l)

	first := true
	for i := 0; i < len(similar); i++ {

		if similar[i].Value < 1900 {

			score := strconv.FormatFloat(similar[i].Value, 'f', 2, 32)

			if first == true {
				first = false
			} else {
				content.WriteString(",")
			}

			content.WriteString("{\"id\":\"" + similar[i].Key.(string) + "\",\"score\":" + score + "}")
		}
	}

	content.WriteString("]")

	w.Header().Add("Access-Control-Allow-Origin", "*")

	fmt.Fprintf(w, content.String())

}

func server(esPort, port, l, k int) (e error) {

	es := NewElasticSearch(esPort)

	http.HandleFunc("/similar_html", func(w http.ResponseWriter, r *http.Request) {
		similarHTMLHandler(w, r, es, l, k)
	})

	http.HandleFunc("/similar", func(w http.ResponseWriter, r *http.Request) {
		similarJSONHandler(w, r, es, l, k)
	})

	fmt.Println("localhost:8585/similar?id=0000000003")

	http.ListenAndServe(":"+strconv.Itoa(port), nil)

	return
}
