# must do pip install mysql-connector-python python-dotenv
# using python 3.6.8
from mysql.connector import connect, MySQLConnection
import csv
import datetime as dt
import os
from dotenv import load_dotenv, find_dotenv


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
    sector TEXT)
    """
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute(query)


def drop_table():
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute("DROP TABLE IF EXISTS jobs")


def populate_table():
    csv_file = os.path.join(os.path.dirname(__file__), os.pardir, 'data/monster-job-data.csv')
    with open(csv_file, encoding='utf8') as file:
        reader = csv.reader(file)
        columns = next(reader)
        query_template = 'INSERT INTO jobs({0}) values ({1})'
        query = query_template.format(','.join(columns), ','.join(['%s'] * len(columns)))
        print(query)
        # get DB connection
        db_cxn = get_db_connection(True)
        cursor = db_cxn.cursor()
        count = 0
        for data in reader:
            # Convert date to MYSQL date format. NULL if no date.
            data[2] = None if data[2] == '' else dt.datetime.strptime(data[2], '%m/%d/%Y').strftime('%Y-%m-%d')
            # Convert has_expired to boolean
            data[3] = True if data[3] == 'Yes' else False
            count = count + 1
            if count % 100 == 0:
                print('Inserted {} rows'.format(count))
            cursor.execute(query, data)
        db_cxn.commit()


if __name__ == '__main__':
    create_db()
    create_table()
    populate_table()
