#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: $0 <new-name>"
  exit 1
fi

NEW_NAME="$1"
ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"

echo "Replacing 'rib' with '$NEW_NAME' in root files..."

for file in "$ROOT_DIR"/.* "$ROOT_DIR"/*; do
  [ -f "$file" ] || continue
  sed -i '' "s/rib/$NEW_NAME/g" "$file"
done

echo "Done."
