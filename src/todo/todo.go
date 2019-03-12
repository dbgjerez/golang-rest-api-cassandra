package todo

type Todo struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
