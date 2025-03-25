import React, { useState, useEffect } from 'react';
import { Card, Button, List, message } from 'antd';
import { ethers } from 'ethers';

export default function VoterPanel({ contract }) {
  const [candidates, setCandidates] = useState([]);
  const [hasVoted, setHasVoted] = useState(false);

  // ���غ�ѡ������
  const loadCandidates = async () => {
    if (contract) {
      const count = await contract.getCandidateCount();
      const candidates = [];
      for (let i = 0; i < count; i++) {
        const candidate = await contract.candidates(i);
        candidates.push({
          id: candidate.id.toNumber(),
          name: candidate.name,
          votes: candidate.voteCount.toNumber()
        });
      }
      setCandidates(candidates);
    }
  };

  // ����ͶƱ
  const handleVote = async (candidateId) => {
    try {
      const tx = await contract.vote(candidateId);
      await tx.wait();
      message.success("ͶƱ�ɹ���");
      setHasVoted(true);
      loadCandidates();
    } catch (error) {
      message.error(`ͶƱʧ��: ${error.message}`);
    }
  };

  useEffect(() => {
    loadCandidates();
  }, [contract]);

  return (
    <Card title="List of candidates" bordered={false}>
      <List
        itemLayout="horizontal"
        dataSource={candidates}
        renderItem={item => (
          <List.Item
            actions={[
              <Button 
                type="primary" 
                onClick={() => handleVote(item.id)}
                disabled={hasVoted}
              >
                {hasVoted ? "��ͶƱ" : "ͶƱ"}
              </Button>
            ]}
          >
            <List.Item.Meta
              title={item.name}
              description={`��ǰƱ��: ${item.votes}`}
            />
          </List.Item>
        )}
      />
    </Card>
  );
}