package plugins

import (
	"github.com/thoj/go-ircevent"
	"strings"
)

// Help provides basic usage instructions
func Help(conn *irc.Connection, r string, event *irc.Event) {

	var query string
	var response string

	a := strings.Split(event.Message(), " ")

	if len(a) > 1 {
		query = a[1]
	}

	switch query {
	case "acronym":
		response = "string ; V.E.R.A. -- Virtual Entity of Relevant Acronyms"
	case "dict":
		response = "string ; Queries WordNet, a large lexical database of English"
	case "fu":
		response = "nil or string ; FoaaS"
	case "news":
		response = "nil or string ; Prints a recent article title from random garbage news source"
	case "quote":
		response = "add string to save ; get [id] to fetch quote"
	case "stock":
		response = "string ; Stock price at previous day closing"
	case "trump":
		response = "string ; Tronald Dump"
	case "urban":
		response = "string ; Urban Dictionary"
	case "weather":
		response = "zip code ; Returns the current temperature, weather condition, humidity, wind, 'feels like' temperature, barometric pressure, and visibility"
	case "wiki":
		response = "string ; Wikipedia"
	default:
		response = "Commands are: acronym, drama, dict, fu, news, stock, trump, urban, weather, wiki"
	}

	conn.Privmsg(r, response)
}
