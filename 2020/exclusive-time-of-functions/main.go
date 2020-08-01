package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/golang/go/src/strconv"
)

// Time: O(n)
// Space: O(n)
func exclusiveTime(n int, logs []string) []int {
	var (
		funcTimes = make([]int, n)
		lg        = parseLog(logs[0])
		lastLog   = lg
		funcStack = []log{lg}
	)
	for i := 1; i < len(logs); i++ {

		// The top of the stack always points to the function
		// whose time must be increased. Careful! The stack
		// can be empty!
		var peek log
		if len(funcStack) > 0 {
			peek = funcStack[len(funcStack)-1]
		}

		lg = parseLog(logs[i])
		// If the current log is a start, we calculate time up to
		// current timestamp-1, but only if the stack is not empty.
		if lg.isStart {
			if len(funcStack) > 0 {
				// Depending on what the last log was, the calculation
				// is different.
				if lastLog.isStart {
					funcTimes[peek.funcID] += lg.timestamp - lastLog.timestamp
				} else {
					funcTimes[peek.funcID] += lg.timestamp - lastLog.timestamp - 1
				}
			}
			// Always push starts, because the next calculated time
			// applies to this new function.
			funcStack = append(funcStack, lg)
		} else {
			// Depending on what the last log was, the calculation
			// is different.
			if lastLog.isStart {
				funcTimes[peek.funcID] += lg.timestamp - lastLog.timestamp + 1
			} else {
				funcTimes[peek.funcID] += lg.timestamp - lastLog.timestamp
			}
			// Always pop on ends, because the next calculated time applies
			// to the element below the latest start.
			funcStack = funcStack[:len(funcStack)-1]
		}
		lastLog = lg
	}

	return funcTimes
}

type log struct {
	funcID    int
	isStart   bool
	timestamp int
}

func parseLog(s string) log {
	parts := strings.Split(s, ":")
	funcID, _ := strconv.Atoi(parts[0])
	timestamp, _ := strconv.Atoi(parts[2])
	return log{funcID: funcID, isStart: parts[1] == "start", timestamp: timestamp}
}

func main() {
	ts := []struct {
		n        int
		logs     []string
		expected []int
	}{
		{
			n:        2,
			logs:     []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"},
			expected: []int{3, 4},
		},
		{
			n:        1,
			logs:     []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"},
			expected: []int{8},
		},
	}
	for _, tc := range ts {
		actual := exclusiveTime(tc.n, tc.logs)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.n, tc.logs, tc.expected, actual)
		}
	}
}
