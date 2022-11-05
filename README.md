# API Spec

<!-- ## Authentication

All API must use this authentication

**Request** :
- Header :
    - X-Api-Key : "your secret api key" -->
## Register User
**Request** :
- Endpoint : `/register`
- Method : POST
- Header : 
    - Content-Type: application/json | application/x-www-form-urlencoded
    - Accept: application/json
- Body :
```json
{
    "user_name":"string, unique",
    "email":"string, unique",
    "full_name":"string",
    "password":"string, minlength(8)",
    "phone_number":"string, unique",
    "age":"number",
    "gender":"sting",
    "address":"string",
}
```
**Response** :
```json
{
    "error":"boolean",
    "code":"number",
    "message":"string",
    "data":{
        "id":"string",
        "user_name":"string",
        "email":"string",
        "full_name":"string",
        "phone_number":"string, unique",
        "age":"number", 
        "gender":"string",
        "address":"string",
        "profil_photo_url":"string",
        "createdAt":"date",
    }
}
```
**========================================================================================**
## Check username/email availability
**Request** :
- Endpoint : `/user/exists?email={email}` || `/user/exists?user_name={user_name}`
- Method : GET
- Header : 
    - Accept: application/json

**Response** :
```json
{
    "error":"boolean",
    "code":"number",
    "message":"string",
}
```
## Login User
**Request** :
- Endpoint : `/login`
- Method : POST
- Header : 
    - Content-Type: application/json | application/x-www-form-urlencoded
    - Accept: application/json
- Body :
```json
{
    "email":"string, unique",
    "password":"string, minlength(8)",
}
```
**Response** :
```json
{
    "error":"boolean",
    "code":"number",
    "message":"string",
    "data":{
        "token":"string"
    }
}
```
**========================================================================================**
## Get User Profile
**Request** :
- Endpoint : `/user/{user_name}`
- Method : GET
- Header : 
    - Accept: application/json 
    
**Response** :
```json
{
    "error":"boolean",
    "code":"number",
    "message":"string",
    "data":{
        "id":"string",
        "user_name":"string",
        "email":"string",
        "full_name":"string",
        "phone_number":"string, unique",
        "age":"number", 
        "gender":"string",
        "address":"string",
        "profil_photo_url":"string",
        "createdAt":"date",
        "updatedAt":"date",
        "Posts":[
            {},
        ]
    }
}
```
**========================================================================================**
## Update User Profile 
**Request** :
- Endpoint : `/user/{user_name}`
- Method : PUT
- Header : 
    - Accept: application/json
    - Content-Type: application/json | application/x-www-form-urlencoded
    - Authorization: Bearer Token
- Body :
```json
{
    "email":"string, unique",
    "user_name":"string, unique",
    "full_name":"string",
    "phone_number":"string, unique",
    "age":"number", 
    "gender":"string",
    "address":"string",
}
```
**Response** :
```json
{
    "error":"boolean",
    "code":"number",
    "message":"string",
    "data":{
        "id":"string",
        "user_name":"string",
        "email":"string",
        "full_name":"string",
        "phone_number":"string, unique",
        "age":"number", 
        "gender":"string",
        "address":"string",
        "updatedAt":"date",
    }
}
```
**========================================================================================**
## Update User Profile Photo
**Request** :
- Endpoint : `/user/{user_name}/UpdateProfilPhoto`
- Method : PATCH
- Header : 
    - Accept: application/json
    - Content-Type: multipart/form-data
    - Authorization: Bearer Token
- Body :
```json
{
    "profil_photo":"file|.jpeg, .jpg, .png"
}
```
**Response** :
```json
{
    "error":"boolean",
    "code":"number",
    "message":"string",
}
```
**========================================================================================**
## Delete User Account
**Request** :
- Endpoint : `/user/{user_name}`
- Method : DELETE
- Header : 
    - Accept: application/json
    - Authorization: Bearer Token
    
**Response** :
```json
{
    "error":"boolean",
    "code":"number",
    "message":"string",
}
```
**=================================================================================================================**
## Create Post

**Request** :
- Method : POST
- Endpoint : `/post`
- Header :
    - Content-Type: multipart/form-data 
    - Accept: application/json
- Body :

```json 
{
    "title" : "string",
    "description" : "string",
    "location" : "string",
    "photo":"file",
    "donation_type":"string",
}
```

**Response** :

```json 
{
    "error":"boolean",
    "code":"number",
    "message":"string",
    "data":{
        "post_id":"string",
        "title" : "string",
        "description" : "string",
        "location" : "string",
        "photo_url":"string",
        "donation_type":"string",
        "user_id":"string",
        "status":"string",
        "createdAt":"string",
    }
}
```
**=================================================================================================================**
## Get Post Feed

**Request** :
- Method : GET
- Endpoint : `/`
- Header :
    - Accept: application/json

**Response** :

```json 
{
    "error":"boolean",
    "code":"number",
    "message":"string",
    "data":[
        {
            "post_id":"string",
            "title" : "string",
            "description" : "string",
            "location" : "string",
            "photo_url":"string",
            "donation_type":"string",
            "user_id":"string",
            "status":"string",
            "createdAt":"string",
            "updatedAt":"string",
        },
        {
            "post_id":"string",
            "title" : "string",
            "description" : "string",
            "location" : "string",
            "photo_url":"string",
            "donation_type":"string",
            "user_id":"string",
            "status":"string",
            "createdAt":"string",
            "updatedAt":"string",
        },
    ]
}
```
**=================================================================================================================**
## Update Post

**Request** :
- Method : PUT
- Endpoint : `/post/:postId`
- Header :
    - Content-Type: multipart/form-data 
    - Accept: application/json
    - Authorization: Bearer Token
- Body :

```json 
{
    "title" : "string",
    "description" : "string",
    "location" : "string",
    "photo":"file",
    "donation_type":"string",
}
```

**Response** :

```json 
{
    "error":"boolean",
    "code":"number",
    "message":"string",
    "data":{
        "post_id":"string",
        "title" : "string",
        "description" : "string",
        "location" : "string",
        "photo_url":"string",
        "donation_type":"string",
        "user_id":"string",
        "status":"string",
        "updatedAt":"string"
    }
}
```
**=================================================================================================================**

## Delete Post
**Request** :
- Endpoint : `/post/:postId`
- Method : DELETE
- Header : 
    - Accept: application/json
    - Authorization: Bearer Token
    
**Response** :
```json
{
    "error":"boolean",
    "code":"number",
    "message":"string",
}
```
