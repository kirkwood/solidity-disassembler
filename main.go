package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com.old/ethereum/go-ethereum/crypto/sha3"
	"io/ioutil"
	"os"
)

func prepareInput(b []byte) ([]byte, error) {
	if len(b) > 1 && b[0] == '0' && b[1] == 'x' {
		b = b[2:]
	}

	bytecode := make([]byte, hex.DecodedLen(len(b)))
	_, err := hex.Decode(bytecode, b)
	return bytecode, err
}

func decode(code []byte) {
	i := 0
	L := len(code)
	for i < L {
		op := code[i]
		inst := InstructionSet[op]
		fmt.Printf("% 4d 0x%08x 0x%02x %-12s", i, i, op, inst.Mnemonic)
		if inst.ConsumeCount > 0 {
			consumed := make([]byte, inst.ConsumeCount)
			if i+inst.ConsumeCount+1 >= L {
				copy(consumed, code[i+1:])
			} else {
				copy(consumed, code[i+1:i+inst.ConsumeCount+1])
			}
			fmt.Printf(" 0x%02x", consumed)
		}
		fmt.Printf("\n")
		i += 1 + inst.ConsumeCount
	}
}

func main() {
	flag.Parse()

	for _, filename := range flag.Args() {
		input, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}
		b, err := prepareInput(input)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		decode(b)
	}
}
