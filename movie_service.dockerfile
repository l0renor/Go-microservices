FROM obraun/vss-micro-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o movie_service/movie_service movie_service/movie.go

FROM alpine
COPY --from=builder /app/movie_service/movie_service /app/movie_service
EXPOSE 8091
ENTRYPOINT [ "/app/movie_service" ]