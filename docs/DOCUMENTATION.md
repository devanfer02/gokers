## API DOCUMENTATION

### List Entities
1. [Student](#student)
2. [Lecturer](#lecturer)
3. [Course](#course)
4. Krs

### Student
1. ##### Register Student 
    endpoint : ```/auth/student/register```     
    method : ```POST METHOD```        
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

2. ##### Login Student 
    endpoint : ```/auth/student/login```    
    method : ```POST METHOD```     
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

3. ##### Logout Student 
    endpoint : ```/auth/student/logout```     
    method : ```POST METHOD```    
    json response : 
    ```
    {
        "status": "success",
        "code": 200,
        "message": "student successfully logout"
    }
    ```

### Lecturer
1. ##### Register Lecturer
    endpoint : ```/auth/lecturer/register```      
    method :  ```POST METHOD```      
    json body request : 
    ```
    {
        "name": "Devan",
        "email": "dvn@ub.ac.id",
        "password": "abcd12345",
        "major": "Teknik Informatika",
        "faculty": "FILKOM"
    }
    ```
    json response : 
    ```
    {
        "status": "success",
        "code": 200,
        "message": "student registered to system",
        "data": {
            "ndn": "002351501001"
        }
    }
    ```

### Course
1. ##### Register Course 
    endpoint : ```/course```       
    method : ```POST METHOD```       
    json body request for major :
    ```
    {
        "name": "Algoritma dan Struktur Data",
        "code": "CIE61010",
        "sks": 3,
        "type": "major",
        "faculty": "FILKOM",
        "major": "Pendidikan Teknologi Informasi"
    }
    ```
    json response for major : 
    ```
    {
        "status": "success",
        "code": 200,
        "message": "course registered to system",
        "data": {
            "record_data": {
                "id": "f8c26e8e-6f2b-44dc-8201-93da7f61e4b5",
                "created_at": "2023-08-21T18:51:23.23+07:00",
                "updated_at": "2023-08-21T18:51:23.23+07:00",
                "name": "Algoritma dan Struktur Data",
                "code": "CIE61010",
                "sks": 3,
                "type": "major",
                "faculty": "FILKOM",
                "major": "Pendidikan Teknologi Informasi"
            }
        }
    }
    ```
    json request body for faculty : 
    ```
    {
        "name": "Pemrograman Dasar",
        "code": "COM60014",
        "sks": 4,
        "type": "faculty",
        "faculty": "FILKOM"
    }
    ```
    json response for faculty : 
    ```
    {
        "status": "success",
        "code": 200,
        "message": "course registered to system",
        "data": {
            "record_data": {
                "id": "ddb386f0-a132-4e12-8d72-b6f99a986318",
                "created_at": "2023-08-21T18:53:07.341+07:00",
                "updated_at": "2023-08-21T18:53:07.341+07:00",
                "name": "Pemrograman Dasar",
                "code": "COM60014",
                "sks": 4,
                "type": "faculty",
                "faculty": "FILKOM",
                "major": ""
            }
        }
    }
    ```
    json request body for university : 
    ```
    {
        "name": "Bahasa Inggris",
        "code": "UBU60004",
        "sks": 2,
        "type": "university"
    }
    ```
    json response for university : 
    ```
    {
        "status": "success",
        "code": 200,
        "message": "course registered to system",
        "data": {
            "record_data": {
                "id": "e8530331-e6a1-4baa-8d80-b01d68f003e0",
                "created_at": "2023-08-21T18:54:32.721+07:00",
                "updated_at": "2023-08-21T18:54:32.721+07:00",
                "name": "Bahasa Inggris",
                "code": "UBU60004",
                "sks": 2,
                "type": "university",
                "faculty": "",
                "major": ""
            }
        }
    }
    ```
    json response conflict : 
    ```
    {
        "status": "error",
        "code": 409,
        "message": "course code conflicted",
        "data": "course code already exist"
    }
    ```
    json response type doesnt exist : 
    ```
    {
        "status": "error",
        "code": 400,
        "message": "types or faculty or major doesnt exist",
        "data": "types must at least be university or faculty or major, requested types: region"
    }
    ```

2. ##### Get All Courses 
    endpoint : ```/course```    
    method : ```GET METHOD```      
    json response : 
    ```
    {
        "status": "success",
        "code": 200,
        "message": "sucessfully fetch course",
        "data": [
            {
                "id": "ddb386f0-a132-4e12-8d72-b6f99a986318",
                "created_at": "2023-08-21T11:53:07.341Z",
                "updated_at": "2023-08-21T11:53:07.341Z",
                "name": "Pemrograman Dasar",
                "code": "COM60014",
                "sks": 4,
                "type": "faculty",
                "faculty": "FILKOM",
                "major": ""
            },
            {
                "id": "e8530331-e6a1-4baa-8d80-b01d68f003e0",
                "created_at": "2023-08-21T11:54:32.721Z",
                "updated_at": "2023-08-21T11:54:32.721Z",
                "name": "Bahasa Inggris",
                "code": "UBU60004",
                "sks": 2,
                "type": "university",
                "faculty": "",
                "major": ""
            },
            {
                "id": "f8c26e8e-6f2b-44dc-8201-93da7f61e4b5",
                "created_at": "2023-08-21T11:51:23.23Z",
                "updated_at": "2023-08-21T11:51:23.23Z",
                "name": "Algoritma dan Struktur Data",
                "code": "CIE61010",
                "sks": 3,
                "type": "major",
                "faculty": "FILKOM",
                "major": "Pendidikan Teknologi Informasi"
            }
        ]
    }
    ```

3. ##### Get Course By ID 
    endpoint : ```/course/{id}```     
    method : ```GET METHOD```
    json response : 
    ```
    {
        "status": "success",
        "code": 200,
        "message": "sucessfully fetch course",
        "data": {
            "id": "f8c26e8e-6f2b-44dc-8201-93da7f61e4b5",
            "created_at": "2023-08-21T11:51:23.23Z",
            "updated_at": "2023-08-21T11:51:23.23Z",
            "name": "Algoritma dan Struktur Data",
            "code": "CIE61010",
            "sks": 3,
            "type": "major",
            "faculty": "FILKOM",
            "major": "Pendidikan Teknologi Informasi"
        }
    }
    ```

4. ##### Update Course 
    endpoint : ```/course/{id}```     
    method : ```PATCH METHOD```     
    json request body : 
    ```
    {
        "name": "Pemrograman Dasar",
        "code": "UBU60004",
        "sks": 4,
        "type": "faculty",
        "faculty": "FILKOM"
    }
    ```
    json response : 
    ```
    {
        "status": "success",
        "code": 200,
        "message": "course data updated",
        "data": {
            "record_data": {
                "id": "ddb386f0-a132-4e12-8d72-b6f99a986318",
                "created_at": "0001-01-01T00:00:00Z",
                "updated_at": "2023-08-21T19:03:37.969+07:00",
                "name": "Pemrograman Dasar",
                "code": "UBU60004",
                "sks": 5,
                "type": "faculty",
                "faculty": "FILKOM",
                "major": ""
            }
        }
    }
    ```

5. ##### Delete Course 
    endpoint : ```/course/{id}```     
    method : ```DELETE METHOD```   
    json response : 
    ```
    {
        "status": "success",
        "code": 200,
        "message": "successfully delete course",
        "data": {
            "deleted_record_data": {
                "id": "ddb386f0-a132-4e12-8d72-b6f99a986318",
                "created_at": "0001-01-01T00:00:00Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "name": "",
                "code": "",
                "sks": 0,
                "type": "",
                "faculty": "",
                "major": ""
            }
        }
    }
    ```

6. ##### Get Courses By Param
    endpoint : ```/course?type={}&faculty={}&major={}```       
    method : ```GET```
    params :   
    1. type : must be string, 3 types, university or faculty or major
    2. faculty : must be string
    3. major : must be string

    json response : 
    ```
    {
        "status": "success",
        "code": 200,
        "message": "sucessfully fetch course",
        "data": [
            {
                "id": "0d7e3e21-edbd-485e-863c-735dee873d57",
                "created_at": "2023-08-21T12:30:13.031Z",
                "updated_at": "2023-08-21T12:30:13.031Z",
                "name": "Pemrograman Web dan Teknologi Informasi",
                "code": "CIE62017",
                "sks": 4,
                "type": "major",
                "faculty": "FILKOM",
                "major": "Pendidikan Teknologi Informasi"
            },
            {
                "id": "f8c26e8e-6f2b-44dc-8201-93da7f61e4b5",
                "created_at": "2023-08-21T11:51:23.23Z",
                "updated_at": "2023-08-21T11:51:23.23Z",
                "name": "Algoritma dan Struktur Data",
                "code": "CIE61010",
                "sks": 3,
                "type": "major",
                "faculty": "FILKOM",
                "major": "Pendidikan Teknologi Informasi"
            }
        ]
    }
    ```