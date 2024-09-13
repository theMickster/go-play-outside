package settings

import "os"

type AuthSettings struct {
	Scope    string
	Audience string
	TenantId string
}

type ApplicationSettings struct {
	ApplicationId string
}

func RetrieveAuthSettings() AuthSettings {
	return AuthSettings{
		Scope:    os.Getenv("ShawskyRecordsScope"),
		Audience: os.Getenv("ShawskyRecordsAudience"),
		TenantId: os.Getenv("ShawskyTenantId"),
	}
}

func RetrieveApplicationSettings() ApplicationSettings {
	return ApplicationSettings{
		ApplicationId: os.Getenv("ShawskyRecordsApplicationId"),
	}
}

func RetrieveApplicationPort() string {
	port := ":8081"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = ":" + val
	}
	return port
}
