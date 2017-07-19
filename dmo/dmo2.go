package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/ruffrey/recurrent-nn-char-go/mat32"
)

// meta params

// discard is the rounds to discard when warming up the network
var discard = 50

// N is the total reservoir units/cells
var N = 200

// algorithm vars

// K is the total input cells
var K = 0

// L is the total output cells
var L = 0

// xn - x(n) - N-dimensional reservoir state
var xn *mat32.Mat

// W is the NxN reservoir weight matrix
var W *mat32.Mat

// n is time; also current step?
var n int

// Win is the KxN input weight matrix
var Win *mat32.Mat

// un - u(n) is the K-dimensional input signal
var un *mat32.Mat

// Wfb is the NxL feedback weight matrix
var Wfb *mat32.Mat

// yn - y(n) is the L-dimensional output signal
var yn *mat32.Mat

var g = &mat32.Graph{}

var inputCellToValue []char
var inputValueToCell map[char]int
var outputCellToValue []char

// zn is the extended system state
func z(n int) {

}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Reading test data
	testDataFile := os.Args[1]
	fmt.Println("Reading test data file", testDataFile)
	buf, err := ioutil.ReadFile(testDataFile)
	if err != nil {
		fmt.Println("Failed opening test data file, should be first argument but got=", testDataFile)
		panic(err)
	}
	testData := strings.Split(string(buf), "")
	fmt.Println("Preparing test data maps, chars=", len(testData))

	// format test data
	testInputCells := make([]int, len(testData)) // input cell id is the index
	for charIndex, c := range testData {
		character := char(c)
		if _, hasInput := inputValueToCell[character]; !hasInput {
			inputCellToValue = append(inputCellToValue, character)
			inputValueToCell[character] = len(inputCellToValue) - 1
			outputCellToValue = append(outputCellToValue, char(character))
		}
		testInputCells[charIndex] = inputValueToCell[character]
	}
	K = len(inputCellToValue)
	L = len(outputCellToValue)
	fmt.Println("inputValueToCell=", inputValueToCell)
	fmt.Println("inputCellToValue=", inputCellToValue)
	fmt.Println("outputValues=", outputCellToValue)
	// fmt.Println("testInputCells=", testInputCells)

	// Run the algorithm - training mode

	// n is time
	for n = 0; n < 10; n++ {
		xn[n+1] = sigmoid()
	}
}

/* helper utilities */
type char string // single character
