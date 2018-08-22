# Get content from URL flogo activity
This activity allows your flogo application to get content from a URL with support for basicAuth


## Installation

```bash
flogo install github.com/abasse/flogogetcontent
```

## Schema
Inputs and Outputs:

```json
  { "inputs":[
        {
          "name": "basicAuthUser",
          "type": "string",
          "required": false
        },
        {
          "name": "basicAuthPassword",
          "type": "string",
          "required": false
        },
        {
          "name": "URL",
          "type": "string",
          "required": true
        }
      ],
      "outputs": [
        {
            "name": "result",
            "type": "string"
        },
        {
            "name": "status",
            "type": "string"
        },
        {
            "name": "header",
            "type": "string"
        }
      ]
  }
```
