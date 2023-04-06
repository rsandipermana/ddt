# Test for Recruitment Process

This is a repository for completing the tests given in the recruitment process. There are two tests that need to be completed, which are:

1. Finding the largest number in an array
2. Counting the occurrences of a string in another string

## Requirements

To run this project, you will need:

- Go 1.16 or higher

## Usage

To use this project, follow these steps:

1. Clone this repository.
2. Go to the root directory of this repository.
3. Run `go run main.go` to start the API server.
4. Send a request to the API server using your preferred client (e.g. `curl`).

## API Endpoints

There are two API endpoints available:

### GET /max

This endpoint is used to find the largest number in an array.

#### Request
```json
{
    "numbers": [12, 41, 24, 2, 1, 4]
}
```

#### Response
```json
{
    "max": 41
}
```
#### CURL
```javascript
curl --location --request GET 'http://localhost:8080/max' \
--header 'Content-Type: application/json' \
--data-raw '{
    "numbers": [12, 41, 24, 2, 1, 4]
}'
```

### POST /count

This endpoint is used to count the occurrences of a string in another string.

#### Request
```json
{
    "target": "makan",
    "text": "selama berpuasa, Pdak dibolehkan makan dan minum. Puasa diawali dengan makan sahur dan diakhir dengan berbuka. Disunahkan untuk berbuka dengan makanan yang manis, misal dengan makan kurma."
}
```

#### Response
```json
{
    "count": 4
}
```

### CURL
```javascript
curl --location --request POST 'http://localhost:8080/count' \
--header 'Content-Type: application/json' \
--data-raw '{
    "target": "makan",
    "text": "selama berpuasa, Pdak dibolehkan makan dan minum. Puasa diawali dengan makan sahur dan diakhir dengan berbuka. Disunahkan untuk berbuka dengan makanan yang manis, misal dengan makan kurma."
}'
```

## Architecture

This project is implemented using the following architecture:
- Interface
- Clean Architecture

## License
This project is licensed under the MIT License - see the LICENSE file for details.
