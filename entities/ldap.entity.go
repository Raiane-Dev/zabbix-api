package entities

type LLDObject struct {
	ItemID          string            `json:"itemid,omitempty"`
	Delay           string            `json:"delay,omitempty"`
	HostID          string            `json:"hostid,omitempty"`
	InterfaceID     string            `json:"interfaceid,omitempty"`
	Key             string            `json:"key_,omitempty"`
	Name            string            `json:"name,omitempty"`
	Type            int               `json:"type,omitempty"`
	Url             string            `json:"url,omitempty"`
	AllowTraps      int               `json:"allow_traps,omitempty"`
	AuthType        int               `json:"authtype,omitempty"`
	Description     string            `json:"description,omitempty"`
	Error           string            `json:"error,omitempty"`
	FollowRedirects int               `json:"follow_redirects,omitempty"`
	Headers         map[string]string `json:"headers,omitempty"`
	HttpProxy       string            `json:"http_proxy,omitempty"`
	IPMISensor      string            `json:"ipmi_sensor,omitempty"`
	JmxEndpoint     string            `json:"jmx_endpoint,omitempty"`
	Lifetime        string            `json:"lifetime,omitempty"`
	MasterItemID    int               `json:"master_itemid,omitempty"`
	OutputFormat    int               `json:"output_format,omitempty"`
	Params          string            `json:"params,omitempty"`
	Parameters      string            `json:"parameters,omitempty"`
	Password        string            `json:"password,omitempty"`
	PostType        int               `json:"post_type,omitempty"`
	Posts           string            `json:"posts,omitempty"`
	PrivateKey      string            `json:"privatekey,omitempty"`
	PublicKey       string            `json:"publickey,omitempty"`
	QueryFields     []string          `json:"query_fields,omitempty"`
	RequestMethod   int               `json:"request_method,omitempty"`
	RetrieveMode    int               `json:"retrieve_mode,omitempty"`
	SNMPOid         string            `json:"snmp_oid,omitempty"`
	SSLCertFile     string            `json:"ssl_cert_file,omitempty"`
	SSLKeyFile      string            `json:"ssl_key_file,omitempty"`
	SSLKeyPassword  string            `json:"ssl_key_password,omitempty"`
	State           int               `json:"state,omitempty"`
	Status          int               `json:"status,omitempty"`
	StatusCodes     string            `json:"status_codes,omitempty"`
	TemplateID      string            `json:"templateid,omitempty"`
	Timeout         string            `json:"timeout,omitempty"`
	TrapperHosts    string            `json:"trapper_hosts,omitempty"`
	Username        string            `json:"username,omitempty"`
	UUID            string            `json:"uuid,omitempty"`
	VerifyHost      int               `json:"verify_host,omitempty"`
	VerifyPeer      int               `json:"verify_peer,omitempty"`
}

