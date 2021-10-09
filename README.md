# Appointy Technical Task
This repo marks for the Appointy Technical Task using Golang

- [X] Create an User
- [X] Get a user using id
- [X] Create a Post
- [X] Get a post using id
- [X] List all posts of a user

### Create an User
The user can create account using the API by going to '/users'. Supports only POST requests.

### Get a user using id
Returns user details using the API by going to ‘/users/{id}’. Supports only GET requests. Throws error if the user does not exist.

### Create a Post
The user can create posts using the API by going to '/posts'. Supports only POST requests. Throws an error if the user with given ID does not exists.

### Get a post using id
Returns post content using the API by going to ‘/posts/{id}’. Supports only GET requests. Throws an error if the post with given ID does not exists.

### List all posts of a user
Returns array containing all posts of a user using the API by going to ‘/posts/users/{id}’. Supports only GET requests. Throws an error if the user with given ID does not exists.