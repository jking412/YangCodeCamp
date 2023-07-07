import React, { useEffect, useState } from 'react';
import './CssExam.css';
import { Layout, theme, Dropdown, Space, Avatar, Col, Row, Select, Button } from 'antd';
import { DownOutlined, UserOutlined, LeftOutlined } from '@ant-design/icons';
import { getQuestionsById, submitQuestion } from '../request';
import {useNavigate} from 'react-router-dom';

const { Header, Content, Footer } = Layout;
const CssExam = ({questionId,courseId}) => {

    const navigator = useNavigate();

    const {
        token: { colorBgContainer },
    } = theme.useToken();
    const [size, setSize] = useState('large');
    const [html_textareaValue, sethtml_TextareaValue] = useState('');
    const [css_textareaValue, setcss_TextareaValue] = useState('');
    const [textareaValue, setTextareaValue] = useState('');
    const [html_description, sethtml_Description] = useState('');
    const [css_description, setcss_Description] = useState('');
    const [html_content, sethtml_Content] = useState('');
    const [css_content, setcss_Content] = useState('');
    const [answer, setAnswer] = useState('');
    const [goalbackgroundColor, setGoalBackgroundColor] = useState('#CECECE');
    const [previewbackgroundColor, setPreviewBackgroundColor] = useState('#EFEFEF');
    const [showOption, setShowOption] = useState('Preview');

    const [nextQuestionId,setNextQuestionId] = useState();
    const [id,setId] = useState(questionId);


    const CssOptions = [
        {
            value: 'file.html',
            label: 'file.html',
        },
        {
            value: 'file.css',
            label: 'file.css',
        },
    ]

    const [selectedOption, setSelectedOption] = useState(CssOptions[0].value);

    const html_handleTextareaChange = (event) => {
        sethtml_TextareaValue(event.target.value);
        setTextareaValue('<style>' + css_textareaValue + '</style>' + event.target.value);
    };

    const css_handleTextareaChange = (event) => {
        setcss_TextareaValue(event.target.value);
        setTextareaValue('<style>' + event.target.value + '</style>' + html_textareaValue);
    };

    const handleChange = (selectedOption) => {
        setSelectedOption(selectedOption);
    };



    useEffect(() => {
        var resp = getQuestionsById(id);
        resp.then(r => {
            // if (r.data.details[0].type == 0) (
            //     sethtml_Description(r.data.details[0].description),
            //     sethtml_Content(r.data.details[0].content)
            // )
            // if (r.data.details[1].type == 1) (
            //     setcss_Description(r.data.details[1].description),
            //     setcss_Content(r.data.details[1].content)
            // )
            sethtml_Description(r.data.details[0].description);
            sethtml_Content(r.data.details[0].content);
            setcss_Description(r.data.details[0].description);
            setcss_Content(r.data.details[1].content);
            sethtml_TextareaValue(r.data.details[0].content);
            setcss_TextareaValue(r.data.details[1].content);
            setAnswer('<style>' + r.data.details[1].answer + '</style>' + r.data.details[0].answer)
            setNextQuestionId(r.data.next_question_id);
        })
    }, [id])

    const handleClick = () => {
        navigator(`/course/${courseId}`);
    }

    const Reset = () => {
        sethtml_TextareaValue(html_content);
        setcss_TextareaValue(css_content);
    }

    const Next =()=>{
        if(nextQuestionId==0){
            return;
        }
        navigator(`/question/${nextQuestionId}`);
        setId(nextQuestionId);
    }

    const Goal = () => {
        setShowOption('Goal');
        setGoalBackgroundColor('#EFEFEF');
        setPreviewBackgroundColor('#CECECE');
    }

    const Preview = () => {
        setShowOption('Preview');
        setGoalBackgroundColor('#CECECE');
        setPreviewBackgroundColor('#EFEFEF');
    }

    const Push = () => {

        let obj = {
            "code": [
                {
                    "type": 0,
                    "content": html_textareaValue
                },
                {
                    "type": 1,
                    "content": css_textareaValue
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
                    <div className='CssFile'>
                        <sapn>
                            <Select
                                value={selectedOption}

                                onChange={handleChange}
                                options={CssOptions}
                                defaultValue="file.html"
                                style={{
                                    width: 120,
                                }}

                            />
                        </sapn>
                    </div>
                </Col>
                <Col span={20}></Col>
                {/* <Col span={2}>
                    <div className='Goal'><sapn style={{ marginLeft: 20 }}>Goal</sapn></div>
                </Col>
                <Col span={2}>
                    <div className='Preview'><span style={{ marginLeft: 12 }}>Preview</span></div>
                </Col> */}
                {/* <Col span={1}></Col> */}
            </Row>
            <Content
                className="site-layout"
            >
                <div className='Content'>
                    {selectedOption === 'file.html' && <div className='Exegesis' style={{ fontSize: 15 }}>{html_description}</div>}
                    {selectedOption === 'file.html' && <textarea className='Code' value={html_textareaValue} style={{ fontSize: 15 }} onChange={html_handleTextareaChange}></textarea>}
                    {selectedOption === 'file.css' && <div className='Exegesis' style={{ fontSize: 15 }}>{css_description}</div>}
                    {selectedOption === 'file.css' && <textarea className='Code' value={css_textareaValue} style={{ fontSize: 15 }} onChange={css_handleTextareaChange}></textarea>}
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
                    <div className='Goal' onClick={Goal} style={{ backgroundColor: goalbackgroundColor }}>Goal</div>
                    <div className='Preview' onClick={Preview} style={{ backgroundColor: previewbackgroundColor }}>Preview</div>
                    {showOption === 'Preview' && <div className="previewFrame" style={{ width: '100%', height: '100%', paddingLeft: 10, fontSize: 15 }} dangerouslySetInnerHTML={{ __html: textareaValue }}></div>}
                    {showOption === 'Goal' && <div className="previewFrame" style={{ width: '100%', height: '100%', paddingLeft: 10, fontSize: 15 }} dangerouslySetInnerHTML={{ __html: answer }}></div>}
                </div>
            </Content>
        </Layout >
    );

};




export default CssExam;