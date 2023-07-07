import React, { useEffect } from 'react';
import {useNavigate} from 'react-router-dom';
// import {MailOutlined, SettingOutlined, AppstoreOutlined} from '@ant-design/icons';
import{CheckCircleOutlined,ClockCircleOutlined,MinusOutlined} from '@ant-design/icons';
import { Menu, Switch, Row, Col } from 'antd';
import { useState } from 'react';
import './menu.css'
import { getAllQuestionsByChapterId } from "../request";

function getItem(label, key, icon, children) {
    return {
        key,
        icon,
        children,
        label,
    };
}

//
// const items = [
//     getItem('Introduction', 'sub1', '', [
//         getItem('Option 1', '1'),
//         getItem('Option 2', '2'),
//         getItem('Option 3', '3'),
//         getItem('Option 4', '4'),
//     ]),
//     getItem('Navigation Two', 'sub2', '', [
//         getItem('Option 5', '5'),
//         getItem('Option 6', '6'),
//         getItem('Submenu', 'sub3', null, [getItem('Option 7', '7'), getItem('Option 8', '8')]),
//     ]),
//     getItem('Navigation Three', 'sub4', '', [
//         getItem('Option 9', '9'),
//         getItem('Option 10', '10'),
//         getItem('Option 11', '11'),
//         getItem('Option 12', '12'),
//     ]),
// ];

const YangMenu = ({ chapters }) => {
    const [theme, setTheme] = useState('dark');
    const [openKeys, setOpenKeys] = useState([]);
    const [items, setItems] = useState([]);

    const navigator = useNavigate();


    useEffect(() => {
        let newItems = [];
        for (let i = 0; i < chapters.length; i++) {
            let chapter = chapters[i];
            let tempItem = getItem(chapter.name, chapter.id, '', []);
            newItems.push(tempItem);
        }
        setItems(newItems);
    }, [chapters])

    const changeTheme = (value) => {
        setTheme(value ? 'dark' : 'light');
    };
    const handleClick = (e) => {
        let id = e.key / 10000;
        navigator(`/question/${id}`)
    };

    const Unfinished = 0;
    const Success = 1;
    const Ondoing = 2;

    const handleOpenChange = (e) => {
        let id;
        if (e.length > openKeys.length) {
            for (let i = 0; i < e.length; i++) {
                if (!openKeys.includes(e[i])) {
                    id = e[i];
                }
            }
            getAllQuestionsByChapterId(id).then((res) => {
                let newItems = [...items]
                let childrenItems = [];
                for (let i = 0; i < res.data.questions.length; i++) {
                    let question = res.data.questions[i];
                    let tempItem;
                    if(res.data.questions[i].status == Unfinished){
                        tempItem = getItem(question.name, question.id * 10000, <MinusOutlined />);
                    }
                    else if(res.data.questions[i].status == Success){
                        tempItem = getItem(question.name, question.id * 10000, <CheckCircleOutlined />);
                    }
                    else if(res.data.questions[i].status == Ondoing){
                        tempItem = getItem(question.name, question.id * 10000, <ClockCircleOutlined />);
                    }

                    childrenItems.push(tempItem);
                }
                for (let i = 0; i < newItems.length; i++) {
                    if (newItems[i].key == e) {
                        newItems[i].children = childrenItems;
                    }
                }
                console.log(newItems)
                setItems(newItems);
                setOpenKeys(e);
            })
            return
        }
        setOpenKeys(e);
    }


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
                            onClick={handleClick}
                            onOpenChange={handleOpenChange}
                            onSelect={handleClick}
                            openKeys={openKeys}
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