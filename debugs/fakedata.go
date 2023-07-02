package main

import (
	"YangCodeCamp/db"
	"YangCodeCamp/model"
	"YangCodeCamp/pkg/config"
)

func main() {
	config.InitConfig("config")
	db.Init()
	generateClass()
	generateChapter()
	generateQuestion()
}

func generateClass() {
	classes := []*model.Class{
		{
			Name:        "HTML",
			Description: "HTML TEST",
		},
		{
			Name:        "CSS",
			Description: "CSS TEST",
		},
		{
			Name:        "JavaScript",
			Description: "JavaScript TEST",
		},
	}
	err := db.Mysql.Create(&classes).Error
	if err != nil {
		panic(err)
	}
}

func generateChapter() {
	chapters := []*model.Chapter{
		{
			Name:    "HTML-1",
			ClassID: 1,
		},
		{
			Name:    "HTML-2",
			ClassID: 1,
		},
		{
			Name:    "HTML-3",
			ClassID: 1,
		},
	}
	err := db.Mysql.Create(&chapters).Error
	if err != nil {
		panic(err)
	}
}

func generateQuestion() {
	questions := []*model.Question{
		{
			ChapterID: 1,
			Answer:    `<div>this is div</div>`,
		},
		{
			ChapterID: 1,
			Answer:    `<h1>this is h1</h1>`,
		},
		{
			ChapterID: 1,
			Answer:    `<p>this is p</p>`,
		},
	}
	err := db.Mysql.Create(&questions).Error
	if err != nil {
		panic(err)
	}
}
