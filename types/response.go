package types

type Response struct {
	Status  string `json: "code"`
	Message string `json: "message"`
}
