NODE_MODULES = 'node_modules'

.phony: assets build deps lint test

assets:
	npm run assets

build:
	ifeq ("$(wildcard $(NODE_MODULES))", "")
		@echo "please run make deps"
	else
		npm run build
		@echo "Anubis is now built to ./var/anubis"
	endif

deps:
	npm ci
	go mod download

lint:
	go vet ./...
	staticcheck ./...

test:
	npm run test