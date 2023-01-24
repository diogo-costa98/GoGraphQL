# GoGraphQL

## SQL File 
The sql file included creates 3 tables: users, questions and options.

It populates these tables as:
- Users: creates 3 users (these have usernames: user, anotheruser and finaluser, and encrypted passwords)
- Questions: creates 7 questions (3 for one user, 4 for another and 5 for the final)
- Options: creates between 2 and 4 options for each of the 7 questions

## User Credentials
1.  Username: user        | Password: password
2.  Username: anotheruser | Password: anotherpassword
3.  Username: finaluser   | Password: finalpassword

## Login
```graphql
#Login Mutation
mutation {
  login(input: {username: "anotheruser", password: "anotherpassword"})
}
```
### JSON Response
```json
{
  "data": {
    "login": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFub3RoZXJ1c2VyIn0.FtTnj1HGMP1Tc6EntpUropaxBdMJKhwlQKSFKcMoD-c"
  }
}
```

## Get Questions

```graphql
#Questions Query
{
  questions{
    id
    body
    options {
      id
      body
      correct
    }
  }
}
```
JSON Header
```json
{
  "Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImZpbmFsdXNlciJ9.zTvn_BH2Yi6aiB95YUJZ4jLUXFnmKI00C7rd9fFQ9lg"
}
```
<sub>*Note: the default page size is 10, and the default page is 0*</sub>



### JSON Response
```json
{
  "data": {
    "questions": [
      {
        "id": "8",
        "body": "How much is 3+1?",
        "options": [
          {
            "id": "23",
            "body": "3+1=4",
            "correct": true
          },
          {
            "id": "24",
            "body": "3+1=5",
            "correct": false
          }
        ]
      },
      {
        "id": "9",
        "body": "How much is 3+2?",
        "options": [
          {
            "id": "25",
            "body": "3+2=3",
            "correct": false
          },
          {
            "id": "26",
            "body": "3+2=2",
            "correct": false
          },
          {
            "id": "27",
            "body": "3+2=5",
            "correct": true
          },
          {
            "id": "28",
            "body": "3+2=4",
            "correct": false
          }
        ]
      },
      {
        "id": "10",
        "body": "How much is 3+3?",
        "options": [
          {
            "id": "29",
            "body": "3+3=5",
            "correct": false
          },
          {
            "id": "30",
            "body": "3+3=6",
            "correct": true
          }
        ]
      },
      {
        "id": "11",
        "body": "How much is 3+4?",
        "options": [
          {
            "id": "31",
            "body": "3+4=5",
            "correct": false
          },
          {
            "id": "32",
            "body": "3+4=6",
            "correct": false
          },
          {
            "id": "33",
            "body": "3+4=7",
            "correct": true
          }
        ]
      },
      {
        "id": "12",
        "body": "How much is 3+5?",
        "options": [
          {
            "id": "34",
            "body": "3+5=8",
            "correct": true
          },
          {
            "id": "35",
            "body": "3+5=9",
            "correct": false
          }
        ]
      }
    ]
  }
}
```

## Get Questions (with pagination)

```graphql
#Questions Query
{
  questions(page:"1",pageSize:"2"){
    id
    body
    options {
      id
      body
      correct
    }
  }
}
```
JSON Header
```json
{
  "Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImZpbmFsdXNlciJ9.zTvn_BH2Yi6aiB95YUJZ4jLUXFnmKI00C7rd9fFQ9lg"
}
```



### JSON Response
```json
{
  "data": {
    "questions": [
      {
        "id": "10",
        "body": "How much is 3+3?",
        "options": [
          {
            "id": "29",
            "body": "3+3=5",
            "correct": false
          },
          {
            "id": "30",
            "body": "3+3=6",
            "correct": true
          }
        ]
      },
      {
        "id": "11",
        "body": "How much is 3+4?",
        "options": [
          {
            "id": "31",
            "body": "3+4=5",
            "correct": false
          },
          {
            "id": "32",
            "body": "3+4=6",
            "correct": false
          },
          {
            "id": "33",
            "body": "3+4=7",
            "correct": true
          }
        ]
      }
    ]
  }
}
```

## Delete Question

```graphql
#Delete Question Mutation
mutation { 
	deleteQuestion(id: 7){
    id
    body
  }
}
```
### JSON Response
```json
{
  "data": {
    "deleteQuestion": {
      "id": "7",
      "body": "How much is 2+4?",
      "options": []
    }
  }
}
```

## Update Question

```graphql
#Update Question Mutation
mutation {
  updateQuestion(
    id: 5
    input: {body: "How much is 20+3?", options: [{body: "20+3=4", correct: false}, {body: "20+3=5", correct: false}, {body: "20+3=23", correct: true}]}
  ) {
    id
    body
    options {
      id
      body
      correct
    }
  }
}
```
### JSON Response
```json
{
  "data": {
    "updateQuestion": {
      "id": "5",
      "body": "How much is 20+3?",
      "options": [
        {
          "id": "36",
          "body": "20+3=4",
          "correct": false
        },
        {
          "id": "37",
          "body": "20+3=5",
          "correct": false
        },
        {
          "id": "38",
          "body": "20+3=23",
          "correct": true
        }
      ]
    }
  }
}
```
