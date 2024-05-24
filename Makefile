DIR = $(shell pwd)#pwd:获得当前路径
CONFIG_PATH = $(DIR)/config
IDL_PATH = $(DIR)/idl
API_PATH = $(DIR)/api
KITEX_GEN_PATH=$(DIR)/kitex_gen
RPC = $(DIR)/rpc
API=api
MODULE= bocchi


.PHONY: env-up
env-up:
	sh init.sh
	go mod tidy
	docker-compose up -d

.PHONY: env-down
env-down:
	docker-compose down

SERVICES := api user party itinerary
service = $(word 1, $@)

.PHONY: $(SERVICES)
$(SERVICES):
	@echo "$(PERFIX) Automatic run server" \

	@if [ $(service) != $(API) ];then \
    		go run $(RPC)/$(service) ; fi

	@if [ $(service) = $(API) ];then \
		go run $(API_PATH) ; fi


.PHONY: start-all
start-all:
	sh start.sh

KSERVICES := user party itinerary interaction
.PHONY: kgen
kgen:
	@for kservice in $(KSERVICES); do \
		kitex -module ${MODULE} ${IDL_PATH}/$$kservice.thrift; \
    	cd ${RPC};cd $$kservice;kitex -module ${MODULE} -service $$kservice -use ${KITEX_GEN_PATH} ${IDL_PATH}/$$kservice.thrift; \
    	cd ../../; \
    done \


.PHONY: hzgen
hzgen:
	cd ${API_PATH}; \
	hz update -idl ${IDL_PATH}/api.thrift; \
	swag init; \


