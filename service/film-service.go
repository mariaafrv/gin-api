package service

import "github.com/mariaafrv/gin-api/model"

type FilmService interface {
	Save(model.Film) model.Film
	FindAll() []model.Film
}

type filmService struct {
	films []model.Film
}

func New() FilmService {
	return &filmService{}
}

func (service *filmService) Save(film model.Film) model.Film {
	service.films = append(service.films, film)
	return film
}

func (service *filmService) FindAll() []model.Film {
	return service.films
}
