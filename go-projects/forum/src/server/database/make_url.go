package database

func MakeURL(title string, result string) string {
	if title == "" {
		return result
	}
	if title[0:1] == "\"" || title[0:1] == "€" || title[0:1] == "#" {
		return MakeURL(title[1:], result+"(spec)")
	}
	if title[0:1] == "?" {
		return MakeURL(title[1:], result+"(question_mark)")
	}
	if title[0:1] == " " {
		return MakeURL(title[1:], result+"&")
	}
	return MakeURL(title[1:], result+title[0:1])
}
