package ascii

/*Map Colonizer populates a map with the ascii-art characters according to their decimal key-value pairs*/

func MapColonizer(asciiChrs map[int][]string, rows []string) map[int][]string {
	dec := 31 // Sets to decimal value just before 'sp'
	for _, stringElement := range rows {
		if stringElement == "" { // Separates map key by empty string element in slice
			dec++
		} else {
			asciiChrs[dec] = append(asciiChrs[dec], stringElement) // Appends key by decimal value to value of each string element needed for ascii art
		}
	}
	return asciiChrs
}
