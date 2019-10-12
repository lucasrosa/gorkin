# Gorkin
[![Build Status](https://travis-ci.org/lucasrosa/gorkin.svg?branch=master)](https://travis-ci.org/lucasrosa/gorkin) 
[![Go Report Card](https://goreportcard.com/badge/github.com/lucasrosa/gorkin)](https://goreportcard.com/report/github.com/lucasrosa/sgorkin) 
[![codecov](https://codecov.io/gh/lucasrosa/gorkin/branch/master/graph/badge.svg)](https://codecov.io/gh/lucasrosa/gorkin)


## Running locally with SAM Local

### Install SAM Local
```npm install -g aws-sam-local```

### Build the Go Binary
```make build```

### Start up sam local
```sam local start-api```

### Call the endpoint
GET http://127.0.0.1:3000/folders?folder=folder1/


### Generating pre-signed URL for S3 in Go
https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-presigned-urls.html
https://github.com/awsdocs/aws-doc-sdk-examples/tree/master/go/example_code/s3


### Example of local request
GET http://127.0.0.1:3000/folders

### Example of local response
```
{
    "id": "root",
    "name": "root",
    "type": "folder",
    "children": {
        "folder1": {
            "id": "d41d8cd98f00b204e9800998ecf8427e",
            "name": "folder1",
            "type": "folder",
            "children": {
                "folder1_1": {
                    "id": "d41d8cd98f00b204e9800998ecf8427e",
                    "name": "folder1_1",
                    "type": "folder",
                    "children": {
                        "folder1_1_1": {
                            "id": "d41d8cd98f00b204e9800998ecf8427e",
                            "name": "folder1_1_1",
                            "type": "folder",
                            "children": {
                                "12362714-dzone-refcard215-microservices.pdf": {
                                    "id": "73a82eb7b09a2c5499cc76c2aefe52e2",
                                    "name": "12362714-dzone-refcard215-microservices.pdf",
                                    "type": "file",
                                    "children": null
                                },
                                "halo.txt": {
                                    "id": "9b1529ddfd06b2046b2615f58ad2829f",
                                    "name": "halo.txt",
                                    "type": "file",
                                    "children": null
                                }
                            }
                        }
                    }
                }
            }
        },
        "folder2": {
            "id": "d41d8cd98f00b204e9800998ecf8427e",
            "name": "folder2",
            "type": "folder",
            "children": {
                "BP-Diet-Roadmap-2019.pdf": {
                    "id": "5053efd24093110335a9a3c3c6dd17f8",
                    "name": "BP-Diet-Roadmap-2019.pdf",
                    "type": "file",
                    "children": null
                },
                "folder2_1": {
                    "id": "d41d8cd98f00b204e9800998ecf8427e",
                    "name": "folder2_1",
                    "type": "folder",
                    "children": {}
                }
            }
        },
        "foto.png": {
            "id": "1ca5cf539336c45b48ae2369a56b40bf",
            "name": "foto.png",
            "type": "file",
            "children": null
        }
    }
}
```