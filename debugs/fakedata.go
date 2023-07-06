package main

import (
	"YangCodeCamp/db"
	"YangCodeCamp/model"
	"YangCodeCamp/pkg/config"
)

func main() {
	config.InitConfig("config")
	db.Init()
	db.Mysql.Exec("delete from question_details")
	db.Mysql.Exec("delete from questions")
	db.Mysql.Exec("delete from chapters")
	db.Mysql.Exec("delete from courses")
	generateClass()
	//generateChapter()
	//generateQuestion()
}

func generateClass() {
	courses := make([]*model.Course, 0)

	{
		htmlCourse := &model.Course{
			Name: "Learn HTML by building simple web pages",
		}

		chapters := make([]*model.Chapter, 0)

		htmlChapter1 := &model.Chapter{
			Name: "Learn Basic HTML By Building a Cat Photo App",
		}

		Chapter1Questions := make([]*model.Question, 0)

		Chapter1Question1 := &model.Question{
			ID:             111,
			NextQuestionID: 112,
			Name:           "study to learn \"h1\" tag",
			Type:           model.HTML,
			Detail: []*model.QuestionDetail{
				{
					Type: model.HTML,
					Description: `Headings are <h1> defined by - <h6> tags.

<h1> Defines the largest title. <h6> Define a minimal title.

Try adding the h1 tag below`,
					Content: `<html>
    <body>
<!--        给下面的内容添加h1-->
    CatPhoto
    </body>
</html>`,
					Answer: `<html>
    <body>
<!--        给下面的内容添加h1-->
    <h1>CatPhoto</h1>
    </body>
</html>`,
				},
			},
		}

		Chapter1Question2 := &model.Question{
			ID:             112,
			NextQuestionID: 113,
			Name:           "study to learn \"p\" tag",
			Type:           model.HTML,
			Detail: []*model.QuestionDetail{
				{
					Type: model.HTML,
					Description: `Paragraphs are <p> defined by tags.

Try adding the p tag to the following`,
					Content: `<html>
    <body>
    <h1>CatPhoto</h1>
<!--    给下面的内容添加一个p标签-->
    this is cute cat
    </body>
</html>`,
					Answer: `<html>
    <body>
    <h1>CatPhoto</h1>
<!--    给下面的内容添加一个p标签-->
    <p>this is cute cat</p>
    </body>
</html>`,
				},
			},
		}

		Chapter1Question3 := &model.Question{
			ID:             113,
			NextQuestionID: 114,
			Name:           "study to learn \"img\" tag",
			Type:           model.HTML,
			Detail: []*model.QuestionDetail{
				{
					Type: model.HTML,
					Description: `In HTML, an image is defined by<img> tags.

<img> is an empty label, meaning that it contains only attributes and no closed label.

To display an image on a page, you need to use the source attribute (src). src means "source". The value of the source property is the URL address of the image.

Now let's get a picture of a kitten`,
					Content: `<html>
    <body>
    <h1>CatPhoto</h1>
    <p>this is cute cat</p>
<!--    添加链接https://cdn.freecodecamp.org/curriculum/cat-photo-app/relaxing-cat.jpg-->
    <img src="">
    </body>
</html>`,
					Answer: `<html>
    <body>
    <h1>CatPhoto</h1>
    <p>this is cute cat</p>
<!--    添加链接https://cdn.freecodecamp.org/curriculum/cat-photo-app/relaxing-cat.jpg-->
    <img src="https://cdn.freecodecamp.org/curriculum/cat-photo-app/relaxing-cat.jpg">
    </body>
</html>`,
				},
			},
		}

		Chapter1Question4 := &model.Question{
			ID:             114,
			NextQuestionID: 0,
			Name:           "study to learn \"ul\" and \"li\" tag",
			Type:           model.HTML,
			Detail: []*model.QuestionDetail{
				{
					Type: model.HTML,
					Description: `An unordered list is a list of items marked with bold dots (typically small black circles).
Use <ul>and <li> to add features that make kittens`,
					Content: `<html>
    <body>
    <h1>CatPhoto</h1>
    <p>this is cute cat</p>
    <img src="https://cdn.freecodecamp.org/curriculum/cat-photo-app/relaxing-cat.jpg">
<!--    添加ul和li标签，以无序列表的形式显示以下内容：-->
    cute
    adorable
    docile
    </body>
</html>`,
					Answer: `<html>
    <body>
    <h1>CatPhoto</h1>
    <p>this is cute cat</p>
    <img src="https://cdn.freecodecamp.org/curriculum/cat-photo-app/relaxing-cat.jpg">
<!--    添加ul和li标签，以无序列表的形式显示以下内容：-->
    <ul>
        <li>cute</li>
        <li>adorable</li>
        <li>docile</li>
    </ul>
    </body>
</html>`,
				},
			},
		}

		Chapter1Questions = append(Chapter1Questions, Chapter1Question1)
		Chapter1Questions = append(Chapter1Questions, Chapter1Question2)
		Chapter1Questions = append(Chapter1Questions, Chapter1Question3)
		Chapter1Questions = append(Chapter1Questions, Chapter1Question4)
		htmlChapter1.Question = Chapter1Questions

		htmlChapter2 := &model.Chapter{
			Name: "Learn HTML Forms By Building a Registration Form",
		}

		htmlChapter3 := &model.Chapter{
			Name: "Learn HTML5 Elements By Building a Text Adventure Game",
		}

		chapters = append(chapters, htmlChapter1)
		chapters = append(chapters, htmlChapter2)
		chapters = append(chapters, htmlChapter3)
		htmlCourse.Chapters = chapters
		courses = append(courses, htmlCourse)
	}

	// css practice
	{
		cssCourse := &model.Course{
			Name: "Learn CSS By Building simple web page",
		}

		chapters := make([]*model.Chapter, 0)

		cssChapter1 := &model.Chapter{
			Name: "Learn CSS Selectors By Building a Set of Color Markers",
		}

		Chapter1Questions := make([]*model.Question, 0)

		Chapter1Question1 := &model.Question{
			ID:             211,
			NextQuestionID: 212,
			Name:           "study to learn how to add css",
			Type:           model.CSS,
			Detail: []*model.QuestionDetail{
				{
					Type: model.HTML,
					Description: `CSS (Cascading Style Sheets) is a computer language used to add styles (fonts, spacing, colors, etc.) to structured documents, such as HTML documents or XML applications

Now follow the prompts to add CSS`,
					Content: `<html lang="en">
<body>
<!--将下面的内容居中-->
<h1>CSS Color Markers</h1>
</body>
</html>`,
					Answer: `<html lang="en">
<body>
<!--将下面的内容居中-->
<h1>CSS Color Markers</h1>
</body>
</html>`,
				},
				{
					Type: model.CSS,
					Content: `/*按照下面的格式为html中的内容居中*/
/*p {*/
/*    text-align: right;*/
/*}*/
`,
					Answer: `/*按照下面的格式为html中的内容居中*/
/*p {*/
/*    text-align: right;*/
/*}*/

h1 {
    text-align: center;
}`,
				},
			},
		}

		Chapter1Question2 := &model.Question{
			ID:             212,
			NextQuestionID: 213,
			Name:           "study to learn class selector",
			Type:           model.CSS,
			Detail: []*model.QuestionDetail{
				{
					Type: model.HTML,
					Description: `The class selector is used to specify a style for a group of elements. 
Now follow the prompts to add class selector`,
					Content: `<html lang="en">
<body>
<h1>CSS Color Markers</h1>
<!--将下面的div的class设为color-marker，然后在css中添加color-marker的样式，就可以看到效果了。-->
<div>this is red</div>
</body>
</html>`,
					Answer: `<html lang="en">
<body>
<h1>CSS Color Markers</h1>
<!--将下面的div的class设为color-marker，然后在css中添加color-marker的样式，就可以看到效果了。-->
<div class="color-marker">this is red</div>
</body>
</html>`,
				},
				{
					Type: model.CSS,
					Content: `h1 {
    text-align: center;
}

/*把class为color-marker的background-color设置为red*/
`,
					Answer: `h1 {
    text-align: center;
}

/*把class为color-marker的background-color设置为red*/
.color-marker{
    background-color: red;
}`,
				},
			},
		}

		Chapter1Question3 := &model.Question{
			ID:             213,
			NextQuestionID: 0,
			Name:           "study to learn id selector",
			Type:           model.CSS,
			Detail: []*model.QuestionDetail{
				{
					Type: model.HTML,
					Description: `The id selector is used to specify a style for a single, unique element.
now follow the prompts to add id selector`,
					Content: `<html lang="en">
<body>
<h1>CSS Color Markers</h1>
<!--将下面的div的class设为color-marker，然后在css中添加color-marker的样式，就可以看到效果了。-->
<div class="color-marker">this is red</div>
<!--将下面的div的id设为blue-marker，然后在css中添加blue-marker的样式，就可以看到效果了。-->
<div>this is blue</div>
</body>
</html>`,
					Answer: `<html lang="en">
<body>
<h1>CSS Color Markers</h1>
<!--将下面的div的class设为color-marker，然后在css中添加color-marker的样式，就可以看到效果了。-->
<div class="color-marker">this is red</div>
<!--将下面的div的id设为blue-marker，然后在css中添加blue-marker的样式，就可以看到效果了。-->
<div id="blue-marker">this is blue</div>
</body>
</html>`,
				},
				{
					Type: model.CSS,
					Content: `h1 {
    text-align: center;
}

.color-marker{
    background-color: red;
}

/*将id为blue-marker的元素的背景色设置为蓝色*/`,
					Answer: `h1 {
    text-align: center;
}

.color-marker{
    background-color: red;
}

/*将id为blue-marker的元素的背景色设置为蓝色*/
#blue-marker{
    background-color: blue;
}`,
				},
			},
		}

		Chapter1Questions = append(Chapter1Questions, Chapter1Question1)
		Chapter1Questions = append(Chapter1Questions, Chapter1Question2)
		Chapter1Questions = append(Chapter1Questions, Chapter1Question3)
		cssChapter1.Question = Chapter1Questions

		cssChapter2 := &model.Chapter{
			Name: "Learn CSS Variables By Building a Cafe Menu",
		}

		cssChapter3 := &model.Chapter{
			Name: "Learn CSS flex By Building Personal Website",
		}

		chapters = append(chapters, cssChapter1)
		chapters = append(chapters, cssChapter2)
		chapters = append(chapters, cssChapter3)
		cssCourse.Chapters = chapters
		courses = append(courses, cssCourse)
	}

	// js practice
	{
		jsCourse := &model.Course{
			Name: "Learn JavaScript by practice",
		}

		chapters := make([]*model.Chapter, 0)

		jsChapter1 := &model.Chapter{
			Name: "Learn Basic JavaScript",
		}

		Chapter1Questions := make([]*model.Question, 0)

		Chapter1Question1 := &model.Question{
			ID:             311,
			NextQuestionID: 0,
			Name:           "Null detection",
			Type:           model.JavaScript,
			Detail: []*model.QuestionDetail{
				{
					Type: model.JavaScript,
					Description: `请你编写一个函数，用于帮助开发人员测试他们的代码。它应该接受一个数组arr并返回一个数组result。
如果arr[i]==null，那么result[i]=true.
如果arr[i]！=null，那么result[i]=false.
example:
arr:[
  1,    1,    null, null,
  1,    null, 1,    1,
  null, null
]
result:[
  false, false, true,
  true,  false, true,
  false, false, true,
  true
]
`,
					Content: `function solve(arr){

}`,
					CheckMessage: `function main(){
    let arr = new Array(10)
    let expect = new Array(10)
    // 随机填充某些元素为空
    for(let i = 0; i < arr.length; i++){
        if(Math.random() > 0.5) {
            arr[i] = null
            expect[i] = true
        }else{
            arr[i] = 1
            expect[i] = false
        }
    }
    let result = solve(arr)
    let resultFlag = true
    for(let i = 0; i < arr.length; i++){
        if(result[i] !== expect[i]){
            resultFlag = false
            break
        }
    }
    if(resultFlag === true){
        console.log(resultFlag)
    }else{
        console.log('input:')
        console.log(arr)
        console.log('expect:')
        console.log(expect)
        console.log('result:')
        console.log(result)
    }
}


main()`,
					Answer: "true",
				},
			},
		}

		Chapter1Questions = append(Chapter1Questions, Chapter1Question1)
		jsChapter1.Question = Chapter1Questions

		jsChapter2 := &model.Chapter{
			Name: "Learn JavaScript Algorithm",
		}

		chapters = append(chapters, jsChapter1)
		chapters = append(chapters, jsChapter2)
		jsCourse.Chapters = chapters
		courses = append(courses, jsCourse)
	}

	// python choice
	{
		pythonCourse := &model.Course{
			Name: "Learn Python by doing rich choice questions",
		}

		chapters := make([]*model.Chapter, 0)

		pythonChapter1 := &model.Chapter{
			Name: "Learn Basic Python",
		}

		Chapter1Questions := make([]*model.Question, 0)

		Chapter1Question1 := &model.Question{
			ID:             411,
			NextQuestionID: 412,
			Name:           "function",
			Type:           model.PythonChoice,
			Detail: []*model.QuestionDetail{
				{
					Type:    model.PythonChoice,
					Content: `It is slang that means "The following code is really cool."`,
				},
				{
					Type:    model.PythonChoice,
					Content: `It indicates the start of a function.`,
				},
				{
					Type:    model.PythonChoice,
					Content: `It indicates that the following indented section of code is to be stored for later.`,
				},
				{
					Type:    model.PythonChoice,
					Content: `It indicates the start of a function, and the following indented section of code is to be stored for later.`,
				},
				{
					Type:        model.Other,
					Description: `What is the purpose of the "def" keyword in Python?`,
					Content:     "link",
					Answer:      "4",
				},
			},
		}

		Chapter1Question2 := &model.Question{
			ID:             412,
			NextQuestionID: 413,
			Name:           "loops and iterations",
			Type:           model.PythonChoice,
			Detail: []*model.QuestionDetail{
				{
					Type: model.PythonChoice,
					Content: `0
1 
2`,
				},
				{
					Type: model.PythonChoice,
					Content: `0
1
2
3 `,
				},
				{
					Type: model.PythonChoice,
					Content: `1
2`,
				},
				{
					Type: model.PythonChoice,
					Content: `1
2
3`,
				},
				{
					Type: model.Other,
					Description: `What will the following code print out?:

1 n = 0
2 while True:
3     if n == 3:
4         break
5     print(n)
6     n = n + 1`,
					Content: `link`,
					Answer:  "1",
				},
			},
		}

		Chapter1Question3 := &model.Question{
			ID:             413,
			NextQuestionID: 0,
			Name:           "Python lists",
			Type:           model.PythonChoice,
			Detail: []*model.QuestionDetail{
				{
					Type:    model.PythonChoice,
					Content: `banana`,
				},
				{
					Type:    model.PythonChoice,
					Content: `a`,
				},
				{
					Type:    model.PythonChoice,
					Content: `b`,
				},
				{
					Type:    model.PythonChoice,
					Content: `True`,
				},
				{
					Type: model.Other,
					Description: `What is the value of x after running this code:

1 fruit = "banana"
2 x = fruit[1]`,
					Content: `link`,
					Answer:  `2`,
				},
			},
		}

		Chapter1Questions = append(Chapter1Questions, Chapter1Question1)
		Chapter1Questions = append(Chapter1Questions, Chapter1Question2)
		Chapter1Questions = append(Chapter1Questions, Chapter1Question3)
		pythonChapter1.Question = Chapter1Questions

		pythonChapter2 := &model.Chapter{
			Name: "Learn Advanced Python",
		}

		pythonChapter3 := &model.Chapter{
			Name: "Learn Python by developing a web application",
		}

		chapters = append(chapters, pythonChapter1)
		chapters = append(chapters, pythonChapter2)
		chapters = append(chapters, pythonChapter3)
		pythonCourse.Chapters = chapters
		courses = append(courses, pythonCourse)
	}

	err := db.Mysql.Create(&courses).Error
	if err != nil {
		panic(err)
	}
}
