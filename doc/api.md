# Restful API

- `v1URL` = http://localhost:8080/api/v1

- **Response Format**

  ```json
  resp = {
    "msg": "",
    "result": "" // response data
  }
  ```

- **Status Code**

  | Status Code | Description         |
  | ----------- | ------------------- |
  | 200         | OK                  |
  | 429         | Too Many Requests   |
  | 500         | Service Unavailable |

* **pairAPI**

  | API Method | API URL          | Desc     | Req Params | Resp Result |
  | ---------- | ---------------- | -------- | ---------- | ----------- |
  | GET        | v1URL/limit/pair | 卡友配對 |            |             |

