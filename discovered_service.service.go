package zabbix_api

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func DiscoveredServiceGet(body entities.DiscoveredServiceGet) (response entities.Response[[]entities.DiscoveredServiceObject]) {

	req := entities.ZabbixPattern[entities.DiscoveredServiceGet]{
		JsonVersion: "2.0",
		Method:      "dservice.get",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &response)
	return
}
