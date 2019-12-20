docker exec item-service_item-service_1 /bin/bash -c "go test -v $(go list ./src/... | grep -v /vendor/) -cover"

