# Bubble-tea

To manipulate AWS Lambda, API Gateway, DynamoDB and IAM.

From https://docs.aws.amazon.com/fr_fr/lambda/latest/dg/golang-package.html to zip the code for AWS Lambda
go to /cmd
go build -tags lambda.norpc -o bootstrap main.go
~\Go\Bin\build-lambda-zip.exe -o myFunction.zip bootstrap

### Curls to test APIs

curl -X POST \
 -H "Content-Type: application/json" \
 -d '{"name":"MangoTea", "price": 6, "flavor": "Mango"}' \
 https://4d3tddb5sl.execute-api.eu-west-3.amazonaws.com/beta

curl -X POST \
 -H "Content-Type: application/json" \
 -d '{"name":"Bubble Tea", "price": 6, "flavor": "Caramel"}' \
 https://4d3tddb5sl.execute-api.eu-west-3.amazonaws.com/beta

- get boba
  curl https://4d3tddb5sl.execute-api.eu-west-3.amazonaws.com/beta/?name=MangoTea

curl -X PUT \
 -H "Content-Type: application/json" \
 -d '{"name":"Bubble Tea", "price": 5, "flavor": "Strawberry"}' \
 https://4d3tddb5sl.execute-api.eu-west-3.amazonaws.com/beta/

curl -X DELETE \
 https://4d3tddb5sl.execute-api.eu-west-3.amazonaws.com/beta\?name=MangoTea
