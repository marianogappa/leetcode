#!/bin/bash

EXERCISE=$1

mkdir -p scratch/$EXERCISE
cd scratch/$EXERCISE
touch ${EXERCISE}.go
cat <<'EOF' > ${EXERCISE}.go
package main

import "fmt"

// Time: O()
// Space: O()

func main() {
    ts := []struct{
        input int
        expected int
    } {
        {
            input: 1,
            expected: 1,
        },
    }
    for _, tc := range ts {
        actual := 1
        if tc.expected != actual {
            fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
        }
    }
}
EOF

/usr/local/bin/code .
