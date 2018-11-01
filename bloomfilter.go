package datastructure

import (
	"hash"
	"hash/fnv"
)

type BloomFilter struct {
	size   uint
	bitset []bool
	hashes []hash.Hash64
}

func CreateBloomFilter(size uint) *BloomFilter {
	return &BloomFilter{
		bitset: make([]bool, size),
		size:   size,
		hashes: []hash.Hash64{fnv.New64(), fnv.New64a()},
	}
}

func (bloomfilter *BloomFilter) Add(item string) {
	bytes := []byte(item)
	for _, hash := range bloomfilter.hashes {
		hashValue := GetHash(hash, bytes)
		position := uint(hashValue) % bloomfilter.size
		///fmt.Printf("%d,%d,%d\n", uint(hashValue), position, bloomfilter.size)
		bloomfilter.bitset[position] = true
	}
	//fmt.Printf("%v\n", bloomfilter.bitset)
}

func (bloomfilter *BloomFilter) Exists(item string) bool {
	exists := true
	bytes := []byte(item)
	for _, hash := range bloomfilter.hashes {
		hashValue := GetHash(hash, bytes)
		position := uint(hashValue) % bloomfilter.size
		//fmt.Printf("%d,%d,%d\n", uint(hashValue), position, bloomfilter.size)
		if !bloomfilter.bitset[position] {
			exists = false
			break
		}
	}
	return exists
}

func GetHash(hash hash.Hash64, data []byte) uint64 {
	hash.Write(data)
	result := hash.Sum64()
	hash.Reset()
	return result
}
