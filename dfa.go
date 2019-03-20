package dfa

import "fmt"

type transitionInput struct {
	srcState int
	input    string
}

type DFA struct {
	initState    int
	currentState int
	totalStates  []int
	finalStates  []int
	transition   map[transitionInput]int
	inputMap     map[string]bool
}

//New a new DFA
func NewDFA(initState int, isFinal bool) *DFA {
	retDFA := &DFA{
		transition:   make(map[transitionInput]int),
		inputMap:     make(map[string]bool),
		initState:    initState,
		currentState: initState}

	retDFA.AddState(initState, isFinal)
	return retDFA
}

//AddState adds new state in this DFA
func (d *DFA) AddState(state int, isFinal bool) {
	if state == -1 {
		fmt.Println("Cannot add state as -1, it is dead state")
		return
	}

	d.totalStates = append(d.totalStates, state)
	if isFinal {
		d.finalStates = append(d.finalStates, state)
	}
}

//AddTransition adds new transition function into DFA
func (d *DFA) AddTransition(srcState int, input string, dstStateint int) {
	find := false

	for _, v := range d.totalStates {
		if v == srcState {
			find = true
		}
	}

	if !find {
		fmt.Println("No such state:", srcState, " in current DFA")
		return
	}

	//find input if exist in DFA input List
	if _, ok := d.inputMap[input]; !ok {
		//not exist, new input in this DFA
		d.inputMap[input] = true
	}

	targetTrans := transitionInput{srcState: srcState, input: input}
	d.transition[targetTrans] = dstStateint
}

func (d *DFA) Input(testInput string) int {
	intputTrans := transitionInput{srcState: d.currentState, input: testInput}
	if val, ok := d.transition[intputTrans]; ok {
		d.currentState = val
		return val
	} else {
		return -1 //dead state
	}
}

//To verify current state if it is final state
func (d *DFA) Verify() bool {
	for _, val := range d.finalStates {
		if val == d.currentState {
			return true
		}
	}
	return false
}

//Reset DFA state to initilize state, but all state and transition function will remain
func (d *DFA) Reset() {
	d.currentState = d.initState
}

//Verify if list of input could be accept by DFA or not
func (d *DFA) VerifyInputs(inputs []string) bool {
	for _, v := range inputs {
		d.Input(v)
	}
	return d.Verify()
}

//To print detail transition table contain of current DFA
func (d *DFA) PrintTransitionTable() {
	fmt.Println("===================================================")
	//list all inputs
	var inputList []string
	for key, _ := range d.inputMap {
		fmt.Printf("\t%s|", key)
		inputList = append(inputList, key)
	}

	fmt.Printf("\n")
	fmt.Println("---------------------------------------------------")

	for _, state := range d.totalStates {
		fmt.Printf("%d |", state)
		for _, key := range inputList {
			checkInput := transitionInput{srcState: state, input: key}
			if dstState, ok := d.transition[checkInput]; ok {
				fmt.Printf("\t %d|", dstState)
			} else {
				fmt.Printf("\tNA|")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Println("---------------------------------------------------")
	fmt.Println("===================================================")
}
