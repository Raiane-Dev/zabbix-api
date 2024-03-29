package zabbix_api

import (
	"github.com/Raiane-Dev/zabbix-api.git/entities"
	"github.com/Raiane-Dev/zabbix-api.git/utils"
)

func bus(body any) (response any) {

	request := &entities.IntegrationRPC{
		Host:          "",
		Authorization: "",
		Params:        body,
	}

	rest, err := utils.RemoteProcedure(request, &response)
	if rest.IsError() || err != nil {
		return nil
	}

	response = rest.Body()
	return
}
