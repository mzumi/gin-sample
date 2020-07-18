package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mzumi/gin-sample/model"
)

func ArticleCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := struct {
			Title string `json:"title"`
			Text  string `json:"text"`
		}{}

		c.Bind(&requestBody)

		article := model.Article{}
		article.Create(requestBody.Title, requestBody.Text)

		c.JSON(http.StatusOK, struct {
			Title string `json:"title"`
			Text  string `json:"text"`
		}{
			Title: article.Title,
			Text:  article.Text,
		})
	}
}
