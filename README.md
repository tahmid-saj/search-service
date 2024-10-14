# search-service

Typeahead search service using a Trie data structure approach. Developed using Go / Gin and AWS DynamoDB.

<br/>

Sample trie table in AWS DynamoDB:

<img width="686" alt="image" src="https://github.com/user-attachments/assets/2478c40a-3494-4efd-854e-91ca6aaf12b7">

<br/>
<br/>

API output when searching:

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
