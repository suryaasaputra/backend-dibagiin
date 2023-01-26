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
    "phone_number":"string",
    "gender":"sting",
    "address":"string",
    "lat":"float64",
    "lng":"float64",
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
        "phone_number":"string",
        "gender":"string",
        "address":"string",
        "profil_photo_url":"string",
        "createdAt":"date",
        "lat":"float64",
        "lng":"float64",
    }
}
```
**========================================================================================**

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
    "email":"string",
    "password":"string",
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
        "phone_number":"string",
        "address":"string",
        "profil_photo_url":"string",
        "lat":"float64",
        "lng":"float64",
        "token":"string",
        "login_time":"date"
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
    - Authorization: Bearer Token
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
        "phone_number":"string", 
        "gender":"string",
        "address":"string",
        "profil_photo_url":"string",
        "createdAt":"date",
        "updatedAt":"date",
        "Donation":[
            {
                "id":"string",
                "title":"string",
                "description":"string",
                "weight":"int",
                "photo_url":"string",
                "location":"string",
                "status":"string",
                "taker_id":"string",
                "created_at":"time",
                "updated_at":"time",
            },
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
    "phone_number":"string",
    "gender":"string",
    "address":"string",
    "lat":"float64",
    "lng":"float64",
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
        "profil_photo_url":"string",
        "phone_number":"string",
        "age":"number", 
        "gender":"string",
        "address":"string",
        "lat":"float64",
        "lng":"float64",
        "updated_at":"date",
    }
}
```
**========================================================================================**
## Update User Profile Photo
**Request** :
- Endpoint : `/user/{user_name}/ProfilPhoto`
- Method : PUT
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
    "data":{
        "id":"string",
        "user_name":"string",
        "email":"string",
        "full_name":"string",
        "profil_photo_url":"string",
        "phone_number":"string",
        "gender":"string",
        "address":"string",
        "updated_at":"date",
    }
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
## Create Donation

**Request** :
- Method : POST
- Endpoint : `/donation`
- Header :
    - Content-Type: multipart/form-data 
    - Accept: application/json
- Body :

```json 
{
    "title" : "string",
    "description" : "string",
    "weight" : "int",
    "location" : "string",
    "lat" : "float64",
    "lng" : "float64",
    "donation_photo":"file|.jpeg, .jpg, .png",
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
        "title" : "string",
        "description" : "string",
        "weight" : "string",
        "location" : "string",
        "lat":"float64",
        "lng":"float64",
        "photo_url":"string",
        "user_id":"string",
        "status":"string",
        "created_at":"time",
    }
}
```
**=================================================================================================================**
## Get All Donation

**Request** :
- Method : GET
- Endpoint : `/donation`
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
            "id":"string",
            "title" : "string",
            "description" : "string",
            "weight" : "string",
            "location" : "string",
            "lat":"float64",
            "lng":"float64",
            "photo_url":"string",
            "user_id":"string",
            "status":"string",
            "created_at":"date",
            "updated_at":"date",
        },
        {
            "id":"string",
            "title" : "string",
            "description" : "string",
            "location" : "string",
            "lat":"float64",
            "lng":"float64",
            "photo_url":"string",
            "user_id":"string",
            "status":"string",
            "created_at":"date",
            "updated_at":"date",
        }
    ]
}
```
**=================================================================================================================**
## Get Donation Detail

**Request** :
- Method : GET
- Endpoint : `/donation/:donationId`
- Header :
    - Accept: application/json

**Response** :

```json 
{
    "error":"boolean",
    "code":"number",
    "message":"string",
    "data":
        {
            "id":"string",
            "title" : "string",
            "description" : "string",
            "weight" : "string",
            "location" : "string",
            "lat":"float64",
            "lng":"float64",
            "photo_url":"string",
            "user_id":"string",
            "status":"string",
            "created_at":"date",
            "updated_at":"date",
        },
}
```
**=================================================================================================================**
## Edit Donation

**Request** :
- Method : PUT
- Endpoint : `/donation/:donationId`
- Header :
    - Content-Type: application/json | application/x-www-form-urlencoded
    - Accept: application/json
    - Authorization: Bearer Token
- Body :

```json 
{
    "title" : "string",
    "description" : "string",
    "weight" : "string",
    "location" : "string",
    "lat" : "float64",
    "lng" : "float64",
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
        "title" : "string",
        "description" : "string",
        "weight":"string",
        "location" : "string",
        "lat":"float64",
        "lng":"float64",
        "photo_url":"string",
        "user_id":"string",
        "status":"string",
        "updated_at":"time",
    }
}
```
**=================================================================================================================**

## Delete Donation
**Request** :
- Endpoint : `/donation/:donationId`
- Method : DELETE
- Header : 
    - Accept: application/json
    - Authorization: Bearer Token
- Body :

```json 
{
    "message" : "string",
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
**=================================================================================================================**

## Send Donation Request
**Request** :
- Endpoint : `/donation/:donationId`
- Method : POST
- Header : 
    - Accept: application/json
    - Authorization: Bearer Token
    
**Response** :
```json
{
    "error":"boolean",
    "code":"number",
    "message":"string",
     "data":{
        "id":"string",
        "user_id" : "string",
        "donation_id" : "string",
        "donator_id" : "string",
        "message" : "string",
        "status" : "string",
        "created_at":"time",
    }
}
```
**=================================================================================================================**
## Get All Received Donation Request

**Request** :
- Method : GET
- Endpoint : `/request`
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
            "id":"string",
            "user_id" : "string",
            "donation_id" : "string",
            "donator_id" : "string",
            "status" : "string",
            "message" : "string",
            "created_at":"date",
            "updated_at":"date",
        },
        {
            "id":"string",
            "user_id" : "string",
            "donation_id" : "string",
            "donator_id" : "string",
            "status" : "string",
            "message" : "string",
            "created_at":"date",
            "updated_at":"date",
        }
    ]
}
```
**=================================================================================================================**
## Confirm Donation Request

**Request** :
- Method : POST
- Endpoint : `/request/:requestId`
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
**=================================================================================================================**
## Reject Donation Request

**Request** :
- Method : DELETE
- Endpoint : `/request/:requestId`
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
**=================================================================================================================**
## Get All Submitted Donation Request

**Request** :
- Method : GET
- Endpoint : `/donation/request`
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
            "id":"string",
            "user_id" : "string",
            "donation_id" : "string",
            "donator_id" : "string",
            "status" : "string",
            "message" : "string",
            "created_at":"date",
            "updated_at":"date",
        },
        {
            "id":"string",
            "user_id" : "string",
            "donation_id" : "string",
            "donator_id" : "string",
            "status" : "string",
            "message" : "string",
            "created_at":"date",
            "updated_at":"date",
        }
    ]
}
```
**=================================================================================================================**
## Cancel Send Request Donation

**Request** :
- Method : DELETE
- Endpoint : `/donation/request/:requestId`
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
**=================================================================================================================**
## Get All Notification

**Request** :
- Method : GET
- Endpoint : `/notification`
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
            "id":"string",
            "user_id" : "string",
            "donation_id" : "string",
            "donator_id" : "string",
            "donation_request_id" : "string",
            "type" : "string",
            "message" : "string",
            "created_at":"date",
            "updated_at":"date",
        },
        {
            "id":"string",
            "user_id" : "string",
            "donation_id" : "string",
            "donator_id" : "string",
            "donation_request_id" : "string",
            "type" : "string",
            "message" : "string",
            "created_at":"date",
            "updated_at":"date",
        }
    ]
}
```
