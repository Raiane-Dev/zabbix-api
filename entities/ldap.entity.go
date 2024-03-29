package entities

type LLDObject struct {
	ItemID          string            `json:"itemid"`
	Delay           string            `json:"delay"`
	HostID          string            `json:"hostid"`
	InterfaceID     string            `json:"interfaceid"`
	Key             string            `json:"key_"`
	Name            string            `json:"name"`
	Type            int               `json:"type"`
	Url             string            `json:"url"`
	AllowTraps      int               `json:"allow_traps"`
	AuthType        int               `json:"authtype"`
	Description     string            `json:"description"`
	Error           string            `json:"error"`
	FollowRedirects int               `json:"follow_redirects"`
	Headers         map[string]string `json:"headers"`
	HttpProxy       string            `json:"http_proxy"`
	IPMISensor      string            `json:"ipmi_sensor"`
	JmxEndpoint     string            `json:"jmx_endpoint"`
	Lifetime        string            `json:"lifetime"`
	MasterItemID    int               `json:"master_itemid"`
	OutputFormat    int               `json:"output_format"`
	Params          string            `json:"params"`
	Parameters      string            `json:"parameters"`
	Password        string            `json:"password"`
	PostType        int               `json:"post_type"`
	Posts           string            `json:"posts"`
	PrivateKey      string            `json:"privatekey"`
	PublicKey       string            `json:"publickey"`
	QueryFields     []string          `json:"query_fields"`
	RequestMethod   int               `json:"request_method"`
	RetrieveMode    int               `json:"retrieve_mode"`
	SNMPOid         string            `json:"snmp_oid"`
	SSLCertFile     string            `json:"ssl_cert_file"`
	SSLKeyFile      string            `json:"ssl_key_file"`
	SSLKeyPassword  string            `json:"ssl_key_password"`
	State           int               `json:"state"`
	Status          int               `json:"status"`
	StatusCodes     string            `json:"status_codes"`
	TemplateID      string            `json:"templateid"`
	Timeout         string            `json:"timeout"`
	TrapperHosts    string            `json:"trapper_hosts"`
	Username        string            `json:"username"`
	UUID            string            `json:"uuid"`
	VerifyHost      int               `json:"verify_host"`
	VerifyPeer      int               `json:"verify_peer"`
}

type LLDCreate struct {
	Delay           string                 `json:"delay"`
	HostID          string                 `json:"hostid"`
	InterfaceID     string                 `json:"interfaceid"`
	Key             string                 `json:"key_"`
	Name            string                 `json:"name"`
	Type            int                    `json:"type"`
	Url             string                 `json:"url"`
	AllowTraps      int                    `json:"allow_traps"`
	AuthType        int                    `json:"authtype"`
	Description     string                 `json:"description"`
	Error           string                 `json:"error"`
	FollowRedirects int                    `json:"follow_redirects"`
	Headers         map[string]string      `json:"headers"`
	HttpProxy       string                 `json:"http_proxy"`
	IPMISensor      string                 `json:"ipmi_sensor"`
	JmxEndpoint     string                 `json:"jmx_endpoint"`
	Lifetime        string                 `json:"lifetime"`
	MasterItemID    int                    `json:"master_itemid"`
	OutputFormat    int                    `json:"output_format"`
	Params          string                 `json:"params"`
	Parameters      string                 `json:"parameters"`
	Password        string                 `json:"password"`
	PostType        int                    `json:"post_type"`
	Posts           string                 `json:"posts"`
	PrivateKey      string                 `json:"privatekey"`
	PublicKey       string                 `json:"publickey"`
	QueryFields     []string               `json:"query_fields"`
	RequestMethod   int                    `json:"request_method"`
	RetrieveMode    int                    `json:"retrieve_mode"`
	SNMPOid         string                 `json:"snmp_oid"`
	SSLCertFile     string                 `json:"ssl_cert_file"`
	SSLKeyFile      string                 `json:"ssl_key_file"`
	SSLKeyPassword  string                 `json:"ssl_key_password"`
	State           int                    `json:"state"`
	Status          int                    `json:"status"`
	StatusCodes     string                 `json:"status_codes"`
	TemplateID      string                 `json:"templateid"`
	Timeout         string                 `json:"timeout"`
	TrapperHosts    string                 `json:"trapper_hosts"`
	Username        string                 `json:"username"`
	UUID            string                 `json:"uuid"`
	VerifyHost      int                    `json:"verify_host"`
	VerifyPeer      int                    `json:"verify_peer"`
	Filter          LLDRuleFilter          `json:"filter"`
	Preprocessing   []LLDRulePreprocessing `json:"preprocessing"`
	LLDMacroPaths   []LLDMacroPath         `json:"lld_macro_paths"`
	Overrides       []LLDOverride          `json:"overrides"`
}

type LLDUpdate struct {
	ItemID          string                 `json:"itemid"`
	Delay           string                 `json:"delay"`
	HostID          string                 `json:"hostid"`
	InterfaceID     string                 `json:"interfaceid"`
	Key             string                 `json:"key_"`
	Name            string                 `json:"name"`
	Type            int                    `json:"type"`
	Url             string                 `json:"url"`
	AllowTraps      int                    `json:"allow_traps"`
	AuthType        int                    `json:"authtype"`
	Description     string                 `json:"description"`
	Error           string                 `json:"error"`
	FollowRedirects int                    `json:"follow_redirects"`
	Headers         map[string]string      `json:"headers"`
	HttpProxy       string                 `json:"http_proxy"`
	IPMISensor      string                 `json:"ipmi_sensor"`
	JmxEndpoint     string                 `json:"jmx_endpoint"`
	Lifetime        string                 `json:"lifetime"`
	MasterItemID    int                    `json:"master_itemid"`
	OutputFormat    int                    `json:"output_format"`
	Params          string                 `json:"params"`
	Parameters      string                 `json:"parameters"`
	Password        string                 `json:"password"`
	PostType        int                    `json:"post_type"`
	Posts           string                 `json:"posts"`
	PrivateKey      string                 `json:"privatekey"`
	PublicKey       string                 `json:"publickey"`
	QueryFields     []string               `json:"query_fields"`
	RequestMethod   int                    `json:"request_method"`
	RetrieveMode    int                    `json:"retrieve_mode"`
	SNMPOid         string                 `json:"snmp_oid"`
	SSLCertFile     string                 `json:"ssl_cert_file"`
	SSLKeyFile      string                 `json:"ssl_key_file"`
	SSLKeyPassword  string                 `json:"ssl_key_password"`
	State           int                    `json:"state"`
	Status          int                    `json:"status"`
	StatusCodes     string                 `json:"status_codes"`
	TemplateID      string                 `json:"templateid"`
	Timeout         string                 `json:"timeout"`
	TrapperHosts    string                 `json:"trapper_hosts"`
	Username        string                 `json:"username"`
	UUID            string                 `json:"uuid"`
	VerifyHost      int                    `json:"verify_host"`
	VerifyPeer      int                    `json:"verify_peer"`
	Filter          LLDRuleFilter          `json:"filter"`
	Preprocessing   []LLDRulePreprocessing `json:"preprocessing"`
	LLDMacroPaths   []LLDMacroPath         `json:"lld_macro_paths"`
	Overrides       []LLDOverride          `json:"overrides"`
}

type LLDGet struct {
	ItemIDs              []string `json:"itemids"`
	GroupIDs             []string `json:"groupids"`
	HostIDs              []string `json:"hostids"`
	Inherited            bool     `json:"inherited"`
	InterfaceIDs         []string `json:"interfaceids"`
	Monitored            bool     `json:"monitored"`
	TemplateIDs          []string `json:"templateids"`
	SelectFilter         string   `json:"selectFilter"`
	SelectGraphs         string   `json:"selectGraphs"`
	SelectHostPrototypes string   `json:"selectHostPrototypes"`
	SelectHosts          string   `json:"selectHosts"`
	SelectItems          string   `json:"selectItems"`
	SelectTriggers       string   `json:"selectTriggers"`
	SelectLLDMacroPaths  string   `json:"selectLLDMacroPaths"`
	SelectPreprocessing  string   `json:"selectPreprocessing"`
	SelectOverrides      string   `json:"selectOverrides"`
	ZabbixCommun
}

type LLDCopy struct {
	DiscoveryIDs []string `json:"discoveryids"`
	HostIDs      []string `json:"hostids"`
}

type LLDOverride struct {
	Name       string              `json:"name"`
	Step       int                 `json:"step"`
	Stop       int                 `json:"stop"`
	Filter     map[string]string   `json:"filter"`
	Operations []map[string]string `json:"operations"`
}

type LLDRuleFilter struct {
	Conditions  []LLDRuleFilterCondition `json:"conditions"`
	Evaltype    int                      `json:"evaltype"`
	EvalFormula string                   `json:"eval_formula"`
	Formula     string                   `json:"formula"`
}

type LLDRuleFilterCondition struct {
	Macro     string `json:"macro"`
	Value     string `json:"value"`
	FormulaID string `json:"formulaid"`
	Operator  int    `json:"operator"`
}

type LLDMacroPath struct {
	LLDMacro string `json:"lld_macro"`
	Path     string `json:"path"`
}

type LLDRulePreprocessing struct {
	Type               int    `json:"type"`
	Params             string `json:"params"`
	ErrorHandler       int    `json:"error_handler"`
	ErrorHandlerParams string `json:"error_handler_params"`
}
