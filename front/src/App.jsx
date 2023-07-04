import React from 'react';

import Home from './home/home'
import YangHeader from "./header/header";

import './App.css';
import {Layout} from "antd";
import Chapter from "./chapter/chapter";
import {Route, Routes} from "react-router-dom";

const { Header, Content, Footer } = Layout;



function App() {
  return (
    <div className="App">
        <Header>
            <YangHeader></YangHeader>
        </Header>
        <Content>
           <Routes>
                <Route path="/" element={<Home/>}/>
                <Route path="/chapter/:chapterId" element={<Chapter/>} loader={
                    ({params}) => {
                        params.chapterId;
                    }
                }/>
           </Routes>
        </Content>
        <Footer
            style={{
                textAlign: 'center',
                backgroundColor:'white',
            }}
        >

            羊羊助力你的Web学习
        </Footer>
    </div>
  );
}

export default App;
