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

docker: docker-movie docker-room docker-user docker-screening docker-reservation

docker-movie:
	docker build -f movie_service.dockerfile -t movie_service .

docker-room:
	docker build -f room_service.dockerfile -t room_service .

docker-user:
	docker build -f user_service.dockerfile -t user_service .

docker-screening:
	docker build -f screening_service.dockerfile -t screening_service .

docker-reservation:
	docker build -f reservation_service.dockerfile -t reservation_service .

docker-client:
	docker build -f client.dockerfile -t client .
