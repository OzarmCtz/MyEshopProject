# This script is used to clean files in the datasources/mysql folder
# It removes the sql. prefix from all the sql.Null* types except sql.Result
# It also removes the "database/sql" import
# It is supposed to be run from buildmodels.go after all the models have  been generated.
# We need to do this (while waiting for a better solution)
# because sqlc can't generate directly the Null* types.
# sql.Null* types are not compatible with NulLtypes redefined in types.go so that
# we can use them to Marshal/Unmarshal to/from JSON transparently.

import re

path = "../datasources/mysql/"
files = ["models.go", "querier.go", "query.sql.go"]
modelsPattern = r'\bsql\.(?!Result\b)|"database/sql"'
pattern = r'\bsql\.(?!Result\b)'

for file in files:
    pathfile = path+file
    with open(pathfile, 'r') as file:
        text = file.read()

    # Remove only the sql. prefix without removing the rest of the string
    if pathfile == "../datasources/mysql/models.go":
        result = re.sub(modelsPattern, '', text)
    else:
        result = re.sub(pattern, '', text)

    with open(pathfile, 'w') as file:
        file.write(result)

    print("File {} cleaned".format(file.name))
