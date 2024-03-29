package zabbix_api

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func CorrelationGet(body entities.CorrelationGet) (response entities.Response[[]entities.CorrelationObject]) {

	req := entities.ZabbixPattern[entities.CorrelationGet]{
		JsonVersion: "2.0",
		Method:      "correlation.get",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &response)
	return
}

func CorrelationCreate(body entities.CorrelationCreate) (response entities.Response[entities.CorrelationResponse], err error) {

	req := entities.ZabbixPattern[entities.CorrelationCreate]{
		JsonVersion: "2.0",
		Method:      "correlation.create",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}

func CorrelationUpdate(body entities.CorrelationUpdate) (response entities.Response[entities.CorrelationResponse], err error) {

	req := entities.ZabbixPattern[entities.CorrelationUpdate]{
		JsonVersion: "2.0",
		Method:      "correlation.update",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}

func CorrelationDelete(ids []string) (response entities.Response[entities.CorrelationResponse], err error) {

	req := entities.ZabbixPattern[[]string]{
		JsonVersion: "2.0",
		Method:      "correlation.update",
		Params:      ids,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}
