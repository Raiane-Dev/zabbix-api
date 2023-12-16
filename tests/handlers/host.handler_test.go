package handlers_test

import (
	"testing"

	"github.com/Raiane-Dev/zabbix-api.git/entities"
	"github.com/Raiane-Dev/zabbix-api.git/handlers"
)

func TestHostGet(t *testing.T) {

	my_zbx := handlers.V6{}

	resp, err := my_zbx.HostGet(entities.HostGet{})
	if err != nil {
		t.Error(err)

	}

	t.Log(resp, err)
}
