#!/bin/bash


content="package entities

type $2Object struct {
}

type $2Create struct {
}

type $2Update struct {
}

type $2Get struct {

}
"

file="entities/$1.entity.go"

echo "$content" > "$file"

echo "file \"$file\" created with success!"
