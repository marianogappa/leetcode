package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Time: O(len(logs))
// Space: O(n + len(logs))
//
// This exercise is not so much tricky in terms of algorithm complexity but rather tricky in terms of edge cases and
// off-by-one errors. So when you figure out this intuition, be on the lookout for a million edge cases!
//
// Do not consider the option of looping for 1..n. There could be one function with 1 start & 1 end, and a very large n
// for the end time.
//
// Assume all log strings are well-formed, so there aren't any errors parsing it.
// Is it ok to assume that 0 is the first timestamp every time?
//
// In terms of space we need to store one int per function (exclusive time), and we need a stack of executing functions
// which is len(logs) in the worst case.
//
// In terms of time, the most efficient option (which is possible) is to go over the logs once, so it's linear to the
// length of logs.
//
// If the log is a start:
// - Push this log to an execution stack. Keep track of the function id and which time it started.
// - If there was a function executing in the stack when this happened, calculate how much time it executed until now,
//   and add it to these function id's exlusive time. Note that the current timestamp is not included, as it belogs to
//   the new function that takes its place in the stack.
//
// Otherwise (i.e. if the log is an end):
// - Pop the current executing function off the stack (here we assume the current end matches the latest start),
//   calculate its exclusive time and add it to the results. Note that current timestamp IS INCLUDED this time.
// - TRICKY PART! Since we popped this fID, if stack isn't empty, the latest stack entry starts executing again. We
//   already accounted for the time it executed up until it was interrupted, but from this timestamp+1 it will start
//   executing again, so update the latestStartAt to log.timestamp+1.
//
// Return all exclusive times calulated.
func exclusiveTime(n int, logs []string) []int {
	// Edge case: there could be 12974019 functions, but none executed.
	if len(logs) == 0 {
		return []int{}
	}

	exclusiveTimes := make([]int, n)
	executionStack := []execution{}

	for _, rawLog := range logs {
		log := logEntryFromString(rawLog)
		if log.isStart {
			// We might have interrupted an executing function. So save how much time it executed.
			if len(executionStack) > 0 {
				latestExecution := executionStack[len(executionStack)-1]
				exclusiveTimes[latestExecution.fID] += log.timestamp - latestExecution.latestStartAt
			}
			// Push the current starting function into the stack.
			executionStack = append(executionStack, execution{fID: log.fID, latestStartAt: log.timestamp})
		} else {
			// Save how much time the current ending function executed. Then pop it off the stack.
			//
			// N.B. Here we assume the current end's fID matches the latest start fID. This should be correct, but ask!
			latestExecution := executionStack[len(executionStack)-1]
			executionStack = executionStack[:len(executionStack)-1]
			exclusiveTimes[latestExecution.fID] += log.timestamp - latestExecution.latestStartAt + 1

			// If there is a function in the stack that will restart executing on the next timestamp, update that start
			// timestamp.
			if len(executionStack) > 0 {
				executionStack[len(executionStack)-1].latestStartAt = log.timestamp + 1
			}
		}
	}

	return exclusiveTimes
}

type logEntry struct {
	fID       int
	isStart   bool
	timestamp int
}

type execution struct {
	fID           int
	latestStartAt int
}

// Assume all log strings are well-formed, so there aren't any errors parsing it.
func logEntryFromString(s string) logEntry {
	entry := logEntry{}
	parts := strings.Split(s, ":")
	entry.fID, _ = strconv.Atoi(parts[0])
	entry.isStart = parts[1] == "start"
	entry.timestamp, _ = strconv.Atoi(parts[2])

	return entry
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
		{
			n:        2,
			logs:     []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"},
			expected: []int{7, 1},
		},
		{
			n:        1,
			logs:     []string{},
			expected: []int{},
		},
	}
	for _, tc := range ts {
		actual := exclusiveTime(tc.n, tc.logs)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For (%v, %v) expected %v but got %v\n", tc.n, tc.logs, tc.expected, actual)
		}
	}
}
