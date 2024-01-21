
# Calorie Calculator

A simple calorie calculator developed in Golang programming language, using HTMX library for dynamic web page updating and Air.toml framework for automatic server reloading during development.

## Installation

1.  Install Golang: https://golang.org/dl/

2.  Install Air.toml:

- `go get -u github.com/cosmtrek/air`.
3.  Install Templ:
- `go get -u github.com/a-h/templ`
4. Install the project dependencies:
   `go mod tidy`.


## Startup

1.  Run Air.toml to automatically reload the server when changes are made to the code

1.  `air`

2. open a web browser and go to http://localhost:8080.


## Usage

1.  Enter the required data into the calorie calculation form.
2.  Select your gender and activity level.
3. Click the Calculate button.