import React from 'react';

import Home from './home/home'
import YangHeader from "./header/header";

import './App.css';
import {Layout} from "antd";
import Chapter from "./chapter/chapter";
import {Link, Route, Routes} from "react-router-dom";

const { Header, Content, Footer } = Layout;



function App() {
  return (
    <div className="App">
        <Header>
            <Link to="/">
                <YangHeader></YangHeader>
            </Link>
        </Header>
        <Content>
           <Routes>
                <Route path="/" element={<Home/>}/>
                <Route path="/course/:courseId" element={<Chapter/>} loader={
                    ({params}) => {
                        params.courseId;
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
