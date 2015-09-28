# NERDZ API TODO list


The API development follows a pretty well defined logic of implementation that will be briefly described here.
In the package *api*, there will be a file for each entity (e.g., users.go). This source file will contain all the operations
associated to that kind of entity that will be exposed by the API. 

For each operations' source file, there **must** be a related
test source file, which will contain all the unit tests for each implemented operation.

### Implement GET operation on the following entities:
   1. Users (/users)
   2. Projects (/projects)
   3. Comments (UserComments and ProjectComments) (/users/:id/comments or /projects/:id/comments)
   4. Pms (/pms)

### Implement POST/PUT operation on the following entities:
   1. Users (/users)
   2. Projects (/projects)
   3. Comments (UserComments and ProjectComments) (/users/:id/comments or /projects/:id/comments)
   4. Pms (/pms)

### Implement DELETE operation on the following entities:
   1. Users (/users)
   2. Projects (/projects)
   3. Comments (UserComments and ProjectComments) (/users/:id/comments or /projects/:id/comments)
   4. Pms (/pms)

### Supplementary features

+ Write doc/CONTRIBUTING.md specifying the operations' request format and their responses
+ Use of [Osin](https://github.com/RangelReale/osin) to create OAuth 2 authorization server
+ Create database support for OAuth 2 (OAuth2 needs to store lots of values)

And so on...


## What has been done

+ Created types (ORM model)
+ Fetch comments and posts (with related options: from friends only, in a language only and these options can be mixed).
+ Add/Delete/Edit comment/post
+ Rereiving user information (numeric (fast) or complete)
+ Implement the messages interfaces for pm -> write tests
+ Add a method for every user action (follow, update post/comment, create things and so on)
+ Tests for every method
+ ...

Contributed to the [gorm](https://github.com/jinzhu/gorm/) project several times:

- [Add support for primary key different from id](https://github.com/jinzhu/gorm/pull/85)
- [Add support to fields with double quotes](https://github.com/jinzhu/gorm/pull/105)
- [Add default values support](https://github.com/jinzhu/gorm/pull/279)
