import React from 'react';
import {Layout} from 'antd';
import YangHeader from './header';
import YangContent from './content';
import './home.css';
import '../request';


const { Header, Content, Footer } = Layout;


const Home = () => { 
  return (
    <Layout>

      <Header className='Home-Header'>
        <YangHeader></YangHeader>
      </Header>

      <Content className="site-layout">
        <YangContent></YangContent>
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



export default Home;