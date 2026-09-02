package types

type BaiduDeviceRes struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationUrl string `json:"verification_url"`
	Qrcodeurl       string `json:"qrcode_url"`
	ExpiresIn       int    `json:"expires_in"`
	Interval        int    `json:"interval"`
}

type BaiduTokenRes struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type BaiduAuthError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
