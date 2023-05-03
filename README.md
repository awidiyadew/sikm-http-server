# Golang HTTP Server Exercise

## Feat 1: Get Product List
Please make an API to get list of product data from file `data/product.txt`.

```text
Method: GET
API Path: /product/list
Resp Content-Type: application/json
```

Sample Success Response (Status Code 200):
```json
[
    {
        "id": "0001",
        "name": "Kit Kat",
        "price": 3500,
        "qty": 10
    },
    {
        "id": "0002",
        "name": "Oreo",
        "price": 2000,
        "qty": 12
    }
    // other product....
]
```

Notes:
* When there's no product in `data/product.txt` the response should be with status code `200` and with content body `[]` or empty json array.
* When the method is not `GET`, api should response with status code `405 (Method Not Allowed)` with proper error message.

## Feat 2: Add Product
Please make an API to add a product data to file `data/product.txt`.

```text
Method: POST
API Path: /product/add
Request Content-Type: application/json
```

Example of valid request body:

```json
{
    "id": "0006",
    "name": "Kopi ABC",
    "price": 1000,
    "qty": 20
}
```

Example Response:
```
Status code: 200 (OK) or 201 (Created)

{
    "id": "0006",
    "name": "Kopi ABC",
    "price": 1000,
    "qty": 20
}
```

Notes:
* When the method is not `POST`, api should response with status code `405 (Method Not Allowed)` with proper error message.
* Product qty must be above 0, when the qty is less than 0 API should response with status code `400 (Bad Request)` with response body json `{ "message": "qty must be more than 0" }`.


## Directions
This is just an exercise, not a mandatory task for the mentees.
* Please fork this repository to your github account.
* Clone the forked repository from your github account.
* Make a new branch with name `be-02/<your_name>`.
* Push your code to your repository, when it's done please raise a Pull Request to [awidiyadew/sikm-http-server](https://github.com/awidiyadew/sikm-http-server).
* Mentee who can finish this exercise will get an extra "nilai ke-aktifan".
