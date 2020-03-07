from db import get_db_connection, populate_table
from collections import OrderedDict
import os
import csv
import codecs


def create_table():
    query = """
    CREATE TABLE IF NOT EXISTS tagsDescribeJobs (
    tid INT,
    jid VARCHAR(255),
    PRIMARY KEY (tid, jid),
    FOREIGN KEY (tid) REFERENCES jobTags (_id) ON DELETE CASCADE,
    FOREIGN KEY (jid) REFERENCES jobs (_id) ON DELETE CASCADE
    )
    """
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute(query)


def drop_table():
    db_cxn = get_db_connection(True)
    db_cxn.cursor().execute("DROP TABLE IF EXISTS tagsDescribeJobs")


def assign_tags(jobs_data_src, tags_data_src, target):
    """
    Programmatically generate the tagsDescribeJobs data csv by reading the tags csv and the jobs csv.
    Check if the job_title, job_description contains the tag name.
    If so, add the tuple to the csv
    :param jobs_data_src: path to jobs data csv
    :param tags_data_src: path to tags data csv
    :param target: path/name of the output file
    """
    print("Assigning tags...")
    src_csv = os.path.join(os.path.dirname(__file__), os.pardir, jobs_data_src)
    tags_csv = os.path.join(os.path.dirname(__file__), os.pardir, tags_data_src)
    target_csv = os.path.join(os.path.dirname(__file__), os.pardir, target)

    tags_and_ids = []  # array of tuples (tag_name, tag_id)
    # Build the array by reading the tags csv
    with codecs.open(tags_csv, mode='r') as tags_file:
        tags_reader = csv.reader(tags_file)
        tags_columns = next(tags_reader)
        tid = tags_columns.index('_id')
        tag_name = tags_columns.index('tag_name')

        for data in tags_reader:
            tags_and_ids.append((data[tag_name].lower(), data[tid]))

    assigned_tags = OrderedDict()  # set of tuples (tag_id, job_id)
    # Build the set by reading the src csv
    with codecs.open(src_csv, encoding='utf8', errors='ignore') as src_file:
        src_reader = csv.reader(src_file)
        src_columns = next(src_reader)
        job_desc = src_columns.index('job_description')
        job_title = src_columns.index('job_title')
        job_id = src_columns.index('_id')

        for data in src_reader:
            search_txt_title = '' if data[job_title] is None else data[job_title].lower()
            search_txt_desc = '' if data[job_desc] is None else data[job_desc].lower()

            for (tag_name, tid) in tags_and_ids:
                if tag_name in search_txt_title or tag_name in search_txt_desc:
                    assigned_tags[(tid, data[job_id])] = None

    # Write out the final output
    with codecs.open(target_csv, mode='w', encoding='utf8') as output:
        output.write('tid,jid\n')
        for (tid, jid) in assigned_tags.keys():
            output.write('{},{}\n'.format(tid, jid))

    print("Tags assigned successfully!")


def populate_tagsDescribeJobs_data():
    drop_table()
    create_table()
    # Only populates tagsDescribeJobs test data in testing environment.
    if os.getenv("BREAD_ENV") == "testing":
        data_src = 'data/tagsDescribeJobs-test.csv'
        assign_tags('data/jobs-test.csv', 'data/jobTags-test.csv', data_src)

        populate_table('tagsDescribeJobs', data_src)


if __name__ == '__main__':
    populate_tagsDescribeJobs_data()
