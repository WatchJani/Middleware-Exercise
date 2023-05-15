package examples

type Return func() string

func Returner(name string) Return {
	return func() string {
		return name
	}
}
