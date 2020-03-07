from db import get_db_connection, populate_table
import os


def create_table():
    query = """
    CREATE TABLE IF NOT EXISTS jobTags (
    _id INT PRIMARY KEY AUTO_INCREMENT,
    tag_name VARCHAR(1000) UNIQUE, /* TODO: create non-clustered index for searching */
    tag_description TEXT
    )
    """
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute(query)


def drop_table():
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute("DROP TABLE IF EXISTS jobTags")


def populate_jobTags_data():
    drop_table()
    create_table()
    # Only populates jobTags test data in testing environment.
    if os.getenv("BREAD_ENV") == "testing":
        data_src = 'data/jobTags-test.csv'
        populate_table('jobTags', data_src)


if __name__ == '__main__':
    populate_jobTags_data()
