# must do pip install mysql-connector-python
# using python 3.6.8
import mysql.connector
import csv
import datetime as dt
import os

db_connection = mysql.connector.connect(
    host="localhost",
    user="root",
    passwd="password",
    # database="bread_db", <--- uncomment after DB is created
    auth_plugin='mysql_native_password'
)
cursor = db_connection.cursor()


# Execute me first
def create_db():
    cursor.execute("CREATE DATABASE IF NOT EXISTS bread_db")


def create_table():
    query = """
    CREATE TABLE IF NOT EXISTS jobs (
    uniq_id VARCHAR(255) PRIMARY KEY,
    country VARCHAR(1000), 
    country_code VARCHAR(10), 
    date_added DATE,
    has_expired BIT,
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
    cursor.execute(query)


def populate_table():
    csv_file = os.path.join(os.path.dirname(__file__), os.pardir, 'data/monster-job-data.csv')
    with open(csv_file, encoding='utf8') as file:
        reader = csv.reader(file)
        columns = next(reader)
        query_template = 'INSERT INTO jobs({0}) values ({1})'
        query = query_template.format(','.join(columns), ','.join(['%s'] * len(columns)))
        print(query)
        for data in reader:
            # Convert date to MYSQL date format. NULL if no date.
            data[2] = None if data[2] == '' else dt.datetime.strptime(data[2], '%m/%d/%Y').strftime('%Y-%m-%d')
            # Convert has_expired to boolean
            data[3] = True if data[3] == 'Yes' else False
            print(data)
            cursor.execute(query, data)
        db_connection.commit()


if __name__ == '__main__':
    create_db()
    # create_table()    <--- uncomment after DB is created
    # populate_table()  <--- uncomment after DB is created
