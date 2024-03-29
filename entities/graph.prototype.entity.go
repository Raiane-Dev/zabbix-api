package entities

type GraphPrototypeCreate struct {
	GItems         []GraphItemObject `json:"gitems,omitempty"`
	Height         int               `json:"height,omitempty"`
	Name           string            `json:"name,omitempty"`
	Width          int               `json:"width,omitempty"`
	GraphType      int               `json:"graphtype,omitempty"`
	PercentLeft    float32           `json:"percent_left,omitempty"`
	PercentRight   float32           `json:"percent_right,omitempty"`
	Show3D         int               `json:"show_3d,omitempty"`
	ShowLegend     int               `json:"show_legend,omitempty"`
	ShowWorkPeriod int               `json:"show_work_period,omitempty"`
	TemplateID     string            `json:"templateid,omitempty"`
	YaxisMax       string            `json:"yaxismax,omitempty"`
	YAxisMin       string            `json:"yaxismin,omitempty"`
	YMaxItemID     string            `json:"ymax_itemid,omitempty"`
	YMaxType       int               `json:"ymax_type,omitempty"`
	YMinItemID     string            `json:"ymin_itemid,omitempty"`
	YMinType       int               `json:"ymin_type,omitempty"`
	Discover       int               `json:"discover,omitempty"`
	UUID           string            `json:"uuid,omitempty"`
}

type GraphPrototypeUpdate struct {
	GraphID        string            `json:"graphid,omitempty"`
	GItems         []GraphItemObject `json:"gitems,omitempty"`
	Height         int               `json:"height,omitempty"`
	Name           string            `json:"name,omitempty"`
	Width          int               `json:"width,omitempty"`
	GraphType      int               `json:"graphtype,omitempty"`
	PercentLeft    float32           `json:"percent_left,omitempty"`
	PercentRight   float32           `json:"percent_right,omitempty"`
	Show3D         int               `json:"show_3d,omitempty"`
	ShowLegend     int               `json:"show_legend,omitempty"`
	ShowWorkPeriod int               `json:"show_work_period,omitempty"`
	TemplateID     string            `json:"templateid,omitempty"`
	YaxisMax       string            `json:"yaxismax,omitempty"`
	YAxisMin       string            `json:"yaxismin,omitempty"`
	YMaxItemID     string            `json:"ymax_itemid,omitempty"`
	YMaxType       int               `json:"ymax_type,omitempty"`
	YMinItemID     string            `json:"ymin_itemid,omitempty"`
	YMinType       int               `json:"ymin_type,omitempty"`
	Discover       int               `json:"discover,omitempty"`
	UUID           string            `json:"uuid,omitempty"`
}

type GraphPrototypeResponse struct {
	GraphIDs []string `json:"graphids,omitempty"`
}

type GraphPrototypeGet struct {
	DiscoveryIDs         []string `json:"discoveryids,omitempty"`
	GraphIDs             []string `json:"graphids,omitempty"`
	GroupIDs             []string `json:"groupids,omitempty"`
	HostIDs              []string `json:"hostids,omitempty"`
	Inherited            bool     `json:"inherited,omitempty"`
	ItemIDs              []string `json:"itemids,omitempty"`
	Templated            bool     `json:"templated,omitempty"`
	TemplateIDs          []string `json:"templateids,omitempty"`
	SelectDiscoveryRule  string   `json:"selectDiscoveryRule,omitempty"`
	SelectGraphItems     string   `json:"selectGraphItems,omitempty"`
	SelectHostGroups     string   `json:"selectHostGroups,omitempty"`
	SelectHosts          string   `json:"selectHosts,omitempty"`
	SelectItems          string   `json:"selectItems,omitempty"`
	SelectTemplateGroups string   `json:"selectTemplateGroups,omitempty"`
	SelectTemplates      string   `json:"selectTemplates,omitempty"`
	ZabbixCommun
}

type GraphPrototypeObject struct {
	GraphID        string  `json:"graphid,omitempty"`
	Height         int     `json:"height,omitempty"`
	Name           string  `json:"name,omitempty"`
	Width          int     `json:"width,omitempty"`
	GraphType      int     `json:"graphtype,omitempty"`
	PercentLeft    float32 `json:"percent_left,omitempty"`
	PercentRight   float32 `json:"percent_right,omitempty"`
	Show3D         int     `json:"show_3d,omitempty"`
	ShowLegend     int     `json:"show_legend,omitempty"`
	ShowWorkPeriod int     `json:"show_work_period,omitempty"`
	TemplateID     string  `json:"templateid,omitempty"`
	YaxisMax       string  `json:"yaxismax,omitempty"`
	YAxisMin       string  `json:"yaxismin,omitempty"`
	YMaxItemID     string  `json:"ymax_itemid,omitempty"`
	YMaxType       int     `json:"ymax_type,omitempty"`
	YMinItemID     string  `json:"ymin_itemid,omitempty"`
	YMinType       int     `json:"ymin_type,omitempty"`
	Discover       int     `json:"discover,omitempty"`
	UUID           string  `json:"uuid,omitempty"`
}
