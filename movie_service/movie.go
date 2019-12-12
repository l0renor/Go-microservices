package movie_service

import (
	"context"
	"fmt"
)

//depricated
type movie struct {
	title string
}
type movieService struct {
	movies map[int32]string
	nextID func() int32
}

func (m *movieService) CreateMovie(ctx context.Context, req *CreateMovieMsg, rsp *CreateMovieResponseMsg) error {
	id := m.nextID()
	m.movies[id] = req.Name
	rsp.Id = id
	return nil
}

func (m *movieService) DeleteMovie(ctx context.Context, req *DeleteMovieMsg, rsp *DeleteMovieResponseMsg) error {
	id := req.Id
	delete(m.movies, id)
	_, ok := m.movies[id]
	rsp.Success = !ok
	return nil
}

func (m *movieService) GetMovie(ctx context.Context, req *GetMovieMsg, rsp *GetMovieResponseMsg) error {
	id := req.Id
	res, ok := m.movies[id]
	if ok {
		rsp.Title = res
	} else {
		rsp.Title = ""
	}
	return nil
}

func (m *movieService) GetMovies(ctx context.Context, req *GetMoviesMsg, rsp *GetMoviesResponseMsg) error {
	var res []*Tuple
	for k, v := range m.movies {
		res = append(res, &Tuple{
			Title: v,
			Id:    k,
		})
	}
	rsp.Movies = res
}

func idGenerator() func() int32 {
	i := 0
	return func() int32 {
		i++
		return int32(i)
	}
}
