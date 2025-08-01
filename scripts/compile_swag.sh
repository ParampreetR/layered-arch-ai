#!/bin/bash

swag init -g ./cmd/server/main.go
swagger2openapi ./docs/swagger.yaml -o ./docs/openapi.yaml
sh ./scripts/add_servers_yaml.sh ./docs/openapi.yaml
