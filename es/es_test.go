package es_test

import (
	"es/es"
	"log"
	"testing"
	"time"
)

func TestNewESClient(t *testing.T) {
	client, err := es.NewESClient("http://localhost:9200", time.Second*3, nil)

	if err != nil {
		t.Fatal(err)
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"nickname": "pedro",
			},
		},
	}

	var q es.Query

	q.SetIndex("test")
	q.SetQuery(query)

	client.SetQuery(&q)

	res, err := client.Search()

	if err != nil {
		t.Fatal(err)
	}

	log.Println(res)

}
