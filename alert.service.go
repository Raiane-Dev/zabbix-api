package zabbix_api

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func AlertGet(body entities.AlertGet) (response entities.Response[[]entities.AlertObject]) {

	req := entities.ZabbixPattern[entities.AlertGet]{
		JsonVersion: "2.0",
		Method:      "alert.get",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &response)
	return
}
