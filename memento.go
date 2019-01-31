package memento

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Memento struct {
	Name   string
	Dir    string
	Values map[string]string
}

func NewMemento(name, dirname string) (*Memento, error) {
	// Create the struct.
	memento := &Memento{
		Name:   name,
		Dir:    fmt.Sprintf("%s/%s.csv", dirname, name),
		Values: make(map[string]string),
	}

	// Open the file for reading; create if not already created.
	memFile, err := os.OpenFile(memento.Dir, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("Error opening the memento destination file: %s", err)
	}
	defer memFile.Close()

	// Create the CSV reader.
	reader := csv.NewReader(memFile)

	// Read all of the records.
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Error reading the records from the memento file: %s", err)
	}

	// Iterate and deserialize the records.
	for _, record := range records {
		// Check for the expected length.
		if len(record) < 2 {
			continue
		}

		// Set the key/value into memory for this record.
		memento.Values[record[0]] = record[1]
	}

	return memento, nil
}

//==========//
// Getters
//==========//

// Returns the boolean value for the provided key.
func (m *Memento) GetBool(key string, defValue bool) bool {
	val, exists := m.Values[key]
	if !exists {
		return defValue
	}

	return val == "true"
}

// Returns the float64 value for the provided key.
func (m *Memento) GetFloat(key string, defValue float64) float64 {
	val, exists := m.Values[key]
	if !exists {
		return defValue
	}

	floatVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return defValue
	}

	return floatVal
}

// Returns the integer value for the provided key.
func (m *Memento) GetInt(key string, defValue int) int {
	val, exists := m.Values[key]
	if !exists {
		return defValue
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		return defValue
	}

	return intVal
}

// Returns the string value for the provided key.
func (m *Memento) GetString(key string, defValue string) string {
	val, exists := m.Values[key]
	if !exists {
		return defValue
	}

	return val
}

//==========//
// Setters
//==========//

// Sets the boolean value for the provided key.
func (m *Memento) SetBool(key string, value bool) error {
	var boolStr string
	if value {
		boolStr = "true"
	} else {
		boolStr = "false"
	}

	m.Values[key] = boolStr

	err := storeValues(m.Dir, m.Values)
	if err != nil {
		return err
	}

	return nil
}

// Sets the float64 value for the provided key.
func (m *Memento) SetFloat(key string, value float64) error {
	floatStr := strconv.FormatFloat(value, 'E', -1, 64)

	m.Values[key] = floatStr

	err := storeValues(m.Dir, m.Values)
	if err != nil {
		return err
	}

	return nil
}

// Sets the integer value for the provided key.
func (m *Memento) SetInt(key string, value int) error {
	intStr := strconv.Itoa(value)

	m.Values[key] = intStr

	err := storeValues(m.Dir, m.Values)
	if err != nil {
		return err
	}

	return nil
}

// Sets the string value for the provided key.
func (m *Memento) SetString(key string, value string) error {
	m.Values[key] = value

	err := storeValues(m.Dir, m.Values)
	if err != nil {
		return err
	}

	return nil
}

//==========//
// Util
//==========//

func getCSVWriter(dirname string) (*csv.Writer, error) {
	memFile, err := os.OpenFile(dirname, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return nil, fmt.Errorf("Error creating the memento file: %s %s", err, dirname)
	}
	defer memFile.Close()

	return csv.NewWriter(memFile), nil
}

func storeValues(dirname string, values map[string]string) error {
	writer, err := getCSVWriter(dirname)
	if err != nil {
		return fmt.Errorf("Error getting the memento CSV writer: %s", err)
	}
	defer writer.Flush()

	for key, value := range values {
		writer.Write([]string{key, value})
	}

	return nil
}
