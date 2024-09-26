package util

import "os"

// externalizar o segredo. Usar AWS Secrets Manager?
func GetJwtSecret() ([]byte, error) {
	return []byte(os.Getenv("SECRET")), nil
}
