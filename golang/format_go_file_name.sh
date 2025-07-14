#!/bin/bash

if [ $# -eq 0 ]; then
  echo "Usage: $0 <directory>"
  exit 1
fi

target_dir="$1"

if [ ! -d "$target_dir" ]; then
  echo "Error: Directory '$target_dir' does not exist."
  exit 1
fi

find "$target_dir" -type f | while read file; do
  new_name=$(dirname "$file")/$(basename "$file" | sed 's/-/_/g')
  mv "$file" "$new_name"
done