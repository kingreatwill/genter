package util

// TryCatch implements try...catch... logistics.
func TryCatch(try func(), catch ...func(exception interface{})) {
	if len(catch) > 0 {
		// If <catch> is given, it's used to handle the exception.
		defer func() {
			if e := recover(); e != nil {
				catch[0](e)
			}
		}()
	} else {
		// If no <catch> function passed, it filters the exception.
		defer func() {
			recover()
		}()
	}
	try()
}
