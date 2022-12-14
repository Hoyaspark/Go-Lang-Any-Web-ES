package es

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

type Auth struct {
	username string
	password string
}

func NewAuth(username, password string) *Auth {
	return &Auth{
		username, password,
	}
}

type Query struct {
	index string
	query map[string]interface{}
}

func (q *Query) SetIndex(index string) {
	q.index = index
}

func (q *Query) SetQuery(query map[string]interface{}) {
	q.query = query
}

type ElasticSearch struct {
	url    string
	client *http.Client
	auth   *Auth
	mu     *sync.Mutex
	query  *Query
}

func (es *ElasticSearch) SetQuery(query *Query) {
	es.query = query
}

func (es *ElasticSearch) Ping(done chan<- struct{}, ch chan<- error) {

	defer func() {
		done <- struct{}{}
	}()

	res, err := http.Get(es.url)

	if err != nil {
		log.Println(err)
		ch <- err
		return
	}

	log.Println(res)

	if err != nil {
		log.Println(err)
		ch <- err
		return
	}

	return
}

func NewESClient(url string, timeout time.Duration, auth *Auth) (*ElasticSearch, error) {

	es := &ElasticSearch{
		url: url,
		client: &http.Client{
			Timeout: timeout,
		},
		auth:  auth,
		mu:    &sync.Mutex{},
		query: &Query{},
	}

	ch := make(chan error)
	done := make(chan struct{})

	go es.Ping(done, ch)

	for {
		select {
		case <-ch:
			return nil, <-ch
		case <-done:
			return es, nil
		}
	}
}

func (es *ElasticSearch) Search() (*Response, error) {

	b, err := json.Marshal(&es.query.query)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	r, err := http.NewRequest("POST", es.url+"/"+es.query.index+"/_search", bytes.NewBuffer(b))

	r.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if es.auth != nil {
		r.SetBasicAuth(es.auth.username, es.auth.password)
	}

	res, err := es.client.Do(r)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		log.Println("http status code is 400")
		return nil, err
	}

	var result Response

	err = json.NewDecoder(res.Body).Decode(&result)

	defer res.Body.Close()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &result, nil

}
