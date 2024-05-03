DIR = $(shell pwd)#pwd:获得当前路径
CONFIG_PATH = $(DIR)/config
IDL_PATH = $(DIR)/idl
API_PATH = $(DIR)/api
RPC = $(DIR)/rpc
OUTPUT_PATH = $(DIR)/output
API=api
SHELL=/bin/bash

PERFIX = "[Makefile]"

.PHONY: env-up
env-up:
	sh init.sh
	docker-compose up -d
	go mod tidy

.PHONY: env-down
env-down:
	docker-compose down

SERVICES := api user party
service = $(word 1, $@)

.PHONY: $(SERVICES)
$(SERVICES):
	@echo "$(PERFIX) Automatic run server" \

	@if [ $(service) != $(API) ];then \
    		go run $(RPC)/$(service) ; fi

	@if [ $(service) = $(API) ];then \
		go run $(API_PATH) ; fi


.PHONY: build-all
build-all:
	sh start.sh