type LLDCreate struct {
	Delay           string                 `json:"delay,omitempty"`
	HostID          string                 `json:"hostid,omitempty"`
	InterfaceID     string                 `json:"interfaceid,omitempty"`
	Key             string                 `json:"key_,omitempty"`
	Name            string                 `json:"name,omitempty"`
	Type            int                    `json:"type,omitempty"`
	Url             string                 `json:"url,omitempty"`
	AllowTraps      int                    `json:"allow_traps,omitempty"`
	AuthType        int                    `json:"authtype,omitempty"`
	Description     string                 `json:"description,omitempty"`
	Error           string                 `json:"error,omitempty"`
	FollowRedirects int                    `json:"follow_redirects,omitempty"`
	Headers         map[string]string      `json:"headers,omitempty"`
	HttpProxy       string                 `json:"http_proxy,omitempty"`
	IPMISensor      string                 `json:"ipmi_sensor,omitempty"`
	JmxEndpoint     string                 `json:"jmx_endpoint,omitempty"`
	Lifetime        string                 `json:"lifetime,omitempty"`
	MasterItemID    int                    `json:"master_itemid,omitempty"`
	OutputFormat    int                    `json:"output_format,omitempty"`
	Params          string                 `json:"params,omitempty"`
	Parameters      string                 `json:"parameters,omitempty"`
	Password        string                 `json:"password,omitempty"`
	PostType        int                    `json:"post_type,omitempty"`
	Posts           string                 `json:"posts,omitempty"`
	PrivateKey      string                 `json:"privatekey,omitempty"`
	PublicKey       string                 `json:"publickey,omitempty"`
	QueryFields     []string               `json:"query_fields,omitempty"`
	RequestMethod   int                    `json:"request_method,omitempty"`
	RetrieveMode    int                    `json:"retrieve_mode,omitempty"`
	SNMPOid         string                 `json:"snmp_oid,omitempty"`
	SSLCertFile     string                 `json:"ssl_cert_file,omitempty"`
	SSLKeyFile      string                 `json:"ssl_key_file,omitempty"`
	SSLKeyPassword  string                 `json:"ssl_key_password,omitempty"`
	State           int                    `json:"state,omitempty"`
	Status          int                    `json:"status,omitempty"`
	StatusCodes     string                 `json:"status_codes,omitempty"`
	TemplateID      string                 `json:"templateid,omitempty"`
	Timeout         string                 `json:"timeout,omitempty"`
	TrapperHosts    string                 `json:"trapper_hosts,omitempty"`
	Username        string                 `json:"username,omitempty"`
	UUID            string                 `json:"uuid,omitempty"`
	VerifyHost      int                    `json:"verify_host,omitempty"`
	VerifyPeer      int                    `json:"verify_peer,omitempty"`
	Filter          LLDRuleFilter          `json:"filter,omitempty"`
	Preprocessing   []LLDRulePreprocessing `json:"preprocessing,omitempty"`
	LLDMacroPaths   []LLDMacroPath         `json:"lld_macro_paths,omitempty"`
	Overrides       []LLDOverride          `json:"overrides,omitempty"`
}

type LLDUpdate struct {
	ItemID          string                 `json:"itemid,omitempty"`
	Delay           string                 `json:"delay,omitempty"`
	HostID          string                 `json:"hostid,omitempty"`
	InterfaceID     string                 `json:"interfaceid,omitempty"`
	Key             string                 `json:"key_,omitempty"`
	Name            string                 `json:"name,omitempty"`
	Type            int                    `json:"type,omitempty"`
	Url             string                 `json:"url,omitempty"`
	AllowTraps      int                    `json:"allow_traps,omitempty"`
	AuthType        int                    `json:"authtype,omitempty"`
	Description     string                 `json:"description,omitempty"`
	Error           string                 `json:"error,omitempty"`
	FollowRedirects int                    `json:"follow_redirects,omitempty"`
	Headers         map[string]string      `json:"headers,omitempty"`
	HttpProxy       string                 `json:"http_proxy,omitempty"`
	IPMISensor      string                 `json:"ipmi_sensor,omitempty"`
	JmxEndpoint     string                 `json:"jmx_endpoint,omitempty"`
	Lifetime        string                 `json:"lifetime,omitempty"`
	MasterItemID    int                    `json:"master_itemid,omitempty"`
	OutputFormat    int                    `json:"output_format,omitempty"`
	Params          string                 `json:"params,omitempty"`
	Parameters      string                 `json:"parameters,omitempty"`
	Password        string                 `json:"password,omitempty"`
	PostType        int                    `json:"post_type,omitempty"`
	Posts           string                 `json:"posts,omitempty"`
	PrivateKey      string                 `json:"privatekey,omitempty"`
	PublicKey       string                 `json:"publickey,omitempty"`
	QueryFields     []string               `json:"query_fields,omitempty"`
	RequestMethod   int                    `json:"request_method,omitempty"`
	RetrieveMode    int                    `json:"retrieve_mode,omitempty"`
	SNMPOid         string                 `json:"snmp_oid,omitempty"`
	SSLCertFile     string                 `json:"ssl_cert_file,omitempty"`
	SSLKeyFile      string                 `json:"ssl_key_file,omitempty"`
	SSLKeyPassword  string                 `json:"ssl_key_password,omitempty"`
	State           int                    `json:"state,omitempty"`
	Status          int                    `json:"status,omitempty"`
	StatusCodes     string                 `json:"status_codes,omitempty"`
	TemplateID      string                 `json:"templateid,omitempty"`
	Timeout         string                 `json:"timeout,omitempty"`
	TrapperHosts    string                 `json:"trapper_hosts,omitempty"`
	Username        string                 `json:"username,omitempty"`
	UUID            string                 `json:"uuid,omitempty"`
	VerifyHost      int                    `json:"verify_host,omitempty"`
	VerifyPeer      int                    `json:"verify_peer,omitempty"`
	Filter          LLDRuleFilter          `json:"filter,omitempty"`
	Preprocessing   []LLDRulePreprocessing `json:"preprocessing,omitempty"`
	LLDMacroPaths   []LLDMacroPath         `json:"lld_macro_paths,omitempty"`
	Overrides       []LLDOverride          `json:"overrides,omitempty"`
}

