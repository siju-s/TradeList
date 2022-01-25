### SETUP

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
5. To launch the server, navigate to `frontend` folder and run 
```
npm start
```
6. Navigate to [http://localhost:4200](http://localhost:4200) to view the application

## Backend

1. Install [Go](https://go.dev/dl/)
2. To listen to requests, go to backend folder and run

```
go run .
```
3. Navigate to [http://localhost:8081](http://localhost:8081) to check

### TROUBLESHOOTING

1. Permissions denied error on Mac

```
sudo chown -R $(whoami) /usr/local/lib/node_modules
```
