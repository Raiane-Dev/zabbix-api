package entities

type ProblemObject struct {
	EventID       string                `json:"eventid,omitempty"`
	Source        int                   `json:"source,omitempty"`
	Object        int                   `json:"object,omitempty"`
	ObjectID      string                `json:"objectid,omitempty"`
	Clock         string                `json:"clock,omitempty"`
	NS            int                   `json:"ns,omitempty"`
	REventID      string                `json:"r_eventid,omitempty"`
	RClock        string                `json:"r_clock,omitempty"`
	RNS           int                   `json:"r_ns,omitempty"`
	CauseEventID  string                `json:"cause_eventid,omitempty"`
	CorrelationID string                `json:"correlationid,omitempty"`
	UserID        string                `json:"userid,omitempty"`
	Name          string                `json:"name,omitempty"`
	Acknowledge   int                   `json:"acknowledge,omitempty"`
	Severity      int                   `json:"severity,omitempty"`
	Supressed     int                   `json:"supressed,omitempty"`
	Opdata        string                `json:"opdata,omitempty"`
	Urls          []ProblemMediaTypeURL `json:"urls,omitempty"`
}

type ProblemMediaTypeURL struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type ProblemGet struct {
	EventIDs             []string   `json:"eventids,omitempty"`
	GroupIDs             []string   `json:"groupids,omitempty"`
	HostIDs              []string   `json:"hostids,omitempty"`
	ObjectIDs            []string   `json:"objectids,omitempty"`
	Source               int        `json:"source,omitempty"`
	Object               int        `json:"object,omitempty"`
	Acknowledge          bool       `json:"acknowledge,omitempty"`
	Supressed            bool       `json:"supressed,omitempty"`
	Symptom              bool       `json:"symptom,omitempty"`
	Severities           []int      `json:"severities,omitempty"`
	Evaltype             int        `json:"evaltype,omitempty"`
	Tags                 ProblemTag `json:"tags,omitempty"`
	Recent               bool       `json:"recent,omitempty"`
	EventIDFrom          string     `json:"eventidfrom,omitempty"`
	EventIDTill          string     `json:"eventidtill,omitempty"`
	TimeFrom             string     `json:"timefrom,omitempty"`
	TimeFill             string     `json:"timefill,omitempty"`
	SelectAcknowledges   string     `json:"selectAcknowledges,omitempty"`
	SelectTags           string     `json:"selectTags,omitempty"`
	SelectSupressionData string     `json:"selectSupressionData,omitempty"`
	ZabbixCommun
}

type ProblemTag struct {
	Tag      string `json:"tag,omitempty"`
	Value    string `json:"value,omitempty"`
	Operator string `json:"operator,omitempty"`
}
