# recipt-processor-challenge

## Running the application locally
1. Clone the repository that contains the Go application to your local machine.
2. If the application has any dependencies, you need to install them by running the following command in the terminal:
    go mod download
This will download and install all the dependencies required to run the application.
3. Build the application by running the following command in the terminal:
    go build
This will compile the Go code and create an executable file.
4. Run the application by executing the executable file created in the previous step:
  ./"executable-file-name"
Replace "executable-file-name" with the name of the file that was created in the previous step.
5. Access the application: Once the application is running, you can access it by navigating to http://localhost:50052<port> in your web browser, where <port> is the port number on which the application is running.

## Running the application with Docker Compose
1. Clone the repository: Clone the repository that contains the Go application, Dockerfile, and docker-compose files to your local machine 
2. Make sure you have Docker and Docker Compose installed on your machine.
3. Clone the repository to your local machine.
4. Open a terminal and navigate to the root of the repository.
5. Run the following command to start the application:
          make up
6. Once the application is running, you can access the API at 'http://localhost:50052'
7. To stop the application, run the following command:
          make down

## API Documentation

### Process Receipt
    'POST /receipt/process'

### Request
The request body should be a JSON object with the following properties:

| Property | Type | Required | Description |
| -------- | ---- | -------- | ----------- |
| retailer | string | Yes | The name of the retailer. |
| purchaseDate | string | Yes | The date of the purchase in ISO 8601 (yyyy-mm-dd) |
| purchaseTime | string | Yes | The time of the purchase in 24-hour format |
| items | array | Yes | An array of items purchased. |
| total | string | Yes | The total amount of the purchase USD |

An Example rqeuest body is below:
``` json 
{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}
```

### Response
The response body will be a JSON object with the following properties:

| Property | Type | Description |
| -------- | ---- | ----------- |
| id | string | The id of the receipt that was processed. |

An Example response body is below:
``` json 
{
  "id": "02ce65ff-1964-4f9f-9404-1fb04abebd5e"
}
```

### Erorrs 

| Status Code | Description |
| ----------- | ----------- |
| 400 | The request body is invalid. |
| 415 | JSON body required |

An Example error response body is below:
``` json 
{
    "error": {
        "message": "Invalid request parameters. See invalidArgs",
        "status": 400,
        "error": "BAD_REQUEST"
    },
    "invalidArgs": [
        {
            "field": "Retailer",
            "value": "",
            "tag": "required",
            "param": ""
        }
    ]
}
```
``` json 
{
    "message": "/fetch/receipt/process only accepts Content-Type application/json",
    "status": 415,
    "error": "UNSUPPORTED_MEDIA_TYPE"
}
```

### Get Points
    'GET /receipt/:id/points'

This API endpoint returns the number of points earned for the receipt with the given id.

### Request
The 'id' parameter in the endpoint path is the unique identifier of the receipt.

### Response
The response body will be a JSON object with the following properties:

| Property | Type | Description |
| -------- | ---- | ----------- |
| points | number | The number of points earned for the receipt. |

An Example response body is below:
``` json 
{
  "points": 109
}
```

### Errors 

| Status Code | Description |
| ----------- | ----------- |
| 404 | The receipt with the given id was not found. |

An Example error response body is below:
``` json 
   {
        "message": "No receipt found for that id: 02ce65ff-1964-4f9f-904-1fb04abebd5e",
        "status": 404,
        "error": "NOT_FOUND"
    }
```

### Unit Tests
Unit tests have been included for the handler and service layers. You can run them by executing the following command in your terminal:

    go test ./handler
    go test ./services 