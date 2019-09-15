# Running locally with SAM Local

### Install SAM Local
```npm install -g aws-sam-local```

### Build the Go Binary
```make build```

### Start up sam local
```sam local start-api```

### Call the endpoint
GET http://127.0.0.1:3000/folders?folder=folder1/
