from db import get_db_connection, populate_table, get_data_src
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
    table_name = 'jobTags'
    data_src = get_data_src(table_name)
    populate_table(table_name, data_src)


if __name__ == '__main__':
    populate_jobTags_data()
