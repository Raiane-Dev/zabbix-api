package entities

type ValueMapObject struct {
	ValueMapID string          `json:"valuemapid,omitempty"`
	HostID     string          `json:"hostid,omitempty"`
	Name       string          `json:"name,omitempty"`
	Mappings   []ValueMappings `json:"mappings,omitempty"`
	UUID       string          `json:"uuid,omitempty"`
}

type ValueMappings struct {
	Type     int    `json:"type,omitempty"`
	Value    string `json:"value,omitempty"`
	NewValue string `json:"newvalue,omitempty"`
}

type ValueMapCreate struct {
	HostID   string          `json:"hostid,omitempty"`
	Name     string          `json:"name,omitempty"`
	Mappings []ValueMappings `json:"mappings,omitempty"`
	UUID     string          `json:"uuid,omitempty"`
}

type ValueMapUpdate struct {
	ValueMapID string          `json:"valuemapid,omitempty"`
	HostID     string          `json:"hostid,omitempty"`
	Name       string          `json:"name,omitempty"`
	Mappings   []ValueMappings `json:"mappings,omitempty"`
	UUID       string          `json:"uuid,omitempty"`
}

type ValueMapResponse struct {
	ValueMapIDs []string `json:"valuemapids,omitempty"`
}

type ValueMapGet struct {
	ValueMapIDs    []string `json:"valuemapids,omitempty"`
	SelectMappings string   `json:"selectMappings,omitempty"`
	ZabbixCommun
}
