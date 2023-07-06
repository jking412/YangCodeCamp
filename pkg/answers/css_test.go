package answers

import (
	"fmt"
	"github.com/aymerick/douceur/parser"
	"testing"
)

func Test_CSSParse(t *testing.T) {
	cssInput := `.class1{
    color: red;
}

#id1{
    color: blue;
}`
	css, _ := parser.Parse(cssInput)
	fmt.Println(css)
}

func TestCssQuestion_Check(t *testing.T) {
	//	cssInput := `.class1{
	//    color: red;
	//}
	//
	//#id1{
	//    color: blue;
	//}`
	//
	//	cssContent := `
	//#id1{
	//    color: blue;
	//}
	//.class1{
	//    color: red;
	//}
	//
	//`

}
