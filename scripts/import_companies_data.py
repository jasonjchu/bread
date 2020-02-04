from db import get_db_connection, populate_table
import os

def create_table():
    query = """
    CREATE TABLE IF NOT EXISTS companies (
    _id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(225),
    description TEXT
    )
    """
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute(query)


def drop_table():
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute("DROP TABLE IF EXISTS companies")


def populate_companies_data():
    drop_table()
    create_table()
    data_src = 'data/companies-test.csv' if os.getenv("BREAD_ENV") == "testing" else 'data/companies.csv'
    populate_table('companies', data_src)


if __name__ == '__main__':
    populate_companies_data()
