package entities

type IconMapObject struct {
	IconmapID     string `json:"iconmapid,omitempty"`
	DefaultIconID string `json:"default_iconid,omitempty"`
	Name          string `json:"name,omitempty"`
}

type IconMapResponse struct {
	IconMapIDs []string `json:"iconmapids,omitempty"`
}

type IconMapping struct {
	IconID        string `json:"iconid,omitempty"`
	Expression    string `json:"expression,omitempty"`
	InventoryLink int    `json:"inventory_link,omitempty"`
	SortOrder     int    `json:"sortorder,omitempty"`
}

type IconMapCreate struct {
	DefaultIconID string        `json:"default_iconid,omitempty"`
	Name          string        `json:"name,omitempty"`
	Mappings      []IconMapping `json:"mappings,omitempty"`
}

type IconMapUpdate struct {
	IconmapID     string        `json:"iconmapid,omitempty"`
	DefaultIconID string        `json:"default_iconid,omitempty"`
	Name          string        `json:"name,omitempty"`
	Mappings      []IconMapping `json:"mappings,omitempty"`
}

type IconMapGet struct {
	MappingIDs     []string `json:"mappingids,omitempty"`
	SysmapIDs      []string `json:"sysmapids,omitempty"`
	SelectMappings string   `json:"selectMappings,omitempty"`
	ZabbixCommun
}
