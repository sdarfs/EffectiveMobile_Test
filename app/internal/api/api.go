// CreateUser godoc
// @Summary Создать нового пользователя
// @Description Добавляет пользователя в БД, обогащая его данные (возраст, пол, национальность)
// @Tags users
// @Accept json
// @Produce json
// @Param request body entity.CreateRequest true "Данные пользователя"
// @Success 200 {object} models.User
// @Failure 400 {object} entity.ResponseErr
// @Router /create_user [post]

package api

import (
	"errors"
	"fmt"
	"project_mobile/database/models"
	"project_mobile/database/postgres"
	"project_mobile/internal/entity"
	"project_mobile/internal/services"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var request entity.CreateRequest
	var response models.User

	pg, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	err = pg.Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, "can`t parse body")
		return
	}

	response.Name = request.Name
	if response.Name == "" {
		ctx.JSON(http.StatusBadRequest, "can`t find name")
		return
	}
	response.Surname = request.Surname
	if response.Surname == "" {
		ctx.JSON(http.StatusBadRequest, "can`t find surname")
		return
	}
	response.Patronymic = request.Patronymic

	age, err := services.Age(request.Name)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	response.Age = age

	gender, err := services.Gender(request.Name)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	response.Gender = gender

	nationality, err := services.Nationality(request.Name)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	response.Nationality = nationality

	err = pg.InsertUser(response)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, response)

}

func UpdateUser(ctx *gin.Context) {
	var request entity.UpdateRequest

	pg, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	err = pg.Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()

	if err := ctx.BindJSON(&request); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	err = pg.UpdateUser(request.FieldToUpdate, request.NewValue, request.UserId)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	resp := entity.ResponseOk{Message: fmt.Sprintf("user %v has been updated by value %v: %v", request.UserId, request.FieldToUpdate, request.NewValue)}
	ctx.IndentedJSON(http.StatusOK, resp)
}

func DeleteUser(ctx *gin.Context) {
	var request entity.DeleteRequest

	pg, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	err = pg.Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()

	if err := ctx.BindJSON(&request); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	err = pg.DeleteUser(request.UserID)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	resp := entity.ResponseOk{
		Message: fmt.Sprintf("user %v has been deleted", request.UserID),
	}
	ctx.IndentedJSON(http.StatusOK, resp)
}

func GetUsers(ctx *gin.Context) {
	pg, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	err = pg.Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()

	res, err := pg.GetUsers()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, res)
}

func GetUsersFilter(ctx *gin.Context) {
	EmptyFilterError := errors.New("got empry filter")

	pg, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	err = pg.Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()

	filters := ctx.Request.URL.Query()
	if len(filters) == 0 {
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: EmptyFilterError})
		return
	}

	res, err := pg.GetUsersByFilter(filters)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, entity.ResponseErr{Err: err})
		return
	}

	ctx.IndentedJSON(http.StatusOK, res)
}
