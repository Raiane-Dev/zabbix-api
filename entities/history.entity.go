package entities

type HistoryGet struct {
	History  int      `json:"history,omitempty"`
	HostIDs  []string `json:"hostids,omitempty"`
	ItemIDs  []string `json:"itemids,omitempty"`
	TimeFrom string   `json:"timefrom,omitempty"`
	TimeTill string   `json:"timetill,omitempty"`
	ZabbixCommun
}

type HistoryObject struct {
	ItemID string `json:"itemid,omitempty"`
	Clock  string `json:"clock,omitempty"`
	Value  string `json:"value,omitempty"`
	NS     string `json:"ns,omitempty"`
}
