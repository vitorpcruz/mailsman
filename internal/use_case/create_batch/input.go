package create_batch

type Input struct {
	Title   string   `json:"title"`
	Message string   `json:"message"`
	Emails  []string `json:"emails"`
}
