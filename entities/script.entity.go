package entities

type ScriptObject struct {
	ScriptID     string                   `json:"scriptid,omitempty"`
	Name         string                   `json:"name,omitempty"`
	Type         int                      `json:"type,omitempty"`
	Command      string                   `json:"command,omitempty"`
	Scope        int                      `json:"scope,omitempty"`
	ExecuteOn    int                      `json:"execute_on,omitempty"`
	MenuPath     string                   `json:"menu_path,omitempty"`
	AuthType     int                      `json:"authtype,omitempty"`
	Username     string                   `json:"username,omitempty"`
	Password     string                   `json:"password,omitempty"`
	PublicKey    string                   `json:"publickey,omitempty"`
	PrivateKey   string                   `json:"privatekey,omitempty"`
	Port         string                   `json:"port,omitempty"`
	GroupID      string                   `json:"groupid,omitempty"`
	UsrGrpID     string                   `json:"usrgrpid,omitempty"`
	HostAccess   int                      `json:"host_access,omitempty"`
	Confirmation string                   `json:"confirmation,omitempty"`
	Timeout      string                   `json:"timeout,omitempty"`
	Parameters   []ScriptWebhookParameter `json:"parameters,omitempty"`
	Description  string                   `json:"description,omitempty"`
	URL          string                   `json:"url,omitempty"`
	NewWindow    int                      `json:"new_window,omitempty"`
}

type ScriptWebhookParameter struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type ScriptDebug struct {
	Logs []string `json:"logs,omitempty"`
	Ms   string   `json:"ms,omitempty"`
}

type ScriptLogEntry struct {
	Level   int    `json:"level,omitempty"`
	Ms      string `json:"ms,omitempty"`
	Message string `json:"message,omitempty"`
}

type ScriptUpdate struct {
	ScriptID     string                   `json:"scriptid,omitempty"`
	Name         string                   `json:"name,omitempty"`
	Type         int                      `json:"type,omitempty"`
	Command      string                   `json:"command,omitempty"`
	Scope        int                      `json:"scope,omitempty"`
	ExecuteOn    int                      `json:"execute_on,omitempty"`
	MenuPath     string                   `json:"menu_path,omitempty"`
	AuthType     int                      `json:"authtype,omitempty"`
	Username     string                   `json:"username,omitempty"`
	Password     string                   `json:"password,omitempty"`
	PublicKey    string                   `json:"publickey,omitempty"`
	PrivateKey   string                   `json:"privatekey,omitempty"`
	Port         string                   `json:"port,omitempty"`
	GroupID      string                   `json:"groupid,omitempty"`
	UsrGrpID     string                   `json:"usrgrpid,omitempty"`
	HostAccess   int                      `json:"host_access,omitempty"`
	Confirmation string                   `json:"confirmation,omitempty"`
	Timeout      string                   `json:"timeout,omitempty"`
	Parameters   []ScriptWebhookParameter `json:"parameters,omitempty"`
	Description  string                   `json:"description,omitempty"`
	URL          string                   `json:"url,omitempty"`
	NewWindow    int                      `json:"new_window,omitempty"`
}

type ScriptCreate struct {
	Name         string                   `json:"name,omitempty"`
	Type         int                      `json:"type,omitempty"`
	Command      string                   `json:"command,omitempty"`
	Scope        int                      `json:"scope,omitempty"`
	ExecuteOn    int                      `json:"execute_on,omitempty"`
	MenuPath     string                   `json:"menu_path,omitempty"`
	AuthType     int                      `json:"authtype,omitempty"`
	Username     string                   `json:"username,omitempty"`
	Password     string                   `json:"password,omitempty"`
	PublicKey    string                   `json:"publickey,omitempty"`
	PrivateKey   string                   `json:"privatekey,omitempty"`
	Port         string                   `json:"port,omitempty"`
	GroupID      string                   `json:"groupid,omitempty"`
	UsrGrpID     string                   `json:"usrgrpid,omitempty"`
	HostAccess   int                      `json:"host_access,omitempty"`
	Confirmation string                   `json:"confirmation,omitempty"`
	Timeout      string                   `json:"timeout,omitempty"`
	Parameters   []ScriptWebhookParameter `json:"parameters,omitempty"`
	Description  string                   `json:"description,omitempty"`
	URL          string                   `json:"url,omitempty"`
	NewWindow    int                      `json:"new_window,omitempty"`
}

type ScriptExecute struct {
	ScriptID string `json:"scriptid,omitempty"`
	HostID   string `json:"hostid,omitempty"`
	EventID  string `json:"eventid,omitempty"`
}

type ScriptExecuteResponse struct {
	Response string `json:"response,omitempty"`
	Value    string `json:"value,omitempty"`
	Debug    any    `json:"debug,omitempty"`
}

type ScriptGet struct {
	GroupsIDs        []string `json:"groupsids,omitempty"`
	HostIDs          []string `json:"hostids,omitempty"`
	ScriptIDs        []string `json:"scriptids,omitempty"`
	UsrGrpIDs        []string `json:"usrgrpids,omitempty"`
	SelectHostGroups string   `json:"selectHostGroups,omitempty"`
	SelectHosts      string   `json:"selectHosts,omitempty"`
	SelectActions    string   `json:"selectActions,omitempty"`
	ZabbixCommun
}
