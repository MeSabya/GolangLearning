package leetcodesolutions_test

import "testing"

func TestReachable(t *testing.T) {
	canreach := CanReach([]int{4, 2, 3, 0, 3, 1, 2}, 5)
	if canreach != true {
		t.Errorf("unexpected reachability expected: true got:", canreach)
	}

}
