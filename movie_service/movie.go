package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"github.com/ob-vss-ws19/blatt-4-myteam/api"
	"github.com/ob-vss-ws19/blatt-4-myteam/helpers"
	"log"
)

type movieService struct {
	movies    map[int32]string
	nextID    func() int32
	screening api.Screening_Service
}

func (m *movieService) CreateMovie(ctx context.Context, req *api.CreateMovieMsg, rsp *api.CreateMovieResponseMsg) error {
	id := m.nextID()
	m.movies[id] = req.Name
	rsp.Id = id
	return nil
}

func (m *movieService) DeleteMovie(ctx context.Context, req *api.DeleteMovieMsg, rsp *api.DeleteMovieResponseMsg) error {
	_, ok := m.movies[req.GetId()]
	if !ok {
		return errors.NotFound("ERR-NO-MOVIE", "Movie (ID: %d) not found!", req.GetId())
	}
	_, err := m.screening.DeleteMovie(context.TODO(), &api.DeleteMovieReq{MovieID: req.GetId()})
	if err != nil {
		return err
	}
	delete(m.movies, req.GetId())
	return nil
}

func (m *movieService) GetMovie(ctx context.Context, req *api.GetMovieMsg, rsp *api.GetMovieResponseMsg) error {
	title, ok := m.movies[req.GetId()]
	if !ok {
		return errors.NotFound("ERR-NO-MOVIE", "Movie (ID: %d) not found!", req.GetId())
	}
	rsp.Title = title
	return nil
}

func (m *movieService) GetMovies(ctx context.Context, req *api.GetMoviesMsg, rsp *api.GetMoviesResponseMsg) error {
	var movies []*api.Tuple
	for id, title := range m.movies {
		movies = append(movies, &api.Tuple{
			Title: title,
			Id:    id,
		})
	}
	rsp.Movies = movies
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("movie"),
		micro.Version("latest"),
	)

	screening := micro.NewService()
	screening.Init()

	service.Init()

	if err := api.RegisterMovie_ServiceHandler(service.Server(), &movieService{
		movies:    make(map[int32]string),
		nextID:    helpers.IDGenerator(),
		screening: api.NewScreening_Service("screening", screening.Client()),
	}); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
