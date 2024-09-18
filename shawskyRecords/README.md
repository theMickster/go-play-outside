# Shawsky Records

A simple api demonstrating leveraging Azure Functions + custom handlers + Go

## Steps to build + deploy to Azure Fx running on Windows host

1. Open a Terminal Window and set the location to the api folder
`cd shawskyrecords\api`

1. Build the main package changing the output directory 
`go build -o ../function/main.exe main.go` We do this so that the `host.json` deployment file can locate the GO package

1. Double-check that the host.json file is configured to release the Windows executable to the proper .exe
```json
"description": {
    "defaultExecutablePath": "main.exe",
    "arguments": []
}
```

4. Open a Terminal Window and set the location to the azure function folder
`cd shawskyrecords\function`

1. Run the GO + GIN api locally via azure function tool chain
`func start`

1. Use Postman to send requests to the application

1. In to the Azure Extension in VS Code, select the workspace for the Local Project, ensure that the functions are correct, hover over the Local Project root, and click the Deploy to Azure button to the right-hand side of the option