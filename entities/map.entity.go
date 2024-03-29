package entities

type MapObject struct {
	SysmapID             string `json:"sysmapid"`
	Height               int    `json:"height"`
	Name                 string `json:"name"`
	Width                int    `json:"width"`
	BackgroundID         string `json:"backgroundid"`
	ExpandMacros         int    `json:"expand_macros"`
	ExpandProblem        int    `json:"expand_problem"`
	GridAlign            int    `json:"grid_align"`
	GridShow             int    `json:"grid_show"`
	GridSize             int    `json:"grid_size"`
	Highlight            int    `json:"highlight"`
	IconMapID            string `json:"iconmapid"`
	LabelFormat          int    `json:"label_format"`
	LabelLocation        int    `json:"label_location"`
	LabelStringHost      string `json:"label_string_host"`
	LabelStringHostGroup string `json:"label_string_hostgroup"`
	LabelStringImage     string `json:"label_string_image"`
	LabelStringMap       string `json:"label_string_map"`
	LabelStringTrigger   string `json:"label_string_trigger"`
	LabelType            int    `json:"label_type"`
	LabelTypeHost        int    `json:"label_type_host"`
	LabelTypeHostGroup   int    `json:"label_type_hostgroup"`
	LabelTypeImage       int    `json:"label_type_image"`
	LabelTypeMap         int    `json:"label_type_map"`
	LabelTypeTrigger     int    `json:"label_type_trigger"`
	MarkElements         int    `json:"markelements"`
	SeverityMin          int    `json:"severity_min"`
	ShowUnack            int    `json:"show_unack"`
	UserID               string `json:"userid"`
	Private              int    `json:"private"`
	ShowSupressed        int    `json:"show_supressed"`
}

type MapElement struct {
	SelementID        string              `json:"selementid"`
	Elements          []map[string]string `json:"elements"`
	ElementType       int                 `json:"elementtype"`
	IconIDOff         string              `json:"iconid_off"`
	AreaType          int                 `json:"areatype"`
	ElementSubType    int                 `json:"elementsubtype"`
	Evaltype          int                 `json:"evaltype"`
	Height            int                 `json:"height"`
	IconIDDisabled    string              `json:"iconid_disabled"`
	IconIDMaintenance string              `json:"iconid_maintenance"`
	IconIDOn          string              `json:"iconid_on"`
	Label             string              `json:"label"`
	LabelLocation     int                 `json:"label_location"`
	Permission        int                 `json:"permission"`
	SysmapID          string              `json:"sysmapid"`
	Urls              []MapElement        `json:"urls"`
	UseIconMap        int                 `json:"use_iconmap"`
	Viewtype          int                 `json:"viewtype"`
	Width             int                 `json:"width"`
	X                 int                 `json:"x"`
	Y                 int                 `json:"y"`
}

type MapElementHost struct {
	HostID string
}

type MapElementUrl struct {
	SysmapElementURLID string `json:"sysmapelementurlid"`
	Name               string `json:"name"`
	Url                string `json:"url"`
	SelementID         string `json:"selementid"`
}

type MapLink struct {
	LinkID       string           `json:"linkid"`
	SelementID1  string           `json:"selementid1"`
	SelementID2  string           `json:"selementid2"`
	Color        string           `json:"color"`
	Drawtype     int              `json:"drawtype"`
	Label        string           `json:"label"`
	LinkTriggers []MapLinkTrigger `json:"linktriggers"`
	Permission   int              `json:"permission"`
	SysmapID     string           `json:"sysmapid"`
}

type MapLinkTrigger struct {
	LinkTriggerID string `json:"linktriggerid"`
	TriggerID     string `json:"triggerid"`
	Color         string `json:"color"`
	Drawtype      int    `json:"drawtype"`
	LinkID        string `json:"linkid"`
}

type MapURL struct {
	SysmapUrlID string `json:"sysmapurlid"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	ElementType int    `json:"elementtype"`
	SysmapID    string `json:"sysmapid"`
}

type MapUser struct {
	SysmapUserID string `json:"sysmapuserid"`
	UserID       string `json:"userid"`
	Permission   int    `json:"permission"`
}
type MapUserGroup struct {
	SysmapUserGrpID string `json:"sysmapusergrpid"`
	UserGrpID       string `json:"usergrpid"`
	Permission      int    `json:"permission"`
}
type MapShapes struct {
	SysmapShapeID   string `json:"sysmap_shapeid"`
	Type            int    `json:"type"`
	X               int    `json:"x"`
	Y               int    `json:"y"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	Text            string `json:"text"`
	Font            int    `json:"font"`
	FontSize        int    `json:"font_size"`
	FontColor       string `json:"font_color"`
	TextHalign      int    `json:"text_halign"`
	TextValign      int    `json:"text_valign"`
	BorderType      int    `json:"border_type"`
	BorderWidth     int    `json:"border_width"`
	BorderColor     string `json:"border_color"`
	BackgroundColor string `json:"background_color"`
	ZIndex          int    `json:"zindex"`
}

