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
	//generateQuestion()
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
			Question: &model.Question{
				Answer:      `<div>this is div</div>`,
				Content:     `<div>...</div>`,
				Description: "在div标签中输入this is div",
			},
		},
		{
			Name:    "HTML-2",
			ClassID: 1,
			Question: &model.Question{
				Answer:      `<h1>this is h1</h1>`,
				Content:     `<h1>...</h1>`,
				Description: "在h1标签中输入this is h1",
			},
		},
		{
			Name:    "HTML-3",
			ClassID: 1,
			Question: &model.Question{
				Answer:      `<p>this is p</p>`,
				Content:     `<p>...</p>`,
				Description: "在p标签中输入this is p",
			},
		},
		{
			Name:    "CSS-1",
			ClassID: 2,
			Question: &model.Question{
				Answer:      `margin: auto`,
				Content:     `margin: ...`,
				Description: "设置margin为auto",
			},
		},
		{
			Name:    "CSS-2",
			ClassID: 2,
			Question: &model.Question{
				Answer:      `padding: 10px`,
				Content:     `padding: ...`,
				Description: "设置padding为10px",
			},
		},
		{
			Name:    "CSS-3",
			ClassID: 2,
			Question: &model.Question{
				Answer:      `background-color: red`,
				Content:     `background-color: ...`,
				Description: "设置background-color为red",
			},
		},
		{
			Name:    "JavaScript-1",
			ClassID: 3,
			Question: &model.Question{
				Answer:      `console.log("hello world")`,
				Content:     `console.log("...")`,
				Description: "在控制台输出hello world",
			},
		},
		{
			Name:    "JavaScript-2",
			ClassID: 3,
			Question: &model.Question{
				Answer:      `var a = 1`,
				Content:     `var a = ...`,
				Description: "声明一个变量a并赋值为1",
			},
		},
	}
	err := db.Mysql.Create(&chapters).Error
	if err != nil {
		panic(err)
	}
}

//func generateQuestion() {
//	questions := []*model.Question{
//		{
//			ChapterID: 1,
//			Answer:    `<div>this is div</div>`,
//		},
//		{
//			ChapterID: 1,
//			Answer:    `<h1>this is h1</h1>`,
//		},
//		{
//			ChapterID: 1,
//			Answer:    `<p>this is p</p>`,
//		},
//	}
//	err := db.Mysql.Create(&questions).Error
//	if err != nil {
//		panic(err)
//	}
//}
