package entities

type MediaTypeObject[T MediaTypeWebhook | MediaTypeScriptParameters] struct {
	MediaTypeID        string `json:"mediatypeid,omitempty"`
	Name               string `json:"name,omitempty"`
	Type               int    `json:"type,omitempty"`
	ExecPath           string `json:"execpath,omitempty"`
	GsmModem           string `json:"gsmmodem,omitempty"`
	Passwd             string `json:"passwd,omitempty"`
	SMTPEmail          string `json:"smtp_email,omitempty"`
	SMTPHelo           string `json:"smtp_helo,omitempty"`
	SMTPServer         string `json:"smtp_server,omitempty"`
	SMTPPort           int    `json:"smtp_port,omitempty"`
	SMTPSecurity       int    `json:"smtp_security,omitempty"`
	SMTPVerifyHost     int    `json:"smtp_verify_host,omitempty"`
	SMTPVerifyPeer     int    `json:"smtp_verify_peer,omitempty"`
	SMTPAuthentication int    `json:"smtp_authentication,omitempty"`
	Status             int    `json:"status,omitempty"`
	Username           string `json:"username,omitempty"`
	MaxSessions        int    `json:"maxsessions,omitempty"`
	MaxAttempts        int    `json:"maxattempts,omitempty"`
	AttemptInterval    string `json:"attempt_interval,omitempty"`
	ContentType        int    `json:"content_type,omitempty"`
	Script             string `json:"script,omitempty"`
	Timeout            string `json:"timeout,omitempty"`
	ProcessTags        int    `json:"process_tags,omitempty"`
	ShowEventMenu      int    `json:"show_event_menu,omitempty"`
	EventMenuUrl       string `json:"event_menu_url,omitempty"`
	EventMenuName      string `json:"event_menu_name,omitempty"`
	Parameters         []T    `json:"parameters,omitempty"`
	Description        string `json:"description,omitempty"`
}

type MediaTypeWebhook struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type MediaTypeScriptParameters struct {
	SortOrder int    `json:"sortorder,omitempty"`
	Value     string `json:"value,omitempty"`
}

type MediaTypeMessageTemplate struct {
	EventSource int    `json:"eventsource,omitempty"`
	Recovery    int    `json:"recovery,omitempty"`
	Subject     string `json:"subject,omitempty"`
	Message     string `json:"message,omitempty"`
}

type MediaTypeCreate[T MediaTypeWebhook | MediaTypeScriptParameters] struct {
	Name               string                     `json:"name,omitempty"`
	Type               int                        `json:"type,omitempty"`
	ExecPath           string                     `json:"execpath,omitempty"`
	GsmModem           string                     `json:"gsmmodem,omitempty"`
	Passwd             string                     `json:"passwd,omitempty"`
	SMTPEmail          string                     `json:"smtp_email,omitempty"`
	SMTPHelo           string                     `json:"smtp_helo,omitempty"`
	SMTPServer         string                     `json:"smtp_server,omitempty"`
	SMTPPort           int                        `json:"smtp_port,omitempty"`
	SMTPSecurity       int                        `json:"smtp_security,omitempty"`
	SMTPVerifyHost     int                        `json:"smtp_verify_host,omitempty"`
	SMTPVerifyPeer     int                        `json:"smtp_verify_peer,omitempty"`
	SMTPAuthentication int                        `json:"smtp_authentication,omitempty"`
	Status             int                        `json:"status,omitempty"`
	Username           string                     `json:"username,omitempty"`
	MaxSessions        int                        `json:"maxsessions,omitempty"`
	MaxAttempts        int                        `json:"maxattempts,omitempty"`
	AttemptInterval    string                     `json:"attempt_interval,omitempty"`
	ContentType        int                        `json:"content_type,omitempty"`
	Script             string                     `json:"script,omitempty"`
	Timeout            string                     `json:"timeout,omitempty"`
	ProcessTags        int                        `json:"process_tags,omitempty"`
	ShowEventMenu      int                        `json:"show_event_menu,omitempty"`
	EventMenuUrl       string                     `json:"event_menu_url,omitempty"`
	EventMenuName      string                     `json:"event_menu_name,omitempty"`
	Description        string                     `json:"description,omitempty"`
	Parameters         []T                        `json:"parameters,omitempty"`
	MessageTemplates   []MediaTypeMessageTemplate `json:"message_templates,omitempty"`
}

type MediaTypeUpdate[T MediaTypeWebhook | MediaTypeScriptParameters] struct {
	MediaTypeID        string                     `json:"mediatypeid,omitempty"`
	Name               string                     `json:"name,omitempty"`
	Type               int                        `json:"type,omitempty"`
	ExecPath           string                     `json:"execpath,omitempty"`
	GsmModem           string                     `json:"gsmmodem,omitempty"`
	Passwd             string                     `json:"passwd,omitempty"`
	SMTPEmail          string                     `json:"smtp_email,omitempty"`
	SMTPHelo           string                     `json:"smtp_helo,omitempty"`
	SMTPServer         string                     `json:"smtp_server,omitempty"`
	SMTPPort           int                        `json:"smtp_port,omitempty"`
	SMTPSecurity       int                        `json:"smtp_security,omitempty"`
	SMTPVerifyHost     int                        `json:"smtp_verify_host,omitempty"`
	SMTPVerifyPeer     int                        `json:"smtp_verify_peer,omitempty"`
	SMTPAuthentication int                        `json:"smtp_authentication,omitempty"`
	Status             int                        `json:"status,omitempty"`
	Username           string                     `json:"username,omitempty"`
	MaxSessions        int                        `json:"maxsessions,omitempty"`
	MaxAttempts        int                        `json:"maxattempts,omitempty"`
	AttemptInterval    string                     `json:"attempt_interval,omitempty"`
	ContentType        int                        `json:"content_type,omitempty"`
	Script             string                     `json:"script,omitempty"`
	Timeout            string                     `json:"timeout,omitempty"`
	ProcessTags        int                        `json:"process_tags,omitempty"`
	ShowEventMenu      int                        `json:"show_event_menu,omitempty"`
	EventMenuUrl       string                     `json:"event_menu_url,omitempty"`
	EventMenuName      string                     `json:"event_menu_name,omitempty"`
	Description        string                     `json:"description,omitempty"`
	Parameters         []T                        `json:"parameters,omitempty"`
	MessageTemplates   []MediaTypeMessageTemplate `json:"message_templates,omitempty"`
}

type MediaTypeGet struct {
	MediaTypeIDs           []string `json:"mediatypeids,omitempty"`
	MediaIDs               []string `json:"mediaids,omitempty"`
	UserIDs                []string `json:"userids,omitempty"`
	SelectMessageTemplates string   `json:"selectMessageTemplates,omitempty"`
	SelectUsers            string   `json:"selectUsers,omitempty"`
	ZabbixCommun
}
