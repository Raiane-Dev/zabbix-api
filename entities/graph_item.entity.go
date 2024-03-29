package entities

type GraphItemGet struct {
	GraphIDs     []string `json:"graphids"`
	ItemIDs      []string `json:"itemids"`
	Type         int      `json:"type"`
	SelectGraphs string   `json:"selectGraphs"`
	ZabbixCommun
}

type GraphItemOutput struct {
	GItemID   string `json:"gitemid,omitempty"`
	Color     string `json:"color,omitempty"`
	ItemID    string `json:"itemid,omitempty"`
	CalcFNC   int    `json:"calc_fnc,omitempty"`
	Drawtype  int    `json:"drawtype,omitempty"`
	GraphID   string `json:"graphid,omitempty"`
	Sortorder int    `json:"sortorder,omitempty"`
	Type      int    `json:"type,omitempty"`
	YAxisSide int    `json:"yaxisside,omitempty"`
}
