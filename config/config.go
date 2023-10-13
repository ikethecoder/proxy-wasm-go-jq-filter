package config

// -----------------------------------------------------------------------------
// Instance Config
// -----------------------------------------------------------------------------
//go:generate ffjson -noencoder $GOFILE

// Config represents the filter configuration
type Config struct {
	Query string
}

// Load json config from data into conf
func Load(data []byte, conf *Config) error {
	conf.Query = "."

	// err := ffjson.Unmarshal(data, conf)
	// if err != nil {
	// 	return err
	// }

	return nil
}
