# recipt-processor-challenge

## API Documentation

#### Process Receipt
    'POST /receipt/process'

#### Request
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
  "retailer": "Walmart",
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

#### Response
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

#### Erorrs 

| Status Code | Description |
| ----------- | ----------- |
| 400 | The request body is invalid. |
| 404 | The receipt with the given id was not found. |

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

#### Get Points
    'GET /receipt/:id/points'

This API endpoint returns the number of points earned for the receipt with the given id.

#### Request
The 'id' parameter in the endpoint path is the unique identifier of the receipt.

#### Response
The response body will be a JSON object with the following properties:

| Property | Type | Description |
| -------- | ---- | ----------- |
| points | number | The number of points earned for the receipt. |

An Example response body is below:
``` json 
{
  "points": 102
}
```

#### Errors 

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

#### Unit Tests
Unit tests have been included for the handler and service layers. You can run them by executing the following command in your terminal:

    go test ./handler
    go test ./services 