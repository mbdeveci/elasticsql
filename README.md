This repo contains cli wrapper for [github.com/cch123/elasticsql](https://github.com/cch123/elasticsql)

### Usage

`elasticsql "SELECT * FROM my_table m WHERE m.id > 3"`
```javascript

{"query" : {"bool" : {"must" : [{"range" : {"m.id" : {"gt" : "3"}}}]}},"from" : 0,"size" : 1}
```

`elasticsql -p "SELECT * FROM my_table m WHERE m.id > 3"`

```javascript
{
  "query": {
    "bool": {
      "must": [
        {
          "range": {
            "m.id": {
              "gt": "3"
            }
          }
        }
      ]
    }
  },
  "from": 0,
  "size": 1
}
```
