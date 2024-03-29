package entities

type MapObject struct {
	SysmapID             string `json:"sysmapid,omitempty"`
	Height               int    `json:"height,omitempty"`
	Name                 string `json:"name,omitempty"`
	Width                int    `json:"width,omitempty"`
	BackgroundID         string `json:"backgroundid,omitempty"`
	ExpandMacros         int    `json:"expand_macros,omitempty"`
	ExpandProblem        int    `json:"expand_problem,omitempty"`
	GridAlign            int    `json:"grid_align,omitempty"`
	GridShow             int    `json:"grid_show,omitempty"`
	GridSize             int    `json:"grid_size,omitempty"`
	Highlight            int    `json:"highlight,omitempty"`
	IconMapID            string `json:"iconmapid,omitempty"`
	LabelFormat          int    `json:"label_format,omitempty"`
	LabelLocation        int    `json:"label_location,omitempty"`
	LabelStringHost      string `json:"label_string_host,omitempty"`
	LabelStringHostGroup string `json:"label_string_hostgroup,omitempty"`
	LabelStringImage     string `json:"label_string_image,omitempty"`
	LabelStringMap       string `json:"label_string_map,omitempty"`
	LabelStringTrigger   string `json:"label_string_trigger,omitempty"`
	LabelType            int    `json:"label_type,omitempty"`
	LabelTypeHost        int    `json:"label_type_host,omitempty"`
	LabelTypeHostGroup   int    `json:"label_type_hostgroup,omitempty"`
	LabelTypeImage       int    `json:"label_type_image,omitempty"`
	LabelTypeMap         int    `json:"label_type_map,omitempty"`
	LabelTypeTrigger     int    `json:"label_type_trigger,omitempty"`
	MarkElements         int    `json:"markelements,omitempty"`
	SeverityMin          int    `json:"severity_min,omitempty"`
	ShowUnack            int    `json:"show_unack,omitempty"`
	UserID               string `json:"userid,omitempty"`
	Private              int    `json:"private,omitempty"`
	ShowSupressed        int    `json:"show_supressed,omitempty"`
}

type MapElement struct {
	SelementID        string              `json:"selementid,omitempty"`
	Elements          []map[string]string `json:"elements,omitempty"`
	ElementType       int                 `json:"elementtype,omitempty"`
	IconIDOff         string              `json:"iconid_off,omitempty"`
	AreaType          int                 `json:"areatype,omitempty"`
	ElementSubType    int                 `json:"elementsubtype,omitempty"`
	Evaltype          int                 `json:"evaltype,omitempty"`
	Height            int                 `json:"height,omitempty"`
	IconIDDisabled    string              `json:"iconid_disabled,omitempty"`
	IconIDMaintenance string              `json:"iconid_maintenance,omitempty"`
	IconIDOn          string              `json:"iconid_on,omitempty"`
	Label             string              `json:"label,omitempty"`
	LabelLocation     int                 `json:"label_location,omitempty"`
	Permission        int                 `json:"permission,omitempty"`
	SysmapID          string              `json:"sysmapid,omitempty"`
	Urls              []MapElement        `json:"urls,omitempty"`
	UseIconMap        int                 `json:"use_iconmap,omitempty"`
	Viewtype          int                 `json:"viewtype,omitempty"`
	Width             int                 `json:"width,omitempty"`
	X                 int                 `json:"x,omitempty"`
	Y                 int                 `json:"y,omitempty"`
}

type MapElementHost struct {
	HostID string
}

type MapElementUrl struct {
	SysmapElementURLID string `json:"sysmapelementurlid,omitempty"`
	Name               string `json:"name,omitempty"`
	Url                string `json:"url,omitempty"`
	SelementID         string `json:"selementid,omitempty"`
}

type MapLink struct {
	LinkID       string           `json:"linkid,omitempty"`
	SelementID1  string           `json:"selementid1,omitempty"`
	SelementID2  string           `json:"selementid2,omitempty"`
	Color        string           `json:"color,omitempty"`
	Drawtype     int              `json:"drawtype,omitempty"`
	Label        string           `json:"label,omitempty"`
	LinkTriggers []MapLinkTrigger `json:"linktriggers,omitempty"`
	Permission   int              `json:"permission,omitempty"`
	SysmapID     string           `json:"sysmapid,omitempty"`
}

type MapLinkTrigger struct {
	LinkTriggerID string `json:"linktriggerid,omitempty"`
	TriggerID     string `json:"triggerid,omitempty"`
	Color         string `json:"color,omitempty"`
	Drawtype      int    `json:"drawtype,omitempty"`
	LinkID        string `json:"linkid,omitempty"`
}

type MapURL struct {
	SysmapUrlID string `json:"sysmapurlid,omitempty"`
	Name        string `json:"name,omitempty"`
	URL         string `json:"url,omitempty"`
	ElementType int    `json:"elementtype,omitempty"`
	SysmapID    string `json:"sysmapid,omitempty"`
}

type MapUser struct {
	SysmapUserID string `json:"sysmapuserid,omitempty"`
	UserID       string `json:"userid,omitempty"`
	Permission   int    `json:"permission,omitempty"`
}
type MapUserGroup struct {
	SysmapUserGrpID string `json:"sysmapusergrpid,omitempty"`
	UserGrpID       string `json:"usergrpid,omitempty"`
	Permission      int    `json:"permission,omitempty"`
}
type MapShapes struct {
	SysmapShapeID   string `json:"sysmap_shapeid,omitempty"`
	Type            int    `json:"type,omitempty"`
	X               int    `json:"x,omitempty"`
	Y               int    `json:"y,omitempty"`
	Width           int    `json:"width,omitempty"`
	Height          int    `json:"height,omitempty"`
	Text            string `json:"text,omitempty"`
	Font            int    `json:"font,omitempty"`
	FontSize        int    `json:"font_size,omitempty"`
	FontColor       string `json:"font_color,omitempty"`
	TextHalign      int    `json:"text_halign,omitempty"`
	TextValign      int    `json:"text_valign,omitempty"`
	BorderType      int    `json:"border_type,omitempty"`
	BorderWidth     int    `json:"border_width,omitempty"`
	BorderColor     string `json:"border_color,omitempty"`
	BackgroundColor string `json:"background_color,omitempty"`
	ZIndex          int    `json:"zindex,omitempty"`
}

type MapLines struct {
	SysmapShapeID string `json:"sysmap_shapeid,omitempty"`
	X1            int    `json:"x1,omitempty"`
	Y1            int    `json:"y1,omitempty"`
	X2            int    `json:"x2,omitempty"`
	Y2            int    `json:"y2,omitempty"`
	LineType      int    `json:"linetype,omitempty"`
	LineWidth     int    `json:"linewidth,omitempty"`
	LineColor     string `json:"linecolor,omitempty"`
	ZIndex        int    `json:"zindex,omitempty"`
}

type MapGet struct {
	SysmapIDs        []string `json:"sysmapids,omitempty"`
	UserIDs          []string `json:"userids,omitempty"`
	ExpandUrls       string   `json:"expandUrls,omitempty"`
	SelectIconMap    string   `json:"selectIconMap,omitempty"`
	SelectLinks      string   `json:"selectLinks,omitempty"`
	SelectSelements  string   `json:"selectSelements,omitempty"`
	SelectUrls       string   `json:"selectUrls,omitempty"`
	SelectUsers      string   `json:"selectUsers,omitempty"`
	SelectUserGroups string   `json:"selectUserGroups,omitempty"`
	SelectShapes     string   `json:"selectShapes,omitempty"`
	SelectLines      string   `json:"selectLines,omitempty"`
	ZabbixCommun
}

type MapCreate struct {
	Height               int             `json:"height,omitempty"`
	Name                 string          `json:"name,omitempty"`
	Width                int             `json:"width,omitempty"`
	BackgroundID         string          `json:"backgroundid,omitempty"`
	ExpandMacros         int             `json:"expand_macros,omitempty"`
	ExpandProblem        int             `json:"expand_problem,omitempty"`
	GridAlign            int             `json:"grid_align,omitempty"`
	GridShow             int             `json:"grid_show,omitempty"`
	GridSize             int             `json:"grid_size,omitempty"`
	Highlight            int             `json:"highlight,omitempty"`
	IconMapID            string          `json:"iconmapid,omitempty"`
	LabelFormat          int             `json:"label_format,omitempty"`
	LabelLocation        int             `json:"label_location,omitempty"`
	LabelStringHost      string          `json:"label_string_host,omitempty"`
	LabelStringHostGroup string          `json:"label_string_hostgroup,omitempty"`
	LabelStringImage     string          `json:"label_string_image,omitempty"`
	LabelStringMap       string          `json:"label_string_map,omitempty"`
	LabelStringTrigger   string          `json:"label_string_trigger,omitempty"`
	LabelType            int             `json:"label_type,omitempty"`
	LabelTypeHost        int             `json:"label_type_host,omitempty"`
	LabelTypeHostGroup   int             `json:"label_type_hostgroup,omitempty"`
	LabelTypeImage       int             `json:"label_type_image,omitempty"`
	LabelTypeMap         int             `json:"label_type_map,omitempty"`
	LabelTypeTrigger     int             `json:"label_type_trigger,omitempty"`
	MarkElements         int             `json:"markelements,omitempty"`
	SeverityMin          int             `json:"severity_min,omitempty"`
	ShowUnack            int             `json:"show_unack,omitempty"`
	UserID               string          `json:"userid,omitempty"`
	Private              int             `json:"private,omitempty"`
	ShowSupressed        int             `json:"show_supressed,omitempty"`
	Links                []MapLink       `json:"links,omitempty"`
	Selements            []MapElement    `json:"selements,omitempty"`
	URLs                 []MapElementUrl `json:"urls,omitempty"`
	Users                []MapUser       `json:"users,omitempty"`
	UserGroups           []MapUserGroup  `json:"usergroups,omitempty"`
	Shapes               []MapShapes     `json:"shapes,omitempty"`
	Lines                []MapLines      `json:"lines,omitempty"`
}

type MapUpdate struct {
	SysmapID             string          `json:"sysmapid,omitempty"`
	Height               int             `json:"height,omitempty"`
	Name                 string          `json:"name,omitempty"`
	Width                int             `json:"width,omitempty"`
	BackgroundID         string          `json:"backgroundid,omitempty"`
	ExpandMacros         int             `json:"expand_macros,omitempty"`
	ExpandProblem        int             `json:"expand_problem,omitempty"`
	GridAlign            int             `json:"grid_align,omitempty"`
	GridShow             int             `json:"grid_show,omitempty"`
	GridSize             int             `json:"grid_size,omitempty"`
	Highlight            int             `json:"highlight,omitempty"`
	IconMapID            string          `json:"iconmapid,omitempty"`
	LabelFormat          int             `json:"label_format,omitempty"`
	LabelLocation        int             `json:"label_location,omitempty"`
	LabelStringHost      string          `json:"label_string_host,omitempty"`
	LabelStringHostGroup string          `json:"label_string_hostgroup,omitempty"`
	LabelStringImage     string          `json:"label_string_image,omitempty"`
	LabelStringMap       string          `json:"label_string_map,omitempty"`
	LabelStringTrigger   string          `json:"label_string_trigger,omitempty"`
	LabelType            int             `json:"label_type,omitempty"`
	LabelTypeHost        int             `json:"label_type_host,omitempty"`
	LabelTypeHostGroup   int             `json:"label_type_hostgroup,omitempty"`
	LabelTypeImage       int             `json:"label_type_image,omitempty"`
	LabelTypeMap         int             `json:"label_type_map,omitempty"`
	LabelTypeTrigger     int             `json:"label_type_trigger,omitempty"`
	MarkElements         int             `json:"markelements,omitempty"`
	SeverityMin          int             `json:"severity_min,omitempty"`
	ShowUnack            int             `json:"show_unack,omitempty"`
	UserID               string          `json:"userid,omitempty"`
	Private              int             `json:"private,omitempty"`
	ShowSupressed        int             `json:"show_supressed,omitempty"`
	Links                []MapLink       `json:"links,omitempty"`
	Selements            []MapElement    `json:"selements,omitempty"`
	URLs                 []MapElementUrl `json:"urls,omitempty"`
	Users                []MapUser       `json:"users,omitempty"`
	UserGroups           []MapUserGroup  `json:"usergroups,omitempty"`
	Shapes               []MapShapes     `json:"shapes,omitempty"`
	Lines                []MapLines      `json:"lines,omitempty"`
}
