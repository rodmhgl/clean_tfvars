package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	// CLI Arguments
	tfDir := flag.String("tf", "", "Directory containing Terraform configuration (*.tf) files")
	tfvarsFile := flag.String("tfvars", "", "Path to the Terraform variable file (*.tfvars)")
	dryRun := flag.Bool("dry-run", false, "Run without making any changes to the tfvars file")
	verbose := flag.Bool("v", false, "Enable verbose mode")
	flag.Parse()

	// Validate input
	if *tfDir == "" || *tfvarsFile == "" {
		log.Fatal("You must specify both the Terraform directory and the tfvars file")
	}

	// Step 1: Parse declared variables from .tf files
	declaredVars, err := ParseDeclaredVars(*tfDir, *verbose)
	if err != nil {
		log.Fatalf("Error parsing declared variables from .tf files: %v", err)
	}
	if *verbose {
		fmt.Printf("Declared variables: %v\n", declaredVars)
	}

	// Step 2: Parse variables from the .tfvars file
	tfvars, err := ParseTFVarsFile(*tfvarsFile, *verbose)
	if err != nil {
		log.Fatalf("Error parsing variables from the .tfvars file: %v", err)
	}
	if *verbose {
		fmt.Printf("Variables in .tfvars file: %v\n", tfvars)
	}

	// Step 3: Find unused variables
	unusedVars := FindUnusedVars(declaredVars, tfvars, *verbose)
	if *verbose {
		fmt.Printf("Unused variables: %v\n", unusedVars)
	}

	// Step 4: If dry-run mode, just print the changes; otherwise, update the file
	if len(unusedVars) > 0 {
		if *dryRun {
			fmt.Println("[Dry Run] The following variables would be removed from the .tfvars file:")
			for _, varName := range unusedVars {
				fmt.Println(varName)
			}
		} else {
			err = UpdateTFVarsFile(*tfvarsFile, unusedVars, *verbose)
			if err != nil {
				log.Fatalf("Error updating .tfvars file: %v", err)
			}
			fmt.Println(".tfvars file updated.")
		}
	} else {
		fmt.Println("No unused variables found.")
	}
}
