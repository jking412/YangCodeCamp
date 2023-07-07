import axios from 'axios';

var client = axios.create({
    baseURL: 'http://localhost:8000',
})

function getAllCourses(){
    return client.get('/course')
}

function getAllChaptersByCourseId(id){
    return client.get(`/course/${id}/chapter`)
}

function getChapterById(id){
    return client.get(`/chapter/${id}`)
}

function getAllQuestionsByChapterId(id){
    return client.get(`/chapter/${id}/question`)
}


function getQuestionsById(id) {
    return client.get('/question/' + id)
}

function submitQuestion(id, obj) {
    return client.put(`/question/${id}/submit`, obj, {
        headers: {
            'Content-Type': 'application/json'
        }
    })
}


export {
    getAllCourses,
    getAllChaptersByCourseId,
    getChapterById,
    getAllQuestionsByChapterId,
    getQuestionsById,
    submitQuestion,

};