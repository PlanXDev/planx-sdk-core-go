# PlanXSDK-Go
Here is the go version of PlanXSDK.

## First step:
> * Go to the **[PlanX CP](https://cp.planx.io)** to register.
> * Get **AppId** and **SecretKey**.
## Second step:
> * Download the package into your project via the command:
```
    go get github.com/PlanXDev/planx-sdk-core-go@v1.3.0
```
## Example Usage
> * This example get balance information for all token accounts:
```go
    baseUrl := "https://cp-api.planx.io"
    appId := "Go to the PlanX CP website to get"
    secretKey := "Go to the PlanX CP website to get"
    client, _ := api.NewClientWithOptions(core.NewConfig(), credential.NewAccessKeyCredential(baseUrl, appId, secretKey))
    base, err := client.GetAccountBalancesBase(api.CreateGetAccountBalancesBaseRequest())
    if err != nil {
    println(err.Error())
    }
    println(base.String())
    
    // Process your bussiness ...
```

## Other
> * Supports Go 1.17+