package entities

type GraphItemGet struct {
	GraphIDs     []string `json:"graphids,omitempty"`
	ItemIDs      []string `json:"itemids,omitempty"`
	Type         int      `json:"type,omitempty"`
	SelectGraphs string   `json:"selectGraphs,omitempty"`
	ZabbixCommun
}

type GraphItemObject struct {
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
