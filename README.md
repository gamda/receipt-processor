# Receipt Processor Challenge

## Requirements

https://github.com/fetch-rewards/receipt-processor-challenge

## Process

I mainly work in Python and haven't used Go, but this was a good opportunity to learn the
basics of the language. Since go.dev has a sample webservice using gin, I decided to use
that as my starting point.

As suggested, I have an in-memory map to hold the receipts' points, there is no persistence.

I broke the project up in 2 packages: rules and server. I first built the rules. Since I 
was learning Go and its unit-testing framework, I started with the easiest rules first.
Then I found and followed the example mentioned using gin to build the server.

## Assumptions

The rules package handles all the input assuming it will be correctly formatted. If any
errors occur while parsing strings to numbers, the rule's function will return 0.

## Design Decisions

I thought about having a "point" parameter to the rules' functions, this way it would be
easier to change all points in one place in the future. I ultimately decided against it
because that is not a requirement, and it's better to have the points defined as close to
where they are used as possible.

I broke out the bool functions for the yes/no rules out to keep the logic simple. It seems
overkill, but the rules might change and it's better to have the yes/no logic encapsulated 
and not worry about it in the future.

I didn't build unit tests for the server because it is simple enough to test by hand. 
Building these tests would be my next step if this project was to be continued.

## How to run the server

From the `server` folder:

`> go run .`

## How to test the server

Once the server is online on port 8080 these are the POST commands with curl

```
curl http://localhost:8080/receipts/process \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"retailer": "Target", "purchaseDate": "2022-01-01", "purchaseTime": "13:01", "items": [{"shortDescription": "Mountain Dew 12PK", "price": "6.49"},{"shortDescription": "Emils Cheese Pizza", "price": "12.25"},{"shortDescription": "Knorr Creamy Chicken", "price": "1.26"},{"shortDescription": "Doritos Nacho Cheese", "price": "3.35"},{"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ", "price": "12.00"}], "total": "35.35"}'
```

and 

```
curl http://localhost:8080/receipts/process \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"retailer": "M&M Corner Market", "purchaseDate": "2022-03-20", "purchaseTime": "14:33", "items": [{"shortDescription": "Gatorade", "price": "2.25"},{"shortDescription": "Gatorade", "price": "2.25"},{"shortDescription": "Gatorade", "price": "2.25"},{"shortDescription": "Gatorade", "price": "2.25"}], "total": "9.00"}'
```

to create the sample Receipts. Then with the ID given by the server, 

```
curl http://localhost:8080/receipts/{received id}/points
```

## How to run unit tests

From the `rules` folder:

`> go test`
