package zabbix_api

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func DiscoveryRuleGet(body entities.DiscoveryRuleGet) (response entities.Response[[]entities.DiscoveryRuleObject]) {

	req := entities.ZabbixPattern[entities.DiscoveryRuleGet]{
		JsonVersion: "2.0",
		Method:      "drule.get",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &response)
	return
}

func DiscoveryRuleCreate(body entities.DiscoveryRuleCreate) (response entities.Response[entities.DiscoveryRuleResponse], err error) {

	req := entities.ZabbixPattern[entities.DiscoveryRuleCreate]{
		JsonVersion: "2.0",
		Method:      "drule.create",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}

func DiscoveryRuleUpdate(body entities.DiscoveryRuleUpdate) (response entities.Response[entities.DiscoveryRuleResponse], err error) {

	req := entities.ZabbixPattern[entities.DiscoveryRuleUpdate]{
		JsonVersion: "2.0",
		Method:      "drule.update",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}

func DiscoveryRuleDelete(ids []string) (response entities.Response[entities.DiscoveryRuleResponse], err error) {

	req := entities.ZabbixPattern[[]string]{
		JsonVersion: "2.0",
		Method:      "drule.delete",
		Params:      ids,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}
