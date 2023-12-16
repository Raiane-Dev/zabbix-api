package handlers

import (
	"github.com/Raiane-Dev/zabbix-api.git/configs"
	"github.com/Raiane-Dev/zabbix-api.git/entities"
	"github.com/Raiane-Dev/zabbix-api.git/utils"
)

type V6 struct{}

func (V6) bus(body any) (response any) {

	request := &entities.IntegrationRPC{
		Host:          configs.ZABBIX_HOST,
		Authorization: configs.ZABBIX_TOKEN,
		Params:        body,
	}

	rest, err := utils.RemoteProcedure(request, &response)
	if err != nil {
		panic(err)
	}

	response = rest.Body()
	return
}
