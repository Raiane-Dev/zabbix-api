package entities

type HistoryGet struct {
	History  int      `json:"history"`
	HostIDs  []string `json:"hostids"`
	ItemIDs  []string `json:"itemids"`
	TimeFrom string   `json:"timefrom"`
	TimeTill string   `json:"timetill"`
	ZabbixCommun
}

type HistoryOutput struct {
	ItemID string `json:"itemid"`
	Clock  string `json:"clock"`
	Value  string `json:"value"`
	NS     string `json:"ns"`
}
