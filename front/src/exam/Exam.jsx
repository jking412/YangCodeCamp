import React from "react";
import { useEffect, useState } from "react";
import { useParams } from 'react-router-dom'
import JsExam from './JsExam.jsx';
import HtmlExam from './HtmlExam.jsx';
import CssExam from './CssExam.jsx';
import QuestionRadio from './Radio.jsx';
import {getQuestionsById,getChapterById} from '../request.js';


const Exam = () => {
    const params = useParams();
    const [questionType, setQuestionType] = useState();
    const [questionId,setQuestionId] = useState();
    const [chapterId,setChapterId] = useState();
    const [courseId,setCourseId] = useState();

    useEffect(() => {
        getQuestionsById(params.questionId).then(r => {
            setQuestionType(r.data.type);
            setQuestionId(params.questionId);
            setChapterId(r.data.chapter_id);
        })
    },[])

    useEffect(() => {
        getChapterById(chapterId).then(r => {
            setCourseId(r.data.chapter.course_id);
        })
    },[chapterId])

    if (questionType == 0) {
        return <HtmlExam questionId={questionId} courseId={courseId}></HtmlExam>
    }
    else if (questionType == 1) {
        return <CssExam questionId={questionId} courseId={courseId}></CssExam>
    }
    else if (questionType == 2) {
        return <JsExam questionId={questionId} courseId={courseId}></JsExam>
    }
    else if (questionType == 3) {
        return <QuestionRadio questionId={questionId} courseId={courseId}></QuestionRadio>
    }

    
}

export default Exam;