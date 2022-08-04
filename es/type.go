package es

type Response struct {
	took int  `json:"took"`
	body body `json:"hits"`
}

type body struct {
	total    interface{} `json:"total"`
	maxScore float32     `json:"max_score"`
	hits     []hits      `json:"hits"`
}

type hits struct {
	id     int           `json:"_id"`
	score  int           `json:"_score"`
	source []interface{} `json:"_source"`
}
