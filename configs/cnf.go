package configs

type Cnf interface {
	Get(k string) interface{}
	GetStringSlice(k string) []string
	GetString(k string) string
	GetInt(k string) int
	GetBool(k string) bool
}
