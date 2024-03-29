package zabbix_api

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func GraphGet(body entities.GraphGet) (response entities.Response[[]entities.GraphObject]) {

	req := entities.ZabbixPattern[entities.GraphGet]{
		JsonVersion: "2.0",
		Method:      "graph.get",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &response)
	return
}

func GraphCreate(body entities.GraphCreate) (response entities.Response[entities.GraphResponse], err error) {

	req := entities.ZabbixPattern[entities.GraphCreate]{
		JsonVersion: "2.0",
		Method:      "graph.create",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}

func GraphUpdate(body entities.GraphUpdate) (response entities.Response[entities.GraphResponse], err error) {

	req := entities.ZabbixPattern[entities.GraphUpdate]{
		JsonVersion: "2.0",
		Method:      "graph.update",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}

func GraphDelete(ids []string) (response entities.Response[entities.GraphResponse], err error) {

	req := entities.ZabbixPattern[[]string]{
		JsonVersion: "2.0",
		Method:      "graph.delete",
		Params:      ids,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}



