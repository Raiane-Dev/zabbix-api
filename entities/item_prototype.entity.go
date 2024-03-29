package entities

type ItemPrototypeObject struct {
	ItemID          string              `json:"itemid,omitempty"`
	Delay           string              `json:"delay,omitempty"`
	HostID          string              `json:"hostid,omitempty"`
	InterfaceID     string              `json:"interfaceid,omitempty"`
	Key             string              `json:"key_,omitempty"`
	Name            string              `json:"name,omitempty"`
	Type            int                 `json:"type,omitempty"`
	Url             string              `json:"url,omitempty"`
	ValueType       int                 `json:"value_type,omitempty"`
	AllowTraps      int                 `json:"allow_traps,omitempty"`
	AuthType        int                 `json:"authtype,omitempty"`
	Description     string              `json:"description,omitempty"`
	FollowRedirects int                 `json:"follow_redirects,omitempty"`
	Headers         map[string]string   `json:"headers,omitempty"`
	History         string              `json:"history,omitempty"`
	HttpProxy       string              `json:"http_proxy,omitempty"`
	IPMISensor      string              `json:"ipmi_sensor,omitempty"`
	JmxEndpoint     string              `json:"jmx_endpoint,omitempty"`
	LogtimeFmt      string              `json:"logtimefmt,omitempty"`
	MasterItemID    int                 `json:"master_itemid,omitempty"`
	OutputFormat    int                 `json:"output_format,omitempty"`
	Params          string              `json:"params,omitempty"`
	Parameters      []map[string]string `json:"parameters,omitempty"`
	Password        string              `json:"password,omitempty"`
	PostType        int                 `json:"post_type,omitempty"`
	Posts           string              `json:"posts,omitempty"`
	PrivateKey      string              `json:"privatekey,omitempty"`
	PublicKey       string              `json:"publickey,omitempty"`
	QueryFields     []map[string]string `json:"query_fields,omitempty"`
	RequestMethod   int                 `json:"request_method,omitempty"`
	RetrieveMode    int                 `json:"retriev_emode,omitempty"`
	SNMPOid         string              `json:"snmp_oid,omitempty"`
	SSLCertFile     string              `json:"ssl_cert_file,omitempty"`
	SSLKeyFile      string              `json:"ssl_key_file,omitempty"`
	SSLKeyPassword  string              `json:"ssl_key_password,omitempty"`
	Status          int                 `json:"status,omitempty"`
	StatusCodes     string              `json:"status_codes,omitempty"`
	TemplateID      string              `json:"templateid,omitempty"`
	Timeout         string              `json:"timeout,omitempty"`
	TrapperHosts    string              `json:"trapper_hosts,omitempty"`
	Trends          string              `json:"trends,omitempty"`
	Units           string              `json:"units,omitempty"`
	Username        string              `json:"username,omitempty"`
	UUID            string              `json:"uuid,omitempty"`
	ValueMapID      string              `json:"valuemapid,omitempty"`
	VerifyHost      int                 `json:"verify_host,omitempty"`
	VerifyPeer      int                 `json:"verify_peer,omitempty"`
	Discover        int                 `json:"discover,omitempty"`
}

type ItemPrototypePreprocessing struct {
	Type               int    `json:"type,omitempty"`
	Params             string `json:"params,omitempty"`
	ErrorHandler       int    `json:"error_handler,omitempty"`
	ErrorHandlerParams string `json:"error_handler_params,omitempty"`
}

