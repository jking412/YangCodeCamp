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

function getAllQuestionsByChapterId(id){
    return client.get(`/chapter/${id}/question`)
}


export {
    getAllCourses,
    getAllChaptersByCourseId,
    getAllQuestionsByChapterId
};