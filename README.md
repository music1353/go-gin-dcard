#  2020 Dcard Web Backend Intern Assignment

Design a middleware prevent from server overloading and which is used in making Dcard's friends a pair at midnight.



There are some demands in the assignment:
1. Limit the number of requests from the same IP per hour less than 1000
2. Add the remaining of the number of requested (X-RateLimit-Remaining) and the time of reset RateLimit in the response header.
3. If over the limit, then response the status code 429 (Too Many Requests)
4. Can implement with any Database 



## Environment

- `Programming Language`：Golang 1.13.4
- `Backend Framework`：Golang Gin
- `Database`：Redis



## Get Started

1. **Clone the project**

   ~~~git
   git clone https://github.com/music1353/go-gin-dcard.git
   ~~~

2. **Run the Gin Server & Redis** 

   * Run Gin Server

     ~~~
     cd go-gin-dcard
     go run main.go
     ~~~

   * Run Redis

     ~~~
     redis-server
     ~~~

3. **Request to the Pair API**

   Request to the following API.

   ~~~
   http://localhost:8080/api/v1/limit/pair
   ~~~

   If you start the server in localhost, the IP will be represented in `::1`, and it will be a key store in the Redis.

   

   Following is the example of the response：

   **content-length →**45

   **content-type →**application/json; charset=utf-8

   **date →**Sat, 07 Mar 2020 11:54:18 GMT

   **x-ratelimit-remaining →**985

   **x-ratelimit-reset →**37m9s

   

   If you over the limit, remaining will represent with 0

   **content-length →**63

   **content-type →**application/json; charset=utf-8

   **date →**Sat, 07 Mar 2020 12:01:04 GMT

   **x-ratelimit-remaining →**0

   **x-ratelimit-reset →**30m23s

   

## API Doc

Please refer to the document in `doc` folder.