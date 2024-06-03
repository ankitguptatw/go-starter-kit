UNITTESTS=$(shell go list ./... |  grep -v myapp/componentTest)

compile:
	CGO_ENABLED=0 go build -o out/apiServer ./cmd/apiServer/*.go

init-db:
	docker exec -it postgres psql -U postgres -c "create user payment_user password '1234567890';";\
	docker exec -it postgres psql -U postgres -c "create database payment_db owner=payment_user;";\

deps:
	go mod tidy # remove unused and download missing modules
	go mod download # cache modules

start-server-locally:
	CGO_ENABLED=0 go build -o out/apiServer ./cmd/apiServer/*.go
	out/apiServer -configFile=localDevSetup/local.yaml

run-migrations:
	CGO_ENABLED=0 go build -o out/migrations ./cmd/migrations/*.go
	out/migrations -configFile=localDevSetup/local.yaml up

rollback-migrations:
	CGO_ENABLED=0 go build -o out/migrations ./cmd/migrations/*.go
	out/migrations -configFile=localDevSetup/local.yaml down

CMD=up -d

static-checks: fmt lint security-checks

fmt:
	go fmt ./...
	go vet ./...

lint:
	#please install golangci-lint if not installed, brew install golangci-lint
	golangci-lint run -v -c tools/.golangci.yml

security-checks:
	#please install nancy https://github.com/sonatype-nexus-community/nancy
	mkdir -p out/nancy
	go list -json -m all | nancy sleuth --exclude-vulnerability-file tools/.nancy-ignore > out/nancy-output.txt
	cat out/nancy-output.txt

unit-tests:
	@go test $(UNITTESTS)

component-tests:
	go test myapp/componentTest

all-tests: unit-tests component-tests

test-with-coverage:
	go clean -testcache
	mkdir -p out/
	go test ./... -coverprofile=out/coverage.out; \
    go tool cover -html=out/coverage.out -o out/coverage.html; \


pre-commit-hooks:
	. ./tools/pre-commit-hooks-setup.sh
	chmod +x .git/hooks/pre-commit

pre-push-hooks:
	. ./tools/pre-push-hooks-setup.sh
	chmod +x .git/hooks/pre-push

generate-mocks:
	# install with command --> brew install mockery
	mockery --all --keeptree --inpackage --output=_mocks

generate-docs:
	# install godoc: go install -v golang.org/x/tools/cmd/godoc@latest
	godoc -http=:6060 & open http://localhost:6060/pkg/$(go list -m)

# terragrunt deployments
%/plan %/apply %/destroy %/init:
	@printf "Executing terragrunt %s...\n" $(@F)
	@cd $(@D); terragrunt $(@F) -auto-approve --terragrunt-non-interactive

.PHONY: test lint fmt compile deps