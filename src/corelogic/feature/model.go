package feature

// Folder represents a single folder with a list its children names
type Folder struct {
	Children []string `json:"children"`
}

// Object represents either a Folder or a File, and it can recursively contain its own child Objects
type Object struct {
	Key      string             `json:"key"`
	Name     string             `json:"name"`
	Type     string             `json:"type"`
	Children map[string]*Object `json:"children"`
}
