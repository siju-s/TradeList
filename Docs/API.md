<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Get Posts](#get-posts)
- [Create Post](#create-post)
- [Fetch Categories](#fetch-categories)

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

## Create Post

Creates a new post

* **URL**

  /post

* **Method:**

  `PUT`

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
  {
  "SellerId": 1,
  "CategoryId": 1,
  "SubcategoryId": 1,
  "Title": "Post1",
  "Description": "Post1Desc"
  }
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

