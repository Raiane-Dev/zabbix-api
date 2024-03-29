package zabbix_api

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func ActionGet(body entities.ActionGet) (response entities.Response[[]entities.ActionObject]) {

	req := entities.ZabbixPattern[entities.ActionGet]{
		JsonVersion: "2.0",
		Method:      "action.get",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &response)
	return
}

func ActionCreate(body entities.ActionCreate) (response entities.Response[entities.ActionResponse], err error) {

	req := entities.ZabbixPattern[entities.ActionCreate]{
		JsonVersion: "2.0",
		Method:      "action.create",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}

func ActionUpdate(body entities.ActionUpdate) (response entities.Response[entities.ActionResponse], err error) {

	req := entities.ZabbixPattern[entities.ActionUpdate]{
		JsonVersion: "2.0",
		Method:      "action.update",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}

func ActionDelete(ids []string) (response entities.Response[entities.ActionResponse], err error) {

	req := entities.ZabbixPattern[[]string]{
		JsonVersion: "2.0",
		Method:      "action.delete",
		Params:      ids,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}
