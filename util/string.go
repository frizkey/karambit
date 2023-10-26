package util

func CoallesceString(s ...string) string {
	for _, v := range s {
		if v != "" {
			return v
		}
	}
	return ""
}
