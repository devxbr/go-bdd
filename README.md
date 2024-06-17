# Behavior-Driven Development

![BDD Operation](https://cms-cdn.katalon.com/BDD_Operation_39de9dd9ff.png)

## Overview

go-bdd is a simple Behavior-Driven Development (BDD) framework for Go, leveraging Gherkin for defining test cases. 
This repository includes a basic calculator example to demonstrate the functionality.

## Features

- Define BDD scenarios in Gherkin
- Execute tests in Go

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

Clone the repository:

```sh
git clone https://github.com/devxbr/go-bdd.git
cd go-bdd
```

Install dependencies:

```sh
go mod tidy
```

### Usage

Run tests:

```sh
go test
```

## Project Structure

- `calculator.go` - Calculator implementation
- `calculator_test.go` - BDD tests for the calculator
- `features/` - Directory for Gherkin feature files

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a pull request

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For questions, please open an issue in this repository.
