package entities

type TemplateDashboardObject struct {
	DashboardID   string `json:"dashboardid,omitempty"`
	Name          string `json:"name,omitempty"`
	TemplateID    string `json:"templateid,omitempty"`
	DisplayPeriod int    `json:"display_period,omitempty"`
	AutoStart     int    `json:"auto_start,omitempty"`
	UUID          string `json:"uuid,omitempty"`
}

type TemplateDashboardPage struct {
	DashboardPageID string                     `json:"dashboardpageid,omitempty"`
	Name            string                     `json:"name,omitempty"`
	DisplayPeriod   int                        `json:"display_period,omitempty"`
	Widgets         []TemplateDashboardWidgets `json:"widgets,omitempty"`
}

type TemplateDashboardWidgets struct {
	WidgetID string                    `json:"widgetid,omitempty"`
	Type     string                    `json:"type,omitempty"`
	Name     string                    `json:"name,omitempty"`
	X        int                       `json:"x,omitempty"`
	Y        int                       `json:"y,omitempty"`
	Width    int                       `json:"width,omitempty"`
	Height   int                       `json:"height,omitempty"`
	ViewMode int                       `json:"view_mode,omitempty"`
	Fields   []TemplateDashboardFields `json:"fields,omitempty"`
}

type TemplateDashboardFields struct {
	Type  int    `json:"type,omitempty"`
	Name  string `json:"name,omitempty"`
	Value any    `json:"value,omitempty"`
}

type TemplateDashboardCreate struct {
	Name          string                  `json:"name,omitempty"`
	TemplateID    string                  `json:"templateid,omitempty"`
	DisplayPeriod int                     `json:"display_period,omitempty"`
	AutoStart     int                     `json:"auto_start,omitempty"`
	UUID          string                  `json:"uuid,omitempty"`
	Pages         []TemplateDashboardPage `json:"pages,omitempty"`
}

type TemplateDashboardUpdate struct {
	DashboardID   string                  `json:"dashboardid,omitempty"`
	Name          string                  `json:"name,omitempty"`
	TemplateID    string                  `json:"templateid,omitempty"`
	DisplayPeriod int                     `json:"display_period,omitempty"`
	AutoStart     int                     `json:"auto_start,omitempty"`
	UUID          string                  `json:"uuid,omitempty"`
	Pages         []TemplateDashboardPage `json:"pages,omitempty"`
}

type TemplateDashboardResponse struct {
	DashboardIDs []string `json:"dashboardids,omitempty"`
}

type TemplateDashboardGet struct {
	DashboardIDs []string `json:"dashboardids,omitempty"`
	TemplateIDs  []string `json:"templateids,omitempty"`
	SelectPages  string   `json:"selectPages,omitempty"`
	ZabbixCommun
}
