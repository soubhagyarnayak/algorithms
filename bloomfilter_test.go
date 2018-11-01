package datastructure

import "testing"

func TestBloomFilter(t *testing.T) {
	bloomfilter := CreateBloomFilter(10)
	bloomfilter.Add("some random string")
	bloomfilter.Add("another string")
	bloomfilter.Add("dummy string")

	if got := bloomfilter.Exists("a somewhat random string"); got != false {
		t.Errorf("Expected no , but got exists")
	}

	//bloom filter does not guarantee this
	if got := bloomfilter.Exists("another string"); got != true {
		t.Errorf("Expected exists , but got no")
	}
}
