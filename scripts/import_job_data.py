# must do pip install mysql-connector-python python-dotenv
# using python 3.6.8
from mysql.connector import connect, MySQLConnection
import datetime as dt
import os
from dotenv import load_dotenv, find_dotenv
from db import populate_table
import random


load_dotenv(find_dotenv())
database_name = os.getenv('BREAD_DB_NAME', 'bread_db')


def get_db_connection(db_exists: bool) -> MySQLConnection:
    return connect(
        host=os.getenv('BREAD_DB_HOST', 'localhost'),
        user=os.getenv('BREAD_DB_USER', 'root'),
        passwd=os.getenv('BREAD_DB_PASSWD', 'password'),
        database=database_name if db_exists else '',
        auth_plugin='mysql_native_password'
    )


def create_db():
    # always assume no DB exists at this step, since if not exists will leave things unchanged
    db_cxn = get_db_connection(False)
    db_cxn.cursor().execute(
        "CREATE DATABASE IF NOT EXISTS {0}".format(database_name)
    )


def drop_db():
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute(
        "DROP DATABASE IF EXISTS {0}".format(database_name)
    )


def create_table():
    query = """
    CREATE TABLE IF NOT EXISTS jobs (
    _id VARCHAR(255) PRIMARY KEY,
    posted_by INT,
    country VARCHAR(1000), 
    country_code VARCHAR(10), 
    date_added DATE,
    has_expired BOOLEAN,
    job_board VARCHAR(1000), 
    job_description TEXT, 
    job_title VARCHAR(1000), 
    job_type VARCHAR(1000),
    location TEXT,
    organization VARCHAR(1000), 
    page_url VARCHAR(2000), 
    salary VARCHAR(1000), 
    sector TEXT,
    FOREIGN KEY (posted_by) REFERENCES companies (_id) ON DELETE CASCADE
    )
    """
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute(query)


def drop_table():
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute("DROP TABLE IF EXISTS jobs")


def get_companies_count():
    db_cxn = get_db_connection(True)
    cur = db_cxn.cursor()
    cur.execute("SELECT COUNT(*) FROM companies")
    res = cur.fetchall()
    return res[0][0]


def populate_jobs_data():
    drop_table()
    create_table()

    data_src = 'data/jobs-test.csv' if os.getenv("BREAD_ENV") == "testing" else 'data/jobs.csv'
    companies_count = get_companies_count()

    # massaging job data before saving into DB
    def transform_job(job):
        # Convert date to MYSQL date format. NULL if no date.
        job[2] = None if job[2] == '' else dt.datetime.strptime(job[2], '%m/%d/%Y').strftime('%Y-%m-%d')
        # Convert has_expired to boolean
        job[3] = True if job[3] == 'Yes' else False
        # Randomly assign job to a company that exists in the DB
        job[-1] = random.randint(1,companies_count) if companies_count > 0 else 0

    populate_table('jobs', data_src, transform_job)


if __name__ == '__main__':
    populate_jobs_data()
