package zabbix_api

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func GraphPrototypeGet(body entities.GraphPrototypeGet) (response entities.Response[[]entities.GraphPrototypeObject]) {

	req := entities.ZabbixPattern[entities.GraphPrototypeGet]{
		JsonVersion: "2.0",
		Method:      "graphprototype.get",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &response)
	return
}

func GraphPrototypeCreate(body entities.GraphPrototypeCreate) (response entities.Response[entities.GraphPrototypeResponse], err error) {

	req := entities.ZabbixPattern[entities.GraphPrototypeCreate]{
		JsonVersion: "2.0",
		Method:      "graphprototype.create",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}

func GraphPrototypeUpdate(body entities.GraphPrototypeUpdate) (response entities.Response[entities.GraphPrototypeResponse], err error) {

	req := entities.ZabbixPattern[entities.GraphPrototypeUpdate]{
		JsonVersion: "2.0",
		Method:      "graphprototype.update",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}

func GraphPrototypeDelete(ids []string) (response entities.Response[entities.GraphPrototypeResponse], err error) {

	req := entities.ZabbixPattern[[]string]{
		JsonVersion: "2.0",
		Method:      "graphprototype.delete",
		Params:      ids,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}
