SELECT jobs.*, companies.name as company_name
FROM jobs,
     companies
WHERE jobs.company_id = companies._id
  AND company_id = 1;
