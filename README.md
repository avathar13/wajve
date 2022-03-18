# Go Programming Challenge

Hello 👋

## Instructions

* You *shall* complete this challenge using **Go language**,
* You *should* try to show your development process to present a **production-ready** code,
* Please, describe your approach, your technical choices, including architectural, and anything you want us to know in a results.md file,

## The Scenario

Your mission is to implement a single endpoint that will be integrated in a larger **microservice** architecture.

* the endpoint is exposed at `/api/v1/trivia`
* the endpoint allows to fetch data from a local database (see explainations below)
* the query may takes parameters, to filter the results. You are expected to propose any kind of query params you may find useful to provide a rich user experience,
* the endpoint returns a JSON response with an array of matching results

## The database

The initial database is provided in a file [db.json](db.json).
You are expected to integrate the data in a mongodb local instance, and explain how to install, launch and populate the database. This database **must** be used by your code.
The schema of the provided data is :

```
{
  "type": "array",
  "items": {
    "type": "object"
    {
      "text": {
        "type": "string",
        "minLength": 1
      },
      "number": {
        "type": "number",
	     "minimum": 1,
	     "maximum": 1e+150
      },
      "found": {
        "type": "boolean"
      },
      "type": {
        "type": "string"
      }
    }
  }
}
```

Example:

```
[
  {
    "text": "1e+40 is the Eddington–Dirac number.",
    "number": 1e+40,
    "found": true,
    "type": "trivia"
  },
  {
    "text": "60 is the years of marriage until the diamond wedding anniversary.",
    "number": 60,
    "found": true,
    "type": "trivia"
  }
]
```
