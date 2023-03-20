docker_image_for_generating = docker_image_for_generating
pwd = $(PWD)

help: ## Show this help.
## --------------------------------------------------------------------------
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

test: ##        Run tests.
	go test -race -bench=. -benchmem ./...

test_with_docker: ## Run tests for each needs docker.
	go test -tags testdocker -race -bench=. -benchmem ./... -timeout 15m

lint: ##        Run linters for project using golangci-lint on this machine.
	golangci-lint run -v

lint_docker: ## Run linter in the docker.
	docker run --rm -v $(pwd):/app -w /app --env GOTAGS=testdocker golangci/golangci-lint:v1.51.2 golangci-lint run -v

ci: ##          Run CI local, needs docker.
## --------------------------------------------------------------------------
	go mod tidy
	go mod vendor
	make lint_docker
	make test_with_docker

.PHONY: pprof
pprof: ##    Fetch all pprof data in the ./pprof/date/ folder.
	./scripts/fetch_profile.sh

ginkgo_run: ## Run ginkgo tests.
	ACK_GINKGO_DEPRECATIONS=2.6.1 ginkgo run \
		--tags testdocker \
		--succinct \
		--output-interceptor-mode=none -
		-trace -r -p

ginkgo_watch: ## Run ginkgo watch for tests.
	ACK_GINKGO_DEPRECATIONS=2.6.1 ginkgo watch \
		--tags testdocker \
		--vv \
		--output-interceptor-mode=none \
		--trace -r -p
