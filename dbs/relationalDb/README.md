# Relational Databases

Leverage Microsoft SQL Server and Azure SQL to store example data.

## Install Go libs

```powershell
go install github.com/microsoft/go-mssqldb@latest
go get github.com/microsoft/go-mssqldb

go get github.com/joho/godotenv/cmd/godotenv
go install github.com/joho/godotenv/cmd/godotenv@latest
```

## Pubs

Remember the old school pubs database for book sellers? I sure do and I've brought it up to 'snuff with some more modern techniques. 

### Set Environment Variables

```powershell
# Set Environment Variables via PowerShell
[Environment]::SetEnvironmentVariable('PubsSqlServerName', '(local)')
[Environment]::SetEnvironmentVariable('PubsSqlDatabaseName', 'Pubs')

# Retrieve Environment Variables via PowerShell
[Environment]::GetEnvironmentVariable('PubsSqlServerName')
[Environment]::GetEnvironmentVariable('PubsSqlDatabaseName')

```
