package main

import (
	"fmt"
	"hash/fnv"

	"github.com/spaolacci/murmur3"
	cityhash "github.com/zhenjl/cityhash"
)

// Defining the bloom filter data structure
type BloomFilter struct {
	bitArray []bool
	size     int
	k        int
}

// define list of available hash functions
var hashArray [3]string = [3]string{"fnv", "murmur", "cityhash"}

func computeFnvHash(value string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(value))
	return h.Sum32()
}

func computeMurmurHash(value string) uint32 {
	return murmur3.Sum32([]byte(value))
}

func computeCityHash(value string) uint32 {
	return cityhash.CityHash32([]byte(value), uint32(len(value)))
}

// Adding an element to the Bloom filter
func (bf *BloomFilter) Add(element string) {
	var hashOutput uint32
	// Generate hash using k hash functions
	for i := 0; i < bf.k; i++ {
		switch hashArray[i] {
		case "fnv":
			hashOutput = computeFnvHash(element)
		case "murmur":
			hashOutput = computeMurmurHash(element)
		case "cityhash":
			hashOutput = computeCityHash(element)
		default:
			fmt.Println("Invalid hash output")
		}
		// compute the index where bit is to be enabled
		index := hashOutput % uint32(bf.size)
		// set bit to true
		bf.bitArray[index] = true
	}
}

// Check if an element is present in the bloom filter
func (bf *BloomFilter) Contains(element string) bool {
	var hashOutput uint32
	// Generate hash using k hash functions
	for i := 0; i < bf.k; i++ {
		switch hashArray[i] {
		case "fnv":
			hashOutput = computeFnvHash(element)
		case "murmur":
			hashOutput = computeMurmurHash(element)
		case "cityhash":
			hashOutput = computeCityHash(element)
		default:
			fmt.Println("Invalid hash output")
		}
		// compute the index where bit is to be enabled
		index := hashOutput % uint32(bf.size)
		// check if hashed index is ON (true)
		// if any one of the index is OFF (false) then element is not present in the bf
		if !bf.bitArray[index] {
			return false
		}
	}
	// if all bits are ON, then element may be present in the filter
	return true
}

func main() {
	// input size of bloom filter
	m := 10
	// input number of hash functions
	k := 3

	if k > len(hashArray) {
		fmt.Println("Exceeded the max number of hash functions. Max value =", len(hashArray))
		return
	}
	// initializing the bloom filter
	var bf = &BloomFilter{
		bitArray: make([]bool, m),
		k:        k,
		size:     m,
	}
	// Adding elements to bf
	bf.Add("violet")
	bf.Add("red")
	bf.Add("green")
	bf.Add("yellow")
	bf.Add("pink")
	// check element membership
	fmt.Println(bf.Contains("purple"))
	fmt.Println(bf.Contains("yellow"))
}
