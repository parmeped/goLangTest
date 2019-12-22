Small REST API for GoLang + GinGonic testing. 

Intent: 
- Mock db, with initialization test values. // Ok
- Authorised users group. // Ok
- Add authorized user. // Ok
- Send message to another with auth. 
- Get unread messages // Ok. Marked as read when seen!
- Group routes with Auth. // Ok
- Login // Ok
- PseudoUserSession // Ok
- Implement tests


-----------------
Improvements:
- Add panic handlers when accessing resources
- Implement hash for passwords
- Check duplicate users!


-----------------
Design:
- DB should have a collection of users, which would have their own messages.