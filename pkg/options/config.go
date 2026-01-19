package options

type Config struct {
	Port          int      `json:"port"`
	IdLog         string   `json:"idLog"`
	PatternsWords []string `json:"pattenersWords"`
}
