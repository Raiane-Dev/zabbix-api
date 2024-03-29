package entities

type MacroType int

const (
	TextMacro   MacroType = iota // 0 - Text macro
	SecretMacro                  // 1 - Secret macro
	VaultSecret                  // 2 - Vault secret
)

type UserMacroObject struct {
	GlobalMacroID string `json:"globalmacroid,omitempty"`
	Macro         string `json:"macro,omitempty"`
	Value         string `json:"value,omitempty"`
	Type          int    `json:"type,omitempty"`
	Description   string `json:"description,omitempty"`
}

type UserMacroCreate struct {
	Macro       string `json:"macro,omitempty"`
	Value       string `json:"value,omitempty"`
	Type        int    `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

type UserMacroUpdate struct {
	GlobalMacroID string `json:"globalmacroid,omitempty"`
	Macro         string `json:"macro,omitempty"`
	Value         string `json:"value,omitempty"`
	Type          int    `json:"type,omitempty"`
	Description   string `json:"description,omitempty"`
}

type HostMacroObject struct {
	HostMacro   string    `json:"hostmacroid,omitempty"`
	HostID      string    `json:"hostid,omitempty"`
	Macro       string    `json:"macro,omitempty"`
	Value       string    `json:"value,omitempty"`
	Type        MacroType `json:"type,omitempty"`
	Description string    `json:"description,omitempty"`
	Automatic   int       `json:"automatic,omitempty"`
}
