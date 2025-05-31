#!/usr/bin/env bash

awslocal dynamodb create-table \
	--table-name 'master-local' \
	--attribute-definitions '[
		{"AttributeName": "pk", "AttributeType": "S"},
		{"AttributeName": "sk", "AttributeType": "S"}
	]' \
	--key-schema '[
		{"AttributeName": "pk", "KeyType": "HASH"},
		{"AttributeName": "sk", "KeyType": "RANGE"}
	]' \
	--provisioned-throughput '{
		"ReadCapacityUnits": 5,
		"WriteCapacityUnits": 5
	}'

