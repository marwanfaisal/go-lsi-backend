# go-lsi-backend

### Run App ###
- sudo chmod -R 777 dbdata
- go mod tidy <!-- (if needed) -->
- sudo docker-compose up -d

### Stop App ###
- sudo docker-compose down

### App Log ###
- docker logs -f lsi_backend

### Migration SourceDB ###
- docker exec -it lsi_backend bash
- go run db/source/migrate/up/up.go
- exit

### Migration Rollback SourceDB ###
- docker exec -it lsi_backend bash
- go run db/source/migrate/down/down.go
- exit

### Migration DestinationDB ###
- docker exec -it lsi_backend bash
- go run db/destination/migrate/up/up.go
- exit

### Migration Rollback DestinationDB ###
- docker exec -it lsi_backend bash
- go run db/destination/migrate/down/down.go
- exit

### Seed Data ###
- docker exec -it lsi_backend bash
- go run db/seed/product/product.go
- exit

### Run Test ###
- docker exec -it lsi_backend bash
- go clean -testcache && go test ./test/
- exit