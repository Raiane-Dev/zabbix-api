package zabbix_api

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func AuditlogGet(body entities.AuditlogGet) (response entities.Response[[]entities.AuditlogObject]) {

	req := entities.ZabbixPattern[entities.AuditlogGet]{
		JsonVersion: "2.0",
		Method:      "auditlog.get",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &response)
	return
}
