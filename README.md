# GO REST API

Production ready golang REST API.

## Architecture

```
    +-------------+      +-------------+      +-------------+
--->|    HTTP     | ---> |   Service   | ---> | Repository  | ---> DB(postgres)
    +-------------+      +-------------+      +-------------+
```

1. **HTTP (Hypertext Transfer Protocol)**

   - This component represents the presentation layer or the user interface. It receives requests from clients and routes them to the appropriate service for processing.

2. **Service**

   - The service is responsible for business logic and coordination between different parts of the system. It receives requests from the HTTP component, performs necessary operations, and may interact with the data repository to retrieve or modify information.

3. **Repository**
   - The repository represents the data access layer. It is responsible for interacting with persistent storage, such as databases or file systems. The service uses the repository to perform CRUD (Create, Read, Update, Delete) operations on the data.
