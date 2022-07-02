# Create Fiber App

Create Fiber App easily

**Status**: In development 

## Commands

These commands are available.

### `create`

You can create a new fiber project using like that;

**Alias:**: `c`

`create-fiber-app create -n github.com/usernane/myproject`

#### Subcommands

These are subcommands to customize your setup

##### --name

It's a flag for the project name. It's required

**Alias:** "-n"

**Required**: `true`

**Example:** `create-fiber-app create -n github.com/usernane/myproject`

##### --dir

It's a flag to specify the project directory. It will use the current directory if you don't specify a directory.

**Alias:** "-d"

**Required**: `false`

**Example:** `create-fiber-app create -n github.com/usernane/myproject -d /path/to/myproject`

## Todo

- [ ] Creating project templates to create example project files, custom middleware, .etc
- [ ] Adding middleware to the current project (Fiber's itself and custom middleware)
- [ ] Ability to create Fiber project by some project structures
- [ ] Add run command like `create-fiber-app run`