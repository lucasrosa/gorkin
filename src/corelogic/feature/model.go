package feature

type Feature struct {
	Name string `json:"name"`
	Scenarios []string `json:"scenarios"`
}

type Folder struct {
	Name string `json:"name"`
	Children []string `json:"children"`
}