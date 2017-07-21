package snscommon

func ExecUntilSuccess(f func() (interface{}, string)) (res interface{}) {
	for true {
		res, ok = f()
		if ok {
			break
		}
	}
	return
}
