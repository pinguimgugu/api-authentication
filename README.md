# Simple api to authenticate user find by credentials in multiples repositories using golang concurrency
A basic example how to use golang concurrency with a simple mechanism of authentication using multiple repositories to find user by credentials with clean code and easyly grow up

Make sure already have been installed 
 - Docker engine
 - Docker compose

Start application
  # build and runnig all stack 
  ./application serve

  # get all dependency application 
  ./application dep

  # run application 
  ./application run

  # build base image with golang and dependencies manager dep
  ./application build

# Just make this requisition to authenticate user
curl -X POST \
  http://localhost:7006/authenticate/user/ \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 874f52e0-5bb9-4c55-8cf8-9dd3789c963e' \
  -d '{
    "user":"example",
    "password":"example"  
  }'