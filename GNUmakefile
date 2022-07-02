VERSION=1.8.2-beta.1

default: build

build: fmtcheck
	go build

codegen:
	openapi-generator generate -i openapi.yml -g go --additional-properties=packageName=pingone,packageVersion=$(VERSION) -o . --git-repo-id pingone-go --git-user-id patrickcping

.PHONY: build codegen fmtcheck