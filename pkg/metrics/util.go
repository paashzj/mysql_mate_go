package metrics

func promConvMysqlStr2Bool(str string) float64 {
	if str == "Yes" {
		return 1
	}
	return 0
}
