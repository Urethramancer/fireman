package main

import (
	"os/user"
	"path"
	"runtime"

	"code.google.com/p/getopt"
)

const PROGNAME string = "Fireman"
const PROGVERSION string = "0.1.0"

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
	version := getopt.BoolLong("version", 'v', "Show version and exit.")
	allusers := getopt.BoolLong("userlist", 'l', "List all users.")
	completeusers := getopt.BoolLong("complete", 'c', "List all users with e-mail set.")
	incompleteusers := getopt.BoolLong("incomplete", 'i', "List all users missing e-mail.")
	user := getopt.StringLong("user", 'u', "", "User to edit.")
	mail := getopt.StringLong("email", 'm', "", "E-mail address of specified user.")
	pw := getopt.StringLong("password", 'p', "", "Password of specified user.")
	initConfig()

	if *version {
		p("%s v%s", PROGNAME, PROGVERSION)
		return
	}

	if *allusers {
		getUsersOF(0)
		return
	}
	if *completeusers {
		getUsersOF(1)
		return
	}
	if *incompleteusers {
		getUsersOF(2)
		return
	}
	if *user != "" {
		if *mail == "" && *pw == "" {
			p("Do what with %s?", *user)
			return
		}
		if *mail != "" && *pw != "" {
			setDetailsOF(*user, *mail, *pw)
			return
		}
		if *mail != "" {
			setDetailsOF(*user, *mail, "")
			return
		}
		if *pw != "" {
			setDetailsOF(*user, "", *pw)
		}
		return
	}
	getopt.Usage()
}
