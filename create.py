import os
import shutil
import sys

DAYS_PER_YEAR = 25
GO_VERSION = "1.25.4"


def create_year(year: int):
    days = DAYS_PER_YEAR
    if year >= 2025:
        days = DAYS_PER_YEAR // 2
    os.mkdir(f"./{year}")
    for day in range(1, days + 1):
        for challenge in ["A","B"]:
            dst = f"./{year}/Day {day:02d} {challenge}"
            shutil.copytree("./_base/", dst)
            inplace_change(f"{dst}/go.mod", "{{ .Year }}", f"{year}")
            inplace_change(f"{dst}/go.mod", "{{ .Day }}", f"{day}")
            inplace_change(f"{dst}/go.mod", "{{ .Part }}", f"{challenge}")
            inplace_change(f"{dst}/go.mod", "{{ .GoVersion }}", GO_VERSION)

def inplace_change(filename, old_string, new_string):
    # Safely read the input filename using 'with'
    with open(filename) as f:
        s = f.read()
        if old_string not in s:
            print('"{old_string}" not found in {filename}.'.format(**locals()))
            return

    # Safely write the changed content, if found in the file
    with open(filename, "w") as f:
        s = s.replace(old_string, new_string)
        f.write(s)

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Error: Use like create.py <year>")
        exit(1)
    year = sys.argv[1]
    
    create_year(int(year))
