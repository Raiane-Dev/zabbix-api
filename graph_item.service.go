package zabbix_api

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func GraphItemGet(body entities.GraphItemGet) (response entities.Response[[]entities.GraphItemObject]) {

	req := entities.ZabbixPattern[entities.GraphItemGet]{
		JsonVersion: "2.0",
		Method:      "graphitem.get",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &response)
	return
}
