package types

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func DeCodeJSON(b []*Book, jsonData []byte) {
	reader := bytes.NewReader(jsonData)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&b); err != nil {
		fmt.Println("Ошибка декодирования:", err)
	}
}

func EnCodeJSON(b []Book) []byte {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	if err := encoder.Encode(b); err != nil {
		fmt.Println("Ошибка при записи JSON-данных:", err)
		return nil
	}
	return buf.Bytes()
}
