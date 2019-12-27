package main

import (
	"regexp"
)

func parse(input string) string {
	parsedId := ""
	linkRegExp, _ := regexp.Compile(`(http(s)?:\/\/)?((w){3}.)?youtu(be|.be)?(\.com)?\/.+`)
	videoIdRegExp2, _ := regexp.Compile(`^.*(youtu\.be\/|v\/|u\/\w\/|embed\/|watch\?v=|\&v=)([^#\&\?]*).*`)

	parserLink := linkRegExp.FindStringSubmatch(input)
	if (len(parserLink)>0) {
		//fmt.Println("link: "+parserLink[0])
		parserdId := videoIdRegExp2.FindStringSubmatch(parserLink[0])
		if (len(parserdId)>0) {
			//fmt.Println("id: "+parserdId[2])
			parsedId = parserdId[2]
		}
	}

	return parsedId
}
