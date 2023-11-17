package types

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Books []Book

func (b *Books) DeCodeJSON(jsonData []byte) {
	reader := bytes.NewReader(jsonData)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&b); err != nil {
		fmt.Println("Ошибка декодирования:", err)
	}
}

func (b *Books) EnCodeJSON() []byte {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	if err := encoder.Encode(b); err != nil {
		fmt.Println("Ошибка при записи JSON-данных:", err)
		return nil
	}
	return buf.Bytes()
}
