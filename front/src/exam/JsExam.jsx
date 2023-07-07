import React, { useEffect, useState } from 'react';
import './JsExam.css';
import { Layout, theme, Dropdown, Space, Avatar, Col, Row, Select, Button } from 'antd';
import { DownOutlined, UserOutlined, LeftOutlined } from '@ant-design/icons';
import { getQuestionsById, submitQuestion } from '../request';
import {useNavigate} from 'react-router-dom';

const { Header, Content, Footer } = Layout;
const JsExam = ({questionId,courseId}) => {
const [nextQuestionId,setNextQuestionId] = useState();
const [id,setId] = useState(questionId);

    const navigator = useNavigate();


    const {
        token: { colorBgContainer },
    } = theme.useToken();
    const [size, setSize] = useState('large');
    const [description, setDescription] = useState('');
    const [textareaValue, setTextareaValue] = useState('');
    const [content, setContent] = useState('');
    const [answer, setAnswer] = useState('');
    const [resultContent,setResultContent] = useState('');


    const handleTextareaChange = (event) => {
        setTextareaValue(event.target.value);
    };


    useEffect(() => {
        var resp = getQuestionsById(id);
        resp.then(r => {
            setDescription(r.data.details[0].description)
            setContent(r.data.details[0].content)
            setAnswer(r.data.details[0].answer)
            setTextareaValue(r.data.details[0].content)
            setNextQuestionId(r.data.next_question_id);
        })
    }, [id])

    const handleClick = () => {
        navigator(`/course/${courseId}`);
        console.log('ok');
        setDescription('');
    }

    const Reset = () => {
        setTextareaValue(content);
    }

    const Next =()=>{
        if(nextQuestionId==0){
            return;
        }
        navigator(`/question/${nextQuestionId}`);
        setId(nextQuestionId);
        }

    const Push = () => {

        let obj = {
            "code": [
                {
                    "type": 2,
                    "content": textareaValue
                }
            ]
        }


        console.log(obj)
        // let objStr = JSON.stringify(obj)

        submitQuestion(questionId, obj).then(r => {
            setResultContent(r.data.content);
            if (r.data.result == 1) {
                alert('good!');
            }
            else {
                alert('error!')
            }
        })
    }

    return (
        <Layout>
            <div className='Divider'></div>
            <Row className='Exam-Row'>
                <Col span={2}>
                    <div className='Back' onClick={handleClick}><LeftOutlined style={{ fontSize: 20 }} /></div>
                </Col>
                <Col span={2}>
                    <div className='JsFile'>file.js
                    </div>
                </Col>
                <Col span={20}></Col>
            </Row>
            <Content
                className="site-layout"
            >
                <div className='JsContent'>
                    <div className='JsExegesis' style={{ fontSize: 15 }}>{description}</div>
                    <div className='Button'>
                        <Button size={size} onClick={Reset} >
                            Reset
                        </Button>
                        <Button size={size} style={{ marginLeft: 10 }} onClick={Push}>
                            Run your code !
                        </Button>
                        <Button size={size} onClick={Next} style={{marginLeft:350}}>
                            Next Question
                        </Button>
                    </div>
                </div>
                <div className='JsCode'>
                    <textarea className="Js_Code" value={textareaValue} onChange={handleTextareaChange} style={{ fontSize: 15 }}></textarea>
                <div className="console">{resultContent}</div>
                </div>
            </Content>
        </Layout >
    );

};






export default JsExam;