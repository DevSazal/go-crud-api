# CRUD API >> Go Fiber v2

### Requirements:

* Go
* Fiber [v2.x](https://docs.gofiber.io/)
* Git
* MongoDB


Let's clone the repository on your machine.

🎁 The application includes the following files and folders.

- `app` - code for the application written in [Go](https://go.dev/).
- `.env.example` - a sample of .env which can be helpful for configuration.

```bash
# architecture
# deep drive in app directory

app/
├── handlers/
│   └── user.handler.go
├── models/
│   └── user.go
├── repositories/
│   └── user.repository.go
├── services/
│   └── user.service.go
└── utils/
    └── response.go

```


## Installation and Configuration

Let's move to the cloned directory with your terminal.


Let's rename from `.env.example` to `.env` and make sure all the necessary information is correct:

```bash
PORT=3000

MONGO_URI=''

```

Already done? Cool! You are almost ready to enjoy the app. ⛳️


### Be Ready/Install:
```
go mod tidy
```

### Run:
```bash
go run main.go
```

### API Endpoint:

```bash

POST /users
GET /users
Get /users/:id
PUT /users/:id
DELETE /users/:id

# try the api with postman
# port 3000
```


#### 🎯 I know, you liked it.

#### 🧑‍💻 Stay in touch

- Author - [Sazal Ahamed](https://sazal.vercel.app)

#### tada! 🎉





