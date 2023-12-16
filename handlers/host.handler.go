package handlers

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/configs"
	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func (v V6) HostGet(body entities.HostGet) (response []entities.HostOutput, _err *entities.Error) {
	var extract entities.Body[entities.HostOutput]

	req := entities.ZabbixPattern[entities.HostGet]{
		JsonVersion:    "2.0",
		Authentication: configs.ZABBIX_TOKEN,
		Method:         "host.get",
		Params:         body,
	}

	req_res := v.bus(req)
	json.Unmarshal(req_res.([]uint8), &extract.Return)

	response = extract.Return.Result
	if extract.Return.Error.Code != 0 {
		_err = &extract.Return.Error
	}

	return

}
