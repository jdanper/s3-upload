# S3-Upload
Send files to a AWS S3 Bucket

### Test
With `make test` all tests available will run.

### Build
Just run `make build` to build the binary. This step will run the tests before building. 

### Run
After build step, run the application with `./s3-upload`. Make sure the environment variables are set.

### Arguments

`bucket` - S3 bucket to which the files will be sent

`file` - Path of the file

`region` - S3 bucket region (default "us-east-1")

##### Pass `--help` or `-h` to see available options 

### Environment

`LOG_LEVEL` defaults to `1` - options: `1 (DEBUG), 2 (INFO), 3 (WARNING), 4 (ERROR)`

`AWS_ACCESS_KEY_ID`

`AWS_SECRET_ACCESS_KEY`