from db import populate_table, get_db_connection, get_data_src


def build_job_indices():
    query = """
    CREATE INDEX cid_index
    ON jobs (company_id);
    """
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute(query)

if __name__ == '__main__':
    build_job_indices()
