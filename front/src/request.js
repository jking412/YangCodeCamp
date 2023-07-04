import axios from 'axios';

var client = axios.create({
    baseURL: 'http://localhost:8000',
})

function getAllCourses(){
    return client.get('/course')
}



export {
    getAllCourses,
};