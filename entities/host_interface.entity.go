package entities

type HostInterfaceObject struct {
	InterfaceID  string                 `json:"interfaceid"`
	Available    int                    `json:"available"`
	HostID       string                 `json:"hostid"`
	Type         int                    `json:"type"`
	IP           string                 `json:"ip"`
	DNS          string                 `json:"dns"`
	Port         string                 `json:"port"`
	UseIP        int                    `json:"useip"`
	Main         int                    `json:"main"`
	Details      []HostInterfaceDetails `json:"details"`
	DisableUntil string                 `json:"disable_until"`
	Error        string                 `json:"error"`
	ErrorsFrom   string                 `json:"errors_from"`
}

type HostInterfaceCreate struct {
	Available    int                    `json:"available"`
	HostID       string                 `json:"hostid"`
	Type         int                    `json:"type"`
	IP           string                 `json:"ip"`
	DNS          string                 `json:"dns"`
	Port         string                 `json:"port"`
	UseIP        int                    `json:"useip"`
	Main         int                    `json:"main"`
	Details      []HostInterfaceDetails `json:"details"`
	DisableUntil string                 `json:"disable_until"`
	Error        string                 `json:"error"`
	ErrorsFrom   string                 `json:"errors_from"`
}

type HostInterfaceUpdate struct {
	InterfaceID  string                 `json:"interfaceid"`
	Available    int                    `json:"available"`
	HostID       string                 `json:"hostid"`
	Type         int                    `json:"type"`
	IP           string                 `json:"ip"`
	DNS          string                 `json:"dns"`
	Port         string                 `json:"port"`
	UseIP        int                    `json:"useip"`
	Main         int                    `json:"main"`
	Details      []HostInterfaceDetails `json:"details"`
	DisableUntil string                 `json:"disable_until"`
	Error        string                 `json:"error"`
	ErrorsFrom   string                 `json:"errors_from"`
}

type HostInterfaceResponse struct {
	InterfaceIDs []string `json:"interfaceids"`
}

type HostInterfaceGet struct {
	HostIDs      []string `json:"hostids"`
	InterfaceIDs []string `json:"interfaceids"`
	ItemIDs      []string `json:"itemids"`
	TriggerIDs   []string `json:"triggerids"`
	SelectItems  string   `json:"selectItems"`
	LimitSelects int      `json:"limitSelects"`
	ZabbixCommun
}

type HostInterfaceDetails struct {
	Version        int    `json:"version"`
	Bulk           int    `json:"bulk"`
	Community      string `json:"community"`
	MaxRepetitions int    `json:"max_repetitions"`
	SecurityName   string `json:"securityname"`
	SecurityLevel  int    `json:"securitylevel"`
	AuthPassPhrase string `json:"authpassphrase"`
	PrivPassPhrase string `json:"privpassphrase"`
	AuthProtocol   int    `json:"authprotocol"`
	PrivProtocol   int    `json:"privprotocol"`
	ContextName    string `json:"contextname"`
}

type HostInterfaceMassAdd struct {
	Hosts      []map[string]string   `json:"hosts"`
	Interfaces []HostInterfaceObject `json:"interfaces"`
}

type HostInterfaceMassRemove struct {
	HostIDs    []string              `json:"hostids"`
	Interfaces []HostInterfaceObject `json:"interfaces"`
}

type HostInterfaceReplaceHostInterface struct {
	HostID     string                `json:"hostid"`
	Interfaces []HostInterfaceObject `json:"interfaces"`
}
