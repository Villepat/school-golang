
# FORUM

This project is an online discussion board where people can discuss topics of mutual interest, 
in the form of posted messages and interact with the posts. 

The forum is divided into categories (by tags) for the relevant discussions. 

Guests of the forum are granted access to all functions that do not require database alterations. To create a post 
or react to it, one has to register to become a member and be logged in. 

Posts(user-submitted message) are enclosed into a block containing the user's details and the date and time it was submitted. 

Thread(or topic) is a collection of posts. In this forum, members can reply to posts and the replies are consolidated
in a logical reply structure of chronological order.



## USAGE
Run go run . to start server
## FEATURES

- Communication between users. Only registered users will be able to create posts and comments.
- Associating categories to posts. When registered users are creating a post they can associate one or more categories to it
- Liking and disliking posts and comments. (Only registered users will be able to)
- Filtering posts (categories, created posts, liked posts). Number of views are displayed. 

 
## Tech  

**Client** Javascript, CSS, HTML

**Server:** Golang, SQLite3



## AUTHORS

- Adrian, Johannes, Korin, Viktor , Ville  


## Lessons Learned
- HTML
- HTTP
- Sessions and cookies
- SQL language , manipulation of databases
- The basics of encryption

Using and setting up Docker
- Containerizing an application
- Compatibility/Dependency
- Creating images
## Documentation

[Golang Standard Library](https://pkg.go.dev/std)

[SQLite](https://www.sqlite.org/docs.html)

