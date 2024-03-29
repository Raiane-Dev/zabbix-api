package entities

type HousekeepingObject struct {
	HKEventsMode            int    `json:"hk_events_mode,omitempty"`
	HKEventsTrigger         string `json:"hk_events_trigger,omitempty"`
	HKEventsService         string `json:"hk_events_service,omitempty"`
	HKEventsInternal        string `json:"hk_events_internal,omitempty"`
	HKEventsDiscovery       string `json:"hk_events_discovery,omitempty"`
	HKEventsAutoreg         string `json:"hk_events_autoreg,omitempty"`
	HKServicesMode          int    `json:"hk_service_smode,omitempty"`
	HKServices              string `json:"hk_services,omitempty"`
	HKAuditMode             int    `json:"hk_audit_mode,omitempty"`
	HKAudit                 string `json:"hk_audit,omitempty"`
	HKSessionsMode          int    `json:"hk_sessions_mode,omitempty"`
	HKSessions              string `json:"hk_sessions,omitempty"`
	HKHistoryMode           int    `json:"hk_history_mode,omitempty"`
	HkHistoryGlobal         int    `json:"hk_history_global,omitempty"`
	HKHistory               string `json:"hk_history,omitempty"`
	HKTrendsMode            int    `json:"hk_trends_mode,omitempty"`
	HKTrendsGlobal          int    `json:"hk_trends_global,omitempty"`
	DBExtension             string `json:"db_extension,omitempty"`
	CompressionAvailability int    `json:"compression_availability,omitempty"`
	CompressionStatus       int    `json:"compression_status,omitempty"`
	CompressOlder           string `json:"compress_older,omitempty"`
}

type HousekeepingUpdate struct {
	HKEventsMode            int    `json:"hk_events_mode,omitempty"`
	HKEventsTrigger         string `json:"hk_events_trigger,omitempty"`
	HKEventsService         string `json:"hk_events_service,omitempty"`
	HKEventsInternal        string `json:"hk_events_internal,omitempty"`
	HKEventsDiscovery       string `json:"hk_events_discovery,omitempty"`
	HKEventsAutoreg         string `json:"hk_events_autoreg,omitempty"`
	HKServicesMode          int    `json:"hk_service_smode,omitempty"`
	HKServices              string `json:"hk_services,omitempty"`
	HKAuditMode             int    `json:"hk_audit_mode,omitempty"`
	HKAudit                 string `json:"hk_audit,omitempty"`
	HKSessionsMode          int    `json:"hk_sessions_mode,omitempty"`
	HKSessions              string `json:"hk_sessions,omitempty"`
	HKHistoryMode           int    `json:"hk_history_mode,omitempty"`
	HkHistoryGlobal         int    `json:"hk_history_global,omitempty"`
	HKHistory               string `json:"hk_history,omitempty"`
	HKTrendsMode            int    `json:"hk_trends_mode,omitempty"`
	HKTrendsGlobal          int    `json:"hk_trends_global,omitempty"`
	DBExtension             string `json:"db_extension,omitempty"`
	CompressionAvailability int    `json:"compression_availability,omitempty"`
	CompressionStatus       int    `json:"compression_status,omitempty"`
	CompressOlder           string `json:"compress_older,omitempty"`
}

type HousekeepingGet struct {
	Output string `json:"output,omitempty"`
}
