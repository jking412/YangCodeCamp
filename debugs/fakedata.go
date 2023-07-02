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
	courses := []*model.Course{
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
	err := db.Mysql.Create(&courses).Error
	if err != nil {
		panic(err)
	}
}

func generateChapter() {
	chapters := []*model.Chapter{
		{
			Name:     "HTML-1",
			CourseID: 1,
			Question: &model.Question{
				Answer:      `<div>this is div</div>`,
				Content:     `<div>...</div>`,
				Description: "在div标签中输入this is div",
				Type:        model.HTML,
			},
		},
		{
			Name:     "HTML-2",
			CourseID: 1,
			Question: &model.Question{
				Answer:      `<h1>this is h1</h1>`,
				Content:     `<h1>...</h1>`,
				Description: "在h1标签中输入this is h1",
				Type:        model.HTML,
			},
		},
		{
			Name:     "HTML-3",
			CourseID: 1,
			Question: &model.Question{
				Answer:      `<p>this is p</p>`,
				Content:     `<p>...</p>`,
				Description: "在p标签中输入this is p",
				Type:        model.HTML,
			},
		},
		{
			Name:     "CSS-1",
			CourseID: 2,
			Question: &model.Question{
				Answer:      `margin: auto`,
				Content:     `margin: ...`,
				Description: "设置margin为auto",
				Type:        model.CSS,
			},
		},
		{
			Name:     "CSS-2",
			CourseID: 2,
			Question: &model.Question{
				Answer:      `padding: 10px`,
				Content:     `padding: ...`,
				Description: "设置padding为10px",
				Type:        model.CSS,
			},
		},
		{
			Name:     "CSS-3",
			CourseID: 2,
			Question: &model.Question{
				Answer:      `background-color: red`,
				Content:     `background-color: ...`,
				Description: "设置background-color为red",
				Type:        model.CSS,
			},
		},
		{
			Name:     "JavaScript-1",
			CourseID: 3,
			Question: &model.Question{
				Answer:      `console.log("hello world")`,
				Content:     `console.log("...")`,
				Description: "在控制台输出hello world",
				Type:        model.JavaScript,
			},
		},
		{
			Name:     "JavaScript-2",
			CourseID: 3,
			Question: &model.Question{
				Answer:      `var a = 1`,
				Content:     `var a = ...`,
				Description: "声明一个变量a并赋值为1",
				Type:        model.JavaScript,
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
