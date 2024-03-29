package entities

type RoleObject struct {
	RoleID   string `json:"roleid"`
	Name     string `json:"name"`
	Type     int    `json:"type"`
	Readonly int    `json:"readonly"`
}

type RoleCreate struct {
	Name     string      `json:"name"`
	Type     int         `json:"type"`
	Readonly int         `json:"readonly"`
	Rules    []RoleRules `json:"rules"`
}
type RoleUpdate struct {
	RoleID   string      `json:"roleid"`
	Name     string      `json:"name"`
	Type     int         `json:"type"`
	Readonly int         `json:"readonly"`
	Rules    []RoleRules `json:"rules"`
}

type RoleGet struct {
	RoleIDs     []string `json:"roleids"`
	SelectRules string   `json:"selectRules"`
	SelectUsers string   `json:"selectUsers"`
	ZabbixCommun
}

type RoleRules struct {
	Ui                   []RoleUiElement
	UiDefaultAccess      int
	ServicesReadMode     int
	ServicesReadList     []RoleServiceList
	ServicesReadTag      []RoleServiceTag
	ServicesWriteMode    int
	ServicesWriteList    []RoleServiceList
	ServicesWriteTag     []RoleServiceTag
	Modules              []ModuleObject
	ModulesDefaultAccess int
	ApiAcess             int
	ApiMode              int
	Api                  []string
	Actions              []RoleAction
	ActionsDefaultAccess int
}

type RoleUiElement struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type RoleServiceList struct {
	ServiceID string `json:"serviceid"`
}

type RoleServiceTag struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type RoleAction struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
}
