package pipe_filter

import "testing"

func TestPipeFilter(t *testing.T) {
	var data string
	data = "1,2,3"
	splitFilter := NewSplitFilter(",")
	integerFilter := NewIntegerFilter()
	sumFilter := NewSumFilter()
	sp := NewStraightPipeline("p1", splitFilter, integerFilter, sumFilter)
	ret, err := sp.Process(data)
	if err != nil {
		t.Fatal(err)
	}
	if ret != 6 {
		t.Fatalf("the excepted is 6, but the actually value is %d", ret)
	}
	t.Log(ret)
}
