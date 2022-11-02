package authService


type AuthService struct {
	jwtAccessSecret  string
	jwtRefreshSecret string
}

// func NewAuthService() (*AuthService, error) {
// 	return &AuthService{
// 		jwtAccessSecret:  config.JwtAccessSecret,
// 		jwtRefreshSecret: config.JwtRefreshSecret,
// 	}, nil
// }