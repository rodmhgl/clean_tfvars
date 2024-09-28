package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ParseTFVarsFile parses the tfvars file and returns a map of variables
func ParseTFVarsFile(filePath string, verbose bool) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	vars := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "#") {
			// Ignore empty lines and comments
			continue
		}
		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			varName := strings.TrimSpace(parts[0])
			varValue := strings.TrimSpace(parts[1])

			// Remove quotes from the value if present
			varValue = strings.Trim(varValue, `"`)

			vars[varName] = varValue
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if verbose {
		fmt.Printf("Variables found in tfvars file %s: %v\n", filePath, vars)
	}

	return vars, nil
}

// UpdateTFVarsFile removes unused variables from the tfvars file
func UpdateTFVarsFile(filePath string, unusedVars []string, verbose bool) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	unusedVarsSet := make(map[string]bool)
	for _, v := range unusedVars {
		unusedVarsSet[v] = true
	}

	var updatedLines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "#") {
			// Keep comments and empty lines
			updatedLines = append(updatedLines, line)
			continue
		}
		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			varName := strings.TrimSpace(parts[0])
			if _, isUnused := unusedVarsSet[varName]; isUnused {
				if verbose {
					fmt.Printf("Removing unused variable: %s\n", varName)
				}
				continue // Skip this line
			}
		}
		updatedLines = append(updatedLines, line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// Write the updated content back to the tfvars file
	outputFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	if verbose {
		fmt.Println("Updated lines to be written to the tfvars file:")
	}
	for _, line := range updatedLines {
		if verbose {
			fmt.Println(line)
		}
		_, err = outputFile.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
