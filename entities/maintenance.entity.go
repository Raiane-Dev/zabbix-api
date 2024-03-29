package entities

type MaintenanceObject struct {
	MaintenanceID   string `json:"maintenanceid,omitempty"`
	Name            string `json:"name,omitempty"`
	ActiveSince     string `json:"active_since,omitempty"`
	ActiveTill      string `json:"active_till,omitempty"`
	Description     string `json:"description,omitempty"`
	MaintenanceType int    `json:"maintenance_type,omitempty"`
	TagsEvaltype    int    `json:"tags_evaltype,omitempty"`
}

type MaintenanceTimePeriod struct {
	Period         int    `json:"period,omitempty"`
	TimePeriodType int    `json:"timeperiod_type,omitempty"`
	StartDate      string `json:"start_date,omitempty"`
	StartTime      string `json:"start_time,omitempty"`
	Every          int    `json:"every,omitempty"`
	DayofWeek      int    `json:"dayofweek,omitempty"`
	Day            int    `json:"day,omitempty"`
	Month          int    `json:"month,omitempty"`
}

type MaintenanceProblemTag struct {
	Tag      string `json:"tag,omitempty"`
	Operator int    `json:"operator,omitempty"`
	Value    string `json:"value,omitempty"`
}

type MaintenanceCreate struct {
	Name            string                  `json:"name,omitempty"`
	ActiveSince     string                  `json:"active_since,omitempty"`
	ActiveTill      string                  `json:"active_till,omitempty"`
	Description     string                  `json:"description,omitempty"`
	MaintenanceType int                     `json:"maintenance_type,omitempty"`
	TagsEvaltype    int                     `json:"tags_evaltype,omitempty"`
	Groups          []HostGroupObject       `json:"groups,omitempty"`
	Hosts           []HostObject            `json:"hosts,omitempty"`
	TimePeriods     []MaintenanceTimePeriod `json:"timeperiods,omitempty"`
	Tags            []MaintenanceProblemTag `json:"tags,omitempty"`
}

type MaintenanceUpdate struct {
	MaintenanceID   string                  `json:"maintenanceid,omitempty"`
	Name            string                  `json:"name,omitempty"`
	ActiveSince     string                  `json:"active_since,omitempty"`
	ActiveTill      string                  `json:"active_till,omitempty"`
	Description     string                  `json:"description,omitempty"`
	MaintenanceType int                     `json:"maintenance_type,omitempty"`
	TagsEvaltype    int                     `json:"tags_evaltype,omitempty"`
	Groups          []HostGroupObject       `json:"groups,omitempty"`
	Hosts           []HostObject            `json:"hosts,omitempty"`
	TimePeriods     []MaintenanceTimePeriod `json:"timeperiods,omitempty"`
	Tags            []MaintenanceProblemTag `json:"tags,omitempty"`
}

type MaintenanceGet struct {
	GroupIDs          []string `json:"groupids,omitempty"`
	HostIDs           []string `json:"hostids,omitempty"`
	MaintenanceIDs    []string `json:"maintenanceids,omitempty"`
	SelectHostGroups  string   `json:"selectHostGroups,omitempty"`
	SelectHosts       string   `json:"selectHosts,omitempty"`
	SelectTags        string   `json:"selectTags,omitempty"`
	SelectTimePeriods string   `json:"selectTimeperiods,omitempty"`
	ZabbixCommun
}
