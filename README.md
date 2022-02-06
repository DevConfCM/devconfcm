# devconfcm
Website for DevConfCM (devconf.cm)

### How to run the code

#### Build

`git clone https://github.com/DevConfCM/devconfcm.git` 

#### Build frontend code
`cd devconfcm/frontend` 
`npm run build` 

#### Run
`cd devconfcm` 
`go run main.go`

**Note**: Because of how the code is setup, when you make changes to the frontend you need to run `npm run build` from within the frontend directory
then restart the server. During development, if you need need to use react development tools like livereload, then run `npm start` instead and mock api calls.
Once your future is ready, you can make it work with actual api calls by doing `npm run build` and running the application.
