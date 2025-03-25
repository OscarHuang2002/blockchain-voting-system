import React, { useState, useEffect } from 'react';
import { Pie } from '@ant-design/charts';
import { Spin } from 'antd';

export default function LiveResults({ contract }) {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(true);

  const loadResults = async () => {
    if (contract) {
      const count = await contract.getCandidateCount();
      const results = [];
      for (let i = 0; i < count; i++) {
        const candidate = await contract.candidates(i);
        results.push({
          name: candidate.name,
          value: candidate.voteCount.toNumber()
        });
      }
      setData(results);
      setLoading(false);
    }
  };

  useEffect(() => {
    loadResults();
  }, [contract]);

  const config = {
    data,
    angleField: 'value',
    colorField: 'name',
    radius: 0.8,
    label: {
      type: 'spider',
      content: '{name}\n{percentage}'
    }
  };

  return (
    <div style={{ padding: 24, background: '#fff' }}>
      <h2>Live vote statistics</h2>
      {loading ? <Spin /> : <Pie {...config} />}
    </div>
  );
}