# Address Book API

0. Add the CRUD API
1. Define Address book data structure
2. Add database persistence 
3. Implement respository
4. Implement persistence service
5. e2e integration
6. externalize configration using viper - https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66 
7. Define Protobuf contract


## Swagger Documentation 
https://github.com/swaggo/swag

### Generate Swagger templates
`swag init -g pkg/middleware/router.go -o pkg/docs`

Download sway binary from here https://github.com/swaggo/swag/releases

## Kafka Events with Golang
https://www.sohamkamani.com/golang/working-with-kafka/



### gRPC Protoc
protoc -I=. --go_out=/Users/oolumide/dev-tools/rpm-tools/ocp-projects/migrator/migrator-poc/db-sync-service ./contact.proto

