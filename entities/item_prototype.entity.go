package entities

type ItemPrototypeObject struct {
	ItemID          string              `json:"itemid"`
	Delay           string              `json:"delay"`
	HostID          string              `json:"hostid"`
	InterfaceID     string              `json:"interfaceid"`
	Key             string              `json:"key_"`
	Name            string              `json:"name"`
	Type            int                 `json:"type"`
	Url             string              `json:"url"`
	ValueType       int                 `json:"value_type"`
	AllowTraps      int                 `json:"allow_traps"`
	AuthType        int                 `json:"authtype"`
	Description     string              `json:"description"`
	FollowRedirects int                 `json:"follow_redirects"`
	Headers         map[string]string   `json:"headers"`
	History         string              `json:"history"`
	HttpProxy       string              `json:"http_proxy"`
	IPMISensor      string              `json:"ipmi_sensor"`
	JmxEndpoint     string              `json:"jmx_endpoint"`
	LogtimeFmt      string              `json:"logtimefmt"`
	MasterItemID    int                 `json:"master_itemid"`
	OutputFormat    int                 `json:"output_format"`
	Params          string              `json:"params"`
	Parameters      []map[string]string `json:"parameters"`
	Password        string              `json:"password"`
	PostType        int                 `json:"post_type"`
	Posts           string              `json:"posts"`
	PrivateKey      string              `json:"privatekey"`
	PublicKey       string              `json:"publickey"`
	QueryFields     []map[string]string `json:"query_fields"`
	RequestMethod   int                 `json:"request_method"`
	RetrieveMode    int                 `json:"retriev_emode"`
	SNMPOid         string              `json:"snmp_oid"`
	SSLCertFile     string              `json:"ssl_cert_file"`
	SSLKeyFile      string              `json:"ssl_key_file"`
	SSLKeyPassword  string              `json:"ssl_key_password"`
	Status          int                 `json:"status"`
	StatusCodes     string              `json:"status_codes"`
	TemplateID      string              `json:"templateid"`
	Timeout         string              `json:"timeout"`
	TrapperHosts    string              `json:"trapper_hosts"`
	Trends          string              `json:"trends"`
	Units           string              `json:"units"`
	Username        string              `json:"username"`
	UUID            string              `json:"uuid"`
	ValueMapID      string              `json:"valuemapid"`
	VerifyHost      int                 `json:"verify_host"`
	VerifyPeer      int                 `json:"verify_peer"`
	Discover        int                 `json:"discover"`
}

type ItemPrototypePreprocessing struct {
	Type               int    `json:"type"`
	Params             string `json:"params"`
	ErrorHandler       int    `json:"error_handler"`
	ErrorHandlerParams string `json:"error_handler_params"`
}

type ItemPrototypeCreate struct {
	Delay           string                       `json:"delay"`
	HostID          string                       `json:"hostid"`
	InterfaceID     string                       `json:"interfaceid"`
	Key             string                       `json:"key_"`
	Name            string                       `json:"name"`
	Type            int                          `json:"type"`
	Url             string                       `json:"url"`
	ValueType       int                          `json:"value_type"`
	AllowTraps      int                          `json:"allow_traps"`
	AuthType        int                          `json:"authtype"`
	Description     string                       `json:"description"`
	FollowRedirects int                          `json:"follow_redirects"`
	Headers         map[string]string            `json:"headers"`
	History         string                       `json:"history"`
	HttpProxy       string                       `json:"http_proxy"`
	IPMISensor      string                       `json:"ipmi_sensor"`
	JmxEndpoint     string                       `json:"jmx_endpoint"`
	LogtimeFmt      string                       `json:"logtimefmt"`
	MasterItemID    int                          `json:"master_itemid"`
	OutputFormat    int                          `json:"output_format"`
	Params          string                       `json:"params"`
	Parameters      []map[string]string          `json:"parameters"`
	Password        string                       `json:"password"`
	PostType        int                          `json:"post_type"`
	Posts           string                       `json:"posts"`
	PrivateKey      string                       `json:"privatekey"`
	PublicKey       string                       `json:"publickey"`
	QueryFields     []map[string]string          `json:"query_fields"`
	RequestMethod   int                          `json:"request_method"`
	RetrieveMode    int                          `json:"retriev_emode"`
	SNMPOid         string                       `json:"snmp_oid"`
	SSLCertFile     string                       `json:"ssl_cert_file"`
	SSLKeyFile      string                       `json:"ssl_key_file"`
	SSLKeyPassword  string                       `json:"ssl_key_password"`
	Status          int                          `json:"status"`
	StatusCodes     string                       `json:"status_codes"`
	TemplateID      string                       `json:"templateid"`
	Timeout         string                       `json:"timeout"`
	TrapperHosts    string                       `json:"trapper_hosts"`
	Trends          string                       `json:"trends"`
	Units           string                       `json:"units"`
	Username        string                       `json:"username"`
	UUID            string                       `json:"uuid"`
	ValueMapID      string                       `json:"valuemapid"`
	VerifyHost      int                          `json:"verify_host"`
	VerifyPeer      int                          `json:"verify_peer"`
	RuleID          string                       `json:"ruleid"`
	Preprocessing   []ItemPrototypePreprocessing `json:"preprocessing"`
	Tags            []ZabbixTag                  `json:"tags"`
}

type ItemPrototypeUpdate struct {
	ItemID          string                       `json:"itemid"`
	Delay           string                       `json:"delay"`
	HostID          string                       `json:"hostid"`
	InterfaceID     string                       `json:"interfaceid"`
	Key             string                       `json:"key_"`
	Name            string                       `json:"name"`
	Type            int                          `json:"type"`
	Url             string                       `json:"url"`
	ValueType       int                          `json:"value_type"`
	AllowTraps      int                          `json:"allow_traps"`
	AuthType        int                          `json:"authtype"`
	Description     string                       `json:"description"`
	FollowRedirects int                          `json:"follow_redirects"`
	Headers         map[string]string            `json:"headers"`
	History         string                       `json:"history"`
	HttpProxy       string                       `json:"http_proxy"`
	IPMISensor      string                       `json:"ipmi_sensor"`
	JmxEndpoint     string                       `json:"jmx_endpoint"`
	LogtimeFmt      string                       `json:"logtimefmt"`
	MasterItemID    int                          `json:"master_itemid"`
	OutputFormat    int                          `json:"output_format"`
	Params          string                       `json:"params"`
	Parameters      []map[string]string          `json:"parameters"`
	Password        string                       `json:"password"`
	PostType        int                          `json:"post_type"`
	Posts           string                       `json:"posts"`
	PrivateKey      string                       `json:"privatekey"`
	PublicKey       string                       `json:"publickey"`
	QueryFields     []map[string]string          `json:"query_fields"`
	RequestMethod   int                          `json:"request_method"`
	RetrieveMode    int                          `json:"retriev_emode"`
	SNMPOid         string                       `json:"snmp_oid"`
	SSLCertFile     string                       `json:"ssl_cert_file"`
	SSLKeyFile      string                       `json:"ssl_key_file"`
	SSLKeyPassword  string                       `json:"ssl_key_password"`
	Status          int                          `json:"status"`
	StatusCodes     string                       `json:"status_codes"`
	TemplateID      string                       `json:"templateid"`
	Timeout         string                       `json:"timeout"`
	TrapperHosts    string                       `json:"trapper_hosts"`
	Trends          string                       `json:"trends"`
	Units           string                       `json:"units"`
	Username        string                       `json:"username"`
	UUID            string                       `json:"uuid"`
	ValueMapID      string                       `json:"valuemapid"`
	VerifyHost      int                          `json:"verify_host"`
	VerifyPeer      int                          `json:"verify_peer"`
	RuleID          string                       `json:"ruleid"`
	Preprocessing   []ItemPrototypePreprocessing `json:"preprocessing"`
	Tags            []ZabbixTag                  `json:"tags"`
}

type ItemPrototypeResponse struct {
	ItemIDs []string `json:"itemids"`
}

type ItemPrototypeGet struct {
	DiscoveryIDs        []string            `json:"discoveryids"`
	ItemIDs             []string            `json:"itemids"`
	GroupIDs            []string            `json:"groupids"`
	TemplateIDs         []string            `json:"templateids"`
	HostIDs             []string            `json:"hostids"`
	ProxyIDs            []string            `json:"proxyids"`
	InterfaceIDs        []string            `json:"interfaceids"`
	GraphIDs            []string            `json:"graphids"`
	TriggerIDs          []string            `json:"triggerids"`
	WebItems            string              `json:"webitems"`
	Inherited           bool                `json:"inherited"`
	Templated           bool                `json:"templated"`
	Monitored           bool                `json:"monitored"`
	Group               string              `json:"group"`
	Host                string              `json:"host"`
	Evaltype            int                 `json:"evaltype"`
	Tags                []map[string]string `json:"tags"`
	WithTriggers        bool                `json:"with_triggers"`
	SelectHosts         string              `json:"selectHosts"`
	SelectInterfaces    string              `json:"selectInterfaces"`
	SelectTriggers      string              `json:"selectTriggers"`
	SelectGraphs        string              `json:"selectGraphs"`
	SelectDiscoveryRule string              `json:"selectDiscoveryRule"`
	SelectPreprocessing string              `json:"selectPreprocessing"`
	SelectTags          string              `json:"selectTags"`
	SelectValueMap      string              `json:"selectValueMap"`
	ZabbixCommun
}
