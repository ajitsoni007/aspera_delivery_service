package config

type Config struct {
	AWSAccessKeyID     string `env:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
	AWSSessionToken    string `env:"AWS_SESSION_TOKEN"`
}

// func LoadConfig() (*Config, error) {
// 	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
// 	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
// 	sessionToken := os.Getenv("AWS_SESSION_TOKEN")

// 	// Return an error if any of the environment variables are missing
// 	if accessKeyID == "" || secretAccessKey == "" || sessionToken == "" {
// 		return nil, fmt.Errorf("missing required environment variable(s)")
// 	}

// 	// Create a new Config struct and return it
// 	return &Config{
// 		AWSAccessKeyID:     accessKeyID,
// 		AWSSecretAccessKey: secretAccessKey,
// 		AWSSessionToken:    sessionToken,
// 	}, nil
// }
