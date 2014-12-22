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
- Bulk create/edit/delete users from files
- Automatic passwords for bulk users
- E-mail the owners of new accounts when bulk creating

## Usage
List all users:
```bash
fireman -l
```
