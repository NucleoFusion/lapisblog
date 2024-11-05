package jwt

type JWT struct {
	Header  Header  `json:"header"`
	Payload Payload `json:"payload"`
}

type Payload struct {
}

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}
