include .env

MAKEFLAGS += --silent
PROJECTNAME=$(shell basename "$(PWD)")

PROTOS_DIR=$(PWD)/api/proto/v1
PROTOS=api/proto/v1/user_service.proto
PROTOC_GEN=protoc -I. -I${PROTOS_DIR} ${PROTOS}

.PHONY: all

all: clean build kubernates
	$(info build done)
clean:
	go clean
	rm -f bin/users_client
	rm -f bin/users_server
	rm -f bin/users_idgen
	rm -f bin/users_web
	rm -rf ./internal/api/proto/v1/*
	rm -rf ./internal/api/wsdl/v1/*
	rm -rf ./docs/*
	buildah rm -a
	buildah rmi -a -f

gendocs:
	protoc --doc_out=docs --doc_opt=html,docs/users.html ${PROTOS}	

proto:	
	${PROTOC_GEN} --go_out=plugins=grpc:internal

wsdl:
	$(info Generate wsdl go code)
	wsdl2go -i api/wsdl/math.wsdl \
	  -o internal/api/wsdl/v1/math_service.wsdl.go \
	  -p v1  
	
test: 
	go test ${MODULE}/internal/... -coverprofile=docs/coverage.out
	go tool cover -html=docs/coverage.out -o docs/coverage.html

build: proto wsdl gendocs test
	GOOS=linux CGO_ENABLED=0 go build -o bin/users_client cmd/client/main.go
	GOOS=linux CGO_ENABLED=0 go build -o bin/users_server cmd/server/main.go 
	GOOS=linux CGO_ENABLED=0 go build -o bin/users_idgen cmd/id_generator/main.go
	GOOS=linux CGO_ENABLED=0 go build -o bin/users_web cmd/web_client/main.go

kubernates: image

image: image_users_server image_idgen iamge_web
image_users_server:
	./package/image/build_grpc_users.sh
image_idgen:	
	./package/image/build_idgen.sh
iamge_web:
	./package/image/build_web.sh

install:
	@echo "deploy to minikube"
	kubectl apply -f \
		deployments/kubernates/deployment-idgen.yaml
	kubectl apply -f \
		deployments/kubernates/deployment-users-grpc-server.yaml
	kubectl apply -f \
		deployments/kubernates/deployment-webclient.yaml		
	kubectl apply -f \
		deployments/kubernates/services.yaml		