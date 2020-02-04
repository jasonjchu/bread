# must do pip install mysql-connector-python python-dotenv
# using python 3.6.8
from mysql.connector import connect, MySQLConnection
import os
import csv
from dotenv import load_dotenv, find_dotenv
from typing import Callable


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


def populate_table(table_name: str, file_name: str, transform_fn: Callable[..., None] = lambda *args: None):
    csv_file = os.path.join(os.path.dirname(__file__), os.pardir, file_name)
    with open(csv_file, encoding='utf8') as file:
        reader = csv.reader(file)
        columns = next(reader)
        query_template = 'INSERT INTO %s({0}) values ({1})' % table_name
        query = query_template.format(','.join(columns), ','.join(['%s'] * len(columns)))
        # get DB connection
        db_cxn = get_db_connection(True)
        cursor = db_cxn.cursor()
        count = 0
        for data in reader:
            transform_fn(data)
            count = count + 1
            if count % 100 == 0:
                print('Inserted {} rows'.format(count))
            cursor.execute(query, data)
        print('Successfully inserted {0} rows into {1}'.format(count, table_name))
        db_cxn.commit()
