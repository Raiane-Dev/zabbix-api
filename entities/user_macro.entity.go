package entities

type UserMacroObject struct {
	GlobalMacroID string `json:"globalmacroid"`
	Macro         string `json:"macro"`
	Value         string `json:"value"`
	Type          int    `json:"type"`
	Description   string `json:"description"`
}
