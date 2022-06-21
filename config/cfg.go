package config

type Cfg interface {
	Apply(rs interface{}) error
}
