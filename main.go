package main

import (
	"bufio"
	"compress/gzip"
	"container/list"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const LIMIT1 = 10
const STEP1 = 100
const LIMIT2 = 1000000
const MAX_SCORE = 150000.0

func main() {

	action := flag.String("a", "", "action")
	limit := flag.Int("n", 1000, "number of imported images")
	path := flag.String("f", "/Volumes/ondra_zaloha/profi-neuralnet-20M.data.gz", "path to gziped file with data")
	esPort := flag.Int("e", 9200, "Elasticsearch port")
	k := flag.Int("k", 200, "how many most important vector components we will use")
	l := flag.Int("l", 100, "how many results we will retrieve from Elasticsearch")
	p := flag.Int("p", 8585, "server port")

	flag.Parse()

	if *action == "" {
		fmt.Println("No action specified.")
		return
	}

	var e error
	switch *action {
	case "server":
		e = server(*esPort, *p, *l, *k)
	case "parser":
		parser()
	case "import":
		e = importer(*limit, *path, *esPort, *k)
	default:
		fmt.Println("Don't know action", *action)
	}

	if e != nil {
		fmt.Println(e)
	}
}

func importer(limit int, path string, esPort int, maxVectors int) (e error) {

	es := NewElasticSearch(esPort)
	es.deleteIndex()

	file, e := os.Open(path)
	r, e := gzip.NewReader(file)
	br := bufio.NewReader(r)

	var header, strVectors, id string

	i := 0
	for e == nil && i < limit {
		fmt.Print("\r", i)
		i += 1
		header, e = br.ReadString('\n')
		id = ParseNameStr(header)
		strVectors, e = br.ReadString('\n')
		strVectors = strings.TrimSuffix(strVectors, "\n")
		fields := ParseVectorsToTags(strVectors, maxVectors)

		es.saveData(id, fields, strVectors)
	}
	fmt.Println()
	return
}

func parser() (e error) {

	limit1 := LIMIT1
	limit2 := LIMIT2

	loaded := make([]*Vectors, limit1)

	//fmt.Println("start", time.Now())
	path := "/Volumes/ondra_zaloha/profi-neuralnet-20M.data.gz"
	//path = "/Users/ondrejodchazel/projects/diplomka/data/tiny.gz"

	file, e := os.Open(path)
	r, e := gzip.NewReader(file)
	br := bufio.NewReader(r)

	var header, strVectors string

	i := 0
	for i < limit1*STEP1 && e == nil {
		fmt.Print("\rfirst ", i)

		header, e = br.ReadString('\n')
		strVectors, e = br.ReadString('\n')

		var element *Vectors
		element, e = NewVector(header, strVectors)
		if i%STEP1 == 0 {
			loaded[i/STEP1] = element
		}
		i += 1
	}

	file, e = os.Open(path)
	r, e = gzip.NewReader(file)
	br = bufio.NewReader(r)
	i = 0
	for i < limit2 && e == nil {
		fmt.Print("\rsecond ", i)

		header, e = br.ReadString('\n')
		strVectors, e = br.ReadString('\n')

		var element *Vectors
		element, e = NewVector(header, strVectors)
		for j := 0; j < limit1; j++ {
			vec := loaded[j]
			if vec != nil {
				vec.compare(element)
			}
		}
		i += 1
	}

	fmt.Print("\r                    \r")

	for _, v := range loaded {
		v.print()
	}

	return

}

func VectorDistance(v1, v2 [4096]float64, max_score float64) (ret float64) {
	for i := 0; i < 4096; i++ {
		ret += math.Abs(v1[i] - v2[i])
		if ret > max_score {
			return ret
		}
	}
	return
}

func ParseNameStr(line string) string {
	return line[49:59]
}

func ParseName(line string) (i int64, e error) {
	line = line[49:59]
	i, e = strconv.ParseInt(line, 10, 0)
	return
}

func ParseVectorsToTags(line string, limit int) []int64 {

	results := make(PairList, 4096)

	line = strings.TrimSuffix(line, "\n")

	strVectors := strings.Split(line, " ")
	j := 0
	var i int64
	for i = 0; i < 4096; i++ {

		fl, e := strconv.ParseFloat(strVectors[i], 32)

		if e != nil {
			fmt.Println(e)
		}

		if fl > 0.0 {
			results[j] = Pair{int64(j), fl}
			j++

		}
	}

	results = results[0:j]

	//sort.Sort(sort.Reverse(results))

	if len(results) > limit {
		results = results[0:limit]
	}

	ret := make([]int64, len(results))

	for i := 0; i < len(ret); i++ {
		ret[i] = results[i].Key.(int64)
	}

	return ret
}

type Pair struct {
	Key   interface{}
	Value float64
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func ParseVectors(line string) (ret [4096]float64, e error) {
	line = strings.TrimSuffix(line, "\n")
	strVectors := strings.Split(line, " ")
	for i := 0; i < 4096; i++ {
		ret[i], e = strconv.ParseFloat(strVectors[i], 0)
	}
	return
}

type Vectors struct {
	id      int64
	vectors [4096]float64
	list    *list.List
}

func (v *Vectors) print() {
	fmt.Println(v.id)
	for e := v.list.Front(); e != nil; e = e.Next() {

		res := *(e.Value.(*Result))

		fmt.Println(" ", res.id, res.score)
	}
}

func (v *Vectors) compare(comp *Vectors) {
	maxLen := 10

	if v.id == comp.id {
		return
	}
	score := VectorDistance(v.vectors, comp.vectors, MAX_SCORE)

	if score > MAX_SCORE {
		return
	}

	res := NewResult(comp.id, score)

	if v.list.Len() == maxLen {
		lastScore := v.list.Back().Value.(*Result).score
		if lastScore <= score {
			return
		} else {
			v.list.Remove(v.list.Back())
			newLastScore := v.list.Back().Value.(*Result).score
			if newLastScore <= score {
				v.list.PushBack(res)
				return
			}
		}
	}

	el := v.list.Front()
	if el == nil {
		v.list.PushFront(res)
		return
	}

	for el != nil {
		currentScore := el.Value.(*Result).score
		if currentScore > score {
			v.list.InsertBefore(res, el)
			el = nil
		} else {
			el = el.Next()
		}
	}
}

func NewVector(header, vector string) (ret *Vectors, e error) {
	ret = new(Vectors)
	ret.id, e = ParseName(header)
	ret.vectors, e = ParseVectors(vector)
	ret.list = list.New()
	return
}

type Result struct {
	id    int64
	score float64
}

func NewResult(id int64, score float64) (ret *Result) {
	ret = new(Result)
	ret.id = id
	ret.score = score
	return
}
