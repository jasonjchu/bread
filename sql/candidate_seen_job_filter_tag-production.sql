SELECT jobs.*, companies.name as company_name
FROM jobs,
     companies
WHERE jobs._id NOT IN
      (SELECT jid FROM candidateSeenJob WHERE cid = 4)
  AND EXISTS(SELECT tid FROM tagsDescribeJobs WHERE jid = jobs._id AND tid IN (1, 2, 3))
  AND jobs.company_id = companies._id
LIMIT 20;
