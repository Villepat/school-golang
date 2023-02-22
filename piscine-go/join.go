package piscine

func Join(strs []string, sep string) string {
	nyastr := ""
	for i := 0; i < len(strs); i++ { // if we are handling the last element of strs we don't want sep to be added
		if i+1 == len(strs) {
			nyastr += strs[i]
		} else {
			nyastr += strs[i] + sep
		}
	}
	return nyastr
}
