#!/bin/bash


content="package zabbix_api

import (
	\"encoding/json\"

	\"github.com/Raiane-Dev/zabbix-api.git/entities\"
)

func $2Get(body entities.$2Get) (response entities.Response[[]entities.$2Object]) {

	req := entities.ZabbixPattern[entities.$2Get]{
		JsonVersion: \"2.0\",
		Method:      \"$1.get\",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	json.Unmarshal(req_res.([]uint8), &response)
	return
}

func $2Create(body entities.$2Create) (response entities.Response[entities.$2Response], err error) {

	req := entities.ZabbixPattern[entities.$2Create]{
		JsonVersion: \"2.0\",
		Method:      \"$1.create\",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}

func $2Update(body entities.$2Update) (response entities.Response[entities.$2Response], err error) {

	req := entities.ZabbixPattern[entities.$2Update]{
		JsonVersion: \"2.0\",
		Method:      \"$1.update\",
		Params:      body,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}

func $2Delete(ids []string) (response entities.Response[entities.$2Response], err error) {

	req := entities.ZabbixPattern[[]string]{
		JsonVersion: \"2.0\",
		Method:      \"$1.delete\",
		Params:      ids,
	}
	req_res := bus(req)
	if req_res == nil {
		return
	}

	err = json.Unmarshal(req_res.([]uint8), &response)
	return
}


"

file="../$1.service.go"

echo "$content" > "$file"

echo "file \"$file\" created with success!"
