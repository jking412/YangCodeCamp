import React from 'react';
import { MailOutlined, SettingOutlined, AppstoreOutlined } from '@ant-design/icons';
import { Menu, Switch, Row, Col } from 'antd';
import { useState } from 'react';
import './menu.css'

function getItem(label, id, icon, children, type) {
    return {
        id,
        icon,
        children,
        label,
        type,
    };
}

const items = [
    getItem('Introduction', 'sub1', <MailOutlined />, [
        getItem('Option 1', '1'),
        getItem('Option 2', '2'),
        getItem('Option 3', '3'),
        getItem('Option 4', '4'),
    ]),
    getItem('Navigation Two', 'sub2', <AppstoreOutlined />, [
        getItem('Option 5', '5'),
        getItem('Option 6', '6'),
        getItem('Submenu', 'sub3', null, [getItem('Option 7', '7'), getItem('Option 8', '8')]),
    ]),
    getItem('Navigation Three', 'sub4', <SettingOutlined />, [
        getItem('Option 9', '9'),
        getItem('Option 10', '10'),
        getItem('Option 11', '11'),
        getItem('Option 12', '12'),
    ]),
];
const YangMenu = () => {
    const [theme, setTheme] = useState('dark');
    const [current, setCurrent] = useState('1');
    // const [questions,setQuestions] = useState([]);
    const changeTheme = (value) => {
        setTheme(value ? 'dark' : 'light');
    };
    const onClick = (e) => {
        console.log('click ', e);
        setCurrent(e.id);
    };
    return (
        <Row>
            <Col span={4}></Col>
            <Col span={16}>
                <div>
                    <Switch
                        checked={theme === 'dark'}
                        onChange={changeTheme}
                        checkedChildren="Dark"
                        unCheckedChildren="Light"
                    />
                    <br />
                    <br />
                    <div>
                        <Menu
                            theme={theme}
                            onClick={onClick}
                            style={{
                                width: 920,
                            }}
                            selectedKeys={[current]}
                            mode="inline"
                            items={items}
                        />
                    </div>
                </div>
            </Col>
            <Col span={4}></Col>
        </Row>
    );
};

export default YangMenu;