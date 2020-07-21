package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title string
	Text  string
}

func NewArticle(title string, text string) *Article {
	a := &Article{Title: title, Text: text}
	return a
}

func (a *Article) Create() {
	db.Create(a)
}

func (a *Article) Save() {
	db.Save(a)
}

func FirstArticle(id int) *Article {
	article := &Article{}

	db.First(article, id)
	return article
}

func DeleteArticle(id int) {
	a := FirstArticle(id)
	db.Delete(a)
}
