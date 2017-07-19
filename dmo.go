package main

import (
	"C"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)
import (
	"math"
)

// TODO: make output weights a separate thing

// meta params

var reservoirSize = 200 // N
var stepsPerInput = 100 // n
var discardRounds = 50

var saveFile = "echo.json"

type reservoirCell int
type outputCell int
type inputCell int
type weight float32
type char string

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

	// Test file exists; initialize arrays
	// reservoir lists each reservoir id and its output ids
	// index = reservoirCell
	var reservoir = make([][]reservoirCell, reservoirSize)
	// reservoirFiringState is the current state of each reservoirCell
	// index = reservoirCell
	var reservoirFiringState = make([]weight, reservoirSize)
	// reservoirWeights are the strength of the synapses to its postsynaptic
	// reservoir cells.
	// index = reservoirCell
	var reservoirWeights = make([]weight, reservoirSize)
	// inputReservoirCells lists each input cell's inputs to the reservoir
	var inputReservoirCells [][]reservoirCell
	var inputWeights []weight

	// reservoirToOutput is for quick lookup
	var reservoirToOutput = make(map[reservoirCell][]outputCell)
	// outputToReservoir is the reverse lookup for reservoirToOutput
	var outputToReservoir = make(map[outputCell][]reservoirCell)
	// outputCells where the index is its ID and the string is a single character
	var outputValues []char
	// testInputCells lists the inputs to be fired
	var testInputCells []inputCell
	// inputValueToCell is for looking up an input index from value
	var inputValueToCell = make(map[char]inputCell)
	// inputCellToValue is for looking up an input value from index
	var inputCellToValue []char

	// format test data
	testInputCells = make([]inputCell, len(testData))
	for charIndex, c := range testData {
		character := char(c)
		if _, hasInput := inputValueToCell[character]; !hasInput {
			inputCellToValue = append(inputCellToValue, character)
			inputValueToCell[character] = inputCell(len(inputCellToValue) - 1)
			outputValues = append(outputValues, char(character))
		}
		testInputCells[charIndex] = inputValueToCell[character]
	}
	fmt.Println("inputValueToCell=", inputValueToCell)
	fmt.Println("inputCellToValue=", inputCellToValue)
	fmt.Println("outputValues=", outputValues)
	// fmt.Println("testInputCells=", testInputCells)

	totalInputCells := len(inputValueToCell)
	totalOutputCells := len(outputValues)

	inputReservoirCells = make([][]reservoirCell, totalInputCells)
	inputWeights = make([]weight, totalInputCells)

	// Add synapses

	// each input cell connects to each reservoir cell
	for i := 0; i < totalInputCells; i++ {
		for synapse := 0; synapse < len(reservoir); synapse++ {
			// this input cell will fire a random reservoir cell
			inputReservoirCells[i] = append(inputReservoirCells[i], randReservoirCell(reservoirSize))
		}
		// init weight
		inputWeights[i] = randWeight()
	}

	// each output cell receives from the reservoir
	for i := 0; i < totalOutputCells; i++ {
		oc := outputCell(i)
		for r := 0; r < len(reservoir); r++ {
			// this output cell will be fired by a random reservoir cell
			reservoirToOutput[rr] = append(reservoirToOutput[rr], oc)
			outputToReservoir[oc] = append(outputToReservoir[oc], rr)
		}
	}

	// each reservoir cell has a few random synapses
	for i := 0; i < reservoirSize; i++ {
		for synapse := 0; synapse < reservoirPostsynapticCount; synapse++ {
			// this reservoir cell will fire a random other cell; it is a
			// synapse to a different cell
			reservoir[i] = append(reservoir[i], randReservoirCell(reservoirSize))
		}
		reservoirWeights[i] = randWeight()
	}

	ss := &saveState{
		Reservoir:         reservoir,
		ReservoirWeights:  reservoirWeights,
		ReservoirToOutput: reservoirToOutput,
		InputValueToCell:  inputValueToCell,
		InputCellToValue:  inputCellToValue,
		InputWeights:      inputWeights,
		OutputToReservoir: outputToReservoir,

		reservoirFiringState: reservoirFiringState,
		testInputCells:       testInputCells,
	}

	ss.step()
	ss.save()
}

type saveState struct {
	Reservoir         [][]reservoirCell
	ReservoirWeights  []weight
	ReservoirToOutput map[reservoirCell][]outputCell
	InputValueToCell  map[char]inputCell
	InputCellToValue  []char
	InputWeights      []weight
	OutputToReservoir map[outputCell][]reservoirCell

	reservoirFiringState []weight
	testInputCells       []inputCell
}

/*
step runs one round through all the test input cells and backpropagates the
output cell connections.

reservoirFiringState[n + 1] = sigmoid(
	reservoirFiringState[n] * reservoirWeights[n]
	+ reservoirCell[inputReservoirCells[n + 1]] * inputWeights[n+1]
	+ reservoirToOutput[reservoirWeights[n]]
)
*/
func (ss *saveState) step() {
	for i := 0; i < len(ss.testInputCells); i++ {

	}
}

/*
sample reads an output value from the system

	output[n] = sigmoid(
		outputWeight[n] * (reservoirFiringState[n] + inputWeights[n])
  )
*/
func (ss *saveState) sample(seed string) (result string) {

	return result
}

func (ss *saveState) save() {
	buf, err := json.Marshal(ss)
	if err != nil {
		fmt.Println("Failed marshaling json of saveState", err, ss)
		return
	}
	err = ioutil.WriteFile(saveFile, buf, os.ModePerm)
	if err != nil {
		fmt.Println("Failed saving saveState", err, ss)
	}
}

func loadState(filename string) (ss *saveState, err error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return ss, err
	}
	err = json.Unmarshal(buf, &ss)
	return ss, err
}

func randReservoirCell(max int) reservoirCell {
	return reservoirCell(rand.Intn(max))
}

func randWeight() weight {
	return weight(rand.Float32())
}

func sigmoid(t float32) float32 {
	return float32(1 / (1 + math.Exp(float64(-t))))
}
