function solve(arr){
    let result = new Array(arr.length)
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] === null) {
            result[i] = true
        }else{
            result[i] = false
        }
    }
    return result
}

function main(){
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


main()