package model

type ThemeChild struct {
	Id       string `json:"id,omitempty"`
	ParentId string `json:"parentId,omitempty"` // required
	ChildId  string `json:"childId,omitempty"`  // required
	Parent   *Theme `json:"parentTheme,omitempty"`
	Child    *Theme `json:"childTheme,omitempty"`
}
