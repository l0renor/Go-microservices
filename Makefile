all:
	@echo "make start-registry"


start-registry:
	etcd

start-movie:
	MICRO_REGISTRY=etcd go run counter-service/main.go

start-greeter:
	MICRO_REGISTRY=etcd go run greeter-service/main.go

start-client:
	MICRO_REGISTRY=etcd go run client/main.go