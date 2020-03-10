from db import get_db_connection, populate_table, get_data_src
import os

def create_table():
    query = """
    CREATE TABLE IF NOT EXISTS employers (
    _id INT PRIMARY KEY,
    name VARCHAR(225),
    works_at INT,
    FOREIGN KEY (works_at) REFERENCES companies (_id) ON DELETE CASCADE,
    FOREIGN KEY (_id) REFERENCES accounts (_id) ON DELETE CASCADE
    )
    """
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute(query)


def drop_table():
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute("DROP TABLE IF EXISTS employers")


def populate_employers_data():
    drop_table()
    create_table()
    table_name = 'employers'
    data_src = get_data_src(table_name)
    populate_table(table_name, data_src)


if __name__ == '__main__':
    populate_employers_data()
