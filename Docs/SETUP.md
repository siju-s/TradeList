##SETUP
<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Frontend](#frontend)
- [Backend](#backend)
- [Troubleshooting](#troubleshooting)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->
## Frontend

1. Install [Node](https://nodejs.org/en/download/)
2. Install `Angular-CLI` using
```
npm install -g @angular/cli
```
3. Create workspace using 

```
ng new frontend
```
4. Add `Angular material`
```
ng add @angular/material
```
5. Add `Angular devkit`
```
npm install --save-dev @angular-devkit/build-angular
```
6. Add `Angular Flex-Layout`
```
npm i @angular/flex-layout
```
7. To launch the server, navigate to `frontend` folder and run 
```
npm start
```
8. Navigate to [http://localhost:4200](http://localhost:4200) to view the application

## Backend

1. Install [Go](https://go.dev/dl/)
2. To listen to requests, go to backend folder and run

```
go run .
```
3. Navigate to [http://localhost:8081](http://localhost:8081) to check
4. Install [Air](https://github.com/cosmtrek/air) for live reload 
```
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```
5. Run `air` on terminal to start the watcher
6. To get JWT for Golang, run the following command:

```
go get github.com/golang-jwt/jwt
```
7. To get bcrypt for password hashing in Golang, run the following command:

```
go get golang.org/x/crypto/bcrypt
```

## Generic
1. DocToc - Table of contents generation. Install using `npm install -g doctoc`

Then, run `doctoc filename.md` to generate table of contents.

## Troubleshooting

1. Permissions denied error on Mac

```
sudo chown -R $(whoami) /usr/local/lib/node_modules
```
