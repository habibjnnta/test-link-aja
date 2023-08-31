# Test Link Aja Habib Jannata

# Run by docker :
- create database terlebih dahulu 2 : testing_link_aja(test) dan test_link_aja
- edit in connection/db.go (sesuaikan username, password, port, host, dan nama tablenya)
- docker build -t test-link-aja .
- docker compose up

# Run by local computer :
- create database terlebih dahulu 2 : testing_link_aja(test) dan test_link_aja
- edit in connection/db.go (sesuaikan username, password, port, host, dan nama tablenya)
- go mod tidy
- go run . / go run main.go

# Documentation by Postman :
https://documenter.getpostman.com/view/24306180/2s9Y5bQgN4
