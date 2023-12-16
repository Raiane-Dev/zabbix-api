package utils

import (
	"crypto/tls"
	"encoding/json"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
	"github.com/go-resty/resty/v2"
)

func RemoteProcedure(api *entities.IntegrationRPC, result *interface{}) (resp *resty.Response, err error) {
	client := resty.New()
	body, err := json.Marshal(api.Params)
	if err != nil {
		return
	}

	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: true,
	})
	resp, err = client.R().
		SetHeader("Content-Type", "application/json-rpc").
		SetBody(body).
		SetResult(&result).
		SetAuthToken(api.Authorization).
		Post(api.Host)

	return
}
