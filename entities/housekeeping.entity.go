package entities

type HousekeepingObject struct {
	HKEventsMode            int    `json:"hk_events_mode"`
	HKEventsTrigger         string `json:"hk_events_trigger"`
	HKEventsService         string `json:"hk_events_service"`
	HKEventsInternal        string `json:"hk_events_internal"`
	HKEventsDiscovery       string `json:"hk_events_discovery"`
	HKEventsAutoreg         string `json:"hk_events_autoreg"`
	HKServicesMode          int    `json:"hk_service_smode"`
	HKServices              string `json:"hk_services"`
	HKAuditMode             int    `json:"hk_audit_mode"`
	HKAudit                 string `json:"hk_audit"`
	HKSessionsMode          int    `json:"hk_sessions_mode"`
	HKSessions              string `json:"hk_sessions"`
	HKHistoryMode           int    `json:"hk_history_mode"`
	HkHistoryGlobal         int    `json:"hk_history_global"`
	HKHistory               string `json:"hk_history"`
	HKTrendsMode            int    `json:"hk_trends_mode"`
	HKTrendsGlobal          int    `json:"hk_trends_global"`
	DBExtension             string `json:"db_extension"`
	CompressionAvailability int    `json:"compression_availability"`
	CompressionStatus       int    `json:"compression_status"`
	CompressOlder           string `json:"compress_older"`
}

type HousekeepingUpdate struct {
	HKEventsMode            int    `json:"hk_events_mode"`
	HKEventsTrigger         string `json:"hk_events_trigger"`
	HKEventsService         string `json:"hk_events_service"`
	HKEventsInternal        string `json:"hk_events_internal"`
	HKEventsDiscovery       string `json:"hk_events_discovery"`
	HKEventsAutoreg         string `json:"hk_events_autoreg"`
	HKServicesMode          int    `json:"hk_service_smode"`
	HKServices              string `json:"hk_services"`
	HKAuditMode             int    `json:"hk_audit_mode"`
	HKAudit                 string `json:"hk_audit"`
	HKSessionsMode          int    `json:"hk_sessions_mode"`
	HKSessions              string `json:"hk_sessions"`
	HKHistoryMode           int    `json:"hk_history_mode"`
	HkHistoryGlobal         int    `json:"hk_history_global"`
	HKHistory               string `json:"hk_history"`
	HKTrendsMode            int    `json:"hk_trends_mode"`
	HKTrendsGlobal          int    `json:"hk_trends_global"`
	DBExtension             string `json:"db_extension"`
	CompressionAvailability int    `json:"compression_availability"`
	CompressionStatus       int    `json:"compression_status"`
	CompressOlder           string `json:"compress_older"`
}

type HousekeepingGet struct {
	Output string `json:"output"`
}
