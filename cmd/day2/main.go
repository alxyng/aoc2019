package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/nullseed/aoc2019/pkg/intcode"
)

func main() {
	m := intcode.New()

	prog, err := readProgram("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("start: %v\n", prog)

	progAlarmState := restoreTo1202ProgramAlarmState(prog)
	log.Printf("after state restore: %v\n", progAlarmState)

	m.Load(progAlarmState)
	m.Run()
	mem := m.Memory()

	log.Printf("after halt: %v\n", mem)
	log.Printf("value at addr 0: %v", mem[0])

	output := 19690720
	noun, verb := getNounAndVerbForValue(m, prog, output)
	if noun == -1 || verb == -1 {
		log.Fatalf("could not find noun and verb for %v", output)
	}

	log.Printf("for output: %v, use noun: %v, verb: %v", output, noun, verb)
	log.Printf("100 * %v + %v = %v", noun, verb, 100*noun+verb)
}

func readProgram(filename string) (intcode.Program, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	intStrs := strings.Split(string(data), ",")

	prog := make(intcode.Program, len(intStrs))
	for i, intStr := range intStrs {
		intInt, err := strconv.Atoi(intStr)
		if err != nil {
			return nil, fmt.Errorf("error parsing file: %w", err)
		}
		prog[i] = intInt
	}

	return prog, nil
}

func restoreTo1202ProgramAlarmState(p intcode.Program) intcode.Program {
	return modifyProgram(p, 12, 2)
}

func getNounAndVerbForValue(m *intcode.Machine, p intcode.Program, result int) (int, int) {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			m.Load(modifyProgram(p, noun, verb))
			m.Run()
			if m.Memory()[0] == result {
				return noun, verb
			}
		}
	}

	return -1, -1
}

func modifyProgram(p intcode.Program, noun int, verb int) intcode.Program {
	pp := make([]int, len(p))
	copy(pp, []int(p))

	pp[1] = noun
	pp[2] = verb

	return pp
}
