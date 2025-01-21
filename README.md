# gocalculator
The goal of this project is to create a minimal http+json API for a calculator service.

## Running the project
To run this webapp server:
1. Open up your CLI in your code editor
2. Make sure to cd to the root directory of this project
3. Run `go run cmd/gocalculator/main.go` on the terminal

## Project goal and inspiration
On [this](https://github.com/dreamsofcode-io/goprojects/tree/47ad97d536752ff5da6703d8cac78c0edf247d89/02-backend-api) github repo you can find the explanation of what the idea here is.
But basically the goal is to learn some basics on API developement in GO by building a backend web API calculator. This calculator is stateless, meaning that there is no data stored in a database or in memory.

This calculator supports basic operations:
- Addition
- Subtraction
- Multiplication
- Division

The OpenAPI specification found in the [api](/api/api-spec.yaml) directory of this project describes the requirements for the API.

## Credits
Dreams of Code [Youtube video](https://youtu.be/gXmznGEW9vo?t=157)

Requirements and specs [Github repo](https://github.com/dreamsofcode-io/goprojects/tree/47ad97d536752ff5da6703d8cac78c0edf247d89/02-backend-api)