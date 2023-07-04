package main

import (
	"YangCodeCamp/db"
	"YangCodeCamp/model"
	"YangCodeCamp/pkg/config"
)

func main() {
	config.InitConfig("config")
	db.Init()
	db.Mysql.Exec("delete from questions")
	db.Mysql.Exec("delete from chapters")
	db.Mysql.Exec("delete from courses")
	generateClass()
	generateChapter()
	generateQuestion()
}

func generateClass() {
	courses := []*model.Course{
		{
			ID:          1,
			Name:        "HTML",
			Description: "HTML TEST",
			Icon:        "fa-brands fa-html5",
		},
		{
			ID:          2,
			Name:        "CSS",
			Description: "CSS TEST",
			Icon:        "fa-brands fa-css3-alt",
		},
		{
			ID:          3,
			Name:        "JavaScript",
			Description: "JavaScript TEST",
			Icon:        "fa-brands fa-js",
		},
	}
	err := db.Mysql.Create(&courses).Error
	if err != nil {
		panic(err)
	}
}

func generateChapter() {
	//chapters := []*model.Chapter{
	//	{
	//		ID:       1,
	//		Name:     "HTML-1",
	//		CourseID: 1,
	//		Question: &model.Question{
	//			ID:          1,
	//			Answer:      `<div>this is div</div>`,
	//			Content:     `<div>...</div>`,
	//			Description: "在div标签中输入this is div",
	//			Type:        model.HTML,
	//		},
	//		PreChapterID:  0,
	//		NextChapterID: 2,
	//	},
	//	{
	//		ID:       2,
	//		Name:     "HTML-2",
	//		CourseID: 1,
	//		Number:   2,
	//		Question: &model.Question{
	//			ID:          2,
	//			Answer:      `<h1>this is h1</h1>`,
	//			Content:     `<h1>...</h1>`,
	//			Description: "在h1标签中输入this is h1",
	//			Type:        model.HTML,
	//		},
	//		PreChapterID:  1,
	//		NextChapterID: 3,
	//	},
	//	{
	//		ID:       3,
	//		Name:     "HTML-3",
	//		CourseID: 1,
	//		Number:   3,
	//		Question: &model.Question{
	//			ID:          3,
	//			Answer:      `<p>this is p</p>`,
	//			Content:     `<p>...</p>`,
	//			Description: "在p标签中输入this is p",
	//			Type:        model.HTML,
	//		},
	//		PreChapterID:  2,
	//		NextChapterID: 0,
	//	},
	//	{
	//		ID:       4,
	//		Name:     "CSS-1",
	//		CourseID: 2,
	//		Number:   1,
	//		Question: &model.Question{
	//			ID:          4,
	//			Answer:      `margin: auto`,
	//			Content:     `margin: ...`,
	//			Description: "设置margin为auto",
	//			Type:        model.CSS,
	//		},
	//		PreChapterID:  0,
	//		NextChapterID: 5,
	//	},
	//	{
	//		ID:       5,
	//		Name:     "CSS-2",
	//		CourseID: 2,
	//		Number:   2,
	//		Question: &model.Question{
	//			ID:          5,
	//			Answer:      `padding: 10px`,
	//			Content:     `padding: ...`,
	//			Description: "设置padding为10px",
	//			Type:        model.CSS,
	//		},
	//		PreChapterID:  4,
	//		NextChapterID: 6,
	//	},
	//	{
	//		ID:       6,
	//		Name:     "CSS-3",
	//		CourseID: 2,
	//		Number:   3,
	//		Question: &model.Question{
	//			ID:          6,
	//			Answer:      `background-color: red`,
	//			Content:     `background-color: ...`,
	//			Description: "设置background-color为red",
	//			Type:        model.CSS,
	//		},
	//		PreChapterID:  5,
	//		NextChapterID: 0,
	//	},
	//	{
	//		ID:       7,
	//		Name:     "JavaScript-1",
	//		CourseID: 3,
	//		Number:   1,
	//		Question: &model.Question{
	//			ID:          7,
	//			Answer:      `console.log("hello world")`,
	//			Content:     `console.log("...")`,
	//			Description: "在控制台输出hello world",
	//			Type:        model.JavaScript,
	//		},
	//		PreChapterID:  0,
	//		NextChapterID: 8,
	//	},
	//	{
	//		ID:       8,
	//		Name:     "JavaScript-2",
	//		CourseID: 3,
	//		Number:   2,
	//		Question: &model.Question{
	//			ID:          8,
	//			Answer:      `var a = 1`,
	//			Content:     `var a = ...`,
	//			Description: "声明一个变量a并赋值为1",
	//			Type:        model.JavaScript,
	//		},
	//		PreChapterID:  7,
	//		NextChapterID: 9,
	//	},
	//}
	chapters := []*model.Chapter{
		{
			ID:            1,
			Name:          "HTML-First",
			CourseID:      1,
			NextChapterID: 2,
			PreChapterID:  0,
		},
		{
			ID:            2,
			Name:          "CSS-First",
			CourseID:      2,
			NextChapterID: 3,
			PreChapterID:  0,
		},
		{
			ID:            3,
			Name:          "JavaScript-First",
			CourseID:      3,
			NextChapterID: 4,
			PreChapterID:  0,
		},
		{
			ID:            4,
			Name:          "HTML-Second",
			CourseID:      1,
			NextChapterID: 0,
			PreChapterID:  1,
		},
		{
			ID:            5,
			Name:          "CSS-Second",
			CourseID:      2,
			NextChapterID: 0,
			PreChapterID:  2,
		},
		{
			ID:            6,
			Name:          "JavaScript-Second",
			CourseID:      3,
			NextChapterID: 0,
			PreChapterID:  3,
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
			ID: 1,
			Detail: []*model.QuestionDetail{
				{
					ID:      1,
					Content: `<div>...</div>`,
					Answer:  `<div>this is div</div>`,
					Type:    model.HTML,
				},
			},
			Description:    "在div标签中输入this is div",
			NextQuestionID: 2,
			PreQuestionID:  0,
			ChapterID:      1,
		},
		{
			ID: 2,
			Detail: []*model.QuestionDetail{
				{
					ID:      2,
					Content: `<h1>...</h1>`,
					Answer:  `<h1>this is h1</h1>`,
					Type:    model.HTML,
				},
			},
			Description:    "在h1标签中输入this is h1",
			NextQuestionID: 3,
			PreQuestionID:  1,
			ChapterID:      1,
		},
		{
			ID: 3,
			Detail: []*model.QuestionDetail{
				{
					ID:      3,
					Content: `<p>...</p>`,
					Answer:  `<p>this is p</p>`,
					Type:    model.HTML,
				},
			},
			Description:    "在p标签中输入this is p",
			NextQuestionID: 0,
			PreQuestionID:  2,
			ChapterID:      1,
		},
		{
			ID: 4,
			Detail: []*model.QuestionDetail{
				{
					ID:      4,
					Content: `margin: ...`,
					Answer:  `margin: auto`,
					Type:    model.CSS,
				},
			},
			Description:    "设置margin为auto",
			NextQuestionID: 5,
			PreQuestionID:  0,
			ChapterID:      2,
		},
		{
			ID: 5,
			Detail: []*model.QuestionDetail{
				{
					ID:      5,
					Content: `padding: ...`,
					Answer:  `padding: 10px`,
					Type:    model.CSS,
				},
			},
			Description:    "设置padding为10px",
			NextQuestionID: 6,
			PreQuestionID:  4,
			ChapterID:      2,
		},
		{
			ID: 6,
			Detail: []*model.QuestionDetail{
				{
					ID:      6,
					Content: `background-color: ...`,
					Answer:  `background-color: red`,
					Type:    model.CSS,
				},
			},
			Description:    "设置background-color为red",
			NextQuestionID: 0,
			PreQuestionID:  5,
			ChapterID:      2,
		},
		{
			ID: 7,
			Detail: []*model.QuestionDetail{
				{
					ID:      7,
					Content: `console.log(...)`,
					Answer:  `console.log("hello world")`,
					Type:    model.JavaScript,
				},
			},
			Description:    "在控制台输出hello world",
			NextQuestionID: 8,
			PreQuestionID:  0,
			ChapterID:      3,
		},
		{
			ID: 8,
			Detail: []*model.QuestionDetail{
				{
					ID:      8,
					Content: `var a = ...`,
					Answer:  `var a = 1`,
					Type:    model.JavaScript,
				},
			},
			Description:    "声明一个变量a并赋值为1",
			NextQuestionID: 0,
			PreQuestionID:  7,
			ChapterID:      3,
		},
	}
	for _, question := range questions {
		err := db.Mysql.Create(question).Error
		if err != nil {
			panic(err)
		}
	}

}
