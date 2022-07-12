VERSION=1.9.1

OWNER=patrickcping
REPO=pingone-go

DV=davinci
MGMT=management
AUTHZ=authorize
CRED=credentials
FRAUD=fraud
MFA=mfa
RISK=risk
VERIFY=verify

default: build

build: fmtcheck
	go build

codegen: codegen-davinci codegen-management codegen-authorize codegen-credentials codegen-fraud codegen-mfa codegen-risk codegen-verify

codegen-davinci:
	@echo "==> Running codegen-$(DV)..."
	@if [[ -f "pingone-$(DV).yml" ]]; then \
		openapi-generator generate -i pingone-$(DV).yml -g go --additional-properties=packageName=$(DV),packageVersion=$(VERSION),isGoSubmodule=true -o ./$(DV) --git-repo-id $(REPO) --git-user-id $(OWNER); \
		cd ./$(DV); \
		go mod tidy; \
		go mod vendor; \
		go build; \
	else \
		echo "pingone-$(DV).yml missing.  Skipping"; \
	fi
	
codegen-management:
	@echo "==> Running codegen-$(MGMT)..."
	@if [[ -f "pingone-$(MGMT).yml" ]]; then \
		openapi-generator generate -i pingone-$(MGMT).yml -g go --additional-properties=packageName=$(MGMT),packageVersion=$(VERSION),isGoSubmodule=true -o ./$(MGMT) --git-repo-id $(REPO) --git-user-id $(OWNER); \
		cd ./$(MGMT); \
		go mod tidy; \
		go mod vendor; \
		go build; \
	else \
		echo "pingone-$(MGMT).yml missing.  Skipping"; \
	fi
	
codegen-authorize:
	@echo "==> Running codegen-$(AUTHZ)..."
	@if [[ -f "pingone-$(AUTHZ).yml" ]]; then \
		openapi-generator generate -i pingone-$(AUTHZ).yml -g go --additional-properties=packageName=$(AUTHZ),packageVersion=$(VERSION),isGoSubmodule=true -o ./$(AUTHZ) --git-repo-id $(REPO) --git-user-id $(OWNER); \
		cd ./$(AUTHZ); \
		go mod tidy; \
		go mod vendor; \
		go build; \
	else \
		echo "pingone-$(AUTHZ).yml missing.  Skipping"; \
	fi
	
codegen-credentials:
	@echo "==> Running codegen-$(CRED)..."
	@if [[ -f "pingone-$(CRED).yml" ]]; then \
		openapi-generator generate -i pingone-$(CRED).yml -g go --additional-properties=packageName=$(CRED),packageVersion=$(VERSION),isGoSubmodule=true -o ./$(CRED) --git-repo-id $(REPO) --git-user-id $(OWNER); \
		cd ./$(CRED); \
		go mod tidy; \
		go mod vendor; \
		go build; \
	else \
		echo "pingone-$(CRED).yml missing.  Skipping"; \
	fi
	
codegen-fraud:
	@echo "==> Running codegen-$(FRAUD)..."
	@if [[ -f "pingone-$(FRAUD).yml" ]]; then \
		openapi-generator generate -i pingone-$(FRAUD).yml -g go --additional-properties=packageName=$(FRAUD),packageVersion=$(VERSION),isGoSubmodule=true -o ./$(FRAUD) --git-repo-id $(REPO) --git-user-id $(OWNER); \
		cd ./$(FRAUD); \
		go mod tidy; \
		go mod vendor; \
		go build; \
	else \
		echo "pingone-$(FRAUD).yml missing.  Skipping"; \
	fi
	
codegen-mfa:
	@echo "==> Running codegen-$(MFA)..."
	@if [[ -f "pingone-$(MFA).yml" ]]; then \
		openapi-generator generate -i pingone-$(MFA).yml -g go --additional-properties=packageName=$(MFA),packageVersion=$(VERSION),isGoSubmodule=true -o ./$(MFA) --git-repo-id $(REPO) --git-user-id $(OWNER); \
		cd ./$(MFA); \
		go mod tidy; \
		go mod vendor; \
		go build; \
	else \
		echo "pingone-$(MFA).yml missing.  Skipping"; \
	fi
	
codegen-risk:
	@echo "==> Running codegen-$(RISK)..."
	@if [[ -f "pingone-$(RISK).yml" ]]; then \
		openapi-generator generate -i pingone-$(RISK).yml -g go --additional-properties=packageName=$(RISK),packageVersion=$(VERSION),isGoSubmodule=true -o ./$(RISK) --git-repo-id $(REPO) --git-user-id $(OWNER); \
		cd ./$(RISK); \
		go mod tidy; \
		go mod vendor; \
		go build; \
	else \
		echo "pingone-$(RISK).yml missing.  Skipping"; \
	fi
	
codegen-verify:
	@echo "==> Running codegen-$(VERIFY)..."
	@if [[ -f "pingone-$(VERIFY).yml" ]]; then \
		openapi-generator generate -i pingone-$(VERIFY).yml -g go --additional-properties=packageName=$(VERIFY),packageVersion=$(VERSION),isGoSubmodule=true -o ./$(VERIFY) --git-repo-id $(REPO) --git-user-id $(OWNER); \
		cd ./$(VERIFY); \
		go mod tidy; \
		go mod vendor; \
		go build; \
	else \
		echo "pingone-$(VERIFY).yml missing.  Skipping"; \
	fi
	
.PHONY: build codegen codegen-davinci codegen-management codegen-authorize codegen-credentials codegen-fraud codegen-mfa codegen-risk codegen-verify fmtcheck