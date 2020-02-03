from db import get_db_connection, populate_table


def create_table():
    query = """
    CREATE TABLE IF NOT EXISTS companies (
    _id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(225),
    description TEXT
    )
    """
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute(query)


def drop_table():
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute("DROP TABLE IF EXISTS companies")


if __name__ == '__main__':
    create_table()
    populate_table('companies', 'data/companies.csv')
