package main

import (
	"os/user"
	"path"
	"runtime"

	"code.google.com/p/getopt"
)

const PROGNAME string = "Fireman"
const PROGVERSION string = "0.2.0"

func main() {
	u, err := user.Current()
	if err != nil {
		f("Error: Couldn't get current user - %s", err.Error())
	}
	home := u.HomeDir
	system := runtime.GOOS
	if system == "darwin" {
		home = path.Join(home, "Library", "Preferences")
	}
	if system == "freebsd" || system == "linux" {
		home = path.Join(home, "."+BASE)
	}
	cfgname := path.Join(home, CONFIGNAME)

	cfgfile = getopt.StringLong("config", 'C', cfgname, "Name of the configuration file to use.")
	version := getopt.BoolLong("version", 'v', "Show version and config path, then exit.")

	allusers := getopt.BoolLong("userlist", 'l', "List all users.")
	completeusers := getopt.BoolLong("complete", 'c', "List all users with e-mail set.")
	incompleteusers := getopt.BoolLong("incomplete", 'i', "List all users missing e-mail.")

	add := getopt.StringLong("adduser", 'a', "", "Name of user to add.")
	edit := getopt.StringLong("edituser", 'e', "", "User to edit.")
	del := getopt.StringLong("deluser", 'd', "", "Name of user to delete.")
	mail := getopt.StringLong("email", 'm', "", "E-mail address of the specified user.")
	pw := getopt.StringLong("password", 'p', "", "Password of the specified user.")
	name := getopt.StringLong("name", 'n', "", "Full name of the specified user.")

	initConfig()

	if *version {
		p("%s v%s", PROGNAME, PROGVERSION)
		p("Default configuration file is %s", cfgname)
		return
	}

	if *allusers {
		getUsers(0)
		return
	}
	if *completeusers {
		getUsers(1)
		return
	}
	if *incompleteusers {
		getUsers(2)
		return
	}

	if *add != "" {
		if *pw == "" {
			*pw = genPassword(16)
		}
		createUser(*add, *name, *mail, *pw)
		return
	}

	if *del != "" {
		deleteUser(*del)
		return
	}

	if *edit != "" {
		if *name == "" || *mail == "" && *pw == "" {
			p("Do what with %s?", *edit)
			return
		}
		setDetails(*edit, *name, *mail, *pw)
		return
	}
	getopt.Usage()
}
