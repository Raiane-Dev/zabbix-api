package entities

type ReportObject struct {
	ReportID    string `json:"reportid"`
	UserID      string `json:"userid"`
	Name        string `json:"name"`
	DashboardID string `json:"dashboardid"`
	Period      int    `json:"period"`
	Cycle       int    `json:"cycle"`
	StartTime   int    `json:"starttime"`
	WeekDays    int    `json:"weekdays"`
	ActiveSince string `json:"active_since"`
	ActiveTIll  string `json:"active_till"`
	Subject     string `json:"subject"`
	Message     string `json:"message"`
	Status      int    `json:"status"`
	Description string `json:"description"`
	State       int    `json:"state"`
	LastSent    string `json:"lastsent"`
	Info        string `json:"info"`
}

type ReportUser struct {
	UserID       string `json:"userid"`
	AccessUserID string `json:"access_userid"`
	Exclude      int    `json:"exclude"`
}

type ReportUserGroup struct {
	UsrGrpID     string `json:"usrgrpid"`
	AccessUserID string `json:"access_userid"`
}

type ReportCreate struct {
	UserID      string            `json:"userid"`
	Name        string            `json:"name"`
	DashboardID string            `json:"dashboardid"`
	Period      int               `json:"period"`
	Cycle       int               `json:"cycle"`
	StartTime   int               `json:"starttime"`
	WeekDays    int               `json:"weekdays"`
	ActiveSince string            `json:"active_since"`
	ActiveTIll  string            `json:"active_till"`
	Subject     string            `json:"subject"`
	Message     string            `json:"message"`
	Status      int               `json:"status"`
	Description string            `json:"description"`
	State       int               `json:"state"`
	LastSent    string            `json:"lastsent"`
	Info        string            `json:"info"`
	Users       []ReportUser      `json:"users"`
	UserGroups  []ReportUserGroup `json:"user_groups"`
}

type ReportUpdate struct {
	ReportID    string            `json:"reportid"`
	UserID      string            `json:"userid"`
	Name        string            `json:"name"`
	DashboardID string            `json:"dashboardid"`
	Period      int               `json:"period"`
	Cycle       int               `json:"cycle"`
	StartTime   int               `json:"starttime"`
	WeekDays    int               `json:"weekdays"`
	ActiveSince string            `json:"active_since"`
	ActiveTIll  string            `json:"active_till"`
	Subject     string            `json:"subject"`
	Message     string            `json:"message"`
	Status      int               `json:"status"`
	Description string            `json:"description"`
	State       int               `json:"state"`
	LastSent    string            `json:"lastsent"`
	Info        string            `json:"info"`
	Users       []ReportUser      `json:"users"`
	UserGroups  []ReportUserGroup `json:"user_groups"`
}

type ReportResponse struct {
	ReportsIDs []string `json:"reportsids"`
}

type ReportGet struct {
	ReportIDs        []string `json:"reportids"`
	Expired          bool     `json:"expired"`
	SelectUsers      string   `json:"selectUsers"`
	SelectUserGroups string   `json:"selectUserGroups"`
	ZabbixCommun
}
