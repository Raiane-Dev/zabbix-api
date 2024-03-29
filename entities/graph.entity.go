package entities

type GraphCreate struct {
	UUID           string            `json:"uuid,omitempty"`
	Height         int               `json:"height,omitempty"`
	Width          int               `json:"width,omitempty"`
	Name           string            `json:"name,omitempty"`
	Flags          int               `json:"flags,omitempty"`
	Graphtype      int               `json:"graphtype,omitempty"`
	PercentLeft    float32           `json:"percent_left,omitempty"`
	PercentRight   float32           `json:"percent_right,omitempty"`
	Show3D         int               `json:"show_3d,omitempty"`
	ShowLegend     int               `json:"show_legend,omitempty"`
	ShowWorkPeriod int               `json:"show_work_period,omitempty"`
	ShowTriggers   int               `json:"show_triggers,omitempty"`
	TemplateID     string            `json:"templateid,omitempty"`
	YAxisMax       float32           `json:"yaxismax,omitempty"`
	YAxisMin       float32           `json:"yaxismin,omitempty"`
	YmaxItemID     string            `json:"ymax_itemid,omitempty"`
	YMaxType       int               `json:"ymax_type,omitempty"`
	YMinItemID     string            `json:"ymin_itemid,omitempty"`
	YMinType       int               `json:"ymin_type,omitempty"`
	GItems         []GraphItemObject `json:"gitems,omitempty"`
}

type GraphUpdate struct {
	GraphID        string            `json:"graphid,omitempty"`
	UUID           string            `json:"uuid,omitempty"`
	Height         int               `json:"height,omitempty"`
	Width          int               `json:"width,omitempty"`
	Name           string            `json:"name,omitempty"`
	Flags          int               `json:"flags,omitempty"`
	Graphtype      int               `json:"graphtype,omitempty"`
	PercentLeft    float32           `json:"percent_left,omitempty"`
	PercentRight   float32           `json:"percent_right,omitempty"`
	Show3D         int               `json:"show_3d,omitempty"`
	ShowLegend     int               `json:"show_legend,omitempty"`
	ShowWorkPeriod int               `json:"show_work_period,omitempty"`
	ShowTriggers   int               `json:"show_triggers,omitempty"`
	TemplateID     string            `json:"templateid,omitempty"`
	YAxisMax       float32           `json:"yaxismax,omitempty"`
	YAxisMin       float32           `json:"yaxismin,omitempty"`
	YmaxItemID     string            `json:"ymax_itemid,omitempty"`
	YMaxType       int               `json:"ymax_type,omitempty"`
	YMinItemID     string            `json:"ymin_itemid,omitempty"`
	YMinType       int               `json:"ymin_type,omitempty"`
	GItems         []GraphItemObject `json:"gitems,omitempty"`
}

type GraphResponse struct {
	GraphIDs []string
}

type GraphGet struct {
	GraphIDs             []string `json:"graphids,omitempty"`
	GroupIDs             []string `json:"groupids,omitempty"`
	TemplateIDs          []string `json:"templateids,omitempty"`
	HostIDs              []string `json:"hostids,omitempty"`
	ItemIDs              []string `json:"itemids,omitempty"`
	Templated            bool     `json:"templated,omitempty"`
	Inherited            bool     `json:"inherited,omitempty"`
	ExpandName           string   `json:"expandName,omitempty"`
	SelectHostGroups     string   `json:"selectHostGroups,omitempty"`
	SelectTemplateGroups string   `json:"selectTemplateGroups,omitempty"`
	SelectTemplates      string   `json:"selectTemplates,omitempty"`
	SelectHosts          string   `json:"selectHosts,omitempty"`
	SelectItems          string   `json:"selectItems,omitempty"`
	SelectGraphDiscovery string   `json:"selectGraphDiscovery,omitempty"`
	SelectGraphItems     string   `json:"selectGraphItems,omitempty"`
	SelectDiscoveryRule  string   `json:"selectDiscoveryRule,omitempty"`
	ZabbixCommun
}

type GraphObject struct {
	GraphID        string            `json:"graphid,omitempty"`
	Height         string            `json:"height,omitempty"`
	Width          string            `json:"width,omitempty"`
	Name           string            `json:"name,omitempty"`
	Flags          string            `json:"flags,omitempty"`
	Graphtype      string            `json:"graphtype,omitempty"`
	PercentLeft    string            `json:"percent_left,omitempty"`
	PercentRight   string            `json:"percent_right,omitempty"`
	Show3D         string            `json:"show_3d,omitempty"`
	ShowLegend     string            `json:"show_legend,omitempty"`
	ShowWorkPeriod string            `json:"show_work_period,omitempty"`
	ShowTriggers   string            `json:"show_triggers,omitempty"`
	TemplateID     string            `json:"templateid,omitempty"`
	YAxisMax       string            `json:"yaxismax,omitempty"`
	YAxisMin       string            `json:"yaxismin,omitempty"`
	YmaxItemID     string            `json:"ymax_itemid,omitempty"`
	YMaxType       string            `json:"ymax_type,omitempty"`
	YMinItemID     string            `json:"ymin_itemid,omitempty"`
	YMinType       string            `json:"ymin_type,omitempty"`
	GItems         []GraphItemObject `json:"gitems,omitempty"`
}