type MapLines struct {
	SysmapShapeID string `json:"sysmap_shapeid"`
	X1            int    `json:"x1"`
	Y1            int    `json:"y1"`
	X2            int    `json:"x2"`
	Y2            int    `json:"y2"`
	LineType      int    `json:"linetype"`
	LineWidth     int    `json:"linewidth"`
	LineColor     string `json:"linecolor"`
	ZIndex        int    `json:"zindex"`
}

type MapGet struct {
	SysmapIDs        []string `json:"sysmapids"`
	UserIDs          []string `json:"userids"`
	ExpandUrls       string   `json:"expandUrls"`
	SelectIconMap    string   `json:"selectIconMap"`
	SelectLinks      string   `json:"selectLinks"`
	SelectSelements  string   `json:"selectSelements"`
	SelectUrls       string   `json:"selectUrls"`
	SelectUsers      string   `json:"selectUsers"`
	SelectUserGroups string   `json:"selectUserGroups"`
	SelectShapes     string   `json:"selectShapes"`
	SelectLines      string   `json:"selectLines"`
	ZabbixCommun
}

type MapCreate struct {
	Height               int             `json:"height"`
	Name                 string          `json:"name"`
	Width                int             `json:"width"`
	BackgroundID         string          `json:"backgroundid"`
	ExpandMacros         int             `json:"expand_macros"`
	ExpandProblem        int             `json:"expand_problem"`
	GridAlign            int             `json:"grid_align"`
	GridShow             int             `json:"grid_show"`
	GridSize             int             `json:"grid_size"`
	Highlight            int             `json:"highlight"`
	IconMapID            string          `json:"iconmapid"`
	LabelFormat          int             `json:"label_format"`
	LabelLocation        int             `json:"label_location"`
	LabelStringHost      string          `json:"label_string_host"`
	LabelStringHostGroup string          `json:"label_string_hostgroup"`
	LabelStringImage     string          `json:"label_string_image"`
	LabelStringMap       string          `json:"label_string_map"`
	LabelStringTrigger   string          `json:"label_string_trigger"`
	LabelType            int             `json:"label_type"`
	LabelTypeHost        int             `json:"label_type_host"`
	LabelTypeHostGroup   int             `json:"label_type_hostgroup"`
	LabelTypeImage       int             `json:"label_type_image"`
	LabelTypeMap         int             `json:"label_type_map"`
	LabelTypeTrigger     int             `json:"label_type_trigger"`
	MarkElements         int             `json:"markelements"`
	SeverityMin          int             `json:"severity_min"`
	ShowUnack            int             `json:"show_unack"`
	UserID               string          `json:"userid"`
	Private              int             `json:"private"`
	ShowSupressed        int             `json:"show_supressed"`
	Links                []MapLink       `json:"links"`
	Selements            []MapElement    `json:"selements"`
	URLs                 []MapElementUrl `json:"urls"`
	Users                []MapUser       `json:"users"`
	UserGroups           []MapUserGroup  `json:"usergroups"`
	Shapes               []MapShapes     `json:"shapes"`
	Lines                []MapLines      `json:"lines"`
}

type MapUpdate struct {
	SysmapID             string          `json:"sysmapid"`
	Height               int             `json:"height"`
	Name                 string          `json:"name"`
	Width                int             `json:"width"`
	BackgroundID         string          `json:"backgroundid"`
	ExpandMacros         int             `json:"expand_macros"`
	ExpandProblem        int             `json:"expand_problem"`
	GridAlign            int             `json:"grid_align"`
	GridShow             int             `json:"grid_show"`
	GridSize             int             `json:"grid_size"`
	Highlight            int             `json:"highlight"`
	IconMapID            string          `json:"iconmapid"`
	LabelFormat          int             `json:"label_format"`
	LabelLocation        int             `json:"label_location"`
	LabelStringHost      string          `json:"label_string_host"`
	LabelStringHostGroup string          `json:"label_string_hostgroup"`
	LabelStringImage     string          `json:"label_string_image"`
	LabelStringMap       string          `json:"label_string_map"`
	LabelStringTrigger   string          `json:"label_string_trigger"`
	LabelType            int             `json:"label_type"`
	LabelTypeHost        int             `json:"label_type_host"`
	LabelTypeHostGroup   int             `json:"label_type_hostgroup"`
	LabelTypeImage       int             `json:"label_type_image"`
	LabelTypeMap         int             `json:"label_type_map"`
	LabelTypeTrigger     int             `json:"label_type_trigger"`
	MarkElements         int             `json:"markelements"`
	SeverityMin          int             `json:"severity_min"`
	ShowUnack            int             `json:"show_unack"`
	UserID               string          `json:"userid"`
	Private              int             `json:"private"`
	ShowSupressed        int             `json:"show_supressed"`
	Links                []MapLink       `json:"links"`
	Selements            []MapElement    `json:"selements"`
	URLs                 []MapElementUrl `json:"urls"`
	Users                []MapUser       `json:"users"`
	UserGroups           []MapUserGroup  `json:"usergroups"`
	Shapes               []MapShapes     `json:"shapes"`
	Lines                []MapLines      `json:"lines"`
}
