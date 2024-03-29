package zabbix_api

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func AutoregistrationGet(body entities.AutoregistrationGet) (response entities.Response[[]entities.AutoregistrationObject]) {

	req := entities.ZabbixPattern[entities.AutoregistrationGet]{
		JsonVersion: "2.0",
		Method:      "autoregistration.get",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &response)
	return
}

func AutoregistrationUpdate(body entities.AutoregistrationUpdate) (response entities.Response[bool], err error) {

	req := entities.ZabbixPattern[entities.AutoregistrationUpdate]{
		JsonVersion: "2.0",
		Method:      "autoregistration.update",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}
