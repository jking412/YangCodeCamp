import React, { useEffect, useState,createRef } from 'react';
import './Radio.css';
import { Layout, theme, Dropdown, Space, Avatar, Col, Row, Radio, Input } from 'antd';
import { DownOutlined, UserOutlined, LeftOutlined } from '@ant-design/icons';
import '../request';
import { getQuestionsById, submitQuestion } from '../request';
import {useNavigate} from 'react-router-dom';
import ReactPlayer from 'react-player';

const { Header, Content, Footer } = Layout;
const QuestionRadio = ({questionId,courseId}) => {

    const navigator = useNavigate();
    const {
        token: { colorBgContainer },
    } = theme.useToken();
    const [size, setSize] = useState('large');
    const [question1, setQuestion1] = useState('');
    const [question2, setQuestion2] = useState('');
    const [question3, setQuestion3] = useState('');
    const [question4, setQuestion4] = useState('');
    const [value, setValue] = useState(1);
    const [description, setDescription] = useState('');
    const [videoSrc, setVideoSrc] = useState('');
    const [result, setResult] = useState('Click the button below to check your answer.');
    const [onclickChoice, setOnclickChoice] = useState(0);
    const [buttonContent, setButtonContent] = useState('Check your answer!');


    const onChange = (e) => {
        console.log('radio checked', e.target.value);
        setValue(e.target.value);
    };

    useEffect(() => {
        var resp = getQuestionsById(questionId);
        resp.then(r => {
            setQuestion1(r.data.details[0].content)
            setQuestion2(r.data.details[1].content)
            setQuestion3(r.data.details[2].content)
            setQuestion4(r.data.details[3].content)
            setDescription(r.data.details[4].description)
            setVideoSrc(r.data.details[4].content)
        })
    },[])
    console.log(videoSrc)

    const handleClick = () => {
        navigator(`/course/${courseId}`);
    }

    const Push = () => {

        if (onclickChoice == 0) {
            let obj = {
                "code": [
                    {
                        "type": 3,
                        "content": '' + value
                    }
                ]
            }



            // let objStr = JSON.stringify(obj)

            submitQuestion(questionId, obj).then(r => {
                if (r.data.result == 1) {
                    setOnclickChoice(1);
                    setButtonContent('Go to next challenge!');
                    setResult("Yes, you are right. Great!");
                }
                else {
                    setResult("Sorry, that's not the right answer. Give it another try?");
                }
            })
        }
        else {

        }
    }


    return (
        <Layout>
            <div className='Divider'></div>
            <Row className='Exam-Row'>
                <Col span={2} >
                    <div className='Back' onClick={handleClick}><LeftOutlined style={{ fontSize: 20 }} /></div>
                </Col>
                <Col span={2}>
                </Col>
                <Col span={20}></Col>
            </Row>
            <Content
                className="site-layout"
            >
                <div className='RadioContent'>
                    <div className='mainContent'>
                    <ReactPlayer url={videoSrc} controls={true} className='video'/>
                    
                        <div className='questionDescription' style={{ fontSize: 15 }}>{description}</div>
                        <div className='radio'>
                            <Radio.Group onChange={onChange} value={value} >
                                <Space direction="vertical">
                                    <Radio value={1}>{question1}</Radio>
                                    <Radio value={2}>{question2}</Radio>
                                    <Radio value={3}>{question3}</Radio>
                                    <Radio value={4}>{question4}</Radio>
                                </Space>
                            </Radio.Group>
                        </div>
                        <div className='result' style={{ fontSize: 15 }}>{result}</div>
                        <button className='post' onClick={Push}>{buttonContent}</button>
                    </div>
                </div>
            </Content>
        </Layout >
    );

};






export default QuestionRadio;