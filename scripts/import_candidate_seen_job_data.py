from db import get_db_connection, populate_table
import os

def create_table():
    query = """
    CREATE TABLE IF NOT EXISTS candidateSeenJob (
    liked BOOLEAN,
    cid INT,
    jid VARCHAR(255),
    FOREIGN KEY (cid) REFERENCES candidates (_id) ON DELETE CASCADE,
    FOREIGN KEY (jid) REFERENCES jobs (_id) ON DELETE CASCADE
    )
    """
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute(query)


def drop_table():
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute("DROP TABLE IF EXISTS candidateSeenJob")


def populate_candidateSeenJob_data():
    drop_table()
    create_table()
    # Only populates candidateSeenJob test data in testing environment.
    if os.getenv("BREAD_ENV") == "testing":
        data_src = 'data/candidateSeenJob-test.csv'
        populate_table('candidateSeenJob', data_src)


if __name__ == '__main__':
    populate_candidateSeenJob_data()

