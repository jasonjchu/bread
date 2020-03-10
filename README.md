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
### Get all jobs
GET `/jobs`

### Get employer by ID
GET `/employers/{employer_id}`

### Employer Registration
POST `/employers/register`
```
{
    username: String
    password: String
    name: String
    worksAt: Int
}
```

### Employer Login
POST `/employers/login`
```json
{
    username: String
    password: String
}
```

### Candidate Registration
POST `/candidates/register`
```json
{
    username: String
    password: String
    name: String
    program: String
    grad_date: DateTime
    description: String
}
```

### Candidate Login
POST `/candidates/login`
```json
{
    username: String
    password: String
}
```


### Get all companies
GET `/companies{?name=}`

Can also query by name. Does case-insensitive substring match with DB names.

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
- [ ] Milestone 2: Midterm report - Mar. 5, 2020
- [ ] In-class demo: - April 2, 2020
- [ ] Final report and code submission - April 2, 2020

## Contributors:
- Aaron Choo (@Aaronchoo)
- Jason Chu (@jasonjchu)
- Kallen Tu (@kallentu)
- Charles Zhang (@gzcharleszhang)
- Jack Zhang (@Novacer)