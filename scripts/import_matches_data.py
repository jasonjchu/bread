from db import get_db_connection, populate_table
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
    if os.getenv("BREAD_ENV") == "testing":
        data_src = 'data/matches-test.csv'
        populate_table('matches', data_src)


if __name__ == '__main__':
    populate_matches_data()
