package entities

type ModuleObject struct {
	ModuleID     string            `json:"moduleid,omitempty"`
	ID           string            `json:"id,omitempty"`
	RelativePath string            `json:"relative_path,omitempty"`
	Status       int               `json:"status,omitempty"`
	Config       map[string]string `json:"config,omitempty"`
}

type ModuleCreate struct {
	ID           string            `json:"id,omitempty"`
	RelativePath string            `json:"relative_path,omitempty"`
	Status       int               `json:"status,omitempty"`
	Config       map[string]string `json:"config,omitempty"`
}

type ModuleUpdate struct {
	ModuleID     string            `json:"moduleid,omitempty"`
	ID           string            `json:"id,omitempty"`
	RelativePath string            `json:"relative_path,omitempty"`
	Status       int               `json:"status,omitempty"`
	Config       map[string]string `json:"config,omitempty"`
}

type ModuleGet struct {
	ModuleIDs []string `json:"moduleids,omitempty"`
	ZabbixCommun
}
