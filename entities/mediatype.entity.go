package entities

type MediaTypeObject[T MediaTypeWebhook | MediaTypeScriptParameters] struct {
	MediaTypeID        string `json:"mediatypeid"`
	Name               string `json:"name"`
	Type               int    `json:"type"`
	ExecPath           string `json:"execpath"`
	GsmModem           string `json:"gsmmodem"`
	Passwd             string `json:"passwd"`
	SMTPEmail          string `json:"smtp_email"`
	SMTPHelo           string `json:"smtp_helo"`
	SMTPServer         string `json:"smtp_server"`
	SMTPPort           int    `json:"smtp_port"`
	SMTPSecurity       int    `json:"smtp_security"`
	SMTPVerifyHost     int    `json:"smtp_verify_host"`
	SMTPVerifyPeer     int    `json:"smtp_verify_peer"`
	SMTPAuthentication int    `json:"smtp_authentication"`
	Status             int    `json:"status"`
	Username           string `json:"username"`
	MaxSessions        int    `json:"maxsessions"`
	MaxAttempts        int    `json:"maxattempts"`
	AttemptInterval    string `json:"attempt_interval"`
	ContentType        int    `json:"content_type"`
	Script             string `json:"script"`
	Timeout            string `json:"timeout"`
	ProcessTags        int    `json:"process_tags"`
	ShowEventMenu      int    `json:"show_event_menu"`
	EventMenuUrl       string `json:"event_menu_url"`
	EventMenuName      string `json:"event_menu_name"`
	Parameters         []T    `json:"parameters"`
	Description        string `json:"description"`
}

type MediaTypeWebhook struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type MediaTypeScriptParameters struct {
	SortOrder int    `json:"sortorder"`
	Value     string `json:"value"`
}

type MediaTypeMessageTemplate struct {
	EventSource int    `json:"eventsource"`
	Recovery    int    `json:"recovery"`
	Subject     string `json:"subject"`
	Message     string `json:"message"`
}

type MediaTypeCreate[T MediaTypeWebhook | MediaTypeScriptParameters] struct {
	Name               string                     `json:"name"`
	Type               int                        `json:"type"`
	ExecPath           string                     `json:"execpath"`
	GsmModem           string                     `json:"gsmmodem"`
	Passwd             string                     `json:"passwd"`
	SMTPEmail          string                     `json:"smtp_email"`
	SMTPHelo           string                     `json:"smtp_helo"`
	SMTPServer         string                     `json:"smtp_server"`
	SMTPPort           int                        `json:"smtp_port"`
	SMTPSecurity       int                        `json:"smtp_security"`
	SMTPVerifyHost     int                        `json:"smtp_verify_host"`
	SMTPVerifyPeer     int                        `json:"smtp_verify_peer"`
	SMTPAuthentication int                        `json:"smtp_authentication"`
	Status             int                        `json:"status"`
	Username           string                     `json:"username"`
	MaxSessions        int                        `json:"maxsessions"`
	MaxAttempts        int                        `json:"maxattempts"`
	AttemptInterval    string                     `json:"attempt_interval"`
	ContentType        int                        `json:"content_type"`
	Script             string                     `json:"script"`
	Timeout            string                     `json:"timeout"`
	ProcessTags        int                        `json:"process_tags"`
	ShowEventMenu      int                        `json:"show_event_menu"`
	EventMenuUrl       string                     `json:"event_menu_url"`
	EventMenuName      string                     `json:"event_menu_name"`
	Description        string                     `json:"description"`
	Parameters         []T                        `json:"parameters"`
	MessageTemplates   []MediaTypeMessageTemplate `json:"message_templates"`
}

type MediaTypeUpdate[T MediaTypeWebhook | MediaTypeScriptParameters] struct {
	MediaTypeID        string                     `json:"mediatypeid"`
	Name               string                     `json:"name"`
	Type               int                        `json:"type"`
	ExecPath           string                     `json:"execpath"`
	GsmModem           string                     `json:"gsmmodem"`
	Passwd             string                     `json:"passwd"`
	SMTPEmail          string                     `json:"smtp_email"`
	SMTPHelo           string                     `json:"smtp_helo"`
	SMTPServer         string                     `json:"smtp_server"`
	SMTPPort           int                        `json:"smtp_port"`
	SMTPSecurity       int                        `json:"smtp_security"`
	SMTPVerifyHost     int                        `json:"smtp_verify_host"`
	SMTPVerifyPeer     int                        `json:"smtp_verify_peer"`
	SMTPAuthentication int                        `json:"smtp_authentication"`
	Status             int                        `json:"status"`
	Username           string                     `json:"username"`
	MaxSessions        int                        `json:"maxsessions"`
	MaxAttempts        int                        `json:"maxattempts"`
	AttemptInterval    string                     `json:"attempt_interval"`
	ContentType        int                        `json:"content_type"`
	Script             string                     `json:"script"`
	Timeout            string                     `json:"timeout"`
	ProcessTags        int                        `json:"process_tags"`
	ShowEventMenu      int                        `json:"show_event_menu"`
	EventMenuUrl       string                     `json:"event_menu_url"`
	EventMenuName      string                     `json:"event_menu_name"`
	Description        string                     `json:"description"`
	Parameters         []T                        `json:"parameters"`
	MessageTemplates   []MediaTypeMessageTemplate `json:"message_templates"`
}

type MediaTypeGet struct {
	MediaTypeIDs           []string `json:"mediatypeids"`
	MediaIDs               []string `json:"mediaids"`
	UserIDs                []string `json:"userids"`
	SelectMessageTemplates string   `json:"selectMessageTemplates"`
	SelectUsers            string   `json:"selectUsers"`
	ZabbixCommun
}
