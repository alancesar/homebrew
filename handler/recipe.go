package handler

import (
	"github.com/alancesar/homebrew/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RecipeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := api.Recipe{}
		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		recipe := api.BuildRecipe(request)
		response := api.BuildResponse(*recipe)
		ctx.JSON(http.StatusOK, response)
	}
}
