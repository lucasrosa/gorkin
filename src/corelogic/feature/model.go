package feature

type Folder struct {
	Children []string `json:"children"`
}

type Object struct {
	Key      string             `json:"key"`
	Name     string             `json:"name"`
	Type     string             `json:"type"`
	Children map[string]*Object `json:"children"`
}
