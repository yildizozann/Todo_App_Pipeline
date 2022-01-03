#  To Do App Api

This is a to-do application. 
Adding and listing todos

Test Driven Development
Consumer and Provider Testing
Unit Testing


**Server: Golang  
Client: React 
Database: Local MongoDB**

---

# Highlights

1. DB connection string, name and collection name moved to `.env` file as environment variable. Using `github.com/joho/godotenv` to read the environment variables.
2. 
## Application Requirement


# Highlights

1. First docker-compose-up for the mongodb database
2. After that, start todoapi with run go main.go 
3. For the front-end using npm start command
## Application Requirement

### golang server requirement

1. golang https://golang.org/dl/
2. goland fiber for https get and post methods `go get -u github.com/gofiber/fiber/v2`
3. mongo-driver package to connect with mongoDB `go get go.mongodb.org/mongo-driver`
4. github.com/joho/godotenv to access the environment variable.

### react client

From the Application directory

`create-react-app client`

## :computer: Start the application

1. Make sure your mongoDB is started
2. From docker-compose up
3. From todoapi directory, open a terminal and run
   `go run main.go`
4. From client directory,  
   a. install all the dependencies using `npm install`  
   b. start client `npm start`


### Create To-Do

Enter a To-Do input box and click submit.








