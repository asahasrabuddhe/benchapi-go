VERSION := $(shell date -u +%Y%m%dT%H%M%S)

.PHONY: test
test:
	go test .

.PHONY: build
build: test
	docker build -f build/package/Dockerfile -t ajitemsahasrabuddhe/benchapi-go:$(VERSION) -t ajitemsahasrabuddhe/benchapi-go:latest .

.PHONY: push
push:
	docker push ajitemsahasrabuddhe/benchapi-go