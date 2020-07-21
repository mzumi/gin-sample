package handler

import (
	"net/http"
	"strconv"

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

		article := model.NewArticle(requestBody.Title, requestBody.Text)
		article.Create()

		c.JSON(http.StatusOK, struct {
			Title string `json:"title"`
			Text  string `json:"text"`
		}{
			Title: article.Title,
			Text:  article.Text,
		})
	}
}

func ArticleRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, struct {
				Message string `json:"message"`
			}{
				Message: "invalid ID",
			})
			return
		}

		article := model.FirstArticle(id)

		c.JSON(http.StatusOK, struct {
			Title string `json:"title"`
			Text  string `json:"text"`
		}{
			Title: article.Title,
			Text:  article.Text,
		})
	}
}

func ArticleUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, struct {
				Message string `json:"message"`
			}{
				Message: "invalid ID",
			})
			return
		}

		requestBody := struct {
			Title string `json:"title"`
			Text  string `json:"text"`
		}{}

		c.Bind(&requestBody)

		article := model.FirstArticle(id)
		article.Title = requestBody.Title
		article.Text = requestBody.Text

		article.Save()

		c.JSON(http.StatusOK, struct {
			Title string `json:"title"`
			Text  string `json:"text"`
		}{
			Title: article.Title,
			Text:  article.Text,
		})
	}
}

func ArticleDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, struct {
				Message string `json:"message"`
			}{
				Message: "invalid ID",
			})
			return
		}

		model.DeleteArticle(id)

		c.JSON(http.StatusOK, struct{}{})
	}
}
