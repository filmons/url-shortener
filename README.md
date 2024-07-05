# URL Shorter
***
## TEAM MEMBERS

- [Persi MANKITA](https://github.com/persi-man)
- [Filmon SEARE](https://github.com/filmons)
- [Nour MIBANDIKIDI](https://github.com/Nour-Mibandikidi)
- [Sylvianne FEUPE](https://github.com/riade2014)
***
## Table of Contents
- [Description](#description)
- [Installation](#installation)
  - [Prerequisites](#prerequisites)
  - [Backend](#backend)
- [Usage](#usage)

## Description
This is a simple URL shorter project. It is a web application that allows users to shorten URLs. It is built with Golang and Vue.js. The backend is built with Golang and the frontend is built with React.js. The backend is a RESTful API and the frontend is a single page application. The backend provides the following features:
- Shorten a URL
- Get a URL by its short code
- Get a URL by its original URL
- Get all URLs
- Delete a URL by its short code
- Delete a URL by its original URL
- Delete all URLs
- Redirect to the original URL by its short code
- Redirect to the original URL by its original URL
- Get the statistics of a URL by its short code
- Get the statistics of a URL by its original URL
- Get the statistics of all URLs
- Get the statistics of a URL by its short code in a specific time range
- Get the statistics of a URL by its original URL in a specific time range
- Get the statistics of all URLs in a specific time range

The frontend provides the following features:
- Shorten a URL
- Get a URL by its short code
- Get a URL by its original URL
- Get all URLs
- Delete a URL by its short code
- Delete a URL by its original URL
- Delete all URLs

The backend uses the following technologies:
- Golang
- Gin

The frontend uses the following technologies:
- React.js


## Installation
### Prerequisites
- Golang
- Node.js
- npm

### Backend
1. Clone the repository
```bash
go get https://github.com/filmons/url-shortener
```

2. run the project
```bash
./run.sh
```
or 
```bash
bash ./scripts/start.sh
```

3. Run the frontend
```bash
cd url-shortener-frontend
npm install
npm start
```