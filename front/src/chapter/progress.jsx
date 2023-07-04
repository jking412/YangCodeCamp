import React from 'react';
import { Progress, Row, Col } from 'antd';
import './progress.css'


const YangProgress = () => {
    return (
        <div>
            <Row>
                <Col span={4}></Col>
                <Col span={16}>
                <Progress className='progress' percent={50} format={(percent) => `${percent}/100`} status="active"
                    strokeColor={{
                        '0%': '#108ee9',
                        '100%': '#87d068',
                    }} />
                </Col>
                <Col span={4}></Col>
            </Row>
        </div>
    )
}

export default YangProgress;