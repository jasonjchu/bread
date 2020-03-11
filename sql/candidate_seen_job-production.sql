SELECT * FROM candidates WHERE _id in
      (SELECT cid FROM candidateSeenJob WHERE jid='1e2637cb5f7a2c4615a99a26c0566c66' AND liked=True
      AND cid NOT IN (SELECT cid FROM jobSeenCandidate WHERE jid='1e2637cb5f7a2c4615a99a26c0566c66')) LIMIT 200;
