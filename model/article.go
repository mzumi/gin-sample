package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title string
	Text  string
}

func (a *Article) Create(title string, text string) {
	a.Title = title
	a.Text = text

	db.Create(a)
}
