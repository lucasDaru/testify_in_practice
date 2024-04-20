# Testify in Practice

This is a sample project demonstrating the usage of the Testify framework for writing tests in Go.

## Overview

This project showcases how to use the Testify framework to write tests for a simple application that simulates a repository for managing persons.

## Structure

The project consists of the following directories:

- `model`: Contains the definition of the `Person` struct representing a person in the database.
- `repository`: Contains the implementation of repository operations for managing persons.
- `repositorytest`: Contains fixture data and tests for the repository package.
- `go.mod`: Go module file for managing project dependencies.

## Getting Started

To get started with this project, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/seuusername/testify_in_practice.git
   cd testify_in_practice
   ```
2. Install dependencies:The only dependency for this project is the Testify framework. You can install it using the following command:

    ```bash
    go get github.com/stretchr/testify/assert
    ```
3. Run tests:Execute the tests to ensure everything is working correctly:

    ```bash
    go test -v ./...
   ```
    ````
    ?       github.com/lucasDaru/testify_in_practice/model  [no test files]
    === RUN   TestFindPersonByID
    === RUN   TestFindPersonByID/success
    === RUN   TestFindPersonByID/invalid-person-id
    === RUN   TestFindPersonByID/person-id-not-found
    --- PASS: TestFindPersonByID (0.00s)
    --- PASS: TestFindPersonByID/success (0.00s)
    --- PASS: TestFindPersonByID/invalid-person-id (0.00s)
    --- PASS: TestFindPersonByID/person-id-not-found (0.00s)
    PASS
    ok      github.com/lucasDaru/testify_in_practice/repository     0.624s
    ?       github.com/lucasDaru/testify_in_practice/repositorytest [no test files]
    ````
