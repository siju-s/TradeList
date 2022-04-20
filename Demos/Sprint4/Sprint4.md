## Project : Tradelist (Sprint 4)
* Contributors : Siju Sakaria, Yashasvi Mutteneni, Mansi Singh, Sharwari Marathe

## Backend 
The programming language used is Golang. ORM library for Golang used is GORM and the database used is SQLite.

Endpoints: ![Alt text](endpoints_sprint4.png)

* `post/category/{id}`: This endpoint is used to create a post based on category id or fetch the post using HTTP POST or GET request. 
   Once the request is sent, the post details will be stored in the database.
* `post`:This endpoint is used to get all the posts stored in DB
* `categories`:This endpoint is used to get all the post categories
* `subcategories/{id}`:This endpoint is used to get all the post subcategories for a post category
* `login`: This endpoint is used to log the user into the webpage using the username and password entered by the user. If the user enters the correct credential they will be logged in and a token will be created with an expiration time. This endpoint uses the HTTP POST method.
* `home`: This endpoint retrieves the token and displays a welcome message if the token is valid. It uses the HTTP GET method.
* `refresh`: This endpoint is used to create a new token. A new token will be issued only if the old one is within 30 seconds of expiry. It uses the HTTP POST method.
* `logout`: This endpoint deletes the token and the user is logged out. It uses the HTTP POST method.
* `signup`: This endpoint is used to create a new user using HTTP POST
* `forgot`: This endpoint is used to complete the forgot password request. An email is sent to user with login link
* `reset`: This endpoint is used to reset the password

* Unit tests and API tests have been added for the backend




## Frontend
