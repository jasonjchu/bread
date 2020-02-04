from db import create_db, drop_db
from import_job_data import populate_jobs_data
from import_companies_data import populate_companies_data

if __name__ == '__main__':
    drop_db()
    create_db()
    print("Populating company data...")
    populate_companies_data()
    print("Populating jobs data...")
    populate_jobs_data()
    print("Successfully populated all data!")
