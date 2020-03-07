import os
import codecs
import shutil


# Note: throws error if file_name is also being accessed by another process!
def clean_utf8(file_name):
    """
    Read file line by line, removing (ignoring) non utf8 characters
    Write result to temp file, copying to orig file on success
    :param file_name: the relative filename from project root
    """
    csv_file = os.path.join(os.path.dirname(__file__), os.pardir, file_name)
    with codecs.open(csv_file, encoding='utf8', errors='ignore') as old, \
            codecs.open(csv_file + 'temp', mode='w', encoding='utf8') as new:
        new.writelines(old)
    shutil.move(csv_file + 'temp', csv_file)


if __name__ == '__main__':
    clean_utf8('data/jobs-test.csv')
