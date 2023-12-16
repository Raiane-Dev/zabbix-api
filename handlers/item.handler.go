package handlers

import (
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/configs"
	"github.com/Raiane-Dev/zabbix-api.git/entities"
)

func (v V6) GetItem(body entities.ItemGet) (response []entities.ItemOutput, err error) {
	var extract entities.Body[entities.ItemOutput]

	req := entities.ZabbixPattern[entities.ItemGet]{
		JsonVersion:    "2.0",
		Authentication: configs.ZABBIX_TOKEN,
		Method:         "item.get",
		Params:         body,
	}

	req_res := v.bus(req)
	if req_res == nil || err != nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &extract.Return)
	response = extract.Return.Result

	return
}

func (v V6) CreateItem(body entities.ItemCreate) (response []entities.ItemOutput, _err *entities.Error) {
	var extract entities.Body[entities.ItemOutput]

	req := entities.ZabbixPattern[entities.ItemCreate]{
		JsonVersion:    "2.0",
		Authentication: configs.ZABBIX_TOKEN,
		Method:         "item.create",
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
