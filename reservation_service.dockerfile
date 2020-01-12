FROM obraun/vss-micro-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o reservation_service/reservation_service reservation_service/reservation.go

FROM alpine
COPY --from=builder /app/reservation_service/reservation_service /app/reservation_service
EXPOSE 8091
ENTRYPOINT [ "/app/reservation_service" ]