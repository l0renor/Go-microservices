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

func (m *movieService) CreateMovie(ctx context.Context, req *api.CreateMovieReq, resp *api.CreateMovieResp) error {
	id := m.nextID()
	m.movies[id] = req.Name
	resp.MovieID = id
	return nil
}

func (m *movieService) DeleteMovie(ctx context.Context, req *api.DeleteMovieReq, resp *api.DeleteMovieResp) error {
	_, ok := m.movies[req.GetMovieID()]
	if !ok {
		return errors.NotFound("ERR-NO-MOVIE", "Movie (ID: %d) not found!", req.GetMovieID())
	}
	_, err := m.screening.DeleteScreeningsWithMovie(context.TODO(), &api.DeleteScreeningsWithMovieReq{MovieID: req.GetMovieID()})
	if err != nil {
		return err
	}
	delete(m.movies, req.GetMovieID())
	return nil
}

func (m *movieService) GetMovie(ctx context.Context, req *api.GetMovieReq, resp *api.GetMovieResp) error {
	title, ok := m.movies[req.GetMovieID()]
	if !ok {
		return errors.NotFound("ERR-NO-MOVIE", "Movie (ID: %d) not found!", req.GetMovieID())
	}
	resp.Title = title
	return nil
}

func (m *movieService) GetMovies(ctx context.Context, req *api.GetMoviesReq, resp *api.GetMoviesResp) error {
	var movies []*api.Tuple
	for id, title := range m.movies {
		movies = append(movies, &api.Tuple{
			Title:   title,
			MovieID: id,
		})
	}
	resp.Movies = movies
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
