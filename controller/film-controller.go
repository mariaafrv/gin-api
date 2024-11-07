package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mariaafrv/gin-api/model"
	"github.com/mariaafrv/gin-api/service"
)

type FilmController interface {
	FindAll() []model.Film
	Save(ctx *gin.Context) model.Film
}

type controller struct {
	service service.FilmService
}

func New(service service.FilmService) FilmController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []model.Film {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) model.Film {
	var film model.Film
	ctx.BindJSON(&film)
	c.service.Save(film)
	return film
}
