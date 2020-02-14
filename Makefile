WD = $(shell pwd)

INT_IMAGE_NAME = secureworks_integration_dale
APP_IMAGE_NAME = secureworks_assessment_dale
BIN_IMAGE_NAME = secureworks_binary_dale
DATABASE_IMAGE_NAME = secureworks_database_dale
SOURCER_IMAGE_NAME = secureworks_sourcer_dale
VENDOR_IMAGE_NAME = secureworks_vendor_dale
TEST_IMAGE_NAME = secureworks_tester_dale

APP_CONTAINER_NAME = ${APP_IMAGE_NAME}_container

VENDORED_SRC = ${WD}/${BUILD_DIR}/src
VENDOR = ${BUILD_DIR}/vendor

export HOST_IP_DB = ips/GeoLite2-City.mmdb
export CONT_IP_DB = /root/maxmind/ips.mndb
export APP_SRC = src
export CONT_BIN = /root/secureworks
export CONT_GEN_SCRIPT = /root/generate.sh
export GEN_GO = /go/src/secureworks/constants/constants.go
export SQL_DB = /root/secureworks_db
export VENDOR_DIR = /root/vendor_dir
export CREATE_DB_SQL = /root/createdb.sql
export PORT = 8080
export TEST_SHELL = /root/test.sh

all: ## All the things, runs docker container prune -f, be careful
	@docker build                     \
		--build-arg HOST_IP_DB        \
		--build-arg CONT_IP_DB        \
		--build-arg APP_SRC           \
		--build-arg CONT_BIN          \
		--build-arg CONT_GEN_SCRIPT   \
		--build-arg GEN_GO            \
		--build-arg VENDOR_DIR        \
		--build-arg SQL_DB            \
		--build-arg CREATE_DB_SQL     \
		--build-arg PORT              \
		-t ${APP_IMAGE_NAME} .  
	@scripts/kill.sh ${APP_CONTAINER_NAME}
	@docker container prune -f
	@docker run -d -p ${PORT}:${PORT} --name ${APP_CONTAINER_NAME} ${APP_IMAGE_NAME}
.PHONY: all

builder: ## makes the binary only
	@docker build                     \
		--build-arg HOST_IP_DB        \
		--build-arg CONT_IP_DB        \
		--build-arg APP_SRC           \
		--build-arg CONT_BIN          \
		--build-arg CONT_GEN_SCRIPT   \
		--build-arg GEN_GO            \
		--build-arg VENDOR_DIR        \
		--build-arg SQL_DB            \
		--build-arg CREATE_DB_SQL     \
		--build-arg PORT              \
		-t ${BIN_IMAGE_NAME}          \
		--target builder .
.PHONY: builder

database: ## provisions the database only, mostly for testing
	@docker build                     \
		--build-arg SQL_DB            \
		--build-arg CREATE_DB_SQL     \
		-t ${DATABASE_IMAGE_NAME} \
		--target database .
.PHONY: database

fmt: ## vim-go plugin wasn't fmting on save, just do it manually....
	GOPATH=${WD} go fmt secureworks/...	
.PHONY: fmt

test: ## Run the testing suite only
	@docker build                     \
		--build-arg HOST_IP_DB        \
		--build-arg CONT_IP_DB        \
		--build-arg APP_SRC           \
		--build-arg CONT_BIN          \
		--build-arg CONT_GEN_SCRIPT   \
		--build-arg GEN_GO            \
		--build-arg VENDOR_DIR        \
		--build-arg SQL_DB            \
		--build-arg CREATE_DB_SQL     \
		--build-arg PORT              \
		-t ${TEST_IMAGE_NAME}         \
		--target tester .  
.PHONY: test
	
help: ## Print this helpfile
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help

clean:
	@docker image rm -f ${INT_IMAGE_NAME}
	@docker image rm -f ${APP_IMAGE_NAME}
	@docker image rm -f ${BIN_IMAGE_NAME}
	@docker image rm -f ${DATABASE_IMAGE_NAME}
	@docker image rm -f ${SOURCER_IMAGE_NAME}
	@docker image rm -f ${VENDOR_IMAGE_NAME}
	@docker image rm -f ${TEST_IMAGE_NAME}
	@docker builder prune -f
	@docker container prune -f
.PHONY: clean
