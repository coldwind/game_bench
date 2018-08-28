package main

import (
	"log"
	"io/ioutil"
	"strings"
)

func readToken() ([]byte, error) {

	fileBody, err := ioutil.ReadFile("etc/token_list")

	if (err != nil) {
		log.Println(err)
		return []byte{}, err
	}

	return fileBody, nil
}

func tokenToArray(token []byte) ([]string, error) {
	tokenString := string(token)
	tokenData := strings.Split(tokenString, "\n")

	return tokenData, nil
}

func tokenList() ([]string, error) {
	t, err := readToken()

	if err != nil {
		return []string{}, err
	}

	tokenList, err := tokenToArray(t)

	if err != nil {
		return []string{}, err
	}

	return tokenList, nil
}