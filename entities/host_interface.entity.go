package entities

type HostInterfaceObject struct {
	InterfaceID  string                 `json:"interfaceid,omitempty"`
	Available    int                    `json:"available,omitempty"`
	HostID       string                 `json:"hostid,omitempty"`
	Type         int                    `json:"type,omitempty"`
	IP           string                 `json:"ip,omitempty"`
	DNS          string                 `json:"dns,omitempty"`
	Port         string                 `json:"port,omitempty"`
	UseIP        int                    `json:"useip,omitempty"`
	Main         int                    `json:"main,omitempty"`
	Details      []HostInterfaceDetails `json:"details,omitempty"`
	DisableUntil string                 `json:"disable_until,omitempty"`
	Error        string                 `json:"error,omitempty"`
	ErrorsFrom   string                 `json:"errors_from,omitempty"`
}

type HostInterfaceCreate struct {
	Available    int                    `json:"available,omitempty"`
	HostID       string                 `json:"hostid,omitempty"`
	Type         int                    `json:"type,omitempty"`
	IP           string                 `json:"ip,omitempty"`
	DNS          string                 `json:"dns,omitempty"`
	Port         string                 `json:"port,omitempty"`
	UseIP        int                    `json:"useip,omitempty"`
	Main         int                    `json:"main,omitempty"`
	Details      []HostInterfaceDetails `json:"details,omitempty"`
	DisableUntil string                 `json:"disable_until,omitempty"`
	Error        string                 `json:"error,omitempty"`
	ErrorsFrom   string                 `json:"errors_from,omitempty"`
}

type HostInterfaceUpdate struct {
	InterfaceID  string                 `json:"interfaceid,omitempty"`
	Available    int                    `json:"available,omitempty"`
	HostID       string                 `json:"hostid,omitempty"`
	Type         int                    `json:"type,omitempty"`
	IP           string                 `json:"ip,omitempty"`
	DNS          string                 `json:"dns,omitempty"`
	Port         string                 `json:"port,omitempty"`
	UseIP        int                    `json:"useip,omitempty"`
	Main         int                    `json:"main,omitempty"`
	Details      []HostInterfaceDetails `json:"details,omitempty"`
	DisableUntil string                 `json:"disable_until,omitempty"`
	Error        string                 `json:"error,omitempty"`
	ErrorsFrom   string                 `json:"errors_from,omitempty"`
}

type HostInterfaceResponse struct {
	InterfaceIDs []string `json:"interfaceids,omitempty"`
}

type HostInterfaceGet struct {
	HostIDs      []string `json:"hostids,omitempty"`
	InterfaceIDs []string `json:"interfaceids,omitempty"`
	ItemIDs      []string `json:"itemids,omitempty"`
	TriggerIDs   []string `json:"triggerids,omitempty"`
	SelectItems  string   `json:"selectItems,omitempty"`
	LimitSelects int      `json:"limitSelects,omitempty"`
	ZabbixCommun
}

type HostInterfaceDetails struct {
	Version        int    `json:"version,omitempty"`
	Bulk           int    `json:"bulk,omitempty"`
	Community      string `json:"community,omitempty"`
	MaxRepetitions int    `json:"max_repetitions,omitempty"`
	SecurityName   string `json:"securityname,omitempty"`
	SecurityLevel  int    `json:"securitylevel,omitempty"`
	AuthPassPhrase string `json:"authpassphrase,omitempty"`
	PrivPassPhrase string `json:"privpassphrase,omitempty"`
	AuthProtocol   int    `json:"authprotocol,omitempty"`
	PrivProtocol   int    `json:"privprotocol,omitempty"`
	ContextName    string `json:"contextname,omitempty"`
}

type HostInterfaceMassAdd struct {
	Hosts      []map[string]string   `json:"hosts,omitempty"`
	Interfaces []HostInterfaceObject `json:"interfaces,omitempty"`
}

type HostInterfaceMassRemove struct {
	HostIDs    []string              `json:"hostids,omitempty"`
	Interfaces []HostInterfaceObject `json:"interfaces,omitempty"`
}

type HostInterfaceReplaceHostInterface struct {
	HostID     string                `json:"hostid,omitempty"`
	Interfaces []HostInterfaceObject `json:"interfaces,omitempty"`
}
