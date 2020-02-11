from db import get_db_connection, populate_table
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
    # Only populates employer test data in testing environment.
    if os.getenv("BREAD_ENV") == "testing":
        data_src = 'data/employers-test.csv'
        populate_table('employers', data_src)


if __name__ == '__main__':
    populate_employers_data()
