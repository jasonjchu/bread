SELECT * FROM jobs WHERE _id NOT IN (SELECT jid FROM candidateSeenJob WHERE cid = 1) LIMIT 200;
