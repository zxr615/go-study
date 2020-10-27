package customErr

type CustomError struct {
	Name string
	Msg string

}

func (e *CustomError) Error() string  {
	return e.Name + " " + e.Msg + "";
}