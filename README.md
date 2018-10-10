# Highest Product
This repo solves the problem of finding the highest product of 3 elements from a given list.
## Example usage
```go
product, err = HighestProduct([]int{1, 10, 2, 6, 5, 3}) // product becomes 300
```

## Tests
To see it in action, run the tests by cloning the repo and doing
```bash
go test
```
This both runs unit tests and a performance analysis

## Variants
Three variants are supplied which all solve the problem. The first two, `HighestProduct` and `HighestProductSlower` both have O(n) time complexity and O(1) space complexity. They differ in that `HighestProduct` just loops over the input list once, checking every element against a list of the 3 highest elements found previously, whereas `HighestProductSlower` loops over the input list three 3 times, every time storing the highest it hasn't seen yet. `HighestProductSlower` is slower since it loops over the input 3 times and also has to iterate over the length-3 list to check if an element is among the 3 highest. 

The last method, `HighestProductUsingSort` is the slowest of them all. It works by sorting the array and picking out the 3 last elements in the sorted result. This gives it a time complexity if O(n lg(n)). Furthermore, since Go slices are only pointers to an array stored in memory, simply sorting the input would have the side effect of sorting the slice given in the input. To circumvent this we have to make a copy of the slice, which makes the spaces complexity O(n). 

## Performance
When running the tests, a simple performance check is ran to compare the three functions. It tests the functions first on a 1000 slices of length 100 and then a huge slice of length 10000000, all of them containing random integers. A dump is shown below that illustrates the runtime observations above. 
```
Running performance test. This might take some time
Creating lists...
Done creating small lists, creating one huge list
Done creating huge list
--- Testing HighestProduct ---
Many lists: 0.555634ms
Huge list: 385.991263ms
--- Testing HighestProductSlower ---
Many lists: 2.127271ms
Huge list: 1951.559142ms
--- Testing HighestProductUsingSort ---
Many lists: 6.993194ms
Huge list: 25916.590181ms
```

For the 1000 small arrays, `HighestProductUsingSort` is about 10 times slower than the fastest, and for the huge 10 million element array it is a staggering 67 times slower. 
