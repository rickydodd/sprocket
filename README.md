# sprocket
## Brief
sprocket is an API for a learning management system.

## To-Do
- [x] Create a "Hello, world!" command-line application, to get things moving
- [x] Create a basic `.gitignore` file
- [x] Create Makefile to simplify development setup and building
- [x] Add details on how to setup the environment and build the application
- [x] Create a "Hello, world!" end-point, therefore setting up a basic web server
- [ ] Evaluate a suitable database system and implement it
- [ ] Set up authentication and authorisation for the API
- [ ] Renew the README.md To-Do list with new items, based on how the scope changes

## Installing and Building
### Requirements
- `make`
- x86-64 Linux to execute the `sudo make setup` command*
- `bash` to execute the `sudo make setup` command*

\* If dependencies are setup without the use of `sudo make setup` for your operating system and shell combination, then commands `make tidy` and `make build` will be available, as long as your operating system and shell combination has installed and has access to `make`.

### Installing Dependencies
When inside the project directory, the following will install Go version 1.22.0 to `/usr/local` and add the Go bin directory (`/usr/local/go/bin`) to the `.bashrc` file located at `$HOME/.bashrc`,
```bash
sudo make setup
```

### Building
When inside the project directory, the following will build the binary to `./bin`,
```bash
make build
```

