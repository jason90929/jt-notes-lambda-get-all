package getAll

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Text string `json:"text"`
}
