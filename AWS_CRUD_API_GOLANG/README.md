# Create API CRUD

## Overview

This project aims to implement CRUD (Create, Read, Update, Delete) operations using AWS Lambda, API Gateway, DynamoDB, and IAM to manage bubble tea data.

To package the code for AWS Lambda, follow the steps below:

1. Navigate to the `/cmd` directory.
2. Build the code with the following command:
   ```bash
   go build -tags lambda.norpc -o bootstrap main.go
   ```
3. Create a ZIP file containing the Lambda function by running: (from https://docs.aws.amazon.com/fr_fr/lambda/latest/dg/golang-package.html)
   ```bash
   ~/Go/Bin/build-lambda-zip.exe -o myFunction.zip bootstrap
   ```

### Curls to test APIs

Use the following cURL commands to test the APIs:

1. **Create a new boba (Mango Tea):**

   ```bash
   curl -X POST \
    -H "Content-Type: application/json" \
    -d '{"name":"MangoTea", "price": 6, "flavor": "Mango"}' \
    https://4d3tddb5sl.execute-api.eu-west-3.amazonaws.com/beta
   ```

2. **Create another boba (Bubble Tea):**

   ```bash
   curl -X POST \
    -H "Content-Type: application/json" \
    -d '{"name":"Bubble Tea", "price": 6, "flavor": "Caramel"}' \
    https://4d3tddb5sl.execute-api.eu-west-3.amazonaws.com/beta
   ```

3. **Get boba (Mango Tea):**

   ```bash
   curl https://4d3tddb5sl.execute-api.eu-west-3.amazonaws.com/beta/?name=MangoTea
   ```

4. **Update boba (Bubble Tea):**

   ```bash
   curl -X PUT \
    -H "Content-Type: application/json" \
    -d '{"name":"Bubble Tea", "price": 5, "flavor": "Strawberry"}' \
    https://4d3tddb5sl.execute-api.eu-west-3.amazonaws.com/beta/
   ```

5. **Delete boba (Mango Tea):**
   ```bash
   curl -X DELETE \
    https://4d3tddb5sl.execute-api.eu-west-3.amazonaws.com/beta\?name=MangoTea
   ```
