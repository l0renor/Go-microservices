all:
	@echo "make registry"
	@echo "make movie"
	@echo "make room"
	@echo "make user"
	@echo "make screening"
	@echo "make reservation"
	@echo "make client"


registry:
	etcd

movie:
	MICRO_REGISTRY=etcd go run movie_service/main.go

room:
    MICRO_REGISTRY=etcd go run room_service/main.go

user:
	MICRO_REGISTRY=etcd go run user_service/main.go

screening:
    MICRO_REGISTRY=etcd go run screening_service/main.go

reservation:
    MICRO_REGISTRY=etcd go run reservation_service/main.go

client:
	MICRO_REGISTRY=etcd go run client/main.go