SELECT jobs.*, companies.name as company_name
FROM jobs,
     companies
WHERE jobs._id NOT IN
      (SELECT jid FROM candidateSeenJob WHERE cid = 1)
  AND jobs.company_id = companies._id
LIMIT 200;
