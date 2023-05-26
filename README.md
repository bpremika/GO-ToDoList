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
**Task** ```/api/tasks```
- ```GET``` Get all tasks
- ```POST``` Create a task

  ```
  body {
    "Description" : ,
    "Duedate" : ,
    "Position" : ,
    "ListID" : 
  }
  ```
- ```PATCH  /:id``` Update a task

  ```
  body {
    "Description" : ,
    "Duedate" : 
  }
  ```
- ```DELETE  /:id``` Delete a task
- ```PATCH  /movetonewlist/:id (task id)``` Move a task to another list

  ```
  body {
    "list_id" : #new list id
  }
  ```
- ```PATCH  /reorder/:id (task id)``` Reorder a task in a list

  ```
  body {
	"new_position" :  
  }
  ```

**List** ```/api/lists```
- ```GET``` Get all lists
- ```POST``` Create a list 

  ```
  body {
    "Title" : ,
    "Position" : ,
  }
  ```
- ```PATCH  /:id``` Update a list 

  ```
  body {
    "Title" :
  }
  ```
- ```PATCH  /reorder/:id (list id)``` Reorder a list 

  ```
  body {
	"new_position" :  
  }
  ```
- ```DELETE  /:id``` Delete a list, also every tasks in it 

