package entities

type IconMapObject struct {
	IconmapID     string `json:"iconmapid"`
	DefaultIconID string `json:"default_iconid"`
	Name          string `json:"name"`
}

type IconMapResponse struct {
	IconMapIDs []string `json:"iconmapids"`
}

type IconMapping struct {
	IconID        string `json:"iconid"`
	Expression    string `json:"expression"`
	InventoryLink int    `json:"inventory_link"`
	SortOrder     int    `json:"sortorder"`
}

type IconMapCreate struct {
	DefaultIconID string        `json:"default_iconid"`
	Name          string        `json:"name"`
	Mappings      []IconMapping `json:"mappings"`
}

type IconMapUpdate struct {
	IconmapID     string        `json:"iconmapid"`
	DefaultIconID string        `json:"default_iconid"`
	Name          string        `json:"name"`
	Mappings      []IconMapping `json:"mappings"`
}

type IconMapGet struct {
	MappingIDs     []string `json:"mappingids"`
	SysmapIDs      []string `json:"sysmapids"`
	SelectMappings string   `json:"selectMappings"`
	ZabbixCommun
}
