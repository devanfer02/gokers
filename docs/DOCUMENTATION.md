## API DOCUMENTATION

### List Entities
1. [Student](#student)
2. [Lecturer](#lecturer)
3. [Course](#course)
4. [Class](#class)

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
    }
    ```

6. ##### Get Courses By Param
    endpoint : ```/course?type={}&faculty={}&major={}```       
    method : ```GET METHOD```
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

### Class
1. ##### Register Class
    endpoint : ```/class/register```     
    method : ```POST METHOD```       
    json body request : 
    ```
    {
        "course_id": "1ab02384-f491-4b7b-85de-ff64ee2c6d61",
        "lecturer_id": "cdb5ac3b-3ff0-4226-bb86-a6a03947d5c6",
        "class_name": "A Teknik Informatika",
        "class_room": "TIF06 - F2.9"
    }
    ```
    json response : 
    ```
    {
        "status": "success",
        "code": 200,
        "message": "class registered to system",
        "data": {
            "record_data": {
                "id": "ba96d378-7906-4461-8b13-effca58639ad",
                "created_at": "2023-08-25T01:02:05.122+07:00",
                "updated_at": "2023-08-25T01:02:05.122+07:00",
                "course_id": "1ab02384-f491-4b7b-85de-ff64ee2c6d61",
                "lecturer_id": "cdb5ac3b-3ff0-4226-bb86-a6a03947d5c6",
                "class_name": "A Teknik Informatika",
                "class_room": "TIF06 - F2.9",
                "quota": 40
            }
        }
    }
    ```

2. ##### Update Class
    endpoint : ```/class/update/{id}```        
    method : ```PATCH METHOD```     
    json body request : 
    ```
    {
        "course_id": "1ab02384-f491-4b7b-85de-ff64ee2c6d61",
        "lecturer_id": "cdb5ac3b-3ff0-4226-bb86-a6a03947d5c6",
        "class_name": "A Teknik Informatika",
        "class_room": "TIF06 - F2.9",
        "quota": 38
    }
    ```
    json response : 
    ```
    {
        "status": "success",
        "code": 200,
        "message": "class data update",
        "data": {
            "record_data": {
                "id": "ba96d378-7906-4461-8b13-effca58639ad",
                "created_at": "0001-01-01T00:00:00Z",
                "updated_at": "2023-08-25T01:12:24.216+07:00",
                "course_id": "1ab02384-f491-4b7b-85de-ff64ee2c6d61",
                "lecturer_id": "cdb5ac3b-3ff0-4226-bb86-a6a03947d5c6",
                "class_name": "A Teknik Informatika",
                "class_room": "TIF06 - F2.9",
                "quota": 38
            }
        }
    }
    ```

3. ##### Delete Class      
    endpoint : ```/class/delete/{id}```      
    method : ```DELETE METHOD```       
    json response : 
    ```
    {
        "status": "success",
        "code": 200,
        "message": "successfully delete class"
    }
    ```

4. ##### Get Classes Student Entity    
    endpoint : ```/class```      
    method : ```GET METHOD```      
    json response : 
    ```
    {
        "status": "success",
        "code": 200,
        "message": "successfully fetch classes",
        "data": [
            {
                "id": "07b086bb-36ee-42ad-9610-05048c4de6bd",
                "created_at": "2023-08-24T18:44:57.122Z",
                "updated_at": "2023-08-24T18:44:57.122Z",
                "course_info": {
                    "id": "1ab02384-f491-4b7b-85de-ff64ee2c6d61",
                    "created_at": "2023-08-24T03:47:07.492Z",
                    "updated_at": "2023-08-24T04:13:18.635Z",
                    "name": "Pemrograman Dasar",
                    "code": "CIF62001",
                    "sks": 5,
                    "type": "major",
                    "faculty": "FILKOM",
                    "major": "Teknik Informatika",
                    "upper": {
                        "id": "00000000-0000-0000-0000-000000000000",
                        "created_at": "0001-01-01T00:00:00Z",
                        "updated_at": "0001-01-01T00:00:00Z",
                        "course_id": "",
                        "prequisite_id": ""
                    }
                },
                "lecturer_info": {
                    "id": "cdb5ac3b-3ff0-4226-bb86-a6a03947d5c6",
                    "created_at": "2023-08-24T17:47:49.645Z",
                    "updated_at": "2023-08-24T17:47:49.645Z",
                    "name": "Devan",
                    "ndn": "002351501001",
                    "email": "dvn@ub.ac.id",
                    "password": "$2a$10$QrMkWwgaJNf81njJ8ukewejddxFWihIDGCi5.ROgD/BCg3lRnOnSm",
                    "faculty": "FILKOM",
                    "major": "Teknik Informatika",
                    "Class": null
                },
                "course_id": "1ab02384-f491-4b7b-85de-ff64ee2c6d61",
                "lecturer_id": "cdb5ac3b-3ff0-4226-bb86-a6a03947d5c6",
                "class_name": "A Teknik Informatika",
                "class_room": "TIF06 - F2.9",
                "quota": 40
            }
        ]
    }
    ```

5. ##### Get Classes Admin Entity   
    endpoint : ```/class/admin```        
    method : ```GET METHOD```       
    json response : 
    ```
    {
        "status": "success",
        "code": 200,
        "message": "successfully fetch classes",
        "data": [
            {
                "id": "04f045a9-b589-4106-becb-7f9523fa6db1",
                "created_at": "2023-08-24T18:45:43.209Z",
                "updated_at": "2023-08-24T18:45:43.209Z",
                "course_info": {
                    "id": "4a58050f-90fa-4ade-94fc-833203af1030",
                    "created_at": "2023-08-24T03:50:05.787Z",
                    "updated_at": "2023-08-24T03:50:05.787Z",
                    "name": "Pemrograman Dasar",
                    "code": "CIT62001",
                    "sks": 4,
                    "type": "major",
                    "faculty": "FILKOM",
                    "major": "Teknologi Informasi",
                    "upper": {
                        "id": "00000000-0000-0000-0000-000000000000",
                        "created_at": "0001-01-01T00:00:00Z",
                        "updated_at": "0001-01-01T00:00:00Z",
                        "course_id": "",
                        "prequisite_id": ""
                    }
                },
                "lecturer_info": {
                    "id": "cdb5ac3b-3ff0-4226-bb86-a6a03947d5c6",
                    "created_at": "2023-08-24T17:47:49.645Z",
                    "updated_at": "2023-08-24T17:47:49.645Z",
                    "name": "Devan",
                    "ndn": "002351501001",
                    "email": "dvn@ub.ac.id",
                    "password": "$2a$10$QrMkWwgaJNf81njJ8ukewejddxFWihIDGCi5.ROgD/BCg3lRnOnSm",
                    "faculty": "FILKOM",
                    "major": "Teknik Informatika",
                    "Class": null
                },
                "course_id": "4a58050f-90fa-4ade-94fc-833203af1030",
                "lecturer_id": "cdb5ac3b-3ff0-4226-bb86-a6a03947d5c6",
                "class_name": "D Teknologi Informasi",
                "class_room": "TIF06 - F3.9",
                "quota": 40
            },
            {
                "id": "07b086bb-36ee-42ad-9610-05048c4de6bd",
                "created_at": "2023-08-24T18:44:57.122Z",
                "updated_at": "2023-08-24T18:44:57.122Z",
                "course_info": {
                    "id": "1ab02384-f491-4b7b-85de-ff64ee2c6d61",
                    "created_at": "2023-08-24T03:47:07.492Z",
                    "updated_at": "2023-08-24T04:13:18.635Z",
                    "name": "Pemrograman Dasar",
                    "code": "CIF62001",
                    "sks": 5,
                    "type": "major",
                    "faculty": "FILKOM",
                    "major": "Teknik Informatika",
                    "upper": {
                        "id": "00000000-0000-0000-0000-000000000000",
                        "created_at": "0001-01-01T00:00:00Z",
                        "updated_at": "0001-01-01T00:00:00Z",
                        "course_id": "",
                        "prequisite_id": ""
                    }
                },
                "lecturer_info": {
                    "id": "cdb5ac3b-3ff0-4226-bb86-a6a03947d5c6",
                    "created_at": "2023-08-24T17:47:49.645Z",
                    "updated_at": "2023-08-24T17:47:49.645Z",
                    "name": "Devan",
                    "ndn": "002351501001",
                    "email": "dvn@ub.ac.id",
                    "password": "$2a$10$QrMkWwgaJNf81njJ8ukewejddxFWihIDGCi5.ROgD/BCg3lRnOnSm",
                    "faculty": "FILKOM",
                    "major": "Teknik Informatika",
                    "Class": null
                },
                "course_id": "1ab02384-f491-4b7b-85de-ff64ee2c6d61",
                "lecturer_id": "cdb5ac3b-3ff0-4226-bb86-a6a03947d5c6",
                "class_name": "A Teknik Informatika",
                "class_room": "TIF06 - F2.9",
                "quota": 40
            }
        ]
    }
    ```

6. ##### Get Class By ID Student Entity
    endpoint : ```/class/{id}```       
    method : ```GET METHOD```        
    json response : 
    ```
    {
        "status": "success",
        "code": 200,
        "message": "succesfully fetch class",
        "data": {
            "id": "07b086bb-36ee-42ad-9610-05048c4de6bd",
            "created_at": "2023-08-24T18:44:57.122Z",
            "updated_at": "2023-08-24T18:44:57.122Z",
            "course_info": {
                "id": "1ab02384-f491-4b7b-85de-ff64ee2c6d61",
                "created_at": "2023-08-24T03:47:07.492Z",
                "updated_at": "2023-08-24T04:13:18.635Z",
                "name": "Pemrograman Dasar",
                "code": "CIF62001",
                "sks": 5,
                "type": "major",
                "faculty": "FILKOM",
                "major": "Teknik Informatika",
                "upper": {
                    "id": "00000000-0000-0000-0000-000000000000",
                    "created_at": "0001-01-01T00:00:00Z",
                    "updated_at": "0001-01-01T00:00:00Z",
                    "course_id": "",
                    "prequisite_id": ""
                }
            },
            "lecturer_info": {
                "id": "cdb5ac3b-3ff0-4226-bb86-a6a03947d5c6",
                "created_at": "2023-08-24T17:47:49.645Z",
                "updated_at": "2023-08-24T17:47:49.645Z",
                "name": "Devan",
                "ndn": "002351501001",
                "email": "dvn@ub.ac.id",
                "password": "$2a$10$QrMkWwgaJNf81njJ8ukewejddxFWihIDGCi5.ROgD/BCg3lRnOnSm",
                "faculty": "FILKOM",
                "major": "Teknik Informatika",
                "Class": null
            },
            "course_id": "1ab02384-f491-4b7b-85de-ff64ee2c6d61",
            "lecturer_id": "cdb5ac3b-3ff0-4226-bb86-a6a03947d5c6",
            "class_name": "A Teknik Informatika",
            "class_room": "TIF06 - F2.9",
            "quota": 40
        }
    }
    ```
    json response if class.course.major is not same as student.major : 
    ```
    {
        "status": "error",
        "code": 404,
        "message": "class not found"
    }
    ```
