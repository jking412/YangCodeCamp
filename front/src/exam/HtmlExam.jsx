import React, { useEffect, useState } from 'react';
import {useNavigate} from 'react-router-dom';
import './HtmlExam.css';
import { Layout, theme, Dropdown, Space, Avatar, Col, Row, Select, Button } from 'antd';
import { DownOutlined, UserOutlined, LeftOutlined } from '@ant-design/icons';
import '../request';
import { getQuestionsById, submitQuestion } from '../request';

const { Header, Content, Footer } = Layout;
const HtmlExam = ({questionId,courseId}) => {
    const {
        token: { colorBgContainer },
    } = theme.useToken();
    const [size, setSize] = useState('large');
    const [textareaValue, setTextareaValue] = useState('');
    const [description, setDescription] = useState('');
    const [content, setContent] = useState('');
    const [answer, setAnswer] = useState('');
    const [goalbackgroundColor, setGoalBackgroundColor] = useState('#CECECE');
    const [previewbackgroundColor, setPreviewBackgroundColor] = useState('#EFEFEF');
    const [showOption, setShowOption] = useState('Preview');
    const [nextQuestionId,setNextQuestionId] = useState();

    const [id,setId] = useState(questionId);


    const navigator = useNavigate();

    const handleTextareaChange = (event) => {
        setTextareaValue(event.target.value);
    };

    useEffect(() => {
        var resp = getQuestionsById(id);
        resp.then(r => {
            setDescription(r.data.details[0].description);
            setContent(r.data.details[0].content);
            setAnswer(r.data.details[0].answer);
            setTextareaValue(r.data.details[0].content);
            setNextQuestionId(r.data.next_question_id);
        })
    }, [id])

    const handleClick = () => {
        navigator(`/course/${courseId}`);
    }


    const Reset = () => {
        setTextareaValue(content);
    }

    const Next =()=>{
        if(nextQuestionId == 0){
            return
        }
        navigator(`/question/${nextQuestionId}`);
        setId(nextQuestionId);
    }

    const Goal = () => {
        setShowOption('Goal')
        setGoalBackgroundColor('#EFEFEF')
        setPreviewBackgroundColor('#CECECE')
    }

    const Preview = () => {
        setShowOption('Preview')
        setGoalBackgroundColor('#CECECE')
        setPreviewBackgroundColor('#EFEFEF')
    }

    const Push = () => {

        let obj = {
            "code": [
                {
                    "type": 0,
                    "content": textareaValue
                }
            ]
        }


        console.log(obj)

        // let objStr = JSON.stringify(obj)

        submitQuestion(questionId, obj).then(r => {
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
                    <div className='File'>
                        file.html
                    </div>
                </Col>
                <Col span={20}></Col>
            </Row>
            <Content
                className="site-layout"
            >
                <div className='Content'>
                    <div className='Exegesis' style={{ fontSize: 15 }}>{description}</div>
                    <textarea className='Code' value={textareaValue} onChange={handleTextareaChange} style={{ fontSize: 15 }}></textarea>
                    <div className='Button'>
                        <Button size={size} onClick={Reset}>
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
                <div className='View'>
                    <div className='Goal' onClick={Goal} style={{ backgroundColor: goalbackgroundColor }} >Goal</div>
                    <div className='Preview' onClick={Preview} style={{ backgroundColor: previewbackgroundColor }}>Preview</div>
                    {showOption === 'Preview' && <div className="previewFrame" style={{ width: '100%', height: '100%', paddingLeft: 10 }} dangerouslySetInnerHTML={{ __html: textareaValue }}></div>}
                    {showOption === 'Goal' && <div className="previewFrame" style={{ width: '100%', height: '100%', paddingLeft: 10 }} dangerouslySetInnerHTML={{ __html: answer }}></div>}
                </div>
            </Content>
        </Layout >
    );

};





export default HtmlExam;