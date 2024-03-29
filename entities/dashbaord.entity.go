package entities

type DashboardCreate struct {
	Name          string               `json:"name,omitempty"`
	UserID        string               `json:"userid,omitempty"`
	Private       int                  `json:"private,omitempty"`
	DisplayPeriod int                  `json:"display_period,omitempty"`
	AutoStart     int                  `json:"auto_start,omitempty"`
	Pages         []DashboardPage      `json:"pages,omitempty"`
	Users         []DashboardUser      `json:"users,omitempty"`
	UsersGroups   []DashboardUserGroup `json:"usersGroups,omitempty"`
}

type DashboardUpdate struct {
	DashboardID   string               `json:"dashboardid,omitempty"`
	Name          string               `json:"name,omitempty"`
	UserID        string               `json:"userid,omitempty"`
	Private       int                  `json:"private,omitempty"`
	DisplayPeriod int                  `json:"display_period,omitempty"`
	AutoStart     int                  `json:"auto_start,omitempty"`
	Pages         []DashboardPage      `json:"pages,omitempty"`
	Users         []DashboardUser      `json:"users,omitempty"`
	UsersGroups   []DashboardUserGroup `json:"usersGroups,omitempty"`
}

type DashboardResponse struct {
	DashboardIDs []string `json:"dashboardids,omitempty"`
}

type DashboardGet struct {
	DashboardIDs     string `json:"dashboardids,omitempty"`
	SelectPages      string `json:"selectPages,omitempty"`
	SelectUsers      string `json:"selectUsers,omitempty"`
	SelectUserGroups string `json:"selectUserGroups,omitempty"`
	ZabbixCommun
}

type DashboardOutput struct {
	DashboardID   string
	Name          string               `json:"name,omitempty"`
	UserID        string               `json:"userid,omitempty"`
	Private       int                  `json:"private,omitempty"`
	DisplayPeriod int                  `json:"display_period,omitempty"`
	AutoStart     int                  `json:"auto_start,omitempty"`
	Pages         []DashboardPage      `json:"pages,omitempty"`
	Users         []DashboardUser      `json:"users,omitempty"`
	UserGroups    []DashboardUserGroup `json:"usergroups,omitempty"`
}

type DashboardUserGroup struct {
	UsrGrpID   string `json:"usrgrpid,omitempty"`
	Permission int    `json:"permission,omitempty"`
}

type DashboardUser struct {
	UserID     string `json:"userid,omitempty"`
	Permission int    `json:"permission,omitempty"`
}

type DashboardPage struct {
	DashboardPageID string            `json:"dashboard_pageid,omitempty"`
	Name            string            `json:"name,omitempty"`
	DisplayPeriod   int               `json:"display_period,omitempty"`
	Widgets         []DashboardWidget `json:"widgets,omitempty"`
}

type DashboardWidget struct {
	WidgetID   string                `json:"widgetid,omitempty"`
	Type       string                `json:"type,omitempty"`
	NameWidget string                `json:"name,omitempty"`
	X          string                `json:"x,omitempty"`
	Y          string                `json:"y,omitempty"`
	Width      string                `json:"width,omitempty"`
	Height     string                `json:"height,omitempty"`
	ViewMode   string                `json:"view_mode"`
	Fields     []ChildDashboardField `json:"fields,omitempty"`
}

type ChildDashboardField struct {
	Type      string `json:"type,omitempty"`
	NameField string `json:"name,omitempty"`
	Value     string `json:"value"`
}
