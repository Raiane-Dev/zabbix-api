package entities

type TaskObject[T TaskRequestDiagnosticInformation | TaskRequestExecuteNow | TaskRequestRefreshProxyConf] struct {
	TaskID      string         `json:"taskid,omitempty"`
	Type        int            `json:"type,omitempty"`
	Status      int            `json:"status,omitempty"`
	Clock       string         `json:"clock,omitempty"`
	TTL         int            `json:"ttl,omitempty"`
	ProxyHostID string         `json:"proxy_hostid,omitempty"`
	Request     T              `json:"request,omitempty"`
	Result      map[string]any `json:"result,omitempty"`
}

type TaskRequestExecuteNow struct {
	ItemID string `json:"itemid,omitempty"`
}

type TaskRequestRefreshProxyConf struct {
	ProxyHostIDs []string `json:"proxy_hostids,omitempty"`
}

type TaskRequestDiagnosticInformation struct {
	HistoryCache  map[string]any `json:"historycache,omitempty"`
	ValueCache    map[string]any `json:"valuecache,omitempty"`
	Preprocessing map[string]any `json:"preprocessing,omitempty"`
	Alerting      map[string]any `json:"alerting,omitempty"`
	LLD           map[string]any `json:"lld,omitempty"`
}

type TaskRequestStatistic struct {
	Stats string            `json:"stats,omitempty"`
	Top   map[string]string `json:"top,omitempty"`
}

type TaskCreate struct {
}

type TaskUpdate struct {
}

type TaskGet struct {
}
