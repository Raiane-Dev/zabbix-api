package zabbix_api

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func AuthenticationGet(body entities.AuthenticationGet) (response entities.Response[entities.AuthenticationObject]) {

	req := entities.ZabbixPattern[entities.AuthenticationGet]{
		JsonVersion: "2.0",
		Method:      "authentication.get",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &response)
	return
}

func AuthenticationUpdate(body entities.AuthenticationUpdate) (response entities.Response[entities.AuthenticationResponse], err error) {

	req := entities.ZabbixPattern[entities.AuthenticationUpdate]{
		JsonVersion: "2.0",
		Method:      "authentication.update",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}
