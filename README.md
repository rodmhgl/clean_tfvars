# Terraform `.tfvars` Cleaner

This is a Go program designed to compare Terraform configuration files (`.tf`) with a Terraform variable file (`.tfvars`) and remove any variables in the `.tfvars` file that are **not declared** in the configuration. The program is designed to help keep your Terraform variable files clean and free of unused variables.

This is designed to be used in conjunction with `tflint`. `tflint` has a --fix feature which will removed unused variables from your configuration, but unfortunately leaves them in your `tfvars` file. This utility is designed to close that gap. Use `tflint` to remove the unused variables from your config and then this utility to remove the (now undeclared) variables from your `tfvars`.

## Features

- **Parse `.tf` files**: Extracts all declared variables in the `.tf` files.
- **Parse `.tfvars` file**: Reads the `.tfvars` file and identifies all defined variables.
- **Detect unused variables**: Compares the variables in `.tfvars` with the declared variables in `.tf` files.
- **Remove unused variables**: Removes any variables in `.tfvars` that are not declared in the `.tf` files.
- **Dry Run mode**: Preview the changes without making any actual modifications.
- **Verbose mode**: Get detailed logging of the parsing and comparison process.

## Installation

### Prerequisites

- Go 1.16 or later.

### Clone the repository

```bash
git clone https://github.com/yourusername/terraform-tfvars-cleaner.git
cd terraform-tfvars-cleaner
```

## Building the Program

To build the program, follow these steps:

1. Ensure you have **Go 1.16** or later installed on your machine.

2. Clone the repository:

    ```bash
    git clone https://github.com/rodmhgl/clean_tfvars
    cd clean_tfvars
    ```

3. Build the program using the Go compiler:

    ```bash
    go build
    ```

    This will generate an executable file named `terraform-cleanup`.

---

## Running the Program

Once the program is built, you can run it using the following command:

```bash
./terraform-cleanup -tf <path-to-tf-directory> -tfvars <path-to-tfvars-file>
```

## Command-Line Arguments

The program accepts the following command-line arguments:

### `-tf <path>`

- **Description**: Specifies the path to the directory containing the Terraform configuration files (`*.tf`). The directory will be scanned for `.tf` files that declare variables.
- **Example**:
    ```bash
    ./terraform-cleanup -tf ./terraform-config/
    ```

### `-tfvars <path>`

- **Description**: Specifies the path to the `.tfvars` file that contains the variables and their values. This file will be compared against the declared variables in the `.tf` files.
- **Example**:
    ```bash
    ./terraform-cleanup -tfvars ./terraform-config/variables.tfvars
    ```

### `-dry-run`

- **Description**: When this flag is set, the program will not modify the `.tfvars` file. Instead, it will print out the variables that would be removed, allowing you to preview the changes.
- **Example**:
    ```bash
    ./terraform-cleanup -tf ./terraform-config/ -tfvars ./terraform-config/variables.tfvars -dry-run
    ```

### `-v`

- **Description**: Enables verbose logging. Provides detailed output during the parsing and comparison process, including the declared and unused variables.
- **Example**:
    ```bash
    ./terraform-cleanup -tf ./terraform-config/ -tfvars ./terraform-config/variables.tfvars -v
    ```

---

### Example Usage

```bash

