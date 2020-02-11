from db import get_db_connection, populate_table
import os

def create_table():
    query = """
    CREATE TABLE IF NOT EXISTS candidates (
    _id INT PRIMARY KEY,
    name VARCHAR(225),
    program VARCHAR(225),
    grad_date DATE,
    description VARCHAR(225),
    FOREIGN KEY (_id) REFERENCES accounts (_id) ON DELETE CASCADE
    )
    """
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute(query)


def drop_table():
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute("DROP TABLE IF EXISTS candidates")


def populate_candidates_data():
    drop_table()
    create_table()
    # Only populates candidate test data in testing environment.
    if os.getenv("BREAD_ENV") == "testing":
        data_src = 'data/candidates-test.csv'
        populate_table('candidates', data_src)


if __name__ == '__main__':
    populate_candidates_data()
