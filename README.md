# Go Midterm
After completing my finals for school, I thought why not do a midterm for go, like I did for my classes? For that code (done in C++ btw) can be found [here](https://github.com/smmmfrd/school_projects_spring_2026/tree/main)

## 1. Read and Write Files

Read from a file and list the average, mean, and highest occurence of the numbers found in it. (To help out I also wrote something that generates this file.)

## 2. Fetch URLs

Use goroutines to fetch a bunch of URLs and output their status codes

## 3. Mini Redis

Create a mini redis program that just uses memory to do it's stuff

## 4. JSON Loader and Validator

Read JSON into a struct from a file and ensure each key has a value

For the JSON files:
- port should be between 1 and 65535
- max_connections should be a positive integer
- level should only be one of "info", "debug", "warn", "error"
- host and url shouldn't be empty strings
- Maybe timeout has a minimum value

### Continuation Ideas

- [ ] Read in files from folder
- [ ] Pass in a something to do the evaluations of the key values. (making this generic by requiring the business logic from the user)
- [ ] Alphabetically sort both of them (by level obviously) to allow user to have their json files in whatever order.

## 5. Rate limiter

Make a simple HTTP server then implement a rate limiter