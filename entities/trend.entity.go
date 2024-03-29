package entities

import "time"

type TrendObject[T int | float64] struct {
	Clock    time.Time `json:"clock,omitempty"`
	ItemID   int       `json:"itemid,omitempty"`
	Num      int       `json:"num,omitempty"`
	ValueMin T         `json:"value_min,omitempty"`
	ValueAvg T         `json:"value_avg,omitempty"`
	ValueMax T         `json:"value_max,omitempty"`
}

type TrendGet struct {
	ItemIDs     []string  `json:"itemids,omitempty"`
	TimeFrom    time.Time `json:"time_from,omitempty"`
	TimeTill    time.Time `json:"time_till,omitempty"`
	CountOutput bool      `json:"countOutput,omitempty"`
	Limit       int       `json:"limit,omitempty"`
	Output      string    `json:"output,omitempty"`
}
