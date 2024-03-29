package zabbix_api

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func ConfigurationExport(body entities.ConfigurationExport) (response entities.Response[string]) {

	req := entities.ZabbixPattern[entities.ConfigurationExport]{
		JsonVersion: "2.0",
		Method:      "configuration.export",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &response)
	return
}

func ConfigurationImport(body entities.ConfigurationImport) (response entities.Response[bool], err error) {

	req := entities.ZabbixPattern[entities.ConfigurationImport]{
		JsonVersion: "2.0",
		Method:      "configuration.import",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}

func ConfigurationCompare(body entities.ConfigurationImportCompare) (response entities.Response[entities.ConfigurationCompareResponse], err error) {

	req := entities.ZabbixPattern[entities.ConfigurationImportCompare]{
		JsonVersion: "2.0",
		Method:      "configuration.importcompare",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}
