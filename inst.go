package main

import (
	"fmt"
)

type Opcode byte

type Instruction struct {
	Code         Opcode
	Mnemonic     string
	InCount      int
	OutCount     int
	ConsumeCount int
	Description  string
}

var InstructionSet [0x100]Instruction

func init() {
	InstructionSet[0x00] = Instruction{0x00, "STOP", 0, 0, 0, "Halts execution"}
	InstructionSet[0x01] = Instruction{0x01, "ADD", 2, 1, 0, "Addition operation"}
	InstructionSet[0x02] = Instruction{0x02, "MUL", 2, 1, 0, "Multiplication operation"}
	InstructionSet[0x03] = Instruction{0x03, "SUB", 2, 1, 0, "Subtraction operation"}
	InstructionSet[0x04] = Instruction{0x04, "DIV", 2, 1, 0, "Integer division operation"}
	InstructionSet[0x05] = Instruction{0x05, "SDIV", 2, 1, 0, "Signed integer division operation"}
	InstructionSet[0x06] = Instruction{0x06, "MOD", 2, 1, 0, "Modulo remainder operation"}
	InstructionSet[0x07] = Instruction{0x07, "SMOD", 2, 1, 0, "Signed modulo remainder operation"}
	InstructionSet[0x08] = Instruction{0x08, "ADDMOD", 3, 1, 0, "Modulo addition operation"}
	InstructionSet[0x09] = Instruction{0x09, "MULMOD", 3, 1, 0, "Modulo multiplication operation"}
	InstructionSet[0x0a] = Instruction{0x0a, "EXP", 2, 1, 0, "Exponential operation"}
	InstructionSet[0x0b] = Instruction{0x0b, "SIGNEXTEND", 2, 1, 0, "Extend length of two's complement signed integer"}

	InstructionSet[0x10] = Instruction{0x10, "LT", 2, 1, 0, "Less-than comparison"}
	InstructionSet[0x11] = Instruction{0x11, "GT", 2, 1, 0, "Greater-than comparison"}
	InstructionSet[0x12] = Instruction{0x12, "SLT", 2, 1, 0, "Signed less-than comparison"}
	InstructionSet[0x13] = Instruction{0x13, "SGT", 2, 1, 0, "Signed greater-than comparison"}
	InstructionSet[0x14] = Instruction{0x14, "EQ", 2, 1, 0, "Equality comparison"}
	InstructionSet[0x15] = Instruction{0x15, "ISZERO", 1, 1, 0, "Simple not operator"}
	InstructionSet[0x16] = Instruction{0x16, "AND", 2, 1, 0, "Bitwise AND operator"}
	InstructionSet[0x17] = Instruction{0x17, "OR", 2, 1, 0, "Bitwise OR operator"}
	InstructionSet[0x18] = Instruction{0x18, "XOR", 2, 1, 0, "Bitwise XOR operator"}
	InstructionSet[0x19] = Instruction{0x19, "NOT", 1, 1, 0, "Bitwise NOT operator"}
	InstructionSet[0x1a] = Instruction{0x1a, "BYTE", 2, 1, 0, "Retrieve single byte from word"}

	InstructionSet[0x20] = Instruction{0x20, "SHA3", 2, 1, 0, "Compute Keccak-256 hash"}

	InstructionSet[0x30] = Instruction{0x30, "ADDRESS", 0, 1, 0, "Get address of currently executing account"}
	InstructionSet[0x31] = Instruction{0x31, "BALANCE", 1, 1, 0, "Get balance of the given account"}
	InstructionSet[0x32] = Instruction{0x32, "ORIGIN", 0, 1, 0, "Get execution origination address"}
	InstructionSet[0x33] = Instruction{0x33, "CALLER", 0, 1, 0, "Get caller address"}
	InstructionSet[0x34] = Instruction{0x34, "CALLVALUE", 0, 1, 0, "Get deposited value by the instruction/transaction responsible for this execution"}
	InstructionSet[0x35] = Instruction{0x35, "CALLDATALOAD", 1, 1, 0, "Get input data of current environment"}
	InstructionSet[0x36] = Instruction{0x36, "CALLDATASIZE", 0, 1, 0, "Get size of input data in current environment"}
	InstructionSet[0x37] = Instruction{0x37, "CALLDATACOPY", 3, 0, 0, "Copy input data in current environment to memory"}
	InstructionSet[0x38] = Instruction{0x38, "CODESIZE", 0, 1, 0, "Get size of running code in current environment"}
	InstructionSet[0x39] = Instruction{0x39, "CODECOPY", 3, 0, 0, "Copy code running in current environment to memory"}
	InstructionSet[0x3a] = Instruction{0x3a, "GASPRICE", 0, 1, 0, "Get price of gas in current environment"}
	InstructionSet[0x3b] = Instruction{0x3b, "EXTCODESIZE", 1, 1, 0, "Get size of an account's code"}
	InstructionSet[0x3c] = Instruction{0x3c, "EXTCODECOPY", 4, 0, 0, "Copy an account's code to memory"}

	InstructionSet[0x40] = Instruction{0x40, "BLOCKHASH", 1, 1, 0, "Get the hash of one of the 256 most recent complete blocks"}
	InstructionSet[0x41] = Instruction{0x41, "COINBASE", 0, 1, 0, "Get the block's beneficiary address"}
	InstructionSet[0x42] = Instruction{0x42, "TIMESTAMP", 0, 1, 0, "Get the block's timestamp"}
	InstructionSet[0x43] = Instruction{0x43, "NUMBER", 0, 1, 0, "Get the block's number"}
	InstructionSet[0x44] = Instruction{0x44, "DIFFICULTY", 0, 1, 0, "Get the block's difficulty"}
	InstructionSet[0x45] = Instruction{0x45, "GASLIMIT", 0, 1, 0, "Get the block's gas limit"}

	InstructionSet[0x50] = Instruction{0x50, "POP", 1, 0, 0, "Remove item from stack"}
	InstructionSet[0x51] = Instruction{0x51, "MLOAD", 1, 1, 0, "Load word from memory"}
	InstructionSet[0x52] = Instruction{0x52, "MSTORE", 2, 0, 0, "Save word to memory"}
	InstructionSet[0x53] = Instruction{0x53, "MSTORE8", 2, 0, 0, "Save byte to memory"}
	InstructionSet[0x54] = Instruction{0x54, "SLOAD", 1, 1, 0, "Load word from storage"}
	InstructionSet[0x55] = Instruction{0x55, "SSTORE", 2, 0, 0, "Store word to storage"}
	InstructionSet[0x56] = Instruction{0x56, "JUMP", 1, 0, 0, "Alter the program counter"}
	InstructionSet[0x57] = Instruction{0x57, "JUMPI", 2, 0, 0, "Conditionally alter the program counter"}
	InstructionSet[0x58] = Instruction{0x58, "PC", 0, 1, 0, "Get the value of the program counter prior to the increment corresponding to this instruction"}
	InstructionSet[0x59] = Instruction{0x59, "MSIZE", 0, 1, 0, "Get the size of active memory in bytes"}
	InstructionSet[0x5a] = Instruction{0x5a, "GAS", 0, 1, 0, "Get the amount of available gas, including the corresponding reduction for the cost of this instruction"}
	InstructionSet[0x5b] = Instruction{0x5b, "JUMPDEST", 0, 0, 0, "Mark a valid destination for jumps"}

	for i := 0; i < 32; i++ {
		code := Opcode(0x60 + i)
		InstructionSet[code] = Instruction{code, fmt.Sprintf("PUSH%d", i+1), 0, 1, i + 1, fmt.Sprintf("Place %d-byte item on stack", i+1)}
	}

	ord := func(i int) string {
		switch i {
		case 1:
			return "1st"
		case 2:
			return "2nd"
		case 3:
			return "3rd"
		default:
			return fmt.Sprintf("%dth", i)
		}
	}

	for i := 0; i < 16; i++ {
		code := Opcode(0x80 + i)
		InstructionSet[code] = Instruction{code, fmt.Sprintf("DUP%d", i+1), i + 1, i + 2, 0, fmt.Sprintf("Duplicate %s stack item", ord(i+1))}
	}

	for i := 0; i < 16; i++ {
		code := Opcode(0x90 + i)
		InstructionSet[code] = Instruction{code, fmt.Sprintf("SWAP%d", i+1), i + 2, i + 2, 0, fmt.Sprintf("Exchange 1st and %s stack items", ord(i+1))}
	}

	InstructionSet[0xa0] = Instruction{0xa0, "LOG0", 2, 0, 0, "Append log record with no topics"}
	InstructionSet[0xa1] = Instruction{0xa1, "LOG1", 3, 0, 0, "Append log record with one topic"}
	InstructionSet[0xa2] = Instruction{0xa2, "LOG2", 4, 0, 0, "Append log record with two topics"}
	InstructionSet[0xa3] = Instruction{0xa3, "LOG3", 5, 0, 0, "Append log record with three topics"}
	InstructionSet[0xa4] = Instruction{0xa4, "LOG4", 6, 0, 0, "Append log record with four topics"}

	InstructionSet[0xf0] = Instruction{0xf0, "CREATE", 3, 1, 0, "Create a new account with associated code"}
	InstructionSet[0xf1] = Instruction{0xf1, "CALL", 7, 1, 0, "Message-call into an account"}
	InstructionSet[0xf2] = Instruction{0xf2, "CALLCODE", 7, 1, 0, "Message-call into this account with an alternative account's code"}
	InstructionSet[0xf3] = Instruction{0xf3, "RETURN", 2, 0, 0, "Halt execution returning output data"}
	InstructionSet[0xf4] = Instruction{0xf4, "DELEGATECALL", 6, 1, 0, "Message-call into this account with an alternative account's code, but persisting the current values for _sender_ and _value_"}
	InstructionSet[0xff] = Instruction{0xff, "SUICIDE", 1, 0, 0, "Halt execution and register account for later deletion"}

	for i := 0; i < 0x100; i++ {
		if InstructionSet[i].Mnemonic == "" {
			InstructionSet[i] = Instruction{Opcode(i), "INVALID", 0, 0, 0, "Invalid instruction"}
		}
	}
}

// func (o Opcode) String() string {
// 	inst := InstructionSet[o]
// 	if inst.Mnemonic == "" {
// 		return "INVALID"
// 	}
// 	return inst.Mnemonic
// }
