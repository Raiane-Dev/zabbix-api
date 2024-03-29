package zabbix_api

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func DiscoveredHostGet(body entities.DiscoveredHostGet) (response entities.Response[[]entities.DiscoveredHostObject]) {

	req := entities.ZabbixPattern[entities.DiscoveredHostGet]{
		JsonVersion: "2.0",
		Method:      "dhost.get",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &response)
	return
}
