all:
	@echo "make registry"
	@echo "make movie"
	@echo "make room"
	@echo "make user"
	@echo "make screening"
	@echo "make reservation"
	@echo "make client"
	@echo "make docker"

registry:
	etcd

movie:
	MICRO_REGISTRY=etcd go run movie_service/movie.go

room:
	MICRO_REGISTRY=etcd go run room_service/room.go

user:
	MICRO_REGISTRY=etcd go run user_service/user.go

screening:
	MICRO_REGISTRY=etcd go run screening_service/screening.go

reservation:
	MICRO_REGISTRY=etcd go run reservation_service/reservation.go

client:
	MICRO_REGISTRY=etcd go run client/client.go

docker: docker-movie docker-room docker-user docker-screening docker-reservation docker-client

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
