package main

import (
	"YangCodeCamp/db"
	"YangCodeCamp/model"
	"YangCodeCamp/pkg/config"
)

func main() {
	config.InitConfig("config")
	db.Init()
	db.Mysql.Exec("delete from chapters")
	db.Mysql.Exec("delete from courses")
	db.Mysql.Exec("delete from questions")
	generateClass()
	generateChapter()
	//generateQuestion()
}

func generateClass() {
	courses := []*model.Course{
		{
			ID:          1,
			Name:        "HTML",
			Description: "HTML TEST",
		},
		{
			ID:          2,
			Name:        "CSS",
			Description: "CSS TEST",
		},
		{
			ID:          3,
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
			ID:       1,
			Name:     "HTML-1",
			CourseID: 1,
			Question: &model.Question{
				ID:          1,
				Answer:      `<div>this is div</div>`,
				Content:     `<div>...</div>`,
				Description: "在div标签中输入this is div",
				Type:        model.HTML,
			},
		},
		{
			ID:       2,
			Name:     "HTML-2",
			CourseID: 1,
			Question: &model.Question{
				ID:          2,
				Answer:      `<h1>this is h1</h1>`,
				Content:     `<h1>...</h1>`,
				Description: "在h1标签中输入this is h1",
				Type:        model.HTML,
			},
		},
		{
			ID:       3,
			Name:     "HTML-3",
			CourseID: 1,
			Question: &model.Question{
				ID:          3,
				Answer:      `<p>this is p</p>`,
				Content:     `<p>...</p>`,
				Description: "在p标签中输入this is p",
				Type:        model.HTML,
			},
		},
		{
			ID:       4,
			Name:     "CSS-1",
			CourseID: 2,
			Question: &model.Question{
				ID:          4,
				Answer:      `margin: auto`,
				Content:     `margin: ...`,
				Description: "设置margin为auto",
				Type:        model.CSS,
			},
		},
		{
			ID:       5,
			Name:     "CSS-2",
			CourseID: 2,
			Question: &model.Question{
				ID:          5,
				Answer:      `padding: 10px`,
				Content:     `padding: ...`,
				Description: "设置padding为10px",
				Type:        model.CSS,
			},
		},
		{
			ID:       6,
			Name:     "CSS-3",
			CourseID: 2,
			Question: &model.Question{
				ID:          6,
				Answer:      `background-color: red`,
				Content:     `background-color: ...`,
				Description: "设置background-color为red",
				Type:        model.CSS,
			},
		},
		{
			ID:       7,
			Name:     "JavaScript-1",
			CourseID: 3,
			Question: &model.Question{
				ID:          7,
				Answer:      `console.log("hello world")`,
				Content:     `console.log("...")`,
				Description: "在控制台输出hello world",
				Type:        model.JavaScript,
			},
		},
		{
			ID:       8,
			Name:     "JavaScript-2",
			CourseID: 3,
			Question: &model.Question{
				ID:          8,
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
