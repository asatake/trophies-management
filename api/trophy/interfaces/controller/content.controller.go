package controller

import (
	"strconv"

	"github.com/asatake/trophies-management/api/trophy/interfaces/database"
	"github.com/asatake/trophies-management/api/trophy/usecase"
)

type ContentController struct {
	Interactor usecase.ContentInteractor
}

func NewContentController(sqlHandler database.SqlHandler) *ContentController {
	return &ContentController{
		Interactor: usecase.ContentInteractor{
			ContentRepository: &database.ContentRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *ContentController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	content, err := controller.Interactor.ShowContent(id)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, content)

}
