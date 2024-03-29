package entities

type ModuleObject struct {
	ModuleID     string            `json:"moduleid"`
	ID           string            `json:"id"`
	RelativePath string            `json:"relative_path"`
	Status       int               `json:"status"`
	Config       map[string]string `json:"config"`
}

type ModuleCreate struct {
	ID           string            `json:"id"`
	RelativePath string            `json:"relative_path"`
	Status       int               `json:"status"`
	Config       map[string]string `json:"config"`
}

type ModuleUpdate struct {
	ModuleID     string            `json:"moduleid"`
	ID           string            `json:"id"`
	RelativePath string            `json:"relative_path"`
	Status       int               `json:"status"`
	Config       map[string]string `json:"config"`
}

type ModuleGet struct {
	ModuleIDs []string `json:"moduleids"`
	ZabbixCommun
}
