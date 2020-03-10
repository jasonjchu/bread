from db import get_db_connection, populate_table, get_data_src
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
    table_name = 'companies'
    data_src = get_data_src(table_name)
    populate_table(table_name, data_src)


if __name__ == '__main__':
    populate_companies_data()
