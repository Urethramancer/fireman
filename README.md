# Fireman
This is a simple little Go tool I made to manage users on my Openfire server. It's derived from a server tool I made to reset passwords,
but will do a bit more when finished.

## Supported systems
Any Unix-like system which Go runs on is going to run well. The only ones recognised at the moment are FreeBSD, Linux and OS X.
Windows is most likely going to work if somebody writes the setup code (initialising paths for config, mainly), and other
operating systems just need a keyword to test for.


## Features
- List all users
- List complete users (no vital details missing)
- List incomplete users (currently users without e-mail set)
- Set full name, e-mail or password for individual users
- Add/remove individual users
- Add/remove users in bulk, with password generation

### TODO
- E-mail the owners of new accounts when bulk creating

## Building
Packages required:
- code.google.com/p/gcfg
- code.google.com/p/getopt

Currently developed with Go 1.4, so don't bet on it staying compatible with older versions.

## Configuration
The config file is a simple INI-style file with one section, containing two entries. If it's missing it will be created on first run.
To see the location it uses, run Fireman like this:
```sh
fireman -v
```

Default config:
```ini
[main]
server = http://localhost:9090/plugins/userService/
key = <key>
```

Fireman is meant to run on the same server that Openfire is installed on, but you can add other IP addresses to its whitelist if you
need to run Fireman from other locations.

## Usage
Running it without arguments or with the -h flag will show all available flags.

List all users:
```sh
fireman -l
```

List only "complete" users (users with e-mail set):
```sh
fireman -c
```

List users with missing e-mail:
```sh
fireman -i
```
E-mail isn't required by dedault when registering users with Openfire, so this lets you filter those out in case you want to do
anything to them later.

Set e-mail:
```sh
fireman -u <username> -m <e-mail>
```

Set password:
```sh
fireman -u <username> -p <password>
```

All user-edit flags can also be combined to set everything at once:
```sh
fireman -u <username> -n <full name> -m <e-mail> -p <password>
```

Add a user (all fields except username are optional):
```sh
fireman -a <username> -n <full name> -m <e-mail> -p <password>
```

Delete a user:
```sh
fireman -d <username>
```

## Bulk operations
To add many users at once from a list, first make a comma-separated file with username, full name and e-mail.
Example:
```
one,One,one@example.com
two,Two,two@example.com
three,Three,three@example.com
```

Then run the program like this:
```sh
fireman -A <filename>
```

To delete users, simple make a list of usernames, one per line. Then run:
```sh
fireman -D <filename>
```
