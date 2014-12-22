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

// Sets the password of an existing user
func setDetailsOF(user string, email string, pw string) {
	if user == "" {
		return
	}

	var buf bytes.Buffer
	buf.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>")
	buf.WriteString("<user>")
	buf.WriteString("<username>")
	buf.WriteString(user)
	buf.WriteString("</username>")
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
	buf.WriteString("</user>")

	req, _ := http.NewRequest("PUT", cfg.Main.Server+"users"+"/"+user, bytes.NewReader(buf.Bytes()))
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
		p("Successfully changed details for %s", user)
	}
}

// types = 0 for all, 1 for complete, 2 for the ones missing e-mail
func getUsersOF(types int) {
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
