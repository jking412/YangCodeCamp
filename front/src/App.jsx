import React from 'react';

import Home from './home/home'
import YangHeader from "./header/header";
import Chapter from "./chapter/chapter";
import Exam from './exam/Exam';
import './App.css';

import { Layout } from "antd";
import { Link, Route, Routes } from "react-router-dom";

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
                    <Route path="/" element={<Home />} />
                    <Route path="/course/:courseId" element={<Chapter />} loader={
                        ({ params }) => {
                            params.courseId;
                        }
                    } />
                    <Route path='/question/:questionId' element={<Exam/>} loader={
                        ({params}) =>{
                            params.questionId;
                        }
                        } />
                </Routes>
            </Content>
            <Footer
                style={{
                    textAlign: 'center',
                    backgroundColor: 'white',
                    fontFamily:'sans-serif',
                }}
            >

                &nbsp;&nbsp;&nbsp;Yangyang suports your Web learning
            </Footer>
        </div>
    );
}

export default App;
