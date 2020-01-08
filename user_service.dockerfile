FROM obraun/vss-micro-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o user_service/user_service user_service/user.go

FROM alpine
COPY --from=builder /app/user_service/user_service /app/user_service
EXPOSE 8091
ENTRYPOINT [ "/app/user_service" ]