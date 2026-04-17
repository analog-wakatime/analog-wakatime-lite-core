package config

func GetAppPort() string {
	return ":8080"
}

func ConfigGetDatabaseURL() string {
	DATABASE_URL := "postgres://resume_dev_user:resume_password@localhost:7732/resume_db?sslmode=disable"
	return DATABASE_URL
}

func ConfigGetJWTSecret() string {
	JWT_SECRET := "my_super_secret_key"
	return JWT_SECRET
}

func GetRedisURL() string {
	REDIS_URL := "localhost:6380"
	return REDIS_URL
}
