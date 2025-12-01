from datetime import datetime, timedelta
from os import path
import os
import subprocess

RED = "\033[0;31m"
GREEN = "\033[0;32m"
NC = "\033[0m"

DIVIDER = "-----------------------"

YEARS = ["2023","2024"]

def benchmark_year(year):
   
    print(f"* {year} *")
    print(DIVIDER)

    duration_sum = timedelta(0)
    day = 25
    if int(year) >= 2025:
        day = 12
    for i in range(1, day + 1):
        for part in ["A", "B"]:
            dirname = f"{year}/Day {i:02d} {part}"
            starttime = datetime.now()
            cmd = subprocess.run(
                "go test ./",
                cwd=path.join(os.getcwd(), dirname),
                shell=True,
                stdout=open(os.devnull, "wb"),
            )
            duration = datetime.now() - starttime
            duration_sum += duration
            if cmd.returncode == 0:
                print(
                    f"|{dirname}\t|{GREEN}SUCCESS{NC} ({duration.seconds}:{int(duration.microseconds/1000)})\t|"
                )
            else:
                print(f"|{dirname}\t|{RED}FAILURE{NC}\t|")
            print(DIVIDER)
    print(
        f"|(Duration: {duration_sum})\t|"
            )
    print(DIVIDER)
    return duration_sum


subprocess.run("go clean -testcache", shell=True)

total_sum = timedelta(0)
for year in YEARS:
    total_sum += benchmark_year(year)

print(f"Total elapsed time spent running tests: {total_sum}")