type ItemPrototypeCreate struct {
	Delay           string                       `json:"delay,omitempty"`
	HostID          string                       `json:"hostid,omitempty"`
	InterfaceID     string                       `json:"interfaceid,omitempty"`
	Key             string                       `json:"key_,omitempty"`
	Name            string                       `json:"name,omitempty"`
	Type            int                          `json:"type,omitempty"`
	Url             string                       `json:"url,omitempty"`
	ValueType       int                          `json:"value_type,omitempty"`
	AllowTraps      int                          `json:"allow_traps,omitempty"`
	AuthType        int                          `json:"authtype,omitempty"`
	Description     string                       `json:"description,omitempty"`
	FollowRedirects int                          `json:"follow_redirects,omitempty"`
	Headers         map[string]string            `json:"headers,omitempty"`
	History         string                       `json:"history,omitempty"`
	HttpProxy       string                       `json:"http_proxy,omitempty"`
	IPMISensor      string                       `json:"ipmi_sensor,omitempty"`
	JmxEndpoint     string                       `json:"jmx_endpoint,omitempty"`
	LogtimeFmt      string                       `json:"logtimefmt,omitempty"`
	MasterItemID    int                          `json:"master_itemid,omitempty"`
	OutputFormat    int                          `json:"output_format,omitempty"`
	Params          string                       `json:"params,omitempty"`
	Parameters      []map[string]string          `json:"parameters,omitempty"`
	Password        string                       `json:"password,omitempty"`
	PostType        int                          `json:"post_type,omitempty"`
	Posts           string                       `json:"posts,omitempty"`
	PrivateKey      string                       `json:"privatekey,omitempty"`
	PublicKey       string                       `json:"publickey,omitempty"`
	QueryFields     []map[string]string          `json:"query_fields,omitempty"`
	RequestMethod   int                          `json:"request_method,omitempty"`
	RetrieveMode    int                          `json:"retriev_emode,omitempty"`
	SNMPOid         string                       `json:"snmp_oid,omitempty"`
	SSLCertFile     string                       `json:"ssl_cert_file,omitempty"`
	SSLKeyFile      string                       `json:"ssl_key_file,omitempty"`
	SSLKeyPassword  string                       `json:"ssl_key_password,omitempty"`
	Status          int                          `json:"status,omitempty"`
	StatusCodes     string                       `json:"status_codes,omitempty"`
	TemplateID      string                       `json:"templateid,omitempty"`
	Timeout         string                       `json:"timeout,omitempty"`
	TrapperHosts    string                       `json:"trapper_hosts,omitempty"`
	Trends          string                       `json:"trends,omitempty"`
	Units           string                       `json:"units,omitempty"`
	Username        string                       `json:"username,omitempty"`
	UUID            string                       `json:"uuid,omitempty"`
	ValueMapID      string                       `json:"valuemapid,omitempty"`
	VerifyHost      int                          `json:"verify_host,omitempty"`
	VerifyPeer      int                          `json:"verify_peer,omitempty"`
	RuleID          string                       `json:"ruleid,omitempty"`
	Preprocessing   []ItemPrototypePreprocessing `json:"preprocessing,omitempty"`
	Tags            []ZabbixTag                  `json:"tags,omitempty"`
}

type ItemPrototypeUpdate struct {
	ItemID          string                       `json:"itemid,omitempty"`
	Delay           string                       `json:"delay,omitempty"`
	HostID          string                       `json:"hostid,omitempty"`
	InterfaceID     string                       `json:"interfaceid,omitempty"`
	Key             string                       `json:"key_,omitempty"`
	Name            string                       `json:"name,omitempty"`
	Type            int                          `json:"type,omitempty"`
	Url             string                       `json:"url,omitempty"`
	ValueType       int                          `json:"value_type,omitempty"`
	AllowTraps      int                          `json:"allow_traps,omitempty"`
	AuthType        int                          `json:"authtype,omitempty"`
	Description     string                       `json:"description,omitempty"`
	FollowRedirects int                          `json:"follow_redirects,omitempty"`
	Headers         map[string]string            `json:"headers,omitempty"`
	History         string                       `json:"history,omitempty"`
	HttpProxy       string                       `json:"http_proxy,omitempty"`
	IPMISensor      string                       `json:"ipmi_sensor,omitempty"`
	JmxEndpoint     string                       `json:"jmx_endpoint,omitempty"`
	LogtimeFmt      string                       `json:"logtimefmt,omitempty"`
	MasterItemID    int                          `json:"master_itemid,omitempty"`
	OutputFormat    int                          `json:"output_format,omitempty"`
	Params          string                       `json:"params,omitempty"`
	Parameters      []map[string]string          `json:"parameters,omitempty"`
	Password        string                       `json:"password,omitempty"`
	PostType        int                          `json:"post_type,omitempty"`
	Posts           string                       `json:"posts,omitempty"`
	PrivateKey      string                       `json:"privatekey,omitempty"`
	PublicKey       string                       `json:"publickey,omitempty"`
	QueryFields     []map[string]string          `json:"query_fields,omitempty"`
	RequestMethod   int                          `json:"request_method,omitempty"`
	RetrieveMode    int                          `json:"retriev_emode,omitempty"`
	SNMPOid         string                       `json:"snmp_oid,omitempty"`
	SSLCertFile     string                       `json:"ssl_cert_file,omitempty"`
	SSLKeyFile      string                       `json:"ssl_key_file,omitempty"`
	SSLKeyPassword  string                       `json:"ssl_key_password,omitempty"`
	Status          int                          `json:"status,omitempty"`
	StatusCodes     string                       `json:"status_codes,omitempty"`
	TemplateID      string                       `json:"templateid,omitempty"`
	Timeout         string                       `json:"timeout,omitempty"`
	TrapperHosts    string                       `json:"trapper_hosts,omitempty"`
	Trends          string                       `json:"trends,omitempty"`
	Units           string                       `json:"units,omitempty"`
	Username        string                       `json:"username,omitempty"`
	UUID            string                       `json:"uuid,omitempty"`
	ValueMapID      string                       `json:"valuemapid,omitempty"`
	VerifyHost      int                          `json:"verify_host,omitempty"`
	VerifyPeer      int                          `json:"verify_peer,omitempty"`
	RuleID          string                       `json:"ruleid,omitempty"`
	Preprocessing   []ItemPrototypePreprocessing `json:"preprocessing,omitempty"`
	Tags            []ZabbixTag                  `json:"tags,omitempty"`
}

type ItemPrototypeResponse struct {
	ItemIDs []string `json:"itemids,omitempty"`
}

type ItemPrototypeGet struct {
	DiscoveryIDs        []string            `json:"discoveryids,omitempty"`
	ItemIDs             []string            `json:"itemids,omitempty"`
	GroupIDs            []string            `json:"groupids,omitempty"`
	TemplateIDs         []string            `json:"templateids,omitempty"`
	HostIDs             []string            `json:"hostids,omitempty"`
	ProxyIDs            []string            `json:"proxyids,omitempty"`
	InterfaceIDs        []string            `json:"interfaceids,omitempty"`
	GraphIDs            []string            `json:"graphids,omitempty"`
	TriggerIDs          []string            `json:"triggerids,omitempty"`
	WebItems            string              `json:"webitems,omitempty"`
	Inherited           bool                `json:"inherited,omitempty"`
	Templated           bool                `json:"templated,omitempty"`
	Monitored           bool                `json:"monitored,omitempty"`
	Group               string              `json:"group,omitempty"`
	Host                string              `json:"host,omitempty"`
	Evaltype            int                 `json:"evaltype,omitempty"`
	Tags                []map[string]string `json:"tags,omitempty"`
	WithTriggers        bool                `json:"with_triggers,omitempty"`
	SelectHosts         string              `json:"selectHosts,omitempty"`
	SelectInterfaces    string              `json:"selectInterfaces,omitempty"`
	SelectTriggers      string              `json:"selectTriggers,omitempty"`
	SelectGraphs        string              `json:"selectGraphs,omitempty"`
	SelectDiscoveryRule string              `json:"selectDiscoveryRule,omitempty"`
	SelectPreprocessing string              `json:"selectPreprocessing,omitempty"`
	SelectTags          string              `json:"selectTags,omitempty"`
	SelectValueMap      string              `json:"selectValueMap,omitempty"`
	ZabbixCommun
}
