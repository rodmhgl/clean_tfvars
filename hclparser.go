package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// ParseDeclaredVars parses all .tf files and extracts declared variables
func ParseDeclaredVars(directory string, verbose bool) ([]string, error) {
	var declaredVars []string

	// Walk through the directory to find .tf files
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".tf" {
			vars, err := parseDeclaredVarsInFile(path, verbose)
			if err != nil {
				return err
			}
			declaredVars = append(declaredVars, vars...)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return declaredVars, nil
}

// parseDeclaredVarsInFile parses a single .tf file for declared variables
func parseDeclaredVarsInFile(filePath string, verbose bool) ([]string, error) {
	content, err := os.ReadFile(filePath) // Use os.ReadFile instead of ioutil.ReadFile
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	// Regex to find declared variables in the format: variable "name" { ... }
	re := regexp.MustCompile(`variable\s+"([^"]+)"`)
	matches := re.FindAllStringSubmatch(string(content), -1)

	var declaredVars []string
	for _, match := range matches {
		declaredVars = append(declaredVars, match[1]) // match[1] is the variable name
	}

	if verbose {
		fmt.Printf("Variables declared in %s: %v\n", filePath, declaredVars)
	}

	return declaredVars, nil
}
