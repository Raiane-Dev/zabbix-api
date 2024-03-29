package entities

type RoleObject struct {
	RoleID   string `json:"roleid,omitempty"`
	Name     string `json:"name,omitempty"`
	Type     int    `json:"type,omitempty"`
	Readonly int    `json:"readonly,omitempty"`
}

type RoleCreate struct {
	Name     string      `json:"name,omitempty"`
	Type     int         `json:"type,omitempty"`
	Readonly int         `json:"readonly,omitempty"`
	Rules    []RoleRules `json:"rules,omitempty"`
}
type RoleUpdate struct {
	RoleID   string      `json:"roleid,omitempty"`
	Name     string      `json:"name,omitempty"`
	Type     int         `json:"type,omitempty"`
	Readonly int         `json:"readonly,omitempty"`
	Rules    []RoleRules `json:"rules,omitempty"`
}

type RoleGet struct {
	RoleIDs     []string `json:"roleids,omitempty"`
	SelectRules string   `json:"selectRules,omitempty"`
	SelectUsers string   `json:"selectUsers,omitempty"`
	ZabbixCommun
}

type RoleRules struct {
	Ui                   []RoleUiElement   `json:"ui,omitempty"`
	UiDefaultAccess      int               `json:"ui.default_access,omitempty"`
	ServicesReadMode     int               `json:"services.read_mode,omitempty"`
	ServicesReadList     []RoleServiceList `json:"services.read_list,omitempty"`
	ServicesReadTag      []RoleServiceTag  `json:"services.read_tag,omitempty"`
	ServicesWriteMode    int               `json:"services.write_mode,omitempty"`
	ServicesWriteList    []RoleServiceList `json:"services.write_list,omitempty"`
	ServicesWriteTag     []RoleServiceTag  `json:"services.write_tag,omitempty"`
	Modules              []ModuleObject    `json:"modules,omitempty"`
	ModulesDefaultAccess int               `json:"modules.default_access,omitempty"`
	ApiAcess             int               `json:"apiacess,omitempty"`
	ApiMode              int               `json:"apimode,omitempty"`
	Api                  []string          `json:"api,omitempty"`
	Actions              []RoleAction      `json:"actions,omitempty"`
	ActionsDefaultAccess int               `json:"actions.default_access,omitempty"`
}

type RoleUiElement struct {
	Name   string `json:"name,omitempty"`
	Status int    `json:"status,omitempty"`
}

type RoleServiceList struct {
	ServiceID string `json:"serviceid,omitempty"`
}

type RoleServiceTag struct {
	Tag   string `json:"tag,omitempty"`
	Value string `json:"value,omitempty"`
}

type RoleAction struct {
	Name   string `json:"name,omitempty"`
	Status int    `json:"status,omitempty"`
}
