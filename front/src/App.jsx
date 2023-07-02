import React from 'react';
import './App.css';
import { Input } from 'antd';
import { UserOutlined } from '@ant-design/icons';

function App() {
  return (
    <div className="App">
        <In />
      <Input placeholder="Basic usage" />;
      <header className="App-header">
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

function In () {
    return(

        <div>
            <Input size="large" placeholder="large size" prefix={<UserOutlined />} />
            <br />
            <br />
            <Input placeholder="default size" prefix={<UserOutlined />} />
            <br />
            <br />
            <Input size="small" placeholder="small size" prefix={<UserOutlined />} />
        </div>
    );
}

export default App;
