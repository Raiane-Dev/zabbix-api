package zabbix_api

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func DashboardGet(body entities.DashboardGet) (response entities.Response[[]entities.DashboardObject]) {

	req := entities.ZabbixPattern[entities.DashboardGet]{
		JsonVersion: "2.0",
		Method:      "dashboard.get",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &response)
	return
}

func DashboardCreate(body entities.DashboardCreate) (response entities.Response[entities.DashboardResponse], err error) {

	req := entities.ZabbixPattern[entities.DashboardCreate]{
		JsonVersion: "2.0",
		Method:      "dashboard.create",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}

func DashboardUpdate(body entities.DashboardUpdate) (response entities.Response[entities.DashboardResponse], err error) {

	req := entities.ZabbixPattern[entities.DashboardUpdate]{
		JsonVersion: "2.0",
		Method:      "dashboard.update",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}

func DashboardDelete(ids []string) (response entities.Response[entities.DashboardResponse], err error) {

	req := entities.ZabbixPattern[[]string]{
		JsonVersion: "2.0",
		Method:      "dashboard.delete",
		Params:      ids,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}
