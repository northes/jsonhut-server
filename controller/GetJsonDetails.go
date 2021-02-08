package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"jsonhutapi/models"
	"net/http"
)

func GetJsonDetails(ctx *gin.Context) {
	id := ctx.Param("id")

	resultData, err := models.QueryJsonBodyByJsonID(id)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusNotFound, models.ReturnJsonWithoutData{
			Code: 404,
			Msg:  "Record not found",
		})
		return
	}

	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(resultData.JsonBody), &dat); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ReturnJsonWithoutData{
			Code: 500,
			Msg:  "Unsupported data type",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.DetailsReturnJson{
		Code: 200,
		Msg:  "Success",
		Data: models.DetailsReturnData{
			JsonBody: dat,
			Url:      "https://api.jsonhut.com/bins/" + id,
		},
	})
}