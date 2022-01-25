.PHONY: up
up: ##@development Build and start development environment in background.
	docker-compose up --build -d

.PHONY: logs
logs: ##@development Follows development logs [service="svc1 svc2..."].
	docker-compose logs -f --tail=100 $(service)

.PHONY: shell
shell: ##@development Start a shell session within the container.
	docker-compose run --rm app /bin/sh

lint_version ?= v1.40-alpine
.PHONY: lint
lint: ##@development Runs static analysis code.
	docker run --rm \
		-v $(shell pwd):/app \
		-w /app \
		golangci/golangci-lint:$(lint_version) \
		golangci-lint run --timeout 3m

.PHONY: test
test: ##@development Runs the tests.
	docker-compose run --rm app go test ./...

.PHONY: stop
stop: ##@development Stop development environment and remove containers.
	docker-compose down -v --remove-orphans

.PHONY: deploy
deploy: ##@production Deploy to heroku.
	git push heroku main

.PHONY: deploy_logs
deploy_logs:
	heroku logs --tail
