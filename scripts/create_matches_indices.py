from db import populate_table, get_db_connection, get_data_src


def build_matches_indices():
    query = """
    CREATE INDEX jid_index
    ON matches (jid);
    """
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute(query)

if __name__ == '__main__':
    build_matches_indices()
