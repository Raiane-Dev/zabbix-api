package entities

type SlaObject struct {
	SlaID         string  `json:"slaid,omitempty"`
	Name          string  `json:"name,omitempty"`
	Period        int     `json:"period,omitempty"`
	Slo           float32 `json:"slo,omitempty"`
	EffectiveDate int     `json:"effective_date,omitempty"`
	Timezone      string  `json:"timezone,omitempty"`
	Status        int     `json:"status,omitempty"`
	Description   string  `json:"description,omitempty"`
}

type SlaSchedule struct {
	PeriodFrom int `json:"period_from,omitempty"`
	PeriodTo   int `json:"period_to,omitempty"`
}

type SlaExcludedDowntime struct {
	Name       string `json:"name,omitempty"`
	PeriodFrom int    `json:"period_from,omitempty"`
	PeriodTo   int    `json:"period_to,omitempty"`
}

type SlaServiceTag struct {
	Tag      string `json:"tag,omitempty"`
	Operator int    `json:"operator,omitempty"`
	Value    string `json:"value,omitempty"`
}

type SlaCreate struct {
	Name              string                `json:"name,omitempty"`
	Period            int                   `json:"period,omitempty"`
	Slo               float32               `json:"slo,omitempty"`
	Timezone          string                `json:"timezone,omitempty"`
	Status            int                   `json:"status,omitempty"`
	Description       string                `json:"description,omitempty"`
	EffectiveDate     int                   `json:"effective_date,omitempty"`
	ServiceTags       []SlaServiceTag       `json:"service_tags,omitempty"`
	Schedule          []SlaSchedule         `json:"schedule,omitempty"`
	ExcludedDowntimes []SlaExcludedDowntime `json:"excluded_downtimes,omitempty"`
}

type SlaUpdate struct {
	SlaID             string                `json:"slaid,omitempty"`
	Name              string                `json:"name,omitempty"`
	Period            int                   `json:"period,omitempty"`
	Slo               float32               `json:"slo,omitempty"`
	Timezone          string                `json:"timezone,omitempty"`
	Status            int                   `json:"status,omitempty"`
	Description       string                `json:"description,omitempty"`
	EffectiveDate     int                   `json:"effective_date,omitempty"`
	ServiceTags       []SlaServiceTag       `json:"service_tags,omitempty"`
	Schedule          []SlaSchedule         `json:"schedule,omitempty"`
	ExcludedDowntimes []SlaExcludedDowntime `json:"excluded_downtimes,omitempty"`
}

type SlaResponse struct {
	SlaIDs []string `json:"slaids,omitempty"`
}
type SlaGet struct {
	SlaIDs                  []string `json:"slaids,omitempty"`
	ServiceIDs              []string `json:"serviceids,omitempty"`
	SelectSchedule          string   `json:"selectSchedule,omitempty"`
	SelectExcludedDowntimes string   `json:"selectExcludedDowntimes,omitempty"`
	SelectServiceTags       string   `json:"selectServiceTags,omitempty"`
	ZabbixCommun
}

type SlaSliGet struct {
	SlaID      string   `json:"slaid,omitempty"`
	PeriodFrom int      `json:"period_from,omitempty"`
	PeriodTo   int      `json:"period_to,omitempty"`
	Periods    []int    `json:"periods,omitempty"`
	ServiceIDs []string `json:"serviceids,omitempty"`
}

type SlaSliObject struct {
	Periods    []SlaSchedule    `json:"periods,omitempty"`
	ServiceIDs []int            `json:"serviceids,omitempty"`
	Sli        []map[string]any `json:"sli,omitempty"`
}
