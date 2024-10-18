# search-service

Typeahead search service using a Trie data structure approach. Developed using Go / Gin and AWS DynamoDB.

<br/>
<br/>

## Overview

### Design

The search operation of the service uses the trie nodes (stored in DynamoDB) to efficiently find the top frequent searches for a prefix. Similar services can be found <a href="https://whimsical.com/web-microservices-6uqvwWZtcBFsNJB2hepGy1">here</a> and below:

#### Similar services

<img width="834" alt="image" src="https://github.com/user-attachments/assets/b54088e7-870c-46dd-9cf6-2e5ec27d9d5c">

### Examples

#### Sample trie table in AWS DynamoDB

<img width="686" alt="image" src="https://github.com/user-attachments/assets/2478c40a-3494-4efd-854e-91ca6aaf12b7">

#### API input and output when searching

```
// Input
{
    "searchQuery": "alg",
    "tableName": "search-service-trie"
}
```


```
// Output
{

    "ok": true,
    "response": [
        {
            "Result": "algorithms",
            "ResultFrequency": 2
        },
        {
            "Result": "alg",
            "ResultFrequency": 3
        },
        {
            "Result": "algorit",
            "ResultFrequency": 1
        },
        {
            "Result": "algo",
            "ResultFrequency": 2
        }
    ]
}
```
