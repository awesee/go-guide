package array

func ArrayPush(s []string, args ...string) {

	for _, v := range args {
		s = append(s, v)
	}

}
