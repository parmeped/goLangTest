Small REST API for GoLang + GinGonic testing. 

Intent: 
- Mock db, with initialization test values. // Testing phase
- Read a certain collection on the db.
- Authorised users group. // Ok
- Add authorized user. // Ok. But since routes are generated in the beginning, and users live in memory, they aren't added
- Upload files. 
- Send message to another user. 
- Group routes. // Ok
- Implement tests


-----------------
Improvements:
- NewAuthorizedUser not being registered on the route for Auth