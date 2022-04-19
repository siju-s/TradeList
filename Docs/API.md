<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Get Posts](#get-posts)
- [Get Posts for user](#get-posts-for-user)
- [Get Posts for category](#get-posts-for-category)
- [Get Posts for subcategory](#get-posts-for-subcategory)
- [Create Post](#create-post)
- [Fetch Categories](#fetch-categories)
- [Create Job post](#create-job-post)
- [Login](#login)
- [Signup](#signup)
- [Home](#home)
- [Logout](#logout)
- [Refresh](#refresh)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Get Posts

Returns all the posts

* **URL**

  /post

* **Method:**

  `GET`

*  **URL Params**

   **Required:**

   `None`

* **Data Params**

  None

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
```
[
{
"ID": 1,
"CreatedAt": "2022-01-30T20:08:23.115766-05:00",
"UpdatedAt": "2022-01-30T20:08:23.115766-05:00",
"DeletedAt": null,
"SellerId": 1,
"CategoryId": 1,
"SubcategoryId": 1,
"Title": "Post1",
"Description": "Post1Desc",
"IsHidden": false,
"IsFlagged": false,
"IsDeleted": false,
"HasImage": false
}
]
```

* **Sample Call:**

  ```javascript
    $.ajax({
      url: "/posts",
      dataType: "json",
      type : "GET",
      success : function(r) {
        console.log(r);
      }
    });
  ```
## Get Posts for user

Returns all the posts

* **URL**

  /post/user/id

* **Method:**

  `GET`

*  **URL Params**

   **Required:**

   `id` exists in db

* **Data Params**

  None

* **Success Response:**

    * **Code:** 200 <br />
      **Content:**
```
{
    "data": [
        {
            "ID": 17,
            "CreatedAt": "2022-04-18T18:06:58.914174-04:00",
            "UpdatedAt": "2022-04-18T18:06:58.914174-04:00",
            "DeletedAt": null,
            "SellerId": 15,
            "CategoryId": 1,
            "SubcategoryId": 1,
            "Title": "Looking for Web developer",
            "Description": "Web developer required",
            "IsHidden": false,
            "IsFlagged": false,
            "IsDeleted": false,
            "Image": []
        }
    ],
    "message": "Records found",
    "status": 200
}
```

* **Sample Call:**

  ```javascript
    $.ajax({
      url: "/post/user/15",
      dataType: "json",
      type : "GET",
      success : function(r) {
        console.log(r);
      }
    });
  ```
## Edit Post

Edit a post

* **URL**

  /post/id/user/userid
* **Method:**

  `PUT`

* **URL Params**

  **Required:**

  `id` and `postid`


* **Success Response:**

    * **Code:** 200 <br />
      **Content:**
```
{
      "message": "Postid 4 updated",
      "status": 200
}
```
* **Sample input**
```
{
      
        "Title": "test",
        "Description": "Test Description"
}
```  

## Delete posts by user

Returns all the posts

* **URL**

  /post/postid/user/id

* **Method:**

  `DELETE`

*  **URL Params**

   **Required:**

   `user id` anmd `postid` exists in db

* **Data Params**

  None

* **Success Response:**

    * **Code:** 200 <br />
      **Content:**
```
{
    "message": "Postid 4 deleted",
    "status": 200
}
```

* **Sample Call:**

  ```javascript
    $.ajax({
      url: "/post/4/user/1",
      dataType: "json",
      type : "DELETE",
      success : function(r) {
        console.log(r);
      }
    });
  ```

  
## Get Posts for category

Returns all the posts

* **URL**

  /post/category/id

* **Method:**

  `GET`

*  **URL Params**

   **Required:**

   `None`

* **Data Params**

  None

* **Success Response:**

    * **Code:** 200 <br />
      **Content:**
```
{
    "data": [
        {
            "Post": {
                "ID": 1,
                "CreatedAt": "2022-04-01T16:06:27.482064-04:00",
                "UpdatedAt": "2022-04-01T16:06:27.482064-04:00",
                "DeletedAt": null,
                "SellerId": 1,
                "CategoryId": 1,
                "SubcategoryId": 1,
                "Title": "test1",
                "Description": "test1desc",
                "IsHidden": false,
                "IsFlagged": false,
                "IsDeleted": false,
                "Image": null
            },
            "Job": {
                "ID": 1,
                "PostId": 1,
                "SubcategoryId": 1,
                "Salary": 500,
                "Pay": "monthly",
                "Type": "fulltime",
                "Location": "remote",
                "Place": "Gainesville"
            }
        }
    ],
    "message": "Post found",
    "status": 200
}
```

* **Sample Call:**

  ```javascript
    $.ajax({
      url: "/post/category/1",
      dataType: "json",
      type : "GET",
      success : function(r) {
        console.log(r);
      }
    });
  ```
## Get Posts for subcategory

Returns all the posts

* **URL**

  /post/subcategory/id

* **Method:**

  `GET`

*  **URL Params**

   **Required:**

   `None`

* **Data Params**

  None

* **Success Response:**

    * **Code:** 200 <br />
      **Content:**
```
{
    "data": [
        {
            "Post": {
                "ID": 1,
                "CreatedAt": "2022-04-01T16:06:27.482064-04:00",
                "UpdatedAt": "2022-04-01T16:06:27.482064-04:00",
                "DeletedAt": null,
                "SellerId": 1,
                "CategoryId": 1,
                "SubcategoryId": 1,
                "Title": "test1",
                "Description": "test1desc",
                "IsHidden": false,
                "IsFlagged": false,
                "IsDeleted": false,
                "Image": null
            },
            "Job": {
                "ID": 1,
                "PostId": 1,
                "SubcategoryId": 1,
                "Salary": 500,
                "Pay": "monthly",
                "Type": "fulltime",
                "Location": "remote",
                "Place": "Gainesville"
            }
        }
    ],
    "message": "Post found",
    "status": 200
}
```

* **Sample Call:**

  ```javascript
    $.ajax({
      url: "/post/subcategory/1",
      dataType: "json",
      type : "GET",
      success : function(r) {
        console.log(r);
      }
    });
  ```

## Create Post

Creates a new post

* **URL**

  /post (Deprecated, use /post/category/{id}) where id is categoryid

* **Method:**

  `POST`

* **URL Params**

   **Required:**

   `None`

* **Data Params** (Post request)

  `SellerId`, `CategoryId`, `SubCategoryId`, `Title`, `Description`


* **Constraints**
  `CategoryId`, `SellerId` and `SubCategoryId` should exist in the DB


* **Success Response:**

    * **Code:** 200 <br />
      **Content:**
```
[]
```
* **Sample input**
```
files: //Filelist from form data
data: "  {
  "SellerId": 1,
  "CategoryId": 1,
  "SubcategoryId": 1,
  "Title": "Test3",
  "Description": "Test3Desc"
  }"
```
## Fetch Categories

* **URL**

  /categories

* **Method:**

  `GET`

* **URL Params**

  **Required:**

  `None`

* **Data Params** (Post request)

  `None`

* **Constraints**


* **Success Response:**

    * **Code:** 200 <br />
      **Content:**
```
[
    {
        "CategoryId": 1,
        "Name": "Jobs"
    },
    {
        "CategoryId": 2,
        "Name": "Property"
    },
    {
        "CategoryId": 3,
        "Name": "For Sale"
    }
]
```

## Create Job post

Creates a new job posting

* **URL**

  /post/category/1

* **Method:**

  `POST`

* **URL Params**

  **Required:**

  `None`

* **Data Params** (Post request)

  `SellerId`, `CategoryId`, `SubCategoryId`, `Title`, `Description`


* **Constraints**
  `CategoryId`  
  `SellerId`  
   `SubCategoryId`  
  `Pay`  
  `Type`  
  `Location`  
  `Place`  should exist in the DB


* **Success Response:**

    * **Code:** 200 <br />
      **Content:**
```
[]
```
* **Sample input**
```
 {
        "Post": {
            "SellerId": 1,
            "CategoryId": 1,
            "SubcategoryId": 1,
            "Title": "test7",
            "Description": "test7desc"
            },
        "Job": {
            "Salary": 500,
            "Pay": "monthly",
            "Type": "fulltime",
            "Location": "remote",
            "Place": "Gainesville"
            }
    }
```
## Login

User can login using username and password 

* **URL**

  /login

* **Method:**

  `POST`

* **URL Params**

   **Required:**

   `None`

* **Data Params**

  None

* **Success Response:**

    * **Code:** 200 <br />

* **Sample input**
```
{
    "Email":"test2@gmail.com",
    "Password":"test"    
}
```
## Signup

User can login using username and password

* **URL**

  /signup

* **Method:**

  `POST`

* **URL Params**

  **Required:**

  `None`

* **Data Params**

  None

* **Success Response:**

    * **Code:** 200 <br />

* **Sample input**
```
{
    "Contact": {
    "FirstName":"Test2",
    "LastName":"User2",
    "Email":"test2@gmail.com",
    "Password":"test"
    }
}
```
     
## Home

Takes the user to the home page after successfully logging in

* **URL**

  /home

* **Method:**

  `GET`

*  **URL Params**

   **Required:**

   `None`

* **Data Params**

  None

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
```
Welcome User1
```
## Logout

User will be logged out 

* **URL**

  /logout

* **Method:**

  `POST`

*  **URL Params**

   **Required:**

   `None`

* **Data Params**

  None

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 
```
Old cookie deleted. Logged out!
```

## Refresh

New token will be generated 

* **URL**

  /refresh

* **Method:**

  `POST`

*  **URL Params**

   **Required:**

   `None`

* **Data Params**

  None

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** 

