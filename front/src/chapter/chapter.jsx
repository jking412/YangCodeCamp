import React, {useEffect, useState} from 'react'
import YangProgress from './progress'
import YangMenu from './menu'
import { useParams} from "react-router-dom";
import {getAllChaptersByCourseId} from "../request";


const Chapter = () => {

    const params = useParams();
    const [chapters,setChapters] = useState([]);
    const [totalQuestion,setTotalQuestion] = useState(0);
    const [finishedQuestion,setFinishedQuestion] = useState(0);



    useEffect(() => {
        getAllChaptersByCourseId(params.courseId).then((res) => {
            setChapters(res.data.chapters);
            let total_question = 0;
            let finished_question = 0;

            for (let i = 0; i < res.data.chapters.length ; i++) {
                total_question += res.data.chapters[i].total_question;
                finished_question += res.data.chapters[i].finished_question;
            }

            setTotalQuestion(total_question);
            setFinishedQuestion(finished_question);
        })
    },[])


    return (
        <div style={{
            minHeight: 600,
            backgroundColor: 'white',
        }}>
            <YangProgress totalQuestion={totalQuestion} finishedQuestion={finishedQuestion}></YangProgress>
            <YangMenu chapters={chapters}></YangMenu>
        </div>
    );
};

export default Chapter;
