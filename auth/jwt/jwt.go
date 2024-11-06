package jwt

type JWT struct {
	Header  Header  `json:"header"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	UserId    int    `json:"userId"`
	Role      string `json:"role"`
	CreatedAt int64  `json:"createdAt"`
	ValidFor  int64  `json:"validFor"`
}

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}
