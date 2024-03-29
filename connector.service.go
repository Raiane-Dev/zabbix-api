package zabbix_api

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func ConnectorGet(body entities.ConnectorGet) (response entities.Response[[]entities.ConnectorObject]) {

	req := entities.ZabbixPattern[entities.ConnectorGet]{
		JsonVersion: "2.0",
		Method:      "connector.get",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &response)
	return
}

func ConnectorCreate(body entities.ConnectorCreate) (response entities.Response[entities.ConnectorResponse], err error) {

	req := entities.ZabbixPattern[entities.ConnectorCreate]{
		JsonVersion: "2.0",
		Method:      "connector.create",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}

func ConnectorUpdate(body entities.ConnectorUpdate) (response entities.Response[entities.ConnectorResponse], err error) {

	req := entities.ZabbixPattern[entities.ConnectorUpdate]{
		JsonVersion: "2.0",
		Method:      "connector.update",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}

func ConnectorDelete(ids []string) (response entities.Response[entities.ConnectorResponse], err error) {

	req := entities.ZabbixPattern[[]string]{
		JsonVersion: "2.0",
		Method:      "connector.delete",
		Params:      ids,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}
