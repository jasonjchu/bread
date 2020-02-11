from db import create_db, drop_db
from import_job_data import populate_jobs_data
from import_companies_data import populate_companies_data
from import_accounts_data import populate_accounts_data
from import_candidates_data import populate_candidates_data
from import_employers_data import populate_employers_data

if __name__ == '__main__':
    drop_db()
    create_db()
    print("Populating company data...")
    populate_companies_data()
    print("Populating jobs data...")
    populate_jobs_data()
    print("Populating accounts data...")
    populate_accounts_data()
    print("Populating candidates data...")
    populate_candidates_data()
    print("Populating employers data...")
    populate_employers_data()
    print("Successfully populated all data!")
