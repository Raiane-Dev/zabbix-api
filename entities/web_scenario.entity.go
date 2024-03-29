package entities

type WebScenarioObject struct {
	HTTTestID      string   `json:"httptestid,omitempty"`
	HostID         string   `json:"hostid,omitempty"`
	Name           string   `json:"name,omitempty"`
	Agent          string   `json:"agent,omitempty"`
	Authentication int      `json:"authentication,omitempty"`
	Delay          string   `json:"delay,omitempty"`
	Headers        []string `json:"headers,omitempty"`
	HTTPPassword   string   `json:"http_password,omitempty"`
	HTTPProxy      string   `json:"http_proxy,omitempty"`
	HTTPUser       string   `json:"http_user,omitempty"`
	Retries        int      `json:"retries,omitempty"`
	SSLCertFile    string   `json:"ssl_cert_file,omitempty"`
	SSLKeyFile     string   `json:"ssl_key_file,omitempty"`
	SSLKeyPassword string   `json:"ssl_key_password,omitempty"`
	Status         int      `json:"status,omitempty"`
	TemplateID     string   `json:"templateid,omitempty"`
	Variables      []string `json:"variables,omitempty"`
	VerifyHost     int      `json:"verify_host,omitempty"`
	VerifyPeer     int      `json:"verify_peer,omitempty"`
	UUID           string   `json:"uuid,omitempty"`
}

type ScenarioStepObject struct {
	Name            string      `json:"name,omitempty"`
	No              int         `json:"no,omitempty"`
	URL             string      `json:"url,omitempty"`
	FollowRedirects int         `json:"follow_redirects,omitempty"`
	Headers         []string    `json:"headers,omitempty"`
	Posts           interface{} `json:"posts,omitempty"`
	Required        string      `json:"required,omitempty"`
	RetrieveMode    int         `json:"retrieve_mode,omitempty"`
	StatusCodes     string      `json:"status_codes,omitempty"`
	Timeout         string      `json:"timeout,omitempty"`
	Variables       []string    `json:"variables,omitempty"`
	QueryFields     []string    `json:"query_fields,omitempty"`
}

type HttpField struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type WebScenarioCreate struct {
	HostID         string               `json:"hostid,omitempty"`
	Name           string               `json:"name,omitempty"`
	Agent          string               `json:"agent,omitempty"`
	Authentication int                  `json:"authentication,omitempty"`
	Delay          string               `json:"delay,omitempty"`
	Headers        []string             `json:"headers,omitempty"`
	HTTPPassword   string               `json:"http_password,omitempty"`
	HTTPProxy      string               `json:"http_proxy,omitempty"`
	HTTPUser       string               `json:"http_user,omitempty"`
	Retries        int                  `json:"retries,omitempty"`
	SSLCertFile    string               `json:"ssl_cert_file,omitempty"`
	SSLKeyFile     string               `json:"ssl_key_file,omitempty"`
	SSLKeyPassword string               `json:"ssl_key_password,omitempty"`
	Status         int                  `json:"status,omitempty"`
	TemplateID     string               `json:"templateid,omitempty"`
	Variables      []string             `json:"variables,omitempty"`
	VerifyHost     int                  `json:"verify_host,omitempty"`
	VerifyPeer     int                  `json:"verify_peer,omitempty"`
	UUID           string               `json:"uuid,omitempty"`
	Steps          []ScenarioStepObject `json:"steps,omitempty"`
	Tags           []ZabbixTag          `json:"tags,omitempty"`
}

type WebScenarioUpdate struct {
	HTTTestID      string               `json:"httptestid,omitempty"`
	HostID         string               `json:"hostid,omitempty"`
	Name           string               `json:"name,omitempty"`
	Agent          string               `json:"agent,omitempty"`
	Authentication int                  `json:"authentication,omitempty"`
	Delay          string               `json:"delay,omitempty"`
	Headers        []string             `json:"headers,omitempty"`
	HTTPPassword   string               `json:"http_password,omitempty"`
	HTTPProxy      string               `json:"http_proxy,omitempty"`
	HTTPUser       string               `json:"http_user,omitempty"`
	Retries        int                  `json:"retries,omitempty"`
	SSLCertFile    string               `json:"ssl_cert_file,omitempty"`
	SSLKeyFile     string               `json:"ssl_key_file,omitempty"`
	SSLKeyPassword string               `json:"ssl_key_password,omitempty"`
	Status         int                  `json:"status,omitempty"`
	TemplateID     string               `json:"templateid,omitempty"`
	Variables      []string             `json:"variables,omitempty"`
	VerifyHost     int                  `json:"verify_host,omitempty"`
	VerifyPeer     int                  `json:"verify_peer,omitempty"`
	UUID           string               `json:"uuid,omitempty"`
	Steps          []ScenarioStepObject `json:"steps,omitempty"`
	Tags           []ZabbixTag          `json:"tags,omitempty"`
}

type WebScenarioResponse struct {
	HttpTestIDs []string `json:"httptestids,omitempty"`
}

type WebScenarioGet struct {
	GroupIDs       []string    `json:"groupids,omitempty"`
	HostIDs        []string    `json:"hostids,omitempty"`
	HTTPTestIDs    []string    `json:"httptestids,omitempty"`
	Inherited      bool        `json:"inherited,omitempty"`
	Monitored      bool        `json:"monitored,omitempty"`
	Templated      bool        `json:"templated,omitempty"`
	TemplateIDs    []string    `json:"templateids,omitempty"`
	ExpandName     bool        `json:"expandName,omitempty"`
	ExpandStepName bool        `json:"expandStepName,omitempty"`
	EvalType       int         `json:"evaltype,omitempty"`
	Tags           []ZabbixTag `json:"tags,omitempty"`
	SelectHosts    string      `json:"selectHosts,omitempty"`
	SelectSteps    string      `json:"selectSteps,omitempty"`
	SelectTags     string      `json:"selectTags,omitempty"`
}
