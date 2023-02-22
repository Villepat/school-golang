# TODO:

- SEND FROM BACKEND json responses
  - (Viktor) I feel like it's mostly done, but unsure if it's implemented everywhere.
- Create endpoint to GET posts from .db  
- Mutex locks
  - (Viktor) Where? Why?
- Are cookies even being set correctly
- (Korin will attempt to) Handle json responses in frontend
  - (Viktor) It's done for the login, registration and post creation forms.
  If needed in more places refernce how it's done there and try to keep doing it in a similar fashion.
- create master parent struct with info that is needed on every page: login status, username, 
- Random logout when refreshing session? when getusername?
- newline not respected when post or comment is created
- db = Global CONST?
- consistent var naming
- Fixing comments CSS? splitting every comment to a box?



# In Progress: 


# Done:
- Prevent blank page after registration submission
- Provide feedback to the user when one tries to register with an existing username
- Create endpoint to save post  (12/12/22) (Adrian, Viktor)
- Display posts from .db on frontpage (12/12/22) (Adrian)
- Input validation for create post (empty title, limit lenght ETC) (ville)
- Like/dislike button
- Login form validation  
  - (Viktor) Seems to be working as the backend gives a response of wrong credentials if that's the case.