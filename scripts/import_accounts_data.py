from db import get_db_connection, populate_table
import os

def create_table():
    query = """
    CREATE TABLE IF NOT EXISTS accounts (
    _id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(225),
    password VARCHAR(255),
    candidate_id INT,
    employer_id INT,
    FOREIGN KEY (candidate_id) REFERENCES candidates (_id) ON DELETE CASCADE,
    FOREIGN KEY (employer_id) REFERENCES employers (_id) ON DELETE CASCADE
    )
    """
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute(query)


def drop_table():
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute("DROP TABLE IF EXISTS accounts")


def populate_accounts_data():
    drop_table()
    create_table()
    # Only populates account test data in testing environment.
    if os.getenv("BREAD_ENV") == "testing":
        data_src = 'data/accounts-test.csv'

        def transform_account(account):
            account[2] = None if account[2] == '' else account[2]
            account[3] = None if account[3] == '' else account[3]

        populate_table('accounts', data_src, transform_account)


if __name__ == '__main__':
    populate_accounts_data()
