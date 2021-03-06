// passgen - Generates random passwords.
// Copyright (C) 2019 Eric Skoglund

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"crypto/rand"
	"fmt"
	"github.com/pborman/getopt/v2"
	"os"
	"strings"
)

var char_table = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "V", "W", "X", "Y", "Z",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z",
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"!", "\"", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".", "/", ":", ";", "<", ">", "?", "=", "[", "]", "\\", "^", "_", "{", "}", "|",
}

func makePassword(length int, exclude string) string {
	// We make the array bigger so that we have enough
	// bytes if we have to exclude a bunch of them.
	bytes := make([]byte, length*2)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	char_table_len := len(char_table)
	pass := ""
	for _, idx := range bytes {
		next_char := char_table[int(idx)%char_table_len]
		if strings.Index(exclude, next_char) == -1 {
			pass += next_char
			if len(pass) == length {
				return pass
			}
		}
	}

	return pass
}

func main() {
	optHelp := getopt.BoolLong("help", 'h', "Print this help message")
	optLength := getopt.IntLong("length", 'l', 10, "Fixed length of the passwords generated")
	optNumber := getopt.IntLong("number", 'n', 1, "The number of passwords to be generated")
	optExclude := getopt.StringLong("exclude", 'e', "", "Characters to exclude from the password")
	getopt.Parse()

	if *optHelp {
		getopt.Usage()
		os.Exit(0)
	}

	for i := 0; i < *optNumber; i++ {
		pass := makePassword(*optLength, *optExclude)
		fmt.Println(pass)
	}
}
