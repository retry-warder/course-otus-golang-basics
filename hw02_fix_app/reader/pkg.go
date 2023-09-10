package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/retry-warder/course-otus-golang-basics/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
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

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	return data, nil
}
