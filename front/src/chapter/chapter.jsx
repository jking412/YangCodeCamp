import React from 'react'
import { Layout } from 'antd'
import YangHeader from '../home/header'
import YangProgress from './progress'
import YangMenu from './menu'
import '../home/header'
import '../home/header.css'


const { Header, Content, Footer } = Layout;

const Chapter = () => {
    return (
        <Layout>

            <Header>
                <YangHeader></YangHeader>
            </Header>

            <Content style={{
                minHeight:600,
                backgroundColor:'white',
            }}>
                <YangProgress></YangProgress>
                <YangMenu></YangMenu>
            </Content>
            <Footer
                style={{
                    textAlign: 'center',
                    backgroundColor:'white',
                }}
            >
                
                羊羊助力你的Web学习
            </Footer>
        </Layout>
    );
};

export default Chapter;
