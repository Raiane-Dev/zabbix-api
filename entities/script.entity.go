package entities

type ScriptObject struct {
	ScriptID     string                   `json:"scriptid"`
	Name         string                   `json:"name"`
	Type         int                      `json:"type"`
	Command      string                   `json:"command"`
	Scope        int                      `json:"scope"`
	ExecuteOn    int                      `json:"execute_on"`
	MenuPath     string                   `json:"menu_path"`
	AuthType     int                      `json:"authtype"`
	Username     string                   `json:"username"`
	Password     string                   `json:"password"`
	PublicKey    string                   `json:"publickey"`
	PrivateKey   string                   `json:"privatekey"`
	Port         string                   `json:"port"`
	GroupID      string                   `json:"groupid"`
	UsrGrpID     string                   `json:"usrgrpid"`
	HostAccess   int                      `json:"host_access"`
	Confirmation string                   `json:"confirmation"`
	Timeout      string                   `json:"timeout"`
	Parameters   []ScriptWebhookParameter `json:"parameters"`
	Description  string                   `json:"description"`
	URL          string                   `json:"url"`
	NewWindow    int                      `json:"new_window"`
}

type ScriptWebhookParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ScriptDebug struct {
	Logs []string `json:"logs"`
	Ms   string   `json:"ms"`
}

type ScriptLogEntry struct {
	Level   int    `json:"level"`
	Ms      string `json:"ms"`
	Message string `json:"message"`
}

type ScriptUpdate struct {
	ScriptID     string                   `json:"scriptid"`
	Name         string                   `json:"name"`
	Type         int                      `json:"type"`
	Command      string                   `json:"command"`
	Scope        int                      `json:"scope"`
	ExecuteOn    int                      `json:"execute_on"`
	MenuPath     string                   `json:"menu_path"`
	AuthType     int                      `json:"authtype"`
	Username     string                   `json:"username"`
	Password     string                   `json:"password"`
	PublicKey    string                   `json:"publickey"`
	PrivateKey   string                   `json:"privatekey"`
	Port         string                   `json:"port"`
	GroupID      string                   `json:"groupid"`
	UsrGrpID     string                   `json:"usrgrpid"`
	HostAccess   int                      `json:"host_access"`
	Confirmation string                   `json:"confirmation"`
	Timeout      string                   `json:"timeout"`
	Parameters   []ScriptWebhookParameter `json:"parameters"`
	Description  string                   `json:"description"`
	URL          string                   `json:"url"`
	NewWindow    int                      `json:"new_window"`
}

type ScriptCreate struct {
	Name         string                   `json:"name"`
	Type         int                      `json:"type"`
	Command      string                   `json:"command"`
	Scope        int                      `json:"scope"`
	ExecuteOn    int                      `json:"execute_on"`
	MenuPath     string                   `json:"menu_path"`
	AuthType     int                      `json:"authtype"`
	Username     string                   `json:"username"`
	Password     string                   `json:"password"`
	PublicKey    string                   `json:"publickey"`
	PrivateKey   string                   `json:"privatekey"`
	Port         string                   `json:"port"`
	GroupID      string                   `json:"groupid"`
	UsrGrpID     string                   `json:"usrgrpid"`
	HostAccess   int                      `json:"host_access"`
	Confirmation string                   `json:"confirmation"`
	Timeout      string                   `json:"timeout"`
	Parameters   []ScriptWebhookParameter `json:"parameters"`
	Description  string                   `json:"description"`
	URL          string                   `json:"url"`
	NewWindow    int                      `json:"new_window"`
}

type ScriptExecute struct {
	ScriptID string `json:"scriptid"`
	HostID   string `json:"hostid"`
	EventID  string `json:"eventid"`
}

type ScriptExecuteResponse struct {
	Response string `json:"response"`
	Value    string `json:"value"`
	Debug    any    `json:"debug"`
}

type ScriptGet struct {
	GroupsIDs        []string `json:"groupsids"`
	HostIDs          []string `json:"hostids"`
	ScriptIDs        []string `json:"scriptids"`
	UsrGrpIDs        []string `json:"usrgrpids"`
	SelectHostGroups string   `json:"selectHostGroups"`
	SelectHosts      string   `json:"selectHosts"`
	SelectActions    string   `json:"selectActions"`
	ZabbixCommun
}
