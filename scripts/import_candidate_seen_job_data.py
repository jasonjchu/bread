from db import get_db_connection, populate_table, get_data_src
import os

def create_table():
    query = """
    CREATE TABLE IF NOT EXISTS candidateSeenJob (
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
    db_cxn.cursor().execute("DROP TABLE IF EXISTS candidateSeenJob")


def populate_candidateSeenJob_data():
    drop_table()
    create_table()
    table_name = 'candidateSeenJob'
    data_src = get_data_src(table_name)
    populate_table(table_name, data_src)


if __name__ == '__main__':
    populate_candidateSeenJob_data()

