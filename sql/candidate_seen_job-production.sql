SELECT * FROM candidates WHERE _id in
      (SELECT cid FROM candidateSeenJob WHERE jid='002139162b354fa7b2bd9e91e408fa30' AND liked=True
      AND cid NOT IN (SELECT cid FROM jobSeenCandidate WHERE jid='002139162b354fa7b2bd9e91e408fa30')) LIMIT 200;

