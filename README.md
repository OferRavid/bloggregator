# Gator 📰

Gator🐊 is a simple blog aggregator written in Go. It's a CLI tool that allows users to:
- Add RSS feeds from across the internet to be collected.
- Store the collected posts in a PostgreSQL database.
- Follow and unfollow RSS feeds that other users have added.
- View summaries of the aggregated posts in the terminal, with a link to the full post.

It is designed with modularity and extensibility in mind.

## Features ✨

- 🔐 User signup and user management

- ➕ Add and manage RSS/Atom blog feeds

- 📖 Unified feed reader

- 🔁 Periodic background feed refresh

- 🐘 PostgreSQL database support (with potential for other DBs)

## Getting Started

### Prerequisites

- 🧰 Go 1.22+ <img src="imgs/gopher.png" alt="go-pher" width="25"/>
- 🐘 PostgreSQL v15+ (or 🗃️ SQLite3 if using SQLite)

## Installation

### 1. Install Go 1.22 or later
_If you already have `Go` 1.22+ installed, skip to the next step_

There are two options:

**Option 1**: [The webi installer](https://webinstall.dev/golang/) is the simplest way for most people. Just run this in your terminal:

```bash
curl -sS https://webi.sh/golang | sh
```

_Read the output of the command and follow any instructions._

**Option 2**: Use the [official installation instructions](https://go.dev/doc/install).

Run `go version` on your command line to make sure the installation worked. If it did, _move on to step 2_.

**Optional troubleshooting:**

- If you already had Go installed with webi, you should be able to run the same webi command to update it.
- If you already had a version of Go installed a different way, you can use `which go` to find out where it is installed, and remove the old version manually.
- If you're getting a "command not found" error after installation, it's most likely because the directory containing the `go` program isn't in your [`PATH`](https://opensource.com/article/17/6/set-path-linux). You need to add the directory to your `PATH` by modifying your shell's configuration file. First, you need to know _where_ the `go` command was installed. It might be in:

- `~/.local/opt/go/bin` (webi)
- `/usr/local/go/bin` (official installation)
- Somewhere else?

You can ensure it exists by attempting to run `go` using its full filepath. For example, if you think it's in `~/.local/opt/go/bin`, you can run `~/.local/opt/go/bin/go version`. If that works, then you just need to add `~/.local/opt/go/bin` to your `PATH` and reload your shell:

```bash
# For Linux/WSL
echo 'export PATH=$PATH:$HOME/.local/opt/go/bin' >> ~/.bashrc
# next, reload your shell configuration
source ~/.bashrc
```

```bash
# For Mac OS
echo 'export PATH=$PATH:$HOME/.local/opt/go/bin' >> ~/.zshrc
# next, reload your shell configuration
source ~/.zshrc
```

### 2. Setup PostgreSQL

#### 2.1. Installing PostgreSQL v15+:
_If you already have Postgres v15+ installed, skip to 2.2_

**macOS with brew**
   
```bash
brew install postgresql@15
```
   
**Linux / WSL (Debian).** Here are the [docs from Microsoft](https://learn.microsoft.com/en-us/windows/wsl/tutorials/wsl-database#install-postgresql), but simply:
   
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```
Ensure the installation worked. The psql command-line utility is the default client for Postgres. Use it to make sure you're on version 15+ of Postgres:

```bash
psql --version
```

#### 2.2. (Linux only) Update postgres password:

```bash
sudo passwd postgres
```
Enter a password, and be sure you won't forget it. You can just use something easy like postgres.

#### 2.3. Start the Postgres server in the background
- Mac: `brew services start postgresql@15`
- Linux: `sudo service postgresql start`

#### 2.4. Connect to the server.
Enter the psql shell:

- Mac: `psql postgres`
- Linux: `sudo -u postgres psql`

You should see a new prompt that looks like this:
```bash
postgres=#
```

#### 2.5.Create a new database:
```postgreSQL
CREATE DATABASE gator;
```

#### 2.6. Connect to the new database:
```postgreSQL
\c gator
```

You should see a new prompt that looks like this:
```bash
gator=#
```

#### 2.7. Set the user password (Linux only)
```postgreSQL
ALTER USER postgres PASSWORD 'postgres';
```

_For simplicity, I used postgres as the password. Before, we altered the system user's password, now we're altering the database user's password._

Type `exit` to leave the `psql` shell.

### 3. Install gator

#### 3.1. Clone the repository:
```bash
git clone https://github.com/OferRavid/gator.git
cd gator
```

#### 3.2. Install dependencies:
```bash
go mod tidy
```

#### 3.3. Create gator executable
```bash
go install
```
_The executable is installed in the directory named by the GOBIN environment variable, which defaults to $GOPATH/bin or $HOME/go/bin if the GOPATH environment variable is not set. Which should be included in the `PATH`._   
You can now use `gator` as a CLI command

### 4. User management

Gator is a multi-user CLI application. There's no server (other than the database), so it's only intended for local use, but just like games in the 90's and early 2000's, that doesn't mean we can't have multiplayer functionality on a single device!

To manage users we use a single JSON file to keep track of two things:

1. Who is currently logged in.
2. The connection credentials for the PostgreSQL database.

The JSON file has this structure (when prettified):

```bash
{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}
```
_Note: There's no user-based authentication for this app. If someone has the database credentials, they can act as any user._

#### 4.1. Manually create a config file in your home directory, `~/.gatorconfig.json`, with the following content:

- macOS:
```bash
{
  "db_url": "postgres://your_username:@localhost:5432/gator?sslmode=disable"
}
```
- Linux / WSL (Debian):
```bash
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
}
```
_Don't worry about adding current_user_name, that will be set by the application._

The gator CLI is now set.

#### Here are a few commands you can call with `gator`:
`gator register <username>` - registers a new user by adding it to the database.

`gator login <username>` - sets the current user in the config file.

`gator follow <feed_url>` - adds a feed_follow for current user based on the url provided.

`gator agg <time_between_reqs>` - starts a background feed refresh on given time_between_reqs interval.

---

## Purpose of Project
- Demonstrate how to integrate a Go application with a PostgreSQL database.
- Practice using your SQL skills to query and migrate a database (using sqlc and goose, two lightweight tools for typesafe SQL in Go).
- Demonstrate how to write a long-running service that continuously fetches new posts from RSS feeds and stores them in the database.

## License 📄

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---
Created by **Ofer Ravid**