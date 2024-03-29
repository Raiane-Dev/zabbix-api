package entities

type SettingsObject struct {
	DefaultLang                string `json:"default_lang,omitempty"`
	DefaultTimezone            string `json:"default_timezone,omitempty"`
	DefaultTheme               string `json:"default_theme,omitempty"`
	SearchLimit                int    `json:"search_limit,omitempty"`
	MaxOverviewTableSize       int    `json:"max_overview_table_size,omitempty"`
	MaxInTable                 int    `json:"max_in_table,omitempty"`
	ServerCheckInterval        int    `json:"servercheckinterval,omitempty"`
	WorkPeriod                 string `json:"work_period,omitempty"`
	ShowTechnicalErrors        int    `json:"show_technical_errors,omitempty"`
	HistoryPeriod              string `json:"history_period,omitempty"`
	PeriodDefault              string `json:"period_default,omitempty"`
	MaxPeriod                  string `json:"max_period,omitempty"`
	SeverityColor0             string `json:"severity_color_0,omitempty"`
	SeverityColor1             string `json:"severity_color_1,omitempty"`
	SeverityColor2             string `json:"severity_color_2,omitempty"`
	SeverityColor3             string `json:"severity_color_3,omitempty"`
	SeverityColor4             string `json:"severity_color_4,omitempty"`
	SeverityColor5             string `json:"severity_color_5,omitempty"`
	SeverityName0              string `json:"severity_name_0,omitempty"`
	SeverityName1              string `json:"severity_name_1,omitempty"`
	SeverityName2              string `json:"severity_name_2,omitempty"`
	SeverityName3              string `json:"severity_name_3,omitempty"`
	SeverityName4              string `json:"severity_name_4,omitempty"`
	SeverityName5              string `json:"severity_name_5,omitempty"`
	CustomColor                int    `json:"custom_color,omitempty"`
	OkPeriod                   string `json:"ok_period,omitempty"`
	BlinkPeriod                string `json:"blink_period,omitempty"`
	ProblemUnackColor          string `json:"problem_unack_color,omitempty"`
	ProblemAckColor            string `json:"problem_ack_color,omitempty"`
	OkUnackColor               string `json:"ok_unack_color,omitempty"`
	OkAckColor                 string `json:"ok_ack_color,omitempty"`
	ProblemUnackStyle          int    `json:"problem_unack_style,omitempty"`
	ProblemAckStyle            int    `json:"problem_ack_style,omitempty"`
	OkUnackStyle               int    `json:"ok_unack_style,omitempty"`
	OkAckStyle                 int    `json:"ok_ack_style,omitempty"`
	Url                        string `json:"url,omitempty"`
	DiscoveryGroupID           int    `json:"discovery_groupid,omitempty"`
	DefaultInventoryMode       int    `json:"default_inventory_mode,omitempty"`
	AlertUsrGrpID              int    `json:"alert_usrgrpid,omitempty"`
	SNMPTrapLogging            int    `json:"snmptrap_logging,omitempty"`
	LoginAttempts              int    `json:"login_attempts,omitempty"`
	LoginBlock                 string `json:"login_block,omitempty"`
	ValidateURISchemes         int    `json:"validate_uri_schemes,omitempty"`
	UriValidSchemes            int    `json:"uri_valid_schemes,omitempty"`
	XFrameOptions              string `json:"x_frame_options,omitempty"`
	IframeSandboxingEnabled    int    `json:"iframe_sandboxing_enabled,omitempty"`
	IframeSandboxingExceptions int    `json:"iframe_sandboxing_exceptions,omitempty"`
	ConnectTimeout             string `json:"connect_timeout,omitempty"`
	SocketTimeout              string `json:"socket_timeout,omitempty"`
	MediaTypeTestTimeout       string `json:"media_type_test_timeout,omitempty"`
	ItemTestTimeout            string `json:"item_test_timeout,omitempty"`
	ScriptTimeout              string `json:"script_timeout,omitempty"`
	ReportTestTimeout          string `json:"report_test_timeout,omitempty"`
	AuditlogEnabled            int    `json:"auditlog_enabled,omitempty"`
	HaFailoverDelay            string `json:"ha_failover_delay,omitempty"`
	GeomapsTileProvider        string `json:"geomaps_tile_provider,omitempty"`
	GeomapsTileURL             string `json:"geomaps_tile_url,omitempty"`
	GeomapMaxZoom              int    `json:"geomap_max_zoom,omitempty"`
	GeomapsAttribution         string `json:"geomaps_attribution,omitempty"`
	VaultProvider              int    `json:"vault_provider,omitempty"`
}

type SettingsUpdate struct {
	DefaultLang                string `json:"default_lang,omitempty"`
	DefaultTimezone            string `json:"default_timezone,omitempty"`
	DefaultTheme               string `json:"default_theme,omitempty"`
	SearchLimit                int    `json:"search_limit,omitempty"`
	MaxOverviewTableSize       int    `json:"max_overview_table_size,omitempty"`
	MaxInTable                 int    `json:"max_in_table,omitempty"`
	ServerCheckInterval        int    `json:"servercheckinterval,omitempty"`
	WorkPeriod                 string `json:"work_period,omitempty"`
	ShowTechnicalErrors        int    `json:"show_technical_errors,omitempty"`
	HistoryPeriod              string `json:"history_period,omitempty"`
	PeriodDefault              string `json:"period_default,omitempty"`
	MaxPeriod                  string `json:"max_period,omitempty"`
	SeverityColor0             string `json:"severity_color_0,omitempty"`
	SeverityColor1             string `json:"severity_color_1,omitempty"`
	SeverityColor2             string `json:"severity_color_2,omitempty"`
	SeverityColor3             string `json:"severity_color_3,omitempty"`
	SeverityColor4             string `json:"severity_color_4,omitempty"`
	SeverityColor5             string `json:"severity_color_5,omitempty"`
	SeverityName0              string `json:"severity_name_0,omitempty"`
	SeverityName1              string `json:"severity_name_1,omitempty"`
	SeverityName2              string `json:"severity_name_2,omitempty"`
	SeverityName3              string `json:"severity_name_3,omitempty"`
	SeverityName4              string `json:"severity_name_4,omitempty"`
	SeverityName5              string `json:"severity_name_5,omitempty"`
	CustomColor                int    `json:"custom_color,omitempty"`
	OkPeriod                   string `json:"ok_period,omitempty"`
	BlinkPeriod                string `json:"blink_period,omitempty"`
	ProblemUnackColor          string `json:"problem_unack_color,omitempty"`
	ProblemAckColor            string `json:"problem_ack_color,omitempty"`
	OkUnackColor               string `json:"ok_unack_color,omitempty"`
	OkAckColor                 string `json:"ok_ack_color,omitempty"`
	ProblemUnackStyle          int    `json:"problem_unack_style,omitempty"`
	ProblemAckStyle            int    `json:"problem_ack_style,omitempty"`
	OkUnackStyle               int    `json:"ok_unack_style,omitempty"`
	OkAckStyle                 int    `json:"ok_ack_style,omitempty"`
	Url                        string `json:"url,omitempty"`
	DiscoveryGroupID           int    `json:"discovery_groupid,omitempty"`
	DefaultInventoryMode       int    `json:"default_inventory_mode,omitempty"`
	AlertUsrGrpID              int    `json:"alert_usrgrpid,omitempty"`
	SNMPTrapLogging            int    `json:"snmptrap_logging,omitempty"`
	LoginAttempts              int    `json:"login_attempts,omitempty"`
	LoginBlock                 string `json:"login_block,omitempty"`
	ValidateURISchemes         int    `json:"validate_uri_schemes,omitempty"`
	UriValidSchemes            int    `json:"uri_valid_schemes,omitempty"`
	XFrameOptions              string `json:"x_frame_options,omitempty"`
	IframeSandboxingEnabled    int    `json:"iframe_sandboxing_enabled,omitempty"`
	IframeSandboxingExceptions int    `json:"iframe_sandboxing_exceptions,omitempty"`
	ConnectTimeout             string `json:"connect_timeout,omitempty"`
	SocketTimeout              string `json:"socket_timeout,omitempty"`
	MediaTypeTestTimeout       string `json:"media_type_test_timeout,omitempty"`
	ItemTestTimeout            string `json:"item_test_timeout,omitempty"`
	ScriptTimeout              string `json:"script_timeout,omitempty"`
	ReportTestTimeout          string `json:"report_test_timeout,omitempty"`
	AuditlogEnabled            int    `json:"auditlog_enabled,omitempty"`
	HaFailoverDelay            string `json:"ha_failover_delay,omitempty"`
	GeomapsTileProvider        string `json:"geomaps_tile_provider,omitempty"`
	GeomapsTileURL             string `json:"geomaps_tile_url,omitempty"`
	GeomapMaxZoom              int    `json:"geomap_max_zoom,omitempty"`
	GeomapsAttribution         string `json:"geomaps_attribution,omitempty"`
	VaultProvider              int    `json:"vault_provider,omitempty"`
}

type SettingsGet struct {
	DefaultLang                string `json:"default_lang,omitempty"`
	DefaultTimezone            string `json:"default_timezone,omitempty"`
	DefaultTheme               string `json:"default_theme,omitempty"`
	SearchLimit                int    `json:"search_limit,omitempty"`
	MaxOverviewTableSize       int    `json:"max_overview_table_size,omitempty"`
	MaxInTable                 int    `json:"max_in_table,omitempty"`
	ServerCheckInterval        int    `json:"servercheckinterval,omitempty"`
	WorkPeriod                 string `json:"work_period,omitempty"`
	ShowTechnicalErrors        int    `json:"show_technical_errors,omitempty"`
	HistoryPeriod              string `json:"history_period,omitempty"`
	PeriodDefault              string `json:"period_default,omitempty"`
	MaxPeriod                  string `json:"max_period,omitempty"`
	SeverityColor0             string `json:"severity_color_0,omitempty"`
	SeverityColor1             string `json:"severity_color_1,omitempty"`
	SeverityColor2             string `json:"severity_color_2,omitempty"`
	SeverityColor3             string `json:"severity_color_3,omitempty"`
	SeverityColor4             string `json:"severity_color_4,omitempty"`
	SeverityColor5             string `json:"severity_color_5,omitempty"`
	SeverityName0              string `json:"severity_name_0,omitempty"`
	SeverityName1              string `json:"severity_name_1,omitempty"`
	SeverityName2              string `json:"severity_name_2,omitempty"`
	SeverityName3              string `json:"severity_name_3,omitempty"`
	SeverityName4              string `json:"severity_name_4,omitempty"`
	SeverityName5              string `json:"severity_name_5,omitempty"`
	CustomColor                int    `json:"custom_color,omitempty"`
	OkPeriod                   string `json:"ok_period,omitempty"`
	BlinkPeriod                string `json:"blink_period,omitempty"`
	ProblemUnackColor          string `json:"problem_unack_color,omitempty"`
	ProblemAckColor            string `json:"problem_ack_color,omitempty"`
	OkUnackColor               string `json:"ok_unack_color,omitempty"`
	OkAckColor                 string `json:"ok_ack_color,omitempty"`
	ProblemUnackStyle          int    `json:"problem_unack_style,omitempty"`
	ProblemAckStyle            int    `json:"problem_ack_style,omitempty"`
	OkUnackStyle               int    `json:"ok_unack_style,omitempty"`
	OkAckStyle                 int    `json:"ok_ack_style,omitempty"`
	Url                        string `json:"url,omitempty"`
	DiscoveryGroupID           int    `json:"discovery_groupid,omitempty"`
	DefaultInventoryMode       int    `json:"default_inventory_mode,omitempty"`
	AlertUsrGrpID              int    `json:"alert_usrgrpid,omitempty"`
	SNMPTrapLogging            int    `json:"snmptrap_logging,omitempty"`
	LoginAttempts              int    `json:"login_attempts,omitempty"`
	LoginBlock                 string `json:"login_block,omitempty"`
	ValidateURISchemes         int    `json:"validate_uri_schemes,omitempty"`
	UriValidSchemes            int    `json:"uri_valid_schemes,omitempty"`
	XFrameOptions              string `json:"x_frame_options,omitempty"`
	IframeSandboxingEnabled    int    `json:"iframe_sandboxing_enabled,omitempty"`
	IframeSandboxingExceptions int    `json:"iframe_sandboxing_exceptions,omitempty"`
	ConnectTimeout             string `json:"connect_timeout,omitempty"`
	SocketTimeout              string `json:"socket_timeout,omitempty"`
	MediaTypeTestTimeout       string `json:"media_type_test_timeout,omitempty"`
	ItemTestTimeout            string `json:"item_test_timeout,omitempty"`
	ScriptTimeout              string `json:"script_timeout,omitempty"`
	ReportTestTimeout          string `json:"report_test_timeout,omitempty"`
	AuditlogEnabled            int    `json:"auditlog_enabled,omitempty"`
	HaFailoverDelay            string `json:"ha_failover_delay,omitempty"`
	GeomapsTileProvider        string `json:"geomaps_tile_provider,omitempty"`
	GeomapsTileURL             string `json:"geomaps_tile_url,omitempty"`
	GeomapMaxZoom              int    `json:"geomap_max_zoom,omitempty"`
	GeomapsAttribution         string `json:"geomaps_attribution,omitempty"`
	VaultProvider              int    `json:"vault_provider,omitempty"`
	Output                     string `json:"output,omitempty"`
}
