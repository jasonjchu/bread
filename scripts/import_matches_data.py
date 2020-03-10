from db import get_db_connection, populate_table, get_data_src
import os

def create_table():
    # TODO: Build index on uid/jid
    query = """
    CREATE TABLE IF NOT EXISTS matches (
    uid INT,
    jid VARCHAR(255),
    PRIMARY KEY(uid, jid),
    FOREIGN KEY (uid) REFERENCES candidates (_id) ON DELETE CASCADE,
    FOREIGN KEY (jid) REFERENCES jobs (_id) ON DELETE CASCADE
    )
    """
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute(query)


def drop_table():
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute("DROP TABLE IF EXISTS matches")


def populate_matches_data():
    drop_table()
    create_table()
    table_name = 'matches'
    data_src = get_data_src(table_name)
    populate_table(table_name, data_src)


if __name__ == '__main__':
    populate_matches_data()
