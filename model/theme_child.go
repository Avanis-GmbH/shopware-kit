package model

type ThemeChild struct {
	Id       string `json:"id,omitempty"`
	ParentId string `json:"parentId"` // required
	ChildId  string `json:"childId"`  // required
	Parent   *Theme `json:"parentTheme,omitempty"`
	Child    *Theme `json:"childTheme,omitempty"`
}
