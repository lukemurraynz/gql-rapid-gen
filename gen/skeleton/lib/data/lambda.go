package data

type RequestIdentity struct {
	Sub      string                 `json:"sub"`
	Issuer   string                 `json:"issuer"`
	Username string                 `json:"username"`
	Claims   map[string]interface{} `json:"claims"`
}
