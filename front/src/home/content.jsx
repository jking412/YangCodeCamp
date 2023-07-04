import { Col, Row, theme } from 'antd'
import React, { useEffect, useState } from 'react';
import { getAllCourses } from '../request';
import './content.css';


const YangContent = () => {

    const {
        token: { colorBgContainer },
    } = theme.useToken();

    const [Course, setCourse] = useState([]);

    const listItems = Course.map(course =>
        <div key={course.id} >
            <Row className='button'>
                <Col span={4}></Col>
                <Col span={2}><i class={course.icon} 
                style={{
                    fontSize: '30px',
                    marginTop:'20px',
                }}></i></Col>
                <Col span={8}><div className='name'>{course.name}</div></Col>
                <Col span={10}></Col>
            </Row>
        </div>
    );

    useEffect(() => {
        var resp = getAllCourses()
        console.log(resp)
        resp.then(r => {
            setCourse(r.data.courses)
        })
    },[])
    return (
        <div>
            <Row style={{
                minHeight: 100,
            }}>
                <Col span={6}></Col>
                <Col span={12} className='title'>
                    <div className='Course'>Course Selection</div>
                    <div className='Description'>Here you can choose from multiple courses (e.g.HTML,CSS,JS) to practice</div>
                </Col>
                <Col span={6}></Col>
            </Row>
            <Row style={{
                minHeight: 450,
                background: colorBgContainer,
            }}
            >
                <Col span={6}></Col>
                <Col span={12}>
                    <ul>{listItems}</ul>
                </Col>
                <Col span={6}></Col>
            </Row>
        </div>
    )
}


export default YangContent;
