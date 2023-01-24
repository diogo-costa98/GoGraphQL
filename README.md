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
  "Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFub3RoZXJ1c2VyIn0.FtTnj1HGMP1Tc6EntpUropaxBdMJKhwlQKSFKcMoD-c"
}
```
<sub>Note that the default page size is 10, and the default page is 0</sub>



### JSON Response
```json
{
  "data": {
    "questions": [
      {
        "id": "4",
        "body": "How much is 2+1?",
        "options": [
          {
            "id": "10",
            "body": "2+1=1",
            "correct": false
          },
          {
            "id": "11",
            "body": "2+1=2",
            "correct": false
          },
          {
            "id": "12",
            "body": "2+1=3",
            "correct": true
          }
        ]
      },
      {
        "id": "5",
        "body": "How much is 2+2?",
        "options": [
          {
            "id": "14",
            "body": "2+2=2",
            "correct": false
          },
          {
            "id": "15",
            "body": "2+2=3",
            "correct": false
          },
          {
            "id": "13",
            "body": "2+2=1",
            "correct": false
          },
          {
            "id": "16",
            "body": "2+2=4",
            "correct": true
          }
        ]
      },
      {
        "id": "6",
        "body": "How much is 2+3?",
        "options": [
          {
            "id": "17",
            "body": "2+3=4",
            "correct": false
          },
          {
            "id": "18",
            "body": "2+3=5",
            "correct": true
          }
        ]
      },
      {
        "id": "7",
        "body": "How much is 2+4?",
        "options": [
          {
            "id": "19",
            "body": "2+4=4",
            "correct": false
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
