# redeam
Simple golang book management API

# Structure
This project is split up into three distinct portions:
1.  The API itself:  Written in Go and dockerized, the API manages a list of books.
The following endpoints are currently available:
  - AddBook adds a single book to the data store.
  - AdjustRating adjusts the rating of a book.  Ratings can be raised or dropped.
  - GetBook retrieves details about a single book.
  - GetAllBooks retrieves the details about all of the books in the data store.
  - DeleteBook permanently removes a book from the data store.
  - ChangeStatus allows a user to toggle the status of a book as check-in or checked-out.
  - Collapse is just for fun.  If a user, for example, was fined, and attempts to set the book store on fire, they can hit this endpoint.
    A set of random numbers generated will dictate whether his attempt was successful, or it was thwarted by amazing employees.
    
Currently, the API is in a running state but has not been exhaustively tested.

2.  The data store.  The data store of choice was mysql due to it's simplicity and COST.  It is available bia a sidecar Docker container along side the Go API above.

3.  The client:  The client is written in Go.  It builds HTTP requests for each of the API calls and is currently being automatically run as
a go-routine within the API's container.  An environment variable toggles it on or off.  Currently, there are no use-cases within the client
that exercise the situation of "user error".

# TODOs
1.  Add "user error" cases to the client
2.  Flesh out unit tests
3.  Finalize end-to-end functional testing

# Potential Future Improvements
1.  Depracate current API protocol over HTTP in favor of gRPC endpoints using a gRPC Gateway for external users.
2.  Add an authenticaion mechanism for sensitive endpoints
3.  Add cloud deployment services such as Jenkins or Google Cloud build service, helm, and terraform deployments
