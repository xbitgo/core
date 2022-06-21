package config

type Apollo struct {
	Namespace string
	URL       string
}

func (a Apollo) Apply(rs interface{}) error {
	//panic("implement me")
	return nil
}
