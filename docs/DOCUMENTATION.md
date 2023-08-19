## API DOCUMENTATION

### List Entities
1. [Student](#student)
2. [Lecturer](#lecturer)
3. Class
4. Krs

### Student
1. ##### Register Student ```POST METHOD```
    endpoint: ```auth/student/register```    
    json body request : 
    ``` 
    {
        "name": "Devan",
        "email": "dvn@ub.ac.id",
        "password": "abcd12345",
        "major": "Teknik Informatika",
        "faculty": "FILKOM",
        "entrance": "mandiri"
    }
    ```
    json response : 
    ```
    {
        "status": "success",
        "code": 200,
        "message": "student registered to system",
        "data": {
            "nim": "235150107111001"
        }
    }
    ```

2. ##### Login Student ```POST METHOD```
    endpoint : ```auth/student/login```    
    json request body : 
    ```
    {
        "nim": "235150107111001",
        "password": "abcd12345"
    }
    ```
    json response : 
    ```
    {
        "status": "success",
        "code": 202,
        "message": "student successfully login"
    }
    ```

3. ##### Logout Student ```POST METHOD```
    endpoint : ```auth/student/logout```    
    json response : 
    ```
    {
        "status": "success",
        "code": 200,
        "message": "student successfully logout"
    }
    ```

### Lecturer
1. ##### Register Lecturer ```POST METHOD```
    endpoint : ```auth/lecturer/register```     
    json body request : 
    ```
    ```
    json response : 
    ```

    ```