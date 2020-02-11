# bread
Tinder-inspired job searching application for employers and prospective employees.

## Contributors:
- Aaron Choo (@Aaronchoo)
- Jason Chu (@jasonjchu)
- Kallen Tu (@kallentu)
- Charles Zhang (@gzcharleszhang)
- Jack Zhang (@Novacer)

## Goal
The purpose of this project is to demonstrate a thorough understanding of the concepts taught in CS348: Introduction to Database Management.

Tasks are to be completed by the set deadlines.

- [X] Milestone 0: Project Preparation - Jan. 30, 2020
- [ ] Milestone 1: Project Proposal - Feb. 6, 2020
- [ ] Milestone 2: Midterm report - Mar. 5, 2020
- [ ] In-class demo: - April 2, 2020
- [ ] Final report and code submission - April 2, 2020

## API Endpoints
### Get all jobs
GET `/jobs`

### Get employer by ID
GET `/employers/{employer_id}`

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

Run `scripts/import_job_data.py` to create the MySQL database, add and populate the tables with 
data from `data/monster-job-data.csv`.
