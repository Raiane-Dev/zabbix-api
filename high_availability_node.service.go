package zabbix_api

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func HighAvailabilityNodeGet(body entities.HighAvailabilityNodeGet) (response entities.Response[[]entities.HighAvailabilityNodeObject]) {

	req := entities.ZabbixPattern[entities.HighAvailabilityNodeGet]{
		JsonVersion: "2.0",
		Method:      "hanode.get",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &response)
	return
}
