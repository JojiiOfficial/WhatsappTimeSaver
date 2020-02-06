package main

import "golang.org/x/text/language"

func strToTag(stri, mode string) (language.Tag, string) {
	lang, err := language.Parse(stri)
	if err != nil {
		return language.Tag{}, ""
	}
	return lang, "Translate " + mode + " " + lang.String()
}
