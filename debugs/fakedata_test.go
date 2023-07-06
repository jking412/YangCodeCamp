package main

import "testing"

func TestData(t *testing.T) {
	var a = `function solve(arr){\n    let result = new Array(arr.length)\n    for (let i = 0; i < arr.length; i++) {\n        if (arr[i] === null) {\n            result[i] = true\n        }else{\n            result[i] = false\n        }\n    }\n    return result\n}\n`
	var b = ``

	t.Log(a, b)
}
