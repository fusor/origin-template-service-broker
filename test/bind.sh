#!/bin/bash -e

. shared.sh

req="{
  \"plan_id\": \"$planUUID\",
  \"service_id\": \"$serviceUUID\"
}"

curl \
  -X PUT \
  -H 'X-Broker-API-Version: 2.9' \
  -H 'Content-Type: application/json' \
  -d "$req" \
  -v \
  http://localhost:8000/v2/service_instances/$instanceUUID/service_bindings/$bindingUUID
