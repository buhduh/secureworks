WD = $(shell pwd)

INT_CONTAINER_NAME = secureworks_integration
APP_CONTAINER_NAME = secureworks_assessment
BIN_CONTAINER_NAME = secureworks_binary
DATABASE_CONTAINER_NAME = secureworks_database
SOURCER_CONTAINER_NAME = secureworks_sourcer
VENDOR_CONTAINER_NAME = secureworks_vendor
TEST_CONTAINER_NAME = secureworks_tester

BUILD_DIR = build
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
export PORT_STR = :8080
export TEST_SHELL = /root/test.sh

all: $(BUILD_DIR) ## All the things
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
		--build-arg PORT_STR          \
		-t ${APP_CONTAINER_NAME} .  
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
		--build-arg PORT_STR          \
		-t ${BIN_CONTAINER_NAME}      \
		--target builder .
.PHONY: builder

integration: ## generates a shell script to be thrown at the application for final testing
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
		--build-arg PORT_STR          \
		--build-arg TEST_SHELL        \
		-t ${INT_CONTAINER_NAME}      \
		--target integration .

database: ## provisions the database only, mostly for testing
	@docker build                     \
		--build-arg SQL_DB            \
		--build-arg CREATE_DB_SQL     \
		-t ${DATABASE_CONTAINER_NAME} \
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
		-t ${TEST_CONTAINER_NAME}     \
		--target tester .  
.PHONY: test
	
help: ## Print this helpfile
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help

clean:
	@rm -rf $(BUILD_DIR)
	@docker image rm -f ${INT_CONTAINER_NAME}
	@docker image rm -f ${APP_CONTAINER_NAME}
	@docker image rm -f ${BIN_CONTAINER_NAME}
	@docker image rm -f ${DATABASE_CONTAINER_NAME}
	@docker image rm -f ${SOURCER_CONTAINER_NAME}
	@docker image rm -f ${VENDOR_CONTAINER_NAME}
	@docker image rm -f ${TEST_CONTAINER_NAME}
	@docker builder prune -f
.PHONY: clean

$(BUILD_DIR):
	@mkdir $@
