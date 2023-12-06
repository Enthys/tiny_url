-include .env
export

# ===================================================================================#
# HELPERS
# ===================================================================================#

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n "Are you sure? [y/N] " && read ans && [ $${ans:-N} = y ]

# ===================================================================================#
# DEVELOPMENT
# ===================================================================================#

## run/setup: run required commands in order to start development
.PHONY: run/setup
run/setup:
	@if [ ! -e ".env" ]; then echo 'Copying .env.default under name .env.' && cp .env.default .env; fi
	@if [ ! -e "bin/migrate" ]; then echo 'Downloading migrate in ./bin ...'; curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz -C bin migrate; fi

## run/dev: start the development environment
.PHONY: run/dev
run/dev:
	@echo 'Starting development environment...'
	docker-compose up

## db/connect: connecto to the development database using psql
.PHONY: db/connect
db/connect:
	@docker container exec -it tiny_url_database psql ${DB_DSN}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	@docker container exec -it tiny_url ./bin/migrate -path ./migrations -database ${DB_DSN} up


## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new: run/setup
	@echo 'Creating migration files for ${name}'
	@./bin/migrate create -seq -ext=.sql -dir=./migrations ${name}
