package ls

func isHiddenLinux(filename string) (bool, error) {
	if filename[0:1] == "." {
		return true, nil
	} else {
		return false, nil
	}
}
