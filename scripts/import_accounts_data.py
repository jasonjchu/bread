from db import get_db_connection, populate_table, get_data_src
import os

def create_table():
    query = """
    CREATE TABLE IF NOT EXISTS accounts (
    _id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(225) UNIQUE,
    password VARCHAR(255)
    )
    """
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute(query)


def drop_account_and_children():
    db_cxn = get_db_connection(True)
    cursor = db_cxn.cursor()
    # Drop table doesn't trigger on delete cascade in candidate/employer tables, hence will fail
    # unless we drop all foreign key constraints in the children table as well.
    # In this case just drop all the children tables cuz YOLO
    cursor.execute("DROP TABLE IF EXISTS candidates")
    cursor.execute("DROP TABLE IF EXISTS employers")
    cursor.execute("DROP TABLE IF EXISTS accounts")


def populate_accounts_data():
    drop_account_and_children()
    create_table()
    table_name = 'accounts'
    data_src = get_data_src(table_name)
    populate_table(table_name, data_src)


warning = """
WARNING: Running this script on its own will also drop candidates and employers!!!!!
Consider using populate_data.py instead, otherwise you are responsible for also calling:
import_candidates_data.py
import_employers_data.py
Failure to do so will result in a sentence of two (2) academic terms in MC basement.
"""

if __name__ == '__main__':
    print(warning)
    populate_accounts_data()
