package cos418_hw1_1

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

// Sum numbers from channel `nums` and output sum to `out`.
// You should only output to `out` once.
// Do NOT modify function signature.
func sumWorker(nums chan int, out chan int) {
	var result int
	for v := range nums {
		result += v
	}
	out <- result
}

// Asynchronous function to read the numbers
func numberReader(out chan int, nums []int) {
	for _, v := range nums {
		out <- v
	}
	close(out)
}

func getResults(resultChannel chan int, num int) int {
	var result = 0
	for i := 0; i < num; i++ {
		value := <-resultChannel
		result += value
	}
	return result
}

// Read integers from the file `fileName` and return sum of all values.
// This function must launch `num` go routines running
// `sumWorker` to find the sum of the values concurrently.
// You should use `checkError` to handle potential errors.
// Do NOT modify function signature.
func sum(num int, fileName string) int {
	dat, err1 := os.Open(fileName)
	checkError(err1)
	r := bufio.NewReader(dat)
	var nums, err2 = readInts(r)
	checkError(err2)
	var capacity = len(nums) / num
	inChannel := make(chan int)
	outChannel := make(chan int, capacity)
	for i := 0; i < num; i++ {
		go sumWorker(inChannel, outChannel)
	}
	go numberReader(inChannel, nums)
	var result = getResults(outChannel, num)
	return result
}

// Read a list of integers separated by whitespace from `r`.
// Return the integers successfully read with no error, or
// an empty slice of integers and the error that occurred.
// Do NOT modify this function.
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var elems []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return elems, err
		}
		elems = append(elems, val)
	}
	return elems, nil
}
