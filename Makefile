GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

all: test build
build:
	(cd cmd/day1; $(GOBUILD) -o main -v)
	(cd cmd/day2; $(GOBUILD) -o main -v)
test:
	$(GOTEST) ./...
