package zabbix_api

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func DiscoveryCheckGet(body entities.DiscoveryCheckGet) (response entities.Response[[]entities.DiscoveryCheckObject]) {

	req := entities.ZabbixPattern[entities.DiscoveryCheckGet]{
		JsonVersion: "2.0",
		Method:      "dcheck.get",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &response)
	return
}
