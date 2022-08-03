package es

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"sync"
)

type Query struct {
	index string
	query []byte
}

func (q *Query) SetIndex(index string) {
	q.index = index
}

type ElasticSearch struct {
	addr  *net.TCPAddr
	mu    *sync.Mutex
	query *Query
}

func (es *ElasticSearch) Ping() error {
	res, err := http.Get(es.addr.String())

	if err != nil {
		log.Println(err)
		return err
	}
	var buf []byte

	err = json.NewDecoder(res.Body).Decode(&buf)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func NewESClient(host string, port int, zone string) (*ElasticSearch, error) {
	es := &ElasticSearch{
		addr: &net.TCPAddr{
			IP:   net.IP(host),
			Port: port,
			Zone: zone,
		},
		mu:    &sync.Mutex{},
		query: &Query{},
	}

	err := es.Ping()

	if err != nil {
		return nil, err
	}
	return es, nil
}

func (es *ElasticSearch) Do() {

}
