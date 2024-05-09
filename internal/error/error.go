package error

type NewError string

func (err NewError) Error() string {
	return string(err)
}
