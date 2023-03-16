package helper

func CheckIfError(err error) error {
	if err != nil {
		return err
	}
	return nil
}
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
