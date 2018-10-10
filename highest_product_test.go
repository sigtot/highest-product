package highest_product

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var numberLists = [][]int{
	{1, 10, 2, 6, 5, 3},
	{5, 2, 7, 6, 1},
	{4, 2, 1, 6, 9, 4},
	{-4,-7,-1,-2,-9,-10,-15,-5},
	{2,5},
	{1,2,0,-1},
	{2, 2, -1, 2},
	{1, 1, 5},
}

var expected = []int{300, 210, 216, -8, 0, 0, 8, 5}

type highestProductFunc func(numbers []int) (int, error)

func TestHighestProduct(t *testing.T) {
	fmt.Println("Testing HighestProduct")
	test(t, HighestProduct)
}

func TestHighestProductSlower(t *testing.T) {
	fmt.Println("Testing HighestProductSlower")
	test(t, HighestProductSlower)
}

func TestHighestProductUsingSort(t *testing.T) {
	fmt.Println("Testing HighestProductUsingSort")
	test(t, HighestProductUsingSort)
}

func test(t *testing.T, f highestProductFunc) {
	for i := 0; i < len(numberLists); i++ {
		product, err := f(numberLists[i])
		if err != nil {
			fmt.Println("List", numberLists[i], "is too short:", err)
			continue
		}
		fmt.Printf("%d -> %d", numberLists[i], product)
		if product == expected[i] {
			fmt.Printf(" (OK)\n")
		} else {
			t.Fail()
			fmt.Printf(" (FAIL) (Expected %d)\n", expected[i])
		}
	}
}

func TestPerformance(t *testing.T) {
	const ShortListLength 	= 100
	const NumShortLists 	= 1000
	const HugeListLength 	= 100000000

	const NanoToMilli = 1000000

	fmt.Println("Running performance test. This might take some time")


	/* --- Prepare lists --- */
	fmt.Println("Creating lists...")
	// Create 1000 small lists
	var shortNumberLists [NumShortLists][ShortListLength]int
	for i := 0; i < NumShortLists; i++ {
		for j := 0; j < ShortListLength; j++ {
			shortNumberLists[i][j] = rand.Int()
		}
	}

	fmt.Println("Done creating small lists, creating one huge list")

	// Create one huge list
	var hugeList [HugeListLength]int
	for i := 0; i < HugeListLength; i++ {
		hugeList[i] = rand.Int()
	}
	fmt.Println("Done creating huge list")

	/* --- Run performance test --- */

	fmt.Println("--- Testing HighestProduct ---")
	startTime := time.Now().UnixNano()
	for i := 0; i < NumShortLists; i++ {
		_, _ = HighestProduct(shortNumberLists[i][:])
	}
	endTime := time.Now().UnixNano()
	fmt.Printf("Many lists: %fms\n", float64(endTime - startTime) / NanoToMilli)

	startTime = time.Now().UnixNano()
	_, _ = HighestProduct(hugeList[:])
	endTime = time.Now().UnixNano()
	fmt.Printf("Huge list: %fms\n", float64(endTime - startTime) / NanoToMilli)


	fmt.Println("--- Testing HighestProductSlower ---")
	startTime = time.Now().UnixNano()
	for i := 0; i < NumShortLists; i++ {
		_, _ = HighestProductSlower(shortNumberLists[i][:])
	}
	endTime = time.Now().UnixNano()
	fmt.Printf("Many lists: %fms\n", float64(endTime - startTime) / NanoToMilli)

	startTime = time.Now().UnixNano()
	_, _ = HighestProductSlower(hugeList[:])
	endTime = time.Now().UnixNano()
	fmt.Printf("Huge list: %fms\n", float64(endTime - startTime) / NanoToMilli)


	fmt.Println("--- Testing HighestProductUsingSort ---")
	startTime = time.Now().UnixNano()
	for i := 0; i < NumShortLists; i++ {
		_, _ = HighestProductUsingSort(shortNumberLists[i][:])
	}
	endTime = time.Now().UnixNano()
	fmt.Printf("Many lists: %fms\n", float64(endTime - startTime) / NanoToMilli)

	startTime = time.Now().UnixNano()
	_, _ = HighestProductUsingSort(hugeList[:])
	endTime = time.Now().UnixNano()
	fmt.Printf("Huge list: %fms\n", float64(endTime - startTime) / NanoToMilli)
}
