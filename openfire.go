package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type XMLUser struct {
	XMLName  xml.Name `xml:"user"`
	Username string   `xml:"username"`
	Name     string   `xml:"name"`
	Email    string   `xml:"email"`
}

type XMLUsers struct {
	XMLName xml.Name  `xml:"users"`
	User    []XMLUser `xml:"user"`
}

func buildUserXML(buf *bytes.Buffer, username string, name string, email string, pw string) {
	// We're not marshalling to avoid making a bloated payload with empty tags
	buf.WriteString("<user>")
	buf.WriteString("<username>")
	buf.WriteString(username)
	buf.WriteString("</username>")
	if name != "" {
		buf.WriteString("<name>")
		buf.WriteString(name)
		buf.WriteString("</name>")
	}
	if email != "" {
		buf.WriteString("<email>")
		buf.WriteString(email)
		buf.WriteString("</email>")
	}
	if pw != "" {
		buf.WriteString("<password>")
		buf.WriteString(pw)
		buf.WriteString("</password>")
	}
	buf.WriteString("</user>\n")
}

func createUser(username string, name string, email string, pw string, silent bool) {
	var buf bytes.Buffer
	buf.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>\n")
	buildUserXML(&buf, username, name, email, pw)

	req, _ := http.NewRequest("POST", cfg.Main.Server+"users", bytes.NewReader(buf.Bytes()))
	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Authorization", cfg.Main.Key)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		p("Couldn't connect to Openfire server: %s", err.Error())
		return
	}
	defer res.Body.Close()

	if res.StatusCode == 201 {
		if !silent {
			p("Successfully created user %s with password %s", username, pw)
		}
	} else {
		p("Error: %s", res.Status)
	}
}

func deleteUser(username string) {
	req, _ := http.NewRequest("DELETE", cfg.Main.Server+"users/"+username, nil)
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Authorization", cfg.Main.Key)
	p("Delete: %s", req.URL.String())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		p("Couldn't connect to Openfire server: %s", err.Error())
		return
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		p("Successfully deleted user %s.", username)
	} else {
		p("Error: %s", res.Status)
	}
}

// Sets the password of an existing user
func setDetails(username string, name string, email string, pw string) {
	var buf bytes.Buffer
	// We're not marshalling to avoid making a bloated payload with empty tags
	buf.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>")
	buildUserXML(&buf, username, name, email, pw)

	req, _ := http.NewRequest("PUT", cfg.Main.Server+"users"+"/"+username, bytes.NewReader(buf.Bytes()))
	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Authorization", cfg.Main.Key)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		p("Couldn't connect to Openfire server: %s", err.Error())
		return
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		p("Successfully changed details for %s", username)
	}
}

// types = 0 for all, 1 for complete, 2 for the ones missing e-mail
func getUsers(types int) {
	req, _ := http.NewRequest("GET", cfg.Main.Server+"users", nil)
	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Authorization", cfg.Main.Key)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		p("Couldn't connect to Openfire server: %s", err.Error())
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		p("Error requesting userlist from the server.")
		return
	}
	body, _ := ioutil.ReadAll(res.Body)
	var users XMLUsers
	xml.Unmarshal(body, &users)
	for _, e := range users.User {
		n := e.Username + ","
		if e.Name != "" {
			n = e.Username + "," + e.Name
		}
		switch types {
		case 0:
			m := "<missing e-mail>"
			if e.Email != "" {
				m = e.Email
			}
			p("%s,%s", n, m)
		case 1:
			if e.Email != "" {
				p("%s,%s", n, e.Email)
			}
		case 2:
			if e.Email == "" {
				p("%s", n)
			}
		}
	}
}
