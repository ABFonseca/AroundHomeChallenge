# AroundHomeChallenge
Code Challenge for AroundHome interview

In this challenge we're intended to have an API receiving customer requests for Flooring experts and give a list of partners that have expertise
with the correct flooring type, ordered by rating and by proximity to the client.

Any partner that is further than it's announced operating radius or that doesn't work with that type of material is excluded from the list.

## Running API
`go run cmd/main.go`

## Running tests:
`go test ./...`

## Getting coverage
`go test -cover ./...`

## API call
The test partners were placed mostly within 50km of AroundHome office address, so a good staring test is to request from the AroundHome address 
(Lattitude: 52.50879681532554, Longitude: 13.375567271135349)

### call customer request from curl:
On a command line with the api running you can call curl with this example (or change the values of the json):

`curl -H "Content-Type: application/json" -d '{"Material": "wood", "Latitude": 52.50879681532554, "Longitude": 13.375567271135349, "SquareMeters": 20, "Phone": 111111111}' http://localhost:8080/request`

Alternatively you can put a request.json file with a valid request json (see example above) and call curl with:

`curl -H "Content-Type: application/json" -d @request.json http://localhost:8080/request`

### call partner detail from curl:
On a command line with the api running you can call curl with this example (or change the id of the partner)

`curl http://localhost:8080/partner/199`

# Challenge info
For the purpose of this challenge some decisions were taken that in a production code would be done diferently
* The partner list is in a json file
  * even though it's a static json file we modulate the behaviour like if it wasn't and fetch the full list each time
  * The test partners were placed mostly within 50km of AroundHome office address
* No authentication was made in the API
* I assumed the frontend builds the requests (and not the user) so I don't validate the request format or if it has missing fields


