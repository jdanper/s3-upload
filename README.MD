# Audit Event Sink Service
This service inserts all audit events into a database

### Test
With `make test` all tests available will run.

### Build
Just run `make build` to build the binary. This step will run the tests before building. 

### Run
After build step, run the application with `./audit-sink`. Make sure the environment variables are set.

### Environment
`ENV` defaults to `debug` - options: `debug, release

`LOG_LEVEL` defaults to `1` - options: `1 (DEBUG), 2 (INFO), 3 (WARNING), 4 (ERROR)`

`HTTP_PORT` defaults to `8080`

`DATABASE_HOST` defaults to ` `

`DATABASE_PORT` defaults to `20097`

`DATABASE_USERNAME` defaults to ` `

`DATABASE_PASSWORD` defaults to ` `