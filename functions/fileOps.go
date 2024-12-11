package functions

import (
	"fmt"
	"os"
)

func FileAsString(fileName string) string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %s", fileName)
		panic(err)
	}
	return string(data)
}
