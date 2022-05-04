package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/manabie-com/togo/models"

	"github.com/manabie-com/togo/factories"
)

func AddTogoTask(context *gin.Context) {
	var newTodo models.Togo

	if err := context.BindJSON(&newTodo); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, errUser := factories.UserFactory("get", &newTodo)

	if errUser != nil {
		context.JSON(http.StatusBadRequest, errUser.Error())
		return
	}

	if user != nil && user.CountTasks() >= user.LimitTasks {
		context.JSON(http.StatusBadRequest, "tasks has been limit.")
		return
	}

	result, errTodo := factories.TodoFactory("add", &newTodo)

	if errTodo != nil {
		context.JSON(http.StatusInternalServerError, errTodo.Error())
		return
	}

	context.JSON(http.StatusOK, result)
}