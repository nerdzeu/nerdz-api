TODO
====

+ Implement the messages interfaces for pm -> write tests
+ Add a method for every user action (follow, update post/comment, create things and so on)
+ Obviously write test for every method
* Write doc/CONTRIBUTING.md
+ Use of [Osin](https://github.com/RangelReale/osin) to create OAuth 2 authorization server
+ Create database support (OAuth2 needs to store lots of values)
+ Create HTTP REST API, following some standard (oData maybe?)

And so on...


# What has been done

+ Created types (ORM model)
+ Fetch comments and posts (with related options: from friends only, in a language only and these options can be mixed).
+ Add/Delete/Edit comment/post
+ Rereiving user information (numeric (fast) or complete)
+ ...

Contributed to the [gorm](https://github.com/jinzhu/gorm/) project several times:

- [Add support for primary key different from id](https://github.com/jinzhu/gorm/pull/85)
- [Add support to fields with double quotes](https://github.com/jinzhu/gorm/pull/105)
- [Add default values support](https://github.com/jinzhu/gorm/pull/279)
