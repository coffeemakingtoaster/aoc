#!/usr/bin/env sh

# Usage: ./commit.sh 2024 05 A
# Produces commit message: "[2024] day A"

YEAR="$1"
DAY="$2"
PART="$3"

if [ -z "$YEAR" ] || [ -z "$DAY" ] || [ -z "$PART" ]; then
    echo "Usage: $0 <year> <day> <A|B>"
    exit 1
fi

git commit -m "[$YEAR] $DAY $PART"
