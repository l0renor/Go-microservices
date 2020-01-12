FROM obraun/vss-micro-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o client/client client/client.go

FROM alpine
COPY --from=builder /app/client/client /app/client
EXPOSE 8091
ENTRYPOINT [ "/app/client" ]