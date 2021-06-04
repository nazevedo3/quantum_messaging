![alt text](https://images.newscientist.com/wp-content/uploads/2020/08/26152459/26-aug_cosmic-rays-quantum-computers.jpg)

# Welcome to Quantum Messaging!

Quantum Messaging is the first public application that runs on pure quantum computers....just kidding.  It's just a simple REST API that does two things.  The first being it will take a message in the request body and return the sha256 hash.  You can then do a GET request, provide the sha256 hash and it will retrieve the associated message.  Please see below for the API end points along with some examples.

## REST API ENDPOINTS
| Endpoint        | Method           
| ------------- |:-------------:| 
| /api/message      | POST |  |
| /api/hash/:hash      | GET      |   



## Examples
You can use your favorite tool.  Examples below are using CURL but also works with POSTMAN, etc.
#### Example #1 - Sending a message to the `api/message` endpoint
```sh
curl -d 'this is a test' -X POST  <server-IP>:4000/api/message
{"hash":"2e99758548972a8e8822ad47fa1017ff72f06f3ff6a016851f45c398732bc50c"}
```

### Example #2 - Retreiving a message from the assoiciated hash
```sh
curl <server-IP>:4000/api/hash/2e99758548972a8e8822ad47fa1017ff72f06f3ff6a016851f45c398732bc50c
{"message":"this is a test"}
```




