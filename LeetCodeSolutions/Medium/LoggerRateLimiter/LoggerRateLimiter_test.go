package loggerratelimiter

import "testing"

func TestLogger(t *testing.T) {
	Logger * logger = initLogger()
	ret := logger.shouldPrintMessage(1, "foo")
	if ret != true {
		t.Errorf("[TC1]Failed expected: true got:%t", ret)
	}
}
