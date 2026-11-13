#!/usr/bin/env bash

cd "$(dirname "$0")" || exit 1

dirs=([0-9][0-9]/)

if [ -n "$1" ]; then
  if [[ ! "$1" =~ ^[0-9]+$ ]]; then
    echo "Usage: $0 [day]" >&2
    exit 1
  fi
  # 10# reads "08" as decimal, not octal
  printf -v dir "%02d/" "$((10#$1))"
  if [ ! -d "$dir" ]; then
    echo "No folder for day $1" >&2
    exit 1
  fi
  dirs=("$dir")
fi

failed=0

for dir in "${dirs[@]}"; do
  day=${dir%/}
  echo "=== Day $day ==="
  go run "./$day" || failed=1
  echo
done

exit $failed
