package parser

import (
	"errors"
	"fmt"
	"strings"

	"example.com/2022_aoc_d5/stack"
)

// PARSER EXPORTED ================================================================================

type orderT int

const (
	FoFi orderT = iota
	LoFi
)

type Parser struct {
	state           parserState
	order           orderT
	stackWarehouse  []stack.Stack
	stringWarehouse []string
}

func (this *Parser) New(order orderT) error {
	if this.state != start {
		return errors.New("Called \"New()\" on already initialized Parser")
	}
	this.order = order
	this.state = initialized
	return nil
}

func (this *Parser) ParseLine(line string) error {
	if this.state == start {
		return errors.New("Called \"ParseLine()\" on not initialized Parser")

	} else if this.state == done {
		return errors.New("Called \"ParseLine()\" on done Parser")

	} else if this.state == fail {
		return errors.New("Called \"ParseLine()\" on failed Parser")

	} else if this.state == initialized {
		strippedLine := strings.ReplaceAll(line, " ", "")
		if line == "" {
			err := this.transformWarehouse()
			if err != nil {
				this.state = fail
				errorMessage := fmt.Sprintf("Error ocurred in \"ParseLine\" %s", err.Error())
				return errors.New(errorMessage)
			}
			this.state = warehouseDone
		} else if strippedLine[0] == '[' {
			this.stringWarehouse = append(this.stringWarehouse, line)
		} else if strippedLine[0] == '1' {
			// skip for now maybe integrity check later
		}
		return nil

	} else if this.state == warehouseDone {
		instruction, err := parseInstruction(line)
		if err != nil {
			this.state = fail
			return err
		}

		err = this.executeInstruction(instruction)
		if err != nil {
			this.state = fail
			return err
		}

		return nil

	} else {
		this.state = fail
		return errors.New("Called \"ParseLine()\" on Parser with corrupted state")
	}
}

func (this *Parser) Done() (string, error) {
	this.state = done

	var result string

	for i := 0; i < len(this.stackWarehouse); i += 1 {
		char, err := this.stackWarehouse[i].Top()
		if err != nil {
			return "", err
		}
		result += string(char)
	}

	return result, nil
}

// PARSER STATIC ==================================================================================

type parserState uint8

const (
	start parserState = iota
	initialized
	warehouseDone
	done
	fail
)

func (this *Parser) executeInstruction(instruction instructionT) error {
	if instruction.from > len(this.stackWarehouse) && instruction.from < 0 {
		return errors.New("instruction.from is out of range")
	}
	if instruction.to > len(this.stackWarehouse) && instruction.to < 0 {
		return errors.New("instruction.to is out of range")
	}
	if instruction.count > this.stackWarehouse[instruction.from].Len() {
		return errors.New("instruction.count is bigger than crates on stack")
	}

	if this.order == FoFi {
		for i := 0; i < instruction.count; i += 1 {
			value, err := this.stackWarehouse[instruction.from].Top()
			if err != nil {
				errorMessage := fmt.Sprintf("Error occurred while executing instructionT %s : %s", instruction, err.Error())
				return errors.New(errorMessage)
			}

			err = this.stackWarehouse[instruction.from].Pop()
			if err != nil {
				errorMessage := fmt.Sprintf("Error occurred while executing instructionT %s : %s", instruction, err.Error())
				return errors.New(errorMessage)
			}

			err = this.stackWarehouse[instruction.to].Push(value)
			if err != nil {
				errorMessage := fmt.Sprintf("Error occurred while executing instructionT %s : %s", instruction, err.Error())
				return errors.New(errorMessage)
			}
		}
	} else if this.order == LoFi {
		crates := make([]byte, instruction.count, instruction.count)

		for i := 0; i < instruction.count; i += 1 {
			value, err := this.stackWarehouse[instruction.from].Top()
			if err != nil {
				errorMessage := fmt.Sprintf("Error occurred while executing instructionT %s : %s", instruction, err.Error())
				return errors.New(errorMessage)
			}

			crates[instruction.count-1-i] = value

			err = this.stackWarehouse[instruction.from].Pop()
			if err != nil {
				errorMessage := fmt.Sprintf("Error occurred while executing instructionT %s : %s", instruction, err.Error())
				return errors.New(errorMessage)
			}
		}

		for i := 0; i < instruction.count; i += 1 {
			err := this.stackWarehouse[instruction.to].Push(crates[i])
			if err != nil {
				errorMessage := fmt.Sprintf("Error occurred while executing instructionT %s : %s", instruction, err.Error())
				return errors.New(errorMessage)
			}
		}
	}
	return nil
}

func (this *Parser) transformWarehouse() error {
	for i := 1; i < len(this.stringWarehouse); i += 1 {
		if len(this.stringWarehouse[i-1]) != len(this.stringWarehouse[i]) {
			return errors.New("Strings in stringWarehouse are not of equal lentgh")
		}
	}

	for i := 0; i < len(this.stringWarehouse); i += 1 {
		this.stringWarehouse[i] += " "
	}

	if len(this.stringWarehouse[0])%4 != 0 {
		return errors.New("Strings in stringWarehouse are not of length + 1 % 4 == 0")
	}

	numberOfStacks := len(this.stringWarehouse[0]) / 4
	this.stackWarehouse = make([]stack.Stack, numberOfStacks, numberOfStacks)
	for i := 0; i < numberOfStacks; i += 1 {
		this.stackWarehouse[i].New(20)
	}

	for rowIndex := len(this.stringWarehouse) - 1; rowIndex >= 0; rowIndex -= 1 {
		for colIndex := 0; colIndex < len(this.stringWarehouse[rowIndex]); colIndex += 4 {
			stackIndex := colIndex / 4
			this.stackWarehouse[stackIndex].Push(this.stringWarehouse[rowIndex][colIndex+1])
		}
	}

	for i := 0; i < len(this.stackWarehouse); i += 1 {
		for true {
			char, err := this.stackWarehouse[i].Top()
			if err != nil {
				errorMessage := fmt.Sprintf("Error Occurred while transforming the warehouse : %s", err.Error())
				return errors.New(errorMessage)
			}
			if char == ' ' {
				err := this.stackWarehouse[i].Pop()
				if err != nil {
					errorMessage := fmt.Sprintf("Error Occurred while transforming the warehouse : %s", err.Error())
					return errors.New(errorMessage)
				}
			} else {
				break
			}
		}

	}

	return nil
}

// STATIC =========================================================================================

type instructionT struct {
	count int
	from  int
	to    int
}

func parseInstruction(line string) (instructionT, error) {
	var count, to, from int
	_, err := fmt.Sscanf(line, "move %d from %d to %d", &count, &from, &to)
	if err != nil {
		errorMessage := fmt.Sprintf("Error occurred while parsing instruction line \"%s\" : %s", line, err.Error())
		return instructionT{0, 0, 0}, errors.New(errorMessage)
	}
	return instructionT{count, from - 1, to - 1}, nil
}
