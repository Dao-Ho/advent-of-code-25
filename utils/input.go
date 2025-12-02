package utils

import (
    "os"
    "strings"
)

func ReadInput(filepath string) (string, error) {
    data, err := os.ReadFile(filepath)
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(string(data)), nil
}

// Return arr of lines from file
func ReadLines(filepath string) ([]string, error) {
    input, err := ReadInput(filepath)
    if err != nil {
        return nil, err
    }
    return strings.Split(input, "\n"), nil
}
