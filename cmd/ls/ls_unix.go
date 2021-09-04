package ls

// isHiddenUnix checks if the file is hidden in Unix
func isHiddenUnix(filename string) bool {
	if filename[0:1] == "." {
		return true
	} else {
		return false
	}
}
