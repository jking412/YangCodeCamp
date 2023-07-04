import React, { useState } from 'react';
import { Form, Row, Col, Select ,Input} from 'antd';

const { Option } = Select;

const YangSearch = () => {
  const [form] = Form.useForm();
    const [value,setValue] = useState('');
  const options = (
    <Option key="1" value="asd">
      asd
    </Option>
  );

  function handleChange() {}

  function handleSearch() {}


  return (
    <Form form={form} layout="vertical">
      <Row>
        <Col lg={6} md={12} sm={24}>
          <Form.Item>
            <Input placeholder="Basic usage" />;
            </Form.Item>
            <Form.Item>
            <Select
              showSearch
              value={value}
              style={{ width: '100%' }}
              placeholder="输入名称查询"
              onChange={handleChange}
              defaultActiveFirstOption={false}
              showArrow={false}
              filterOption={false}
              onSearch={handleSearch}
              notFoundContent={null}
            >
              {options}
            </Select>
          </Form.Item>
        </Col>
      </Row>
    </Form>
  );
};

export default YangSearch;
