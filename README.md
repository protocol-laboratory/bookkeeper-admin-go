# bookkeeper-admin-go
This project is a Go client library for the Apache Bookkeeper Admin API.<br/>
## Requirements
- Go 1.18+

## Usage
Import the client library:
```
go get -u github.com/protocol-laboratory/bookkeeper-admin-go
```
Create admin client first:
```go
// default connect to localhost:8080
bkAdmin := NewDefaultBookkeeperAdmin()
```
Get bookie info:
```go
bookieInfo, err := bkAdmin.BookieInfo()
```
