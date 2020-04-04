SELECT candidates._id  as canid,
       jobs._id        as jid,
       companies._id   as compid,
       companies.name  as compname,
       jobs.job_title,
       jobs.location
FROM jobs,
     matches,
     candidates,
     companies
WHERE candidates._id = 4
  AND matches.uid = candidates._id
  AND jobs._id = matches.jid
  AND jobs.company_id = companies._id;
