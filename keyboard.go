// package keyboard reads user input from the keyboard.
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//GetFloat reads a floating point number from the keyboard.
//This function returns an error value with the number read.
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