type LLDGet struct {
	ItemIDs              []string `json:"itemids,omitempty"`
	GroupIDs             []string `json:"groupids,omitempty"`
	HostIDs              []string `json:"hostids,omitempty"`
	Inherited            bool     `json:"inherited,omitempty"`
	InterfaceIDs         []string `json:"interfaceids,omitempty"`
	Monitored            bool     `json:"monitored,omitempty"`
	TemplateIDs          []string `json:"templateids,omitempty"`
	SelectFilter         string   `json:"selectFilter,omitempty"`
	SelectGraphs         string   `json:"selectGraphs,omitempty"`
	SelectHostPrototypes string   `json:"selectHostPrototypes,omitempty"`
	SelectHosts          string   `json:"selectHosts,omitempty"`
	SelectItems          string   `json:"selectItems,omitempty"`
	SelectTriggers       string   `json:"selectTriggers,omitempty"`
	SelectLLDMacroPaths  string   `json:"selectLLDMacroPaths,omitempty"`
	SelectPreprocessing  string   `json:"selectPreprocessing,omitempty"`
	SelectOverrides      string   `json:"selectOverrides,omitempty"`
	ZabbixCommun
}

type LLDCopy struct {
	DiscoveryIDs []string `json:"discoveryids,omitempty"`
	HostIDs      []string `json:"hostids,omitempty"`
}

type LLDOverride struct {
	Name       string              `json:"name,omitempty"`
	Step       int                 `json:"step,omitempty"`
	Stop       int                 `json:"stop,omitempty"`
	Filter     map[string]string   `json:"filter,omitempty"`
	Operations []map[string]string `json:"operations,omitempty"`
}

type LLDRuleFilter struct {
	Conditions  []LLDRuleFilterCondition `json:"conditions,omitempty"`
	Evaltype    int                      `json:"evaltype,omitempty"`
	EvalFormula string                   `json:"eval_formula,omitempty"`
	Formula     string                   `json:"formula,omitempty"`
}

type LLDRuleFilterCondition struct {
	Macro     string `json:"macro,omitempty"`
	Value     string `json:"value,omitempty"`
	FormulaID string `json:"formulaid,omitempty"`
	Operator  int    `json:"operator,omitempty"`
}

type LLDMacroPath struct {
	LLDMacro string `json:"lld_macro,omitempty"`
	Path     string `json:"path,omitempty"`
}

type LLDRulePreprocessing struct {
	Type               int    `json:"type,omitempty"`
	Params             string `json:"params,omitempty"`
	ErrorHandler       int    `json:"error_handler,omitempty"`
	ErrorHandlerParams string `json:"error_handler_params,omitempty"`
}
