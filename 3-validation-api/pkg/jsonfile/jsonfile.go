package jsonfile

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonFile struct {
	FilePath string
}

func NewJsonFile(filePath string) *JsonFile {
	return &JsonFile{
		FilePath: filePath,
	}
}

// // ReadJSON читает данные из JSON-файла и декодирует их в структуру
// func (j *JsonFile) ReadJSON(v interface{}) error {
// 	file, err := os.Open(j.FilePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	return json.NewDecoder(file).Decode(v)
// }

// ReadJSONByKey читает значение из JSON-файла по ключу
func (j *JsonFile) ReadJSONByKey(key string, out interface{}) error {
	file, err := os.Open(j.FilePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Читаем JSON-данные
	var data map[string]json.RawMessage
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	// Ищем запись по ключу
	rawData, exists := data[key]
	if !exists {
		return fmt.Errorf("key '%s' not found", key)
	}

	// Декодируем данные в указанный тип
	if err := json.Unmarshal(rawData, out); err != nil {
		return fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return nil
}

// WriteJSONByKey записывает значение в JSON-файл по ключу
func (j *JsonFile) WriteJSONByKey(key string, value interface{}) error {
	// Шаг 1: Открыть файл для чтения
	file, err := os.Open(j.FilePath)
	if err != nil {
		// Если файл не существует, создаем новый объект
		if os.IsNotExist(err) {
			return j.createNewJSON(key, value)
		}
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Шаг 2: Прочитать существующие данные
	var data map[string]interface{}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	// Шаг 3: Обновить данные
	data[key] = value

	// Шаг 4: Перезаписать файл
	return j.overwriteJSON(data)
}

// Вспомогательная функция: Создание нового JSON-файла
func (j *JsonFile) createNewJSON(key string, value interface{}) error {
	data := map[string]interface{}{
		key: value,
	}
	return j.overwriteJSON(data)
}

// Вспомогательная функция: Перезапись JSON-файла
func (j *JsonFile) overwriteJSON(data map[string]interface{}) error {
	file, err := os.Create(j.FilePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Добавляем отступы для удобочитаемости
	return encoder.Encode(data)
}

// // AppendToJSON добавляет новый элемент в JSON-массив и сохраняет его в файл
// func (j *JsonFile) AppendToJSON(filePath string, newItem interface{}) error {
// 	// Шаг 1: Открыть файл для чтения
// 	file, err := os.Open(j.FilePath)
// 	if err != nil {
// 		// Если файл не существует, создаем новый массив
// 		if os.IsNotExist(err) {
// 			return j.WriteJSON([]interface{}{newItem})
// 		}
// 		return err
// 	}
// 	defer file.Close()

// 	// Шаг 2: Прочитать существующие данные
// 	var existingData []interface{}
// 	if err := json.NewDecoder(file).Decode(&existingData); err != nil {
// 		return err
// 	}

// 	// Шаг 3: Добавить новый элемент
// 	existingData = append(existingData, newItem)

// 	// Шаг 4: Перезаписать файл новыми данными
// 	return j.WriteJSON(existingData)
// }

// // WriteJSON записывает данные в JSON-файл
// func (j *JsonFile) WriteJSON(v interface{}) error {
// 	file, err := os.Create(j.FilePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	encoder := json.NewEncoder(file)
// 	encoder.SetIndent("", "  ") // Добавляем отступы для удобочитаемости
// 	return encoder.Encode(v)
// }
