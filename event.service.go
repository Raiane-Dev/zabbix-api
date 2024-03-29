package zabbix_api

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func EventGet(body entities.EventGet) (response entities.Response[[]entities.EventObject]) {

	req := entities.ZabbixPattern[entities.EventGet]{
		JsonVersion: "2.0",
		Method:      "event.get",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &response)
	return
}

func EventAcknowledge(body entities.EventAcknowledge) (response entities.Response[entities.EventResponse], err error) {

	req := entities.ZabbixPattern[entities.EventAcknowledge]{
		JsonVersion: "2.0",
		Method:      "event.acknowledge",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}
