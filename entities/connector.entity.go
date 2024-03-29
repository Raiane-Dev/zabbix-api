package entities

type ConnectorGet struct {
	ConnectorIDs []string `json:"connectorids,omitempty"`
	SelectTags   string   `json:"selectTags,omitempty"`
	ZabbixCommun
}

type ConnectorCreate struct {
	Name           string               `json:"name,omitempty"`
	Url            string               `json:"url,omitempty"`
	DataType       int                  `json:"data_type,omitempty"`
	Protocol       int                  `json:"protocol,omitempty"`
	MaxRecords     int                  `json:"max_records,omitempty"`
	MaxSenders     int                  `json:"max_senders,omitempty"`
	MaxAttempts    int                  `json:"max_attempts,omitempty"`
	Timeout        string               `json:"timeout,omitempty"`
	HttpProxy      string               `json:"http_proxy,omitempty"`
	AuthType       int                  `json:"auth_type,omitempty"`
	Username       string               `json:"username,omitempty"`
	Password       string               `json:"password,omitempty"`
	Token          string               `json:"token,omitempty"`
	VeerifyPeer    int                  `json:"veerify_peer,omitempty"`
	VerifyHost     int                  `json:"verify_host,omitempty"`
	SSLCertFile    string               `json:"ssl_cert_file,omitempty"`
	SSLKeyFile     string               `json:"ssl_key_file,omitempty"`
	SSLKeyPassword string               `json:"ssl_key_password,omitempty"`
	Description    string               `json:"description,omitempty"`
	Status         int                  `json:"status,omitempty"`
	TagsEvaltype   int                  `json:"tags_evaltype,omitempty"`
	Tags           []ConnectorTagFilter `json:"tags,omitempty"`
}

type ConnectorObject struct {
	ConnectorID    string               `json:"connectorid,omitempty"`
	Name           string               `json:"name,omitempty"`
	Url            string               `json:"url,omitempty"`
	Protocol       int                  `json:"protocol,omitempty"`
	DataType       int                  `json:"data_type,omitempty"`
	MaxRecords     int                  `json:"max_records,omitempty"`
	MaxSenders     int                  `json:"max_senders,omitempty"`
	MaxAttempts    int                  `json:"max_attempts,omitempty"`
	Timeout        string               `json:"timeout,omitempty"`
	HttpProxy      string               `json:"http_proxy,omitempty"`
	AuthType       int                  `json:"auth_type,omitempty"`
	Username       string               `json:"username,omitempty"`
	Password       string               `json:"password,omitempty"`
	Token          string               `json:"token,omitempty"`
	VeerifyPeer    int                  `json:"veerify_peer,omitempty"`
	VerifyHost     int                  `json:"verify_host,omitempty"`
	SSLCertFile    string               `json:"ssl_cert_file,omitempty"`
	SSLKeyFile     string               `json:"ssl_key_file,omitempty"`
	SSLKeyPassword string               `json:"ssl_key_password,omitempty"`
	Description    string               `json:"description,omitempty"`
	Status         int                  `json:"status,omitempty"`
	TagsEvaltype   int                  `json:"tags_evaltype,omitempty"`
	Tags           []ConnectorTagFilter `json:"tags,omitempty"`
}

type ConnectorUpdate struct {
	ConnectorID    string               `json:"connectorid,omitempty"`
	Name           string               `json:"name,omitempty"`
	Url            string               `json:"url,omitempty"`
	Protocol       int                  `json:"protocol,omitempty"`
	DataType       int                  `json:"data_type,omitempty"`
	MaxRecords     int                  `json:"max_records,omitempty"`
	MaxSenders     int                  `json:"max_senders,omitempty"`
	MaxAttempts    int                  `json:"max_attempts,omitempty"`
	Timeout        string               `json:"timeout,omitempty"`
	HttpProxy      string               `json:"http_proxy,omitempty"`
	AuthType       int                  `json:"auth_type,omitempty"`
	Username       string               `json:"username,omitempty"`
	Password       string               `json:"password,omitempty"`
	Token          string               `json:"token,omitempty"`
	VeerifyPeer    int                  `json:"veerify_peer,omitempty"`
	VerifyHost     int                  `json:"verify_host,omitempty"`
	SSLCertFile    string               `json:"ssl_cert_file,omitempty"`
	SSLKeyFile     string               `json:"ssl_key_file,omitempty"`
	SSLKeyPassword string               `json:"ssl_key_password,omitempty"`
	Description    string               `json:"description,omitempty"`
	Status         int                  `json:"status,omitempty"`
	TagsEvaltype   int                  `json:"tags_evaltype,omitempty"`
	Tags           []ConnectorTagFilter `json:"tags,omitempty"`
}

type ConnectorTagFilter struct {
	Tag      string `json:"tag,omitempty"`
	Operator int    `json:"operator,omitempty"`
	Value    string `json:"value,omitempty"`
}

type ConnectorResponse struct {
	ConnectorID []string `json:"connectorid,omitempty"`
}
