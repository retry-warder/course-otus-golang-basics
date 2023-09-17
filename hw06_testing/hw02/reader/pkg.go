package hw02_reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	hw02_types "github.com/retry-warder/course-otus-golang-basics/hw06_testing/hw02/types"
)

func ReadJSON(filePath string) ([]hw02_types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	defer f.Close()
	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	var data []hw02_types.Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	return data, nil
}
