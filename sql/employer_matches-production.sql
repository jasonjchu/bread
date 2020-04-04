SELECT candidates._id  as canid,
       jobs._id        as jid,
       companies._id   as compid,
       companies.name  as compname,
       jobs.job_title,
       jobs.location,
       candidates.name as candname,
       candidates.program,
       candidates.grad_date
FROM employers,
     jobs,
     matches,
     candidates,
     companies
WHERE employers._id = 9
  AND jobs.company_id = employers.works_at
  AND jobs.company_id = companies._id
  AND matches.jid = jobs._id
  AND candidates._id = matches.uid;
