package movie_service

import (
	"context"
	"github.com/micro/go-micro/errors"
	"github.com/ob-vss-ws19/blatt-4-myteam/api"
)

type movieService struct {
	movies map[int32]string
	nextID func() int32
}

func (m *movieService) CreateMovie(ctx context.Context, req *api.CreateMovieMsg, rsp *api.CreateMovieResponseMsg) error {
	id := m.nextID()
	m.movies[id] = req.Name
	rsp.Id = id
	return nil
}

func (m *movieService) DeleteMovie(ctx context.Context, req *api.DeleteMovieMsg, rsp *api.DeleteMovieResponseMsg) error {
	id := req.Id
	delete(m.movies, id)
	_, ok := m.movies[id]
	rsp.Success = !ok
	return nil
}

func (m *movieService) GetMovie(ctx context.Context, req *api.GetMovieMsg, rsp *api.GetMovieResponseMsg) error {
	id := req.Id
	res, ok := m.movies[id]
	if ok {
		rsp.Title = res
	} else {
		return errors.NotFound("movie_not_found", "Movie  with id %v not found  not found", req.Id)
	}
	return nil
}

func (m *movieService) GetMovies(ctx context.Context, req *api.GetMoviesMsg, rsp *api.GetMoviesResponseMsg) error {
	var res []*api.Tuple
	for k, v := range m.movies {
		res = append(res, &api.Tuple{
			Title: v,
			Id:    k,
		})
	}
	rsp.Movies = res
	return nil
}

func idGenerator() func() int32 {
	i := 0
	return func() int32 {
		i++
		return int32(i)
	}
}
