# ChinguCentral (WIP)
The home for the Chingu Cohorts. It will host a list of all the cohorts, along with their members, completed projects and forums. Each member will have a profile in which he will be able to display his contact information, it will also show the projects he/she has participated on, a probably in the near future some medals, because who doesn't like medals?

**This is in very early stages of development, a lot of stuff needs to be done. Find a list of critical stuff [here](https://github.com/Oxyrus/ChinguCentral/projects/1)**.

## Setup
You must have Golang v1.6+, Node v6.10+ and PostgreSQL v9.2+.

**At the moment we're still writting the API, React client is on its way (you can collaborate, read [here](https://github.com/Oxyrus/ChinguCentral/projects/1))**.

### Run server

Until we set up a dependency manager for Go, install the following packages manually:

```
$ go get github.com/julienschmidt/httprouter
$ go get github.com/urfave/negroni
$ go get gopkg.in/mgo.v2
$ go get github.com/Oxyrus/ChinguCentral
```

1. Get the application `go get -u github.com/Oxyrus/ChinguCentral`
2. Compile the Go code `cd %GOPATH%/src/github.com/Oxyrus/ChinguCentral && go build`
3. Run the compiled file

### Run client

1. Go client folder `cd client`
2. Install dependecies `npm install`
3. Run client server `npm start`
4. Go to: http://localhost:3000

Enjoy!

## Wireframes
* [Home](https://wireframe.cc/3310As)

## Stack
* API - Golang
* Client - ReactJS
* Database - PostgreSQL

## Roadmap
I have never done a roadmap before, so bear with me.

The idea is to setup the user system (registration/login), the list of cohorts with all the users who are members of them and the projects they have completed.
