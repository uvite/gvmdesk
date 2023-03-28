#  BigQuery SQL Driver for Golang
This is an implementation of the BigQuery Client as a `database/sql/driver` for easy integration and usage.

## Goals of project

This project is meant to be a basic `database/sql` driver implementation for Golang so that developers can easily use 
`*sql.DB` functions with Google's BigQuery database.

__Unlike the original project, this driver__
- does not contain ORM extension
- uses different connection string

## Connection string

The connection string format is URL with `bigquery` scheme and host representing project ID.
Path (optional) represents location (cloud region).  

* `bigquery://projectid/?param1=value&param2=value`
* `bigquery://projectid/location?param1=value&param2=value`

### Common parameters

* `dataset` - dataset ID. When set, it allows to use unqualified tables names in queries.

### Authentication parameters

As this is using the Google Cloud Go SDK, you will need to have your credentials available
via the GOOGLE_APPLICATION_CREDENTIALS environment variable point to your credential JSON file.

Alternatively, you can use one of the following parameters:
* `apiKey` - API key value
* `credentials` - base-64 encoded service account or refresh token JSON credentials  

Examples:  
* `bigquery://projectid/?apiKey=AIzaSyB6XK8IO5AzKZXoioQOVNTFYzbDBjY5hy4`
* `bigquery://projectid/?credentials=eyJ0eXBlIjoiYXV0...`

## Usage

```go
package main

import "github.com/bonitoo-io/go-sql-bigquery"

func main() {
    db, err := sql.Open("bigquery", "bigquery://lunar-1234/?dataset=storeys")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close() 
    ...
}
```

## Contribution

Contributions are welcome.  

## Current Support

* [x] `driver.Conn` implemented
* [x] `driver.Querier` implemented
* [x] `driver.Pinger` implemented
* [x] `driver.DriverContext` implemented
* [x] `driver.QueryerContext` implemented
* [x] `driver.ExecerContext` implemented
* [x] `driver.RowsColumnTypeDatabaseTypeName` implemented
* [x] Prepared Statements - supported via a quick hack
* [ ] Parameterized Queries
