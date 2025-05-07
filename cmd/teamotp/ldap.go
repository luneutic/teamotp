package main

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
)

func LdapAuth(username, password string) bool {
	ldapURL := fmt.Sprintf("ldap://%s:389", config.LdapHost)
	conn, err := ldap.DialURL(ldapURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	upn := fmt.Sprintf("%s@%s", username, config.LdapDomain)
	err = conn.Bind(upn, password)
	if err != nil {
		log.Print(err)
		return false
	}

	return true
}
