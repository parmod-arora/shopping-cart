.PHONY: test build run volumes db setup migrate

# APP_NAME is used as a naming convention for resources to the local environment
APP_NAME := shopping-cart
# APP_PATH is the project/app directory in the container
APP_PATH := /${APP_NAME}

# ----------------------------
# Development environment
# ----------------------------

# Set which .env file to use. On local we load .env
ENV_FILE = .env

# Set compose command
COMPOSE = docker-compose -f docker-compose.yml

# Commands for running docker compose
RUN_COMPOSE = $(COMPOSE) run --rm --service-ports -w $(APP_PATH) $(MOUNT_VOLUME) go
TEST_COMPOSE = $(COMPOSE) run --rm -w $(APP_PATH) $(MOUNT_VOLUME) go

# setup creates/initializes development environment dependencies
setup: db sleep migrate

# test executes project tests in a golang container
NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m

test: setup
	@if $(TEST_COMPOSE) env $(shell cat $(ENV_FILE) | egrep -v '^#' | xargs) \
		make go-test; \
	then printf "\n\n$(OK_COLOR)[Test okay -- `date`]$(NO_COLOR)\n"; \
	else printf "\n\n$(ERROR_COLOR)[Test FAILED -- `date`]$(NO_COLOR)\n"; exit 1; fi

# go-test executes test for all packages
go-test:
	GOFLAGS=-mod=vendor go test -v  -parallel 2 -count=1 -cover -coverprofile=coverage.out -outputdir=. ./...
# go tool cover -html=./coverage.out -o cover.html
# run starts the web server in a golang container
run: MOUNT_VOLUME =  $(if $(strip $(CONTAINER_SUFFIX)),,-v $(shell pwd):$(APP_PATH))
run: setup
	@$(RUN_COMPOSE) env $(shell cat $(ENV_FILE) | egrep -v '^#|^DATABASE_URL' | xargs) \
		go run -mod=vendor cmd/serverd/*.go

# build generates project binary
build:
	go clean -mod=vendor -i -x -cache ./...
	go build -mod=vendor -v -a -i ./cmd/serverd


# compose-down stops and removes all containers and resources associated to docker-compose.yml
compose-down:
	$(COMPOSE) down -v

# dbredo drops and re-migrates the database
dbredo: dbdrop migrate

# ----------------------------
# Private targets
# ----------------------------

# db runs the db service defined in the compose file
db:
	$(COMPOSE) run --rm alpine sh -c "nc -vz db 5432" && exit 0 || $(COMPOSE) up -d db

# migrate runs the db-migrate service defined in the compose file
migrate: MOUNT_VOLUME =  $(if $(strip $(CONTAINER_SUFFIX)),,-v $(shell pwd)/data/migrations:/migrations)
migrate:
	$(COMPOSE) run --rm $(MOUNT_VOLUME) db-migrate \
	sh -c 'sleep 5;./migrate -path /migrations -database $$DATABASE_URL up'

# dbdrop executes db migration DROP scripts on configured database (i.e. DATABASE_URL)
dbdrop:
	$(COMPOSE) run --rm $(MOUNT_VOLUME) db-migrate \
	sh -c './migrate -path /migrations -database $$DATABASE_URL drop'

# sleep is to delay the test from running to ensure all services (i.e. db) are up
sleep:
	sleep 5

# ----------------------------
# Tools
# ----------------------------
# update-vendor updates the vendor folder
update-vendor:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor

SQLBOILER_GOMOD := $(shell grep sqlboiler go.mod)
SQLBOILER_VER := $(word 2,$(strip $(SQLBOILER_GOMOD)))
# run migrate; run proper sqlboiler; then tidy go.mod
generate-models: migrate
	GO111MODULE=on go get github.com/volatiletech/sqlboiler@$(SQLBOILER_VER) && \
	GO111MODULE=on go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql@$(SQLBOILER_VER) && \
	sqlboiler psql --config sqlboiler.toml && \
	GO111MODULE=on go mod tidy

generate-mock:
	# FIXME: installing mockery with go get is deprecated. See https://github.com/vektra/mockery#installation
	GO111MODULE=on go get github.com/vektra/mockery/.../
	mockery -dir ./internal -all -keeptree