Small REST API for GoLang + GinGonic testing. 

Intent: 
- Mock db, with initialization test values. // Testing phase
- Read a certain collection on the db.
- Authorised users group. // Ok
- Add authorized user. // Ok. But since routes are generated in the beginning, and users live in memory, they aren't added
- Upload files with auth. 
- Send message to another with auth. 
- Get unread messages // have to handle login first? how to get user from request
- Group routes. // Ok
- Implement tests


-----------------
Improvements:
- NewAuthorizedUser not being registered on the route for Auth
- Apply generics or some fix to the GetCollection & GetMapCollection methods on repository, they're the same