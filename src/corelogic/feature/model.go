package feature

type Folder struct {
	Children []string `json:"children"`
}

type Object struct {
	ID       string             `json:"id"`
	Name     string             `json:"name"`
	Type     string             `json:"type"`
	Children map[string]*Object `json:"children"`
}
