# REST API Guide


**Authentication and authorization of this REST API app is based on JWT tokens.**
```
All request (except /login & /register) shold have Authorization header with following value:
Bearer + [space] + [JWT token]
```

### User
1) Register new user
- ROUTE: **/api/register**
- METHOD: **POST**
- Required parameters:
```
username, password
```
![Mockup for feature A](https://github.com/Maksim1990/Laravel_Postgres_Product_App_and_REST_API/blob/master/public/example/API/API1.PNG)
2) Login and obtain JWT token
- ROUTE: **/api/login**
- METHOD: **POST**
- Required parameters:
```
username, password
```
![Mockup for feature A](https://github.com/Maksim1990/Laravel_Postgres_Product_App_and_REST_API/blob/master/public/example/API/API2.PNG)

3) Get all users
- ROUTE: **/api/users**
- METHOD: **GET**
![Mockup for feature A](https://github.com/Maksim1990/Laravel_Postgres_Product_App_and_REST_API/blob/master/public/example/API/API2.PNG)

4) Get specific
- ROUTE: **/api/users/{id}**
- METHOD: **GET**
![Mockup for feature A](https://github.com/Maksim1990/Laravel_Postgres_Product_App_and_REST_API/blob/master/public/example/API/API2.PNG)

5) Delete user
- ROUTE: **/api/users/{id}**
- METHOD: **DELETE**
![Mockup for feature A](https://github.com/Maksim1990/Laravel_Postgres_Product_App_and_REST_API/blob/master/public/example/API/API2.PNG)


6) Update user
- ROUTE: **/api/users/{id}/update**
- METHOD: **PUT**
![Mockup for feature A](https://github.com/Maksim1990/Laravel_Postgres_Product_App_and_REST_API/blob/master/public/example/API/API2.PNG)


### Posts
7) Create new post
- ROUTE: **/api/users/{id}**
- METHOD: **POST**
![Mockup for feature A](https://github.com/Maksim1990/Laravel_Postgres_Product_App_and_REST_API/blob/master/public/example/API/API2.PNG)


8) Get all posts
- ROUTE: **/api/posts**
- METHOD: **GET**
![Mockup for feature A](https://github.com/Maksim1990/Laravel_Postgres_Product_App_and_REST_API/blob/master/public/example/API/API2.PNG)

9) Get specific post
- ROUTE: **/api/users**
- METHOD: **GET**
![Mockup for feature A](https://github.com/Maksim1990/Laravel_Postgres_Product_App_and_REST_API/blob/master/public/example/API/API2.PNG)

10) Delete user
- ROUTE: **/api/posts/{id}**
- METHOD: **DELETE**
![Mockup for feature A](https://github.com/Maksim1990/Laravel_Postgres_Product_App_and_REST_API/blob/master/public/example/API/API2.PNG)


11) Update posts
- ROUTE: **/api/users/{id}/update**
- METHOD: **PUT**
![Mockup for feature A](https://github.com/Maksim1990/Laravel_Postgres_Product_App_and_REST_API/blob/master/public/example/API/API2.PNG)

