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
- Set e-mail or password for individual users

### TODO
- Update other user details than e-mail/password
- Add/remove users
- Bulk create/edit/delete users from files
- Automatic passwords for bulk users
- E-mail the owners of new accounts when bulk creating

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
fireman -u <user> -m <e-mail>
```

Set password:
```sh
fireman -u <user> -p <password>
```

All user-edit flags can also be combined to set everything at once:
```sh
fireman -u <user> -m <e-mail> -p <password>
```
