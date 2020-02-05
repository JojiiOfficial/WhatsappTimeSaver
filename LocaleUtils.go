package main

import "golang.org/x/text/language"

func strToTag(stri, mode string) (language.Tag, string) {
	switch stri {
	case "en", "english", "englisch":
		{
			return language.English, "Translate " + mode + " english"
		}
	case "de", "deutsch", "german":
		{
			return language.German, "Translate " + mode + " german"
		}
	case "es", "spanish", "spanisch":
		{
			return language.Spanish, "Translate " + mode + " spanish"
		}
	case "pl", "polish", "polnisch":
		{
			return language.Polish, "Translate " + mode + " polish"
		}
	case "nl", "dutch", "niederlaendisch", "niederländisch":
		{
			return language.Dutch, "Translate" + mode + " dutch"
		}
	case "zu", "zulu":
		{
			return language.Zulu, "Translate" + mode + " zulu"
		}
	case "ru", "russian", "russisch":
		{
			return language.Russian, "Translate " + mode + " russian"
		}
	}
	return language.Tag{}, ""
}
