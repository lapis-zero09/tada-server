# payments
# POST
curl -i -X POST -H "Content-Type: application/json" -d "{ \"placeid\": 123, \"cost\": 23253533 }" http://localhost:8080/api/v1/payments

# PUT
curl -i -X PUT -H "Content-Type: application/json" -d "{ \"placeid\": 365335, \"cost\": 9990 }" http://localhost:8080/api/v1/payments/1

# DELETE
curl -i -X DELETE http://localhost:8080/api/v1/payments/1


# paymentTag
# POST
curl -i -X POST -H "Content-Type: application/json" -d "{ \"paymentId\": 5, \"tagId\": 12 }" http://localhost:8080/api/v1/payment_tags

# PUT
curl -i -X PUT -H "Content-Type: application/json" -d "{ \"paymentId\": 3, \"tagId\": 10 }" http://localhost:8080/api/v1/payment_tags/1

# DELETE
curl -i -X DELETE http://localhost:8080/api/v1/payment_tags/1

