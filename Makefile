DIR = $(shell pwd)#pwd:获得当前路径
CONFIG_PATH = $(DIR)/config
IDL_PATH = $(DIR)/idl
API_PATH = $(DIR)/api
RPC = $(DIR)/rpc
API=api


.PHONY: env-up
env-up:
	sh init.sh
	go mod tidy
	docker-compose up -d

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

