package entities

type MaintenanceObject struct {
	MaintenanceID   string `json:"maintenanceid"`
	Name            string `json:"name"`
	ActiveSince     string `json:"active_since"`
	ActiveTill      string `json:"active_till"`
	Description     string `json:"description"`
	MaintenanceType int    `json:"maintenance_type"`
	TagsEvaltype    int    `json:"tags_evaltype"`
}

type MaintenanceTimePeriod struct {
	Period         int    `json:"period"`
	TimePeriodType int    `json:"timeperiod_type"`
	StartDate      string `json:"start_date"`
	StartTime      string `json:"start_time"`
	Every          int    `json:"every"`
	DayofWeek      int    `json:"dayofweek"`
	Day            int    `json:"day"`
	Month          int    `json:"month"`
}

type MaintenanceProblemTag struct {
	Tag      string `json:"tag"`
	Operator int    `json:"operator"`
	Value    string `json:"value"`
}

type MaintenanceCreate struct {
	Name            string                  `json:"name"`
	ActiveSince     string                  `json:"active_since"`
	ActiveTill      string                  `json:"active_till"`
	Description     string                  `json:"description"`
	MaintenanceType int                     `json:"maintenance_type"`
	TagsEvaltype    int                     `json:"tags_evaltype"`
	Groups          []HostGroupObject       `json:"groups"`
	Hosts           []HostObject            `json:"hosts"`
	TimePeriods     []MaintenanceTimePeriod `json:"timeperiods"`
	Tags            []MaintenanceProblemTag `json:"tags"`
}

type MaintenanceUpdate struct {
	MaintenanceID   string                  `json:"maintenanceid"`
	Name            string                  `json:"name"`
	ActiveSince     string                  `json:"active_since"`
	ActiveTill      string                  `json:"active_till"`
	Description     string                  `json:"description"`
	MaintenanceType int                     `json:"maintenance_type"`
	TagsEvaltype    int                     `json:"tags_evaltype"`
	Groups          []HostGroupObject       `json:"groups"`
	Hosts           []HostObject            `json:"hosts"`
	TimePeriods     []MaintenanceTimePeriod `json:"timeperiods"`
	Tags            []MaintenanceProblemTag `json:"tags"`
}

type MaintenanceGet struct {
	GroupIDs          []string `json:"groupids"`
	HostIDs           []string `json:"hostids"`
	MaintenanceIDs    []string `json:"maintenanceids"`
	SelectHostGroups  string   `json:"selectHostGroups"`
	SelectHosts       string   `json:"selectHosts"`
	SelectTags        string   `json:"selectTags"`
	SelectTimePeriods string   `json:"selectTimeperiods"`
	ZabbixCommun
}
