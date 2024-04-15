#!/bin/bash

source ./.env
export AWS_CREDS=$AWS_CREDS
export AWS_CREDS_SECRET=$AWS_CREDS_SECRET
go run ./cmd/api-server/main.go
