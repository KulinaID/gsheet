package gsheet

type Option struct {
	key   string
	value string
}

func (o *Option) Get() (key, value string) {
	return o.key, o.value
}
