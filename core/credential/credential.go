package credential

type SecretKeyCredential struct {
	PathUrl   string
	AppId     string
	SecretKey string
}

func NewAccessKeyCredential(pathUrl, appId, secretKey string) *SecretKeyCredential {
	return &SecretKeyCredential{
		PathUrl:   pathUrl,
		AppId:     appId,
		SecretKey: secretKey,
	}
}
