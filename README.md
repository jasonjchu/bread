# bread
Tinder-inspired job searching application for employers and prospective employees.

## Project Structure

```
├── app
│   ├── db
│   ├── env
│   ├── handlers                # Requests and responses for the API endpoints.
│   ├── models                  # Data structures and queries for the SQL database.
│   ├── routes                  # Initializes API endpoints.
│   ├── server
│   ├── services                # Data structures for handlers and API executing logic.
│   └── utils
├── cmd                         # Main entrypoint to bread.
├── data                        # CSV test data.
├── reports                     # report.pdf for each milestone.
├── scripts                     # Python scripts for populating data in database tables.
├── sql                         # Sample SQL statements.
│   └── out                     # Expected output for sample SQL statements in sql/
├── bread-ER.pdf                # ER Diagram of database structure.
├── members.txt                 # Contribution list of all members.
└── README.md                   # Documentation on running scripts, setting up server...
```

## API Endpoints
### Get all jobs (can also query by id)
GET `/jobs{?id=}`

Related file: `app/handlers/getJobsHandler`

### Get employer by ID
GET `/employers/{employer_id}`

Related file: `app/handlers/getEmployerHandler`

### Employer Registration
POST `/employers/register`
##### Request Body
```
{
    username: String
    password: String
    name: String
    worksAt: Int
}
```

Related file: `app/services/employerRegisterService`

### Employer Login
POST `/employers/login`
##### Request Body
```
{
    username: String
    password: String
}
```

Related file: `app/services/employerLoginService`

### Get open job postings for employer
GET `/employers/jobs`
##### Request Header
```
{
  user_id: Int
}
```

Related file: `app/services/employerJobsService`

### Get candidates for employer job by ID
GET `/employers/jobs/{job_id}/candidates{?limit=}`

Can specify a limit for Candidates returned. Default limit is 200.

Related file: `app/handlers/getCandidatesForJobHandler`

### Employer Likes Candidate for Job (Swipes Right)
POST `/employers/like`
##### Request Body
```
{
  job_id: String,
  candidate_id: Int
}
```

Related file: `app/services/employerLikesCandidateService`

### Employer Dislikes Candidate for Job (Swipes Left)
POST `/employers/dislike`
##### Request Body
```
{
  job_id: String,
  candidate_id: Int
}
```

Related file: `app/services/employerDislikesCandidateService`

### Get matches for employer
GET `/employers/matches`
##### Request Header
```
{
  user_id: Int
}
```

Related file: `app/services/employerMatchesService`

### Candidate Registration
POST `/candidates/register`
##### Request Body
```
{
    username: String
    password: String
    name: String
    program: String
    grad_date: DateTime
    description: String
}
```

Related file: `app/services/candidateRegisterService`

### Candidate Login
POST `/candidates/login`
##### Request Body
```
{
    username: String
    password: String
}
```

Related file: `app/services/candidateLoginService`

### Get Candidate by Id
GET `/candidates/{id}`

Related file: `app/handlers/getCandidatesByIdHandler`

### Candidate Likes Job (Swipes Right)
POST `/candidates/jobs/{job_id}/like`
##### Request Header
```
{
  user_id: Int
}
```

Related file: `app/services/candidateLikesJobService`

### Candidate Dislikes Job (Swipes Left)
POST `/candidates/jobs/{job_id}/dislike`
##### Request Header
```
{
  user_id: Int
}
```

Related file: `app/services/candidateDislikesJobService`

### Get jobs not seen for candidate by ID
GET `/candidates/jobs{?limit=&tag_ids=}`
##### Request Header
```
{
  user_id: Int
}
```
Can specify a limit for number of jobs returned and tag_ids array is used to filter jobs. Limit has a default value of 200, and no including tags will not apply tag filter.

Related file: `app/handlers/getJobsForCandidatesHandler`

### Get matches for candidate
GET `/candidates/matches`
##### Request Header
```
{
  user_id: Int
}
```

Related file: `app/services/candidatesMatchesService`

### Get all companies
GET `/companies{?name=}`

Can also query by name. Does case-insensitive substring match with DB names.

Related file: `app/handlers/getCompaniesHandler`

### Get all tags available
GET `/tags`

Related file: `app/handlers/getJobTagsHandler`

## Running the API Server
### Requirements
1. Go >= 1.13.7
1. [dep](https://github.com/golang/dep) >= 0.5.4

### Starting the Server
1. `dep ensure`
1. `go run cmd/bread/main.go`
1. The API server will be served at port 8080 by default.

## Running scripts
### Requirements
1. Python 3.6
2. `pip install mysql-connector-python python-dotenv`
3. Script sometimes prints unicode characters to terminal for debugging. Ensure this works for you by setting `export PYTHONIOENCODING='utf-8'
` before running the script.

### Populating Database
```
py scripts/populate_data.py
```
Create the MySQL database, add and populate the tables.

Set environment variable `BREAD_ENV=testing` if you wish
to populate using testing data, otherwise the script
will use production data.

## Milestones
Tasks are to be completed by the set deadlines.

- [X] Milestone 0: Project Preparation - Jan. 30, 2020
- [X] Milestone 1: Project Proposal - Feb. 11, 2020
- [X] Milestone 2: Midterm report - Mar. 10, 2020
- [ ] In-class demo: - April 2, 2020
- [ ] Final report and code submission - April 2, 2020

## Contributors:
- Aaron Choo (@Aaronchoo)
- Jason Chu (@jasonjchu)
- Kallen Tu (@kallentu)
- Charles Zhang (@gzcharleszhang)
- Jack Zhang (@Novacer)