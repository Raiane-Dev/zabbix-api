package entities

type ReportObject struct {
	ReportID    string `json:"reportid,omitempty"`
	UserID      string `json:"userid,omitempty"`
	Name        string `json:"name,omitempty"`
	DashboardID string `json:"dashboardid,omitempty"`
	Period      int    `json:"period,omitempty"`
	Cycle       int    `json:"cycle,omitempty"`
	StartTime   int    `json:"starttime,omitempty"`
	WeekDays    int    `json:"weekdays,omitempty"`
	ActiveSince string `json:"active_since,omitempty"`
	ActiveTIll  string `json:"active_till,omitempty"`
	Subject     string `json:"subject,omitempty"`
	Message     string `json:"message,omitempty"`
	Status      int    `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
	State       int    `json:"state,omitempty"`
	LastSent    string `json:"lastsent,omitempty"`
	Info        string `json:"info,omitempty"`
}

type ReportUser struct {
	UserID       string `json:"userid,omitempty"`
	AccessUserID string `json:"access_userid,omitempty"`
	Exclude      int    `json:"exclude,omitempty"`
}

type ReportUserGroup struct {
	UsrGrpID     string `json:"usrgrpid,omitempty"`
	AccessUserID string `json:"access_userid,omitempty"`
}

type ReportCreate struct {
	UserID      string            `json:"userid,omitempty"`
	Name        string            `json:"name,omitempty"`
	DashboardID string            `json:"dashboardid,omitempty"`
	Period      int               `json:"period,omitempty"`
	Cycle       int               `json:"cycle,omitempty"`
	StartTime   int               `json:"starttime,omitempty"`
	WeekDays    int               `json:"weekdays,omitempty"`
	ActiveSince string            `json:"active_since,omitempty"`
	ActiveTIll  string            `json:"active_till,omitempty"`
	Subject     string            `json:"subject,omitempty"`
	Message     string            `json:"message,omitempty"`
	Status      int               `json:"status,omitempty"`
	Description string            `json:"description,omitempty"`
	State       int               `json:"state,omitempty"`
	LastSent    string            `json:"lastsent,omitempty"`
	Info        string            `json:"info,omitempty"`
	Users       []ReportUser      `json:"users,omitempty"`
	UserGroups  []ReportUserGroup `json:"user_groups,omitempty"`
}

type ReportUpdate struct {
	ReportID    string            `json:"reportid,omitempty"`
	UserID      string            `json:"userid,omitempty"`
	Name        string            `json:"name,omitempty"`
	DashboardID string            `json:"dashboardid,omitempty"`
	Period      int               `json:"period,omitempty"`
	Cycle       int               `json:"cycle,omitempty"`
	StartTime   int               `json:"starttime,omitempty"`
	WeekDays    int               `json:"weekdays,omitempty"`
	ActiveSince string            `json:"active_since,omitempty"`
	ActiveTIll  string            `json:"active_till,omitempty"`
	Subject     string            `json:"subject,omitempty"`
	Message     string            `json:"message,omitempty"`
	Status      int               `json:"status,omitempty"`
	Description string            `json:"description,omitempty"`
	State       int               `json:"state,omitempty"`
	LastSent    string            `json:"lastsent,omitempty"`
	Info        string            `json:"info,omitempty"`
	Users       []ReportUser      `json:"users,omitempty"`
	UserGroups  []ReportUserGroup `json:"user_groups,omitempty"`
}

type ReportResponse struct {
	ReportsIDs []string `json:"reportsids,omitempty"`
}

type ReportGet struct {
	ReportIDs        []string `json:"reportids,omitempty"`
	Expired          bool     `json:"expired,omitempty"`
	SelectUsers      string   `json:"selectUsers,omitempty"`
	SelectUserGroups string   `json:"selectUserGroups,omitempty"`
	ZabbixCommun
}
