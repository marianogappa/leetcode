#!/bin/bash

set -e

EXERCISE=$1

mkdir -p "$EXERCISE"
cd "$EXERCISE"
touch main.py
cat <<'EOF' > main.py
EOF

/usr/local/bin/code .
