



# Pop Up Museum
### Note

🌍 - Endpoints that are available to all users.

👤 - Endpoints that are available to authorized users.

📬 - Endpoints that are available to moderators.

👑 - Endpoints that are available to administrators.

### Navigation:
- [Description](#description-back-to-navigation)
- [Features of API](#features-back-to-navigation)
- [API endpoints](#api-endpoints)
  - [**Auth**](#auth--back-to-navigation)
    - [🌍 POST /sign-up](#-post-sign-up)
    - [🌍 POST /sign-in](#-post-sign-in)
	- [🌍 WebSocket /ws](#-websocket-ws)
  - [**Exhibit**](#exhibit-back-to-navigation)
    - [🌍 GET /exhibit/all](#-get-exhibitall) 
    - [🌍 GET /exhibit/:exhibitId](#-get-exhibitexhibitid)
    - [🌍 GET /exhibit/statuses](#-get-exhibitstatuses)
    - [👑 POST /exhibit/all](#-post-exhibitall) 
    - [👑 POST /exhibit/upload-image](#-post-exhibitupload-image)
    - [👑 PUT /exhibit/:exhibitId](#-put-exhibitexhibitid) 
    - [👑 DELETE /exhibit/:exhibitId](#-delete-exhibitexhibitid) 
  - [**Author**](#author-back-to-navigation)
    - [🌍 GET /author/all](#-get-authorall) 
    - [🌍 GET /author/:authorId](#-get-authorauthorid) 
    - [👑 POST /author/all](#-post-authorall) 
    - [👑 POST /author/upload-image](#-post-authorupload-image)
    - [👑 PUT /author/:authorId](#-put-authorauthorid) 
    - [👑 DELETE /author/:authorId](#-delete-authorauthorid)
  - [**Item**](#item-back-to-navigation)
    - [🌍 GET /item/all](#-get-itemall) 
    - [🌍 GET /item/:itemId](#-get-itemitemid) 
    - [🌍 GET /item/statuses](#-get-itemstatuses)
    - [👑 POST /item/all](#-post-itemall) 
    - [👑 POST /item/upload-image](#-post-itemupload-image)
    - [👑 PUT /item/:itemId](#-put-itemitemid) 
    - [👑 DELETE /item/:itemId](#-delete-itemitemid)
  - [**User**](#user-back-to-navigation)
    - [👤 GET /profile](#-get-profile) 
    - [👤 POST /profile/upload-image](#-post-profileupload-image)
    - [👤 DELETE /profile](#-delete-profile)
  - [**Collection**](#collection-back-to-navigation)
    - [🌍 GET /collection/all](#-get-collectionall)
    - [👤 GET /collection](#-get-collection)
    - [👤 POST /collection](#-post-collection)
    - [👤 DELETE /collection](#-delete-collection)
  - [**Appeal**](#appeal-back-to-navigation)
    - [🌍 GET /appeal/statuses](#-get-appealstatuses)
    - [🌍 POST /appeal](#-post-appeal)
    - [📬 GET /appeal/all](#-get-appealall)
    - [📬 PUT /appeal/:appealId](#-put-appealappealid)
  - [**Reservation**](#reservation-back-to-navigation)
    - [🌍 GET /reservation/statuses](#-get-reservationstatuses)
    - [🌍 POST /reservation](#-post-reservation)
    - [📬 GET /reservation/all](#-get-reservationall)
    - [📬 PUT /reservation/:reservationId](#-put-reservationreservationid)
  - [**Info**](#info--back-to-navigation)
    - [🌍 GET /info](#-get-info)
    - [🌍 GET /faq](#-get-faq)

------------

## Description [`back to navigation`](#navigation)
Hello! This is the API description of the Pop Up Museum project. 

It contains: endpoints, JSON requests and JSON responses. Explanations of non-obvious HTTP requests will be in their short descriptions. 

Features of API and tokens will be located in the "Features" section.

## Features [`back to navigation`](#navigation)
1. **Collections**: there are 2 types of collections - user's and the merging of all existing items and authors. The user collection is the "favorites" and the general collection is a list of all existing items and authors. Nevertheless, both collections are in the same "Collection" section of the documentation. Their HTTP requests might be unclear at first sight, this way I advice you read their descriptions.

2. **JWT Tokens**: all tokens are provided and received through the header field ```Authorization```. There is no ```Bearer``` in both cases, only the token string itself is submitted. ```userId``` and ```isAdmin``` are inserted into tokens. Tokens are valid for 30 minutes. You can create new user and get his user token, however if you want to create an admin/moderator user, you need to switch ```role``` field value from 0 to 1(admin) / 2(moderator) directly in the DB.

3. **Appeals and Reservations**: both authorized users and unauthorized users (guests) can create appeals and reservations. When an entity is creating by an unauthorized user, it is not required to pass the token. When an entity is creating by an authorized user, you must pass the current token in header to associate the created entity with the user's account.

4. **DELETE & POST requests**: there is no need to send the creation/deletion time and date for entities that have such fields, because the date and time values will be inserted on the server side.

5. **DELETE requests**:  all entities that have an isRemoved field can be not just deleted, but restored by submitting a ```{ "isRemoved": false }``` request.

6. **Upload-image enpoints**:

	6.1. Initially the entity is created without a linked image, however you can upload one. To send an image to the endpoint, use form-data request, pass the image in the "file" field.  The image must be no larger than 1MB.

	6.2. To remove an image linked to an entity - send an empty form-data request (without the "file" field) to the upload-image endpoint.

# API endpoints
## Auth:  [`back to navigation`](#navigation)
### 🌍 POST /sign-up
Creates a new user.

**Request body**
```
{
	"email": str(100),
	"password": str(50),
	"phoneNumber": str(25),
	"firstName": str(20),
	"lastName": str(20),
	"middleName": str(20),
	"biography": str(400)
}
```

**Response body**
```
{
	"userId": int
}
```

**Response header**
```
{
	"Authorization": JWT
}
```

### 🌍 POST /sign-in
Request to get a JWT Token based on user credentials.

**Request body**
```
{
	"email": str(100),
	"password": str(50)
}
```

**Response body**
```
{
	"userId": int
}
```

**Response header**
```
{
	"Authorization": JWT
}
```

### 🌍 WebSocket /ws

- The lifetime of the connection is only 5 minutes.
- If the recovery code was incorrect 3 times, the connection is dropped. 
- The "RECOVERY_FAILURE" error can only occur due to server-side issues, although this is very rare.

Events:
```
<client to server>	| <server to client>
			|
START_RECOVERY		| RECOVERY_STARTED -> RECOVERY_CODE_SENDED
			|
SEND_RECOVERY_CODE	| RECOVERY_CODE_SUCCESS -> RECEIVE_NEW_PASSWORD_STARTED
			| or RECOVERY_CODE_FAILURE
			| or if code was wrong 3 times RECOVERY_CODE_NO_ATTEMPTS -> close connection
			|
SEND_NEW_PASSWORD	| RECOVERY_SUCCESS -> close connection
			| or RECOVERY_FAILURE -> close connection

```

A successful example of a request flow:			
```
<client to server>					
{					
	"uuid": "example",		
	"event_type": "START_RECOVERY",	
	"data": "example@example.com"	
}

<server to client>			
{					
	"uuid": "example",		
	"event_type": "RECOVERY_STARTED"	
}
				
{					
	"uuid": "example",		
	"event_type": "RECOVERY_CODE_SENDED"	
}	

<client to server>
{					
	"uuid": "example",		
	"event_type": "SEND_RECOVERY_CODE",	
	"data": "920563"	
}

<server to client>			
{					
	"uuid": "example",		
	"event_type": "RECOVERY_CODE_SUCCESS"
}
				
{					
	"uuid": "example",		
	"event_type": "RECEIVE_NEW_PASSWORD_STARTED"	
}

<client to server>
{					
	"uuid": "example",		
	"event_type": "SEND_NEW_PASSWORD",	
	"data": "my-new-password!"	
}

<server to client>			
{					
	"uuid": "example",		
	"event_type": "RECOVERY_SUCCESS"
}

<connection closed>
```

## Exhibit: [`back to navigation`](#navigation)
### 🌍 GET /exhibit/all
Returns all exhibits.

**Response body**
```
[
	{
		"exhibitId": int,
		"exhibitName": str(100),
		"pathToExhibitImage": str(100),
		"startDate": timestamp,
		"endDate": timestamp,
		"status": int
	},
	{
		"exhibitId": int,
		"exhibitName": str(100),
		"pathToExhibitImage": str(100),
		"startDate": timestamp,
		"endDate": timestamp,
		"status": int
	}
]
```

### 🌍 GET /exhibit/:exhibitId
 Returns an exhibit by the *exhibitId* parameter.
 
**Response body**
```
{
	"exhibitId": int,
	"exhibitName": str(100),
	"pathToExhibitImage": str(100),
	"startDate": timestamp,
	"endDate": timestamp,
	"type": str(50),
	"status": int,
	"description": str(750),
	"location": str(100),
	"website": str(100),
	"authors": [
		{
			"authorId": int,
			"authorName": str(100),
			"pathToAuthorImage": str(100)
		},
		{
			"authorId": int,
			"authorName": str(100),
			"pathToAuthorImage": str(100)
		}
	],
	"items": [
		{
			"itemId": int,
			"authorId": int,
			"itemName": str(100),
			"pathToItemImage": str(100)
		},
		{
			"itemId": int,
			"authorId": int,
			"itemName": str(100),
			"pathToItemImage": str(100)
		}
	]
}
```

### 🌍 GET /exhibit/statuses
Returns statuses of exhibit entity.

**Response body**
```
[
	{
		"statusCode": int,
		"locale": str(20)
	},
	{
		"statusCode": int,
		"locale": str(20)
	}
]
```

### 👑 POST /exhibit/all
Creates a new exhibit.

**Request header**
```
{
	"Authorization": JWT
}
```

**Request body**
```
{
	"exhibitName": str(100),
	"startDate": timestamp,
	"endDate": timestamp,
	"type": str(50),
	"status": int,
	"description": str(750),
	"location": str(100),
	"website": str(100),
	"authors": []int,
	"items": []int
}
```

**Response body**
```
{
	"exhibitId": int
}
```

### 👑 POST /exhibit/upload-image
Updates an image of an exhibit.

**Request header**
```
{
	"Authorization": JWT
}
```

**Request body**
```
form-data

"file": <upload image>
```

**Response body**
```
{
	"status": "ok"
}
```

### 👑 PUT /exhibit/:exhibitId
Updates an exhibit by the *exhibitId* parameter.

**Request header**
```
{
	"Authorization": JWT
}
```

**Request body**
```
{
	"exhibitName": str(100),
	"startDate": timestamp,
	"endDate": timestamp,
	"type": str(50),
	"status": int,
	"description": str(750),
	"location": str(100),
 	"website": str(100),
  	"authors": []int,
	"items": []int
}
```

**Response body**
```
{
	"status": "ok"
}
```

### 👑 DELETE /exhibit/:exhibitId
Removes an exhibit by the *exhibitId* parameter.

**Request header**
```
{
	"Authorization": JWT
}
```

**Request body**
```
{
	"isRemoved": bool
}
```

**Response body**
```
{
	"status": "ok"
}
```



## Author: [`back to navigation`](#navigation)
### 🌍 GET /author/all
Returns all authors.

**Response body**
```
[
	{
		"authorId": int,
		"authorName": str(100),
		"pathToAuthorImage": str(100)
    	},
    	{
		"authorId": int,
		"authorName": str(100),
		"pathToAuthorImage": str(100)
    	}
]
```

### 🌍 GET /author/:authorId
Returns an author by the *authorId* parameter.

**Response body**
```
{
	"authorId": int,
	"authorName": str(100),
	"description": str(500),
	"pathToAuthorImage": str(100),
	"pseudonym": str(50),
	"phoneNumber": str(20),
 	"email": str(30),
	"items": [
		{
			"itemId": int,
			"authorId": int,
			"itemName": str(100),
			"pathToItemImage": str(100)
		},
		{
			"itemId": int,
			"authorId": int,
			"itemName": str(100),
			"pathToItemImage": str(100)
		}
	]
}
```

### 👑 POST /author/all
Creates a new author.

**Request header**
```
{
	"Authorization": JWT
}
```

**Request body**
```
{
	"authorName": str(100),
	"description": str(500),
	"pseudonym": str(50),
	"phoneNumber": str(20),
	"email": str(30)
}
```

**Response body**
```
{
	"authorId": int
}
```

### 👑 POST /author/upload-image
Updates an image of an author.

**Request header**
```
{
	"Authorization": JWT
}
```

**Request body**
```
form-data

"file": <upload image>
```

**Response body**
```
{
	"status": "ok"
}
```

### 👑 PUT /author/:authorId
Updates an author by the *authorId* parameter.

**Request header**
```
{
	"Authorization": JWT
}
```

**Request body**
```
{
	"authorName": str(100),
	"description": str(500),
	"pseudonym": str(50),
	"phoneNumber": str(20),
	"email": str(30)
}
```

**Response body**
```
{
	"status": "ok"
}
```

### 👑 DELETE /author/:authorId
Removes an author by the *authorId* parameter.

**Request header**
```
{
	"Authorization": JWT
}
```

**Request body**
```
{
	"isRemoved": bool
}
```

**Response body**
```
{
	"status": "ok"
}
```



## Item: [`back to navigation`](#navigation)
### 🌍 GET /item/all
Returns all items.

**Response body**
```
[
	{
		"itemId": int,
		"authorId": int,
		"itemName": str(100),
		"pathToItemImage": str(100)
    	},
    	{
		"itemId": int,
		"authorId": int,
		"itemName": str(100),
		"pathToItemImage": str(100)
    	}
]
```

### 🌍 GET /item/:itemId
Returns an item by the *itemId* parameter.

**Response body**
```
{
	"itemId": int,
	"authorId": int,
	"exhibitId": int,
	"itemName": str(100),
	"pathToItemImage": str(100),
	"technique": str(50),
	"height": float,
	"length": float,
	"status": int,
	"created": timestamp
}
```

### 🌍 GET /item/statuses
Returns statuses of item entity.

**Response body**
```
[
	{
		"statusCode": int,
		"locale": str(20)
	},
	{
		"statusCode": int,
		"locale": str(20)
	}
]
```

### 👑 POST /item/all
Creates a new item.

**Request header**
```
{
	"Authorization": JWT
}
```

**Request body**
```
{
	"authorId": int,
	"exhibitId": int,
	"itemName": str(100),
	"technique": str(50),
	"height": float,
	"length": float,
	"status": int,
	"created": timestamp
}
```

**Response body**
```
{
	"itemId": int
}
```

### 👑 POST /item/upload-image
Updates an image of an item.

**Request header**
```
{
	"Authorization": JWT
}
```

**Request body**
```
form-data

"file": <upload image>
```

**Response body**
```
{
	"status": "ok"
}
```

### 👑 PUT /item/:itemId
Updates an item by the *itemId* parameter.

**Request header**
```
{
	"Authorization": JWT
}
```

**Request body**
```
{
	"authorId": int,
	"exhibitId": int,
	"itemName": str(100),
	"technique": str(50),
	"height": float,
	"length": float,
	"status": int,
	"created": timestamp
}
```

**Response body**
```
{
	"status": "ok"
}
```

### 👑 DELETE /item/:itemId
Updates an item by the *itemId* parameter.

**Request header**
```
{
	"Authorization": JWT
}
```

**Request body**
```
{
	"isRemoved": bool
}
```

**Response body**
```
{
	"status": "ok"
}
```


## User: [`back to navigation`](#navigation)
### 👤 GET /profile
Returns a user profile by *JWT token*.

**Request header**
```
{
	"Authorization": JWT
}
```

**Response body**
```
{
	"userId": int,
	"pathToUserImage": str(100),
	"firstName": str(20),
	"lastName": str(20),
	"middleName": str(20),
	"appeals": [
		{
			"appealId": int,
			"userId": int,
			"initials": str(100),
			"contact": str(100),
			"issue": str(100),
			"content": str(500), 
			"status": int,
			"dateCreated": timestamp
		},
		{
			"appealId": int,
			"userId": int,
			"initials": str(100),
			"contact": str(100),
			"issue": str(100),
			"content": str(500), 
			"status": int,
			"dateCreated": timestamp
		}
	],
	"reservations": [
		{
			"reservationId": int,
			"userId": int,
			"firstName": str(50),
			"lastName": str(50),
			"middleName": str(50),
			"date": timestamp,
			"autoBrand": str(50),
			"autoNumber": str(20),
			"status": int
		},
		{
			"reservationId": int,
			"userId": int,
			"firstName": str(50),
			"lastName": str(50),
			"middleName": str(50),
			"date": timestamp,
			"autoBrand": str(50),
			"autoNumber": str(20),
			"status": int
		}
	]
}
```

### 👤 POST /profile/upload-image
Updates an image of an user (profile).

**Request header**
```
{
	"Authorization": JWT
}
```

**Request body**
```
form-data

"file": <upload image>
```

**Response body**
```
{
	"status": "ok"
}
```

### 👤 DELETE /profile
Removes a user by *JWT token*.

**Request header**
```
{
	"Authorization": JWT
}
```

**Response body**
```
{
	"status": "ok"
}
```



## Collection: [`back to navigation`](#navigation)
### 🌍 GET /collection/all
Returns all existing authors and items (merged get all authors & get all items).

**Response body**
```
{
	"authors": [
		{
			"authorId": int,
			"authorName": str(100),
			"pathToAuthorImage": str(100)
    		},
    		{
			"authorId": int,
			"authorName": str(100),
			"pathToAuthorImage": str(100)
		}
	],
	"items": [  
    		{
			"itemId": int,
			"authorId": int,
			"itemName": str(100),
			"pathToItemImage": str(100)
    		},  
    		{
			"itemId": int,
			"authorId": int,
			"itemName": str(100),
			"pathToItemImage": str(100)
		}
	]
}	
```

### 👤 GET /collection
Returns a user's personal collection.

**Request header**
```
{
	"Authorization": JWT
}
```

**Response body**
```
[  
	{
		"itemId": int,
		"authorId": int,
		"itemName": str(100),
		"pathToItemImage": str(100)
	},  
	{
		"itemId": int,
		"authorId": int,
		"itemName": str(100),
		"pathToItemImage": str(100)
	}  
]
```

### 👤 POST /collection
Adds an item into a user's personal collection.

**Request header**
```
{
	"Authorization": JWT
}
```

**Request body**
```
{
	"itemId": int
}
```

**Response body**
```
{
	"itemId": int
}
```

### 👤 DELETE /collection
Removes an item from a user's personal collection.

**Request header**
```
{
	"Authorization": JWT
}
```

**Request body**
```
{
	"itemId": int
}
```

**Response body**
```
{
	"status": "ok"
}
```



## Appeal: [`back to navigation`](#navigation)
### 🌍 GET /appeal/statuses
Returns statuses of appeal entity.

**Response body**
```
[
	{
		"statusCode": int,
		"locale": str(20)
	},
	{
		"statusCode": int,
		"locale": str(20)
	}
]
```

### 🌍 POST /appeal
Creates a new appeal. Provide JWT if user creates an appeal, otherwise the request will be created in public form (from an unauthorized user).

**Request header**
```
{
	"Authorization": JWT (if needed)
}
```

**Request body**
```
{
	"initials": str(100),
	"contact": str(100),
	"issue": str(100),
	"content": str(500)
}
```

**Response body**
```
{
	"appealId": int
}
```

### 📬 GET /appeal/all
Returns all appeals.

**Request header**
```
{
	"Authorization": JWT
}
```

**Response body**
```
[
	{
		"appealId": int,
		"userId": int,
		"initials": str(100),
		"contact": str(100),
		"issue": str(100),
		"content": str(500), 
		"status": int,
		"dateCreated": timestamp
	},
	{
		"appealId": int,
		"userId": int,
		"initials": str(100),
		"contact": str(100),
		"issue": str(100),
		"content": str(500), 
		"status": int,
		"dateCreated": timestamp
	}
]
```


### 📬 PUT /appeal/:appealId
Confirms an appeal by the *appealId* parameter.

**Request header**
```
{
	"Authorization": JWT
}
```

**Request body**
```
{
	"status": int
}
```

**Response body**
```
{
	"status": "ok"
}
```



## Reservation: [`back to navigation`](#navigation)
### 🌍 GET /reservation/statuses
Returns statuses of reservation entity.

**Response body**
```
[
	{
		"statusCode": int,
		"locale": str(20)
	},
	{
		"statusCode": int,
		"locale": str(20)
	}
]
```

### 🌍 POST /reservation
Creates a new reservation. Provide JWT if user creates an appeal, otherwise the request will be created in public form (from an unauthorized user).

**Request header**
```
{
	"Authorization": JWT (if needed)
}
```

**Request body**
```
{
	"firstName": str(50),
	"lastName": str(50),
	"middleName": str(50),
	"date": timestamp,
	"autoBrand": str(50),
	"autoNumber": str(20)
}
```

**Response body**
```
{
	"reservationId": int
}
```

### 📬 GET /reservation/all
Returns all reservations.

**Request header**
```
{
	"Authorization": JWT
}
```

**Response body**
```
[
	{
		"reservationId": int,
		"userId": int,
		"firstName": str(50),
		"lastName": str(50),
		"middleName": str(50),
		"date": timestamp,
		"autoBrand": str(50),
		"autoNumber": str(20),
		"status": int
    	},
    	{
		"reservationId": int,
		"userId": int,
		"firstName": str(50),
		"lastName": str(50),
		"middleName": str(50),
		"date": timestamp,
		"autoBrand": str(50),
		"autoNumber": str(20),
		"status": int
    	}
]
```

### 📬 PUT /reservation/:reservationId
Confirms a reservation by the *reservationId* parameter.

**Request header**
```
{
	"Authorization": JWT
}
```

**Request body**
```
{
	"status": int
}
```

**Response body**
```
{
	"status": "ok"
}
```

## Info:  [`back to navigation`](#navigation)
### 🌍 GET /info
Returns info.

**Response body**
```
{
	"museumPhoneNumber": str(20),
	"email": str(100),
	"workSchedule": str(100),
	"address": str(100),
	"organizerPhoneNumber": str(20),
	"website": str(50)
}
```

### 🌍 GET /faq
Returns all questions.

**Response body**
```
[
	{
		"title": str(50),
		"description": str(300)
	},
	{
		"title": str(50),
		"description": str(300)
	}
]
```

