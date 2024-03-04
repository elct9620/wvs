GO=go
GOCOVER=$(GO) tool cover
GOTEST=$(GO) test

.PHONY: test

test:
	$(GOTEST) -coverprofile=coverage.out -coverpkg=./... ./...

cover: test
	$(GOCOVER) -func=coverage.out

cover-html: test
	$(GOCOVER) -html=coverage.out
