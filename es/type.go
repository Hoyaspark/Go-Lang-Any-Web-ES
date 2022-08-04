package es

type Response struct {
	Took int  `json:"took"`
	Body body `json:"hits"`
}

type body struct {
	Total    interface{} `json:"total"`
	MaxScore float32     `json:"max_score"`
	Hits     []hits      `json:"hits"`
}

type hits struct {
	Id     string      `json:"_id"`
	Score  float32     `json:"_score"`
	Source interface{} `json:"_source"`
}
