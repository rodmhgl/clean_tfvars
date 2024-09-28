package main

import "fmt"

// FindUnusedVars compares declared variables with variables in the tfvars file
func FindUnusedVars(declaredVars []string, tfvars map[string]string, verbose bool) []string {
	var unusedVars []string

	for varName := range tfvars {
		if !contains(declaredVars, varName) {
			unusedVars = append(unusedVars, varName)
			if verbose {
				fmt.Printf("Variable %s is in the tfvars file but not declared in the tf files\n", varName)
			}
		}
	}

	return unusedVars
}

// contains checks if a string is present in a slice
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
