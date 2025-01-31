//Ben Nguyen Notes
This is a repo for Fetch Reward's receipt-processor-challenge (https://github.com/fetch-rewards/receipt-processor-challenge). 

I decided to code with Go because it is the preffered language of the company and I also have not used Go yet so this is a good learning experience for me.

I followed the documentation (https://go.dev/doc/tutorial/web-service-gin) to get a baseline of how to create the api endpoints with Go and gin.

I used POSTMAN to test throughout my process of creating the API

I did not containerize the app

//INSTRUCTIONS

To tess: clone this repo and cd into the fetchReceiptProcessor folder. Make sure Go is installed on your system.
Then type the command $ go get .
Then type the command $ go run . 
The server should now be running on your local device on port 8080

Now you can use curl command or POSTMAN to test the api's.


To test the POST api:

The address to test processing a receipt is http://localhost:8080/receipts/process
Make sure you are doing a POST
Then have the body be in raw json format:
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}

After the POST request is made, you will receive an id like this: {
    "id": "85a55dc9-b36d-4225-b41c-a396e39da4d9"
}

Now you use the id to test the GET api:

The address to test is http://localhost:8080/receipts/{id}/points

In the above case the URL will be http://localhost:8080/receipts/85a55dc9-b36d-4225-b41c-a396e39da4d9/points

Make sure you are making a GET request.

In the above example JSON, the GET request will send back: { "points": 28 }
