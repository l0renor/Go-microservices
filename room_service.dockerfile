FROM obraun/vss-micro-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o room_service/room_service room_service/main.go

FROM alpine
COPY --from=builder /app/room_service/room_service /app/room_service
EXPOSE 8091
ENTRYPOINT [ "/app/room_service" ]