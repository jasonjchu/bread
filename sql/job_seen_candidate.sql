select * from candidates where _id in (select cid from candidateSeenJob where jid='1e2637cb5f7a2c4615a99a26c0566c66' AND liked=True AND cid not in (select cid from jobSeenCandidate where jid='1e2637cb5f7a2c4615a99a26c0566c66'));

