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

### Configuration variables
The `app.env` file houses all non-sensitive configuration variables.

* `PORT`: Port our application will be listening to.
* `HOST`: Host for our server.
* `DB_URL`: Url for our database.

For sensitive variables such as `USERNAME`, `PASSWORD` e.t.c, put them in your system's evironment variables.

`go run main.go`

**Note**: Because of how the code is setup, when you make changes to the frontend you need to run `npm run build` from within the frontend directory
then restart the server. During development, if you need need to use react development tools like livereload, then run `npm start` instead and mock api calls.
Once your feature is ready, you can make it work with actual api calls by doing `npm run build` and running the application.
