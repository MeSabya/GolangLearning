/*https://leetcode.com/problems/logger-rate-limiter/

Design a logger system that receive stream of messages along with its timestamps, each message should be printed if and only if it is not printed in the last 10 seconds.

Given a message and a timestamp (in seconds granularity), return true if the message
should be printed in the given timestamp, otherwise returns false.

It is possible that several messages arrive roughly at the same time.
*/

/*
Logger logger = new Logger();

// logging string "foo" at timestamp 1
logger.shouldPrintMessage(1, "foo"); returns true;

// logging string "bar" at timestamp 2
logger.shouldPrintMessage(2,"bar"); returns true;

// logging string "foo" at timestamp 3
logger.shouldPrintMessage(3,"foo"); returns false;

// logging string "bar" at timestamp 8
logger.shouldPrintMessage(8,"bar"); returns false;

// logging string "foo" at timestamp 10
logger.shouldPrintMessage(10,"foo"); returns false;

// logging string "foo" at timestamp 11
logger.shouldPrintMessage(11,"foo"); returns true;
*/

package main

/*
Idea :
Track already visited strings.
Now we need to check : if a string is already visited, then what was the time stamp?
if a visited string time stamp is more than 10 need to cleaned up.
So what are the DS needed ?
1. To track the visited strings : we need to store the strings in a set.
2. We need the timestamps of the already visited strings along with the message.
That we can store in a structure and that structure we can store in a double linked
list.
*/

import "fmt"

type Log struct {
	timestamp int
	message   string
}

type Logger struct {
	logset   map[string]bool
	logcache map[string]Log
}

func initLogger() *Logger {
	return &Logger{logset: make(map[string]bool), logcache: make(map[string]Log)}
}

func (logger *Logger) shouldPrintMessage(timestamp int, message string) bool {
	if len(logger.logcache) > 0 {
		for msg, log := range logger.logcache {
			if timestamp-log.timestamp >= 10 {
				delete(logger.logcache, msg)
				delete(logger.logset, msg)
			}
		}
	}

	if _, found := logger.logset[message]; !found {
		logger.logcache[message] = Log{timestamp, message}
		logger.logset[message] = true
		return true
	}

	return false

}

func assertEqual(a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	fmt.Println(message)
}

func main() {
	logger := initLogger()
	ret := logger.shouldPrintMessage(1, "foo")
	if ret != true {
		fmt.Printf("[TC1]Failed expected: true got:%t", ret)
	}

	ret = logger.shouldPrintMessage(2, "bar")
	assertEqual(ret, true, "TC2 failed")
	// logging string "foo" at timestamp 3
	ret = logger.shouldPrintMessage(3, "foo")
	assertEqual(ret, false, "TC3 failed")

	// logging string "bar" at timestamp 8
	ret = logger.shouldPrintMessage(8, "bar")
	assertEqual(ret, false, "TC4 failed")
	// logging string "foo" at timestamp 10
	ret = logger.shouldPrintMessage(10, "foo")
	assertEqual(ret, false, "TC5 failed")
	// logging string "foo" at timestamp 11
	ret = logger.shouldPrintMessage(11, "foo")
	assertEqual(ret, true, "TC6 failed")

}
