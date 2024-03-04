GO=go
GOCOVER=$(GO) tool cover
GOTEST=$(GO) test

.PHONY: test

test:
	$(GOTEST) -coverprofile=coverage.out -coverpkg=./... ./...
	$(GOCOVER) -func=coverage.out
	$(GOCOVER) -html=coverage.out
