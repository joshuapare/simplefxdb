# simplefxdb: A Simple Document Database Engine using FX

Welcome to **SimpleFXDB**, a minimalistic filesystem-based document database engine written in Go. This project primarily demonstrates the power and utility of [**FX**](https://github.com/uber-go/fx) in Go for dependency injection.

## Why SimpleFXDB?

While Go is renowned for its simplicity and straight-to-the-point nature, dependency injection frameworks such as [FX](https://github.com/uber-go/fx) can help developers better structure their applications, manage dependencies, and avoid boilerplate. DI frameworks are a game changer in the world of large applications with tons of components, services, and dependencies.

## Features

- Simple query syntax to interact with the database.
- Uses the filesystem structure and JSON to store `collections` with `documents`.
- Dependency injection using [FX](https://github.com/uber-go/fx)

**Note**: This project is for educational purposes and does not include advanced features found in real database engines, such as paging, transactions, and indexes.

## Query Syntax

Here's a glimpse into the simple query syntax you'll be using:

```bash
# Create a new collection called "users"
CREATE users
# List all the collections
LIST collections
# Find all documents in the "users" collection
GET users
# Find all documents in the "users" collection with a "name" property equal to "John"
GET users WHERE name = "John"
# Delete all documents in the "users" collection with a "name" property equal to "John"
DELETE users WHERE name = "John"
```

## Project Components

This project can be broken into the following main components:

1. **Query Parsing Engine**: Interprets and processes our query syntax.
2. **Collection Engine**: Manages creation and access to different collections.
3. **Storage Engine**: Handles the reading and writing of documents to the disk.
4. **Logger**: An essential tool for logging events, errors, and information.
5. **Webserver**: A simple server exposing the query endpoint.

## Getting Started

### Setting Up the Project

1. Clone the repository: `git clone https://github.com/joshuapare/simplefxdb.git`

2. Install the dependencies: `go mod download`

3. Run the project: `go run .`
This will initialize the db engine, storage, and expose an HTTP endpoint for interacting with the database at `localhost:4422`.

### Using the Query Endpoint

The project exposes a simple HTTP endpoint for interacting with the database. You can use a tool like [Postman](https://www.postman.com/) to send requests to the endpoint or `curl` from the command line.

Here are some examples of how you can interact with the database:

```bash
# Create a new collection called "users"
curl -X POST http://localhost:4422/query -d "CREATE users"
```

## Contributions

While this project is primarily educational, feedback, issues, or pull requests are always welcome. Feel free to open an issue or submit a pull request for suggestions or improvements.

## License

This project is open-source and available under the MIT License.
