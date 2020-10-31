package jumpgameIII

import "testing"

func TestReachable(t *testing.T) {
	canreach := CanReach([]int{4, 2, 3, 0, 3, 1, 2}, 5)
	if canreach != true {
		t.Errorf("[TC1] unexpected reachability expected: true got:%t", canreach)
	}

	canreach = CanReach([]int{3, 0, 2, 1, 2}, 2)
	if canreach != false {
		t.Errorf("[TC2] unexpected reachability expected: true got:%t", canreach)
	}

}
