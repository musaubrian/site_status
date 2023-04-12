GOCMD=go
GORUN=$(GOCMD) run .
GOTEST=$(GOCMD) test -v ./...

run:
	@$(GORUN)
test:
	@$(GOTEST)
