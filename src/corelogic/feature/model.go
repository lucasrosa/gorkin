package feature

type Feature struct {
	Name      string   `json:"name"`
	Scenarios []string `json:"scenarios"`
}

type Folder struct {
	Children []string `json:"children"`
}

type Object struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	Children []Object `json:"children"`
}
