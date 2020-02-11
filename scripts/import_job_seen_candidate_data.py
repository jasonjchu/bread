from db import get_db_connection, populate_table
import os

def create_table():
    query = """
    CREATE TABLE IF NOT EXISTS jobSeenCandidate (
    cid INT,
    jid VARCHAR(255),
    liked BOOLEAN,
    FOREIGN KEY (cid) REFERENCES candidates (_id) ON DELETE CASCADE,
    FOREIGN KEY (jid) REFERENCES jobs (_id) ON DELETE CASCADE,
    PRIMARY KEY (cid,jid)
    )
    """
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute(query)


def drop_table():
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute("DROP TABLE IF EXISTS jobSeenCandidate")


def populate_jobSeenCandidate_data():
    drop_table()
    create_table()
    # Only populates jobSeenCandidate test data in testing environment.
    if os.getenv("BREAD_ENV") == "testing":
        data_src = 'data/jobSeenCandidate-test.csv'
        populate_table('jobSeenCandidate', data_src)


if __name__ == '__main__':
    populate_jobSeenCandidate_data()
