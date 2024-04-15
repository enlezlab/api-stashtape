#!/bin/bash

gen:
	go generate ./

run_api_server:
	# go run ./cmd/api-server/main.go
	./scripts/dev.sh
