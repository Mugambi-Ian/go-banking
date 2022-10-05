package service

import (
	"fmt"
	"go-banking/banking-auth/domain"
	"go-banking/banking-auth/dto"
	"go-banking/banking-auth/utils"

	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	Register(dto.RegisterRequest) (*dto.LoginResponse, *utils.AppError)
	Login(dto.LoginRequest) (*dto.LoginResponse, *utils.AppError)
	Verify(urlParams map[string]string) *utils.AppError
	Refresh(request dto.RefreshTokenRequest) (*dto.LoginResponse, *utils.AppError)
}

type DefaultAuthService struct {
	repo            domain.AuthRepository
	rolePermissions domain.RolePermissions
}

func (s DefaultAuthService) Refresh(request dto.RefreshTokenRequest) (*dto.LoginResponse, *utils.AppError) {
	if vErr := request.IsAccessTokenValid(); vErr != nil {
		if vErr.Errors == jwt.ValidationErrorExpired {
			var appErr *utils.AppError
			if appErr = s.repo.RefreshTokenExists(request.RefreshToken); appErr != nil {
				return nil, appErr
			}
			var accessToken string
			if accessToken, appErr = domain.NewAccessTokenFromRefreshToken(request.RefreshToken); appErr != nil {
				return nil, appErr
			}
			return &dto.LoginResponse{AccessToken: accessToken}, nil
		}
		return nil, utils.NewAuthenticationError("invalid token")
	}
	return nil, utils.NewAuthenticationError("cannot generate a new access token until the current one expires")
}

func (s DefaultAuthService) Register(req dto.RegisterRequest) (*dto.LoginResponse, *utils.AppError) {
	err := req.IsValid()
	if err != nil {
		return nil, err
	}

	login, err := req.ToLoginCredentials()
	if err != nil {
		return nil, err
	}

	login, err = s.repo.Register(*login)
	if err != nil {
		return nil, err
	}

	response, err := s.GenerateLoginCredentials(*login)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s DefaultAuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, *utils.AppError) {
	var appErr *utils.AppError
	var login *domain.Login

	if login, appErr = s.repo.FindBy(req.Username, req.Password); appErr != nil {
		return nil, appErr
	}
	response, err := s.GenerateLoginCredentials(*login)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s DefaultAuthService) Verify(urlParams map[string]string) *utils.AppError {
	if jwtToken, err := jwtTokenFromString(urlParams["token"]); err != nil {
		return utils.NewAuthorizationError(err.Error())
	} else {
		if jwtToken.Valid {
			claims := jwtToken.Claims.(*domain.AccessTokenClaims)
			if claims.IsUserRole() {
				if !claims.IsRequestVerifiedWithTokenClaims(urlParams) {
					return utils.NewAuthorizationError("request not verified with the token claims")
				}
			}
			isAuthorized := s.rolePermissions.IsAuthorizedFor(claims.Role, urlParams["routeName"])
			if !isAuthorized {
				return utils.NewAuthorizationError(fmt.Sprintf("%s role is not authorized", claims.Role))
			}
			return nil
		} else {
			return utils.NewAuthorizationError("Invalid token")
		}
	}
}

func jwtTokenFromString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(domain.HMAC_SAMPLE_SECRET), nil
	})
	if err != nil {
		utils.LogError("Error while parsing token: " + err.Error())
		return nil, err
	}
	return token, nil
}

func NewLoginService(repo domain.AuthRepository, permissions domain.RolePermissions) DefaultAuthService {
	return DefaultAuthService{repo, permissions}
}

func (s DefaultAuthService) GenerateLoginCredentials(login domain.Login) (*dto.LoginResponse, *utils.AppError) {

	claims := login.ClaimsForAccessToken()
	authToken := domain.NewAuthToken(claims)

	var refreshToken string
	accessToken, appErr := authToken.NewAccessToken()
	if appErr != nil {
		return nil, appErr
	}

	if refreshToken, appErr = s.repo.GenerateAndSaveRefreshTokenToStore(authToken); appErr != nil {
		return nil, appErr
	}

	return &dto.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
