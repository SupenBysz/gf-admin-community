package kyFuns

func If[R any](condition bool, trueVal, falseVal R) R {
	if condition {
		return trueVal
	} else {
		return falseVal
	}
}
