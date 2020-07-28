VERSION := $(shell date -u +%Y%m%dT%H%M%S)

.PHONY: build
build:
	docker build -f build/package/Dockerfile -t ajitemsahasrabuddhe/benchapi-go:$(VERSION) -t ajitemsahasrabuddhe/benchapi-go:latest .