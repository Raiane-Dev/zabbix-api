package entities

type ProblemObject struct {
	EventID       string                `json:"eventid"`
	Source        int                   `json:"source"`
	Object        int                   `json:"object"`
	ObjectID      string                `json:"objectid"`
	Clock         string                `json:"clock"`
	NS            int                   `json:"ns"`
	REventID      string                `json:"r_eventid"`
	RClock        string                `json:"r_clock"`
	RNS           int                   `json:"r_ns"`
	CauseEventID  string                `json:"cause_eventid"`
	CorrelationID string                `json:"correlationid"`
	UserID        string                `json:"userid"`
	Name          string                `json:"name"`
	Acknowledge   int                   `json:"acknowledge"`
	Severity      int                   `json:"severity"`
	Supressed     int                   `json:"supressed"`
	Opdata        string                `json:"opdata"`
	Urls          []ProblemMediaTypeURL `json:"urls"`
}

type ProblemMediaTypeURL struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ProblemGet struct {
	EventIDs             []string   `json:"eventids"`
	GroupIDs             []string   `json:"groupids"`
	HostIDs              []string   `json:"hostids"`
	ObjectIDs            []string   `json:"objectids"`
	Source               int        `json:"source"`
	Object               int        `json:"object"`
	Acknowledge          bool       `json:"acknowledge"`
	Supressed            bool       `json:"supressed"`
	Symptom              bool       `json:"symptom"`
	Severities           []int      `json:"severities"`
	Evaltype             int        `json:"evaltype"`
	Tags                 ProblemTag `json:"tags"`
	Recent               bool       `json:"recent"`
	EventIDFrom          string     `json:"eventidfrom"`
	EventIDTill          string     `json:"eventidtill"`
	TimeFrom             string     `json:"timefrom"`
	TimeFill             string     `json:"timefill"`
	SelectAcknowledges   string     `json:"selectAcknowledges"`
	SelectTags           string     `json:"selectTags"`
	SelectSupressionData string     `json:"selectSupressionData"`
	ZabbixCommun
}

type ProblemTag struct {
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Operator string `json:"operator"`
}
