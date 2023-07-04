import { Dropdown, Space, Avatar } from 'antd';
import React from 'react';
import { DownOutlined, UserOutlined } from '@ant-design/icons';
import './header.css';


const YangHeader = () => {
    return (
        <div>
            <Space direction="vertical" size={16}>
                <Space wrap size={16} >
                    <Avatar className='Avatar' icon={<UserOutlined />} />

                </Space>

            </Space>

            <span className='Title'>YangCodeCamp</span>

            <Dropdown
                menu={{
                    items,
                }}
            >

                <a onClick={(e) => e.preventDefault()} className='Menu'>
                    <Space>
                        Profile
                        <DownOutlined />
                    </Space>
                </a>
            </Dropdown>
        </div>
    )
}

const items = [
    {
        label: (
            <a target="_blank" rel="noopener noreferrer" href="https://www.antgroup.com">
                Profile
            </a>
        ),
        key: '0',
    },
    {
        label: (
            <a target="_blank" rel="noopener noreferrer" href="https://www.aliyun.com">
                Settings
            </a>
        ),
        key: '1',
    },
    {
        label: (
            <a target="_blank" rel="noopener noreferrer" href="">
                Sign out
            </a>
        ),
        key: '2',
    },
];

export default YangHeader;