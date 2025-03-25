import React, { useState } from 'react';
import { Form, Input, Button, Modal } from 'antd';
import { ethers } from 'ethers';

export default function AdminPanel({ contract }) {
  const [form] = Form.useForm();
  const [visible, setVisible] = useState(false);

  // 添加候选人
  const addCandidate = async (values) => {
    try {
      const tx = await contract.addCandidate(values.name);
      await tx.wait();
      setVisible(false);
      form.resetFields();
    } catch (error) {
      console.error("添加失败:", error);
    }
  };

  return (
    <div>
      <Button 
        type="primary" 
        onClick={() => setVisible(true)}
        style={{ marginBottom: 16 }}
      >
        addCandidate
      </Button>

      <Modal
        title="candidate control"
        visible={visible}
        onCancel={() => setVisible(false)}
        footer={null}
      >
        <Form form={form} onFinish={addCandidate}>
          <Form.Item
            name="name"
            label="candidate name"
            rules={[{ required: true }]}
          >
            <Input />
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit">
              submit
            </Button>
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
}