# ChAMP Backend Assignment-ToDoList
## How to set up project
1. fill .env
```
  - DB_USER
  - DB_PASSWORD
  - DB_NAME
```
2. ```docker compose up ```
3. ```go run main.go```

## API Documentation
Task [/api/tasks]
- Create a task [POST]
  body {
    "Description" : ,
    "Duedate" : ,
    "Position" : ,
    "ListID" : 
  }
- Update a task [PATCH  /:id]
```
  body {
    "Description" : ,
    "Duedate" : 
  }
 ```
- Delete a task [DELETE  /:id]
- Move a task to another list [PATCH  movetonewlist/:id #task id]
 ```
 body {
    "list_id" : #new list id
  }
  ```
- Reorder a task in a list [PATCH  reorder/:id #task id]
  ```
  body {
	"new_position" :  
  }
  ```
List [/api/lists]
- Create a list [POST]
- Update a list [PATCH  /:id]
  ```
  body {
    "Title" :
  }
  ```
- Reorder a list [PATCH  reorder/:id #list id]
  ```
  body {
	"new_position" :  
  }```
- Delete a list, also every tasks in it [DELETE  /:id]

