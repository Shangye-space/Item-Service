docker exec item-service_item-service_1 /bin/bash -c "go test -v $(go list ./src/api/item | grep -v /vendor/) -cover"

