package config

func GetAppPort() string {
	return ":8080"
}

func ConfigGetDatabaseURL() string {
	DATABASE_URL := "postgres://analog_waka_time_lite_core_users:analog_waka_time_lite_core_password@localhost:7732/analog_waka_time_lite_core_db?sslmode=disable"
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
