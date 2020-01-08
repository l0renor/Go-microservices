FROM obraun/vss-micro-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o screening_service/screening_service screening_service/main.go

FROM alpine
COPY --from=builder /app/screening_service/screening_service /app/screening_service
EXPOSE 8091
ENTRYPOINT [ "/app/screening_service" ]