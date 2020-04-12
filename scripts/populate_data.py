from db import create_db, drop_db
from import_job_data import populate_jobs_data
from import_companies_data import populate_companies_data
from import_accounts_data import populate_accounts_data
from import_candidates_data import populate_candidates_data
from import_employers_data import populate_employers_data
from import_candidate_seen_job_data import populate_candidateSeenJob_data
from import_job_seen_candidate_data import populate_jobSeenCandidate_data
from import_matches_data import populate_matches_data
from import_job_tags_data import populate_jobTags_data
from import_tags_describe_jobs_data import populate_tagsDescribeJobs_data
from create_job_indices import build_job_indices
from create_matches_indices import build_matches_indices

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
    print("Populating candidate seen job data...")
    populate_candidateSeenJob_data()
    print("Populating job seen candidate data...")
    populate_jobSeenCandidate_data()
    print("Populating matches data...")
    populate_matches_data()
    print("Populating job tags data...")
    populate_jobTags_data()
    print("Populating tags describe jobs data...")
    populate_tagsDescribeJobs_data()
    print("Successfully populated all data!")
    print("Building indices")
    build_job_indices()
    build_matches_indices()
    print("Successfully built indices on jobs!")
