package zabbix_api

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func ItemGet(body entities.ItemGet) (response entities.Response[[]entities.ItemObject]) {

	req := entities.ZabbixPattern[entities.ItemGet]{
		JsonVersion: "2.0",
		Method:      "item.get",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &response)
	return
}

func ItemCreate(body entities.ItemCreate) (response entities.Response[entities.ItemResponse], err error) {

	req := entities.ZabbixPattern[entities.ItemCreate]{
		JsonVersion: "2.0",
		Method:      "item.create",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}

func ItemUpdate(body entities.ItemUpdate) (response entities.Response[entities.ItemResponse], err error) {

	req := entities.ZabbixPattern[entities.ItemUpdate]{
		JsonVersion: "2.0",
		Method:      "item.update",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}
