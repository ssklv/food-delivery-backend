package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/ssklv/food-delivery-backend/internal/domain"
	"github.com/ssklv/food-delivery-backend/internal/repository"
)

//ошибки добавить

type authUsecase struct {
	repository     AuthRepository
	tokenProvider  TokenProvider
	passwordHasher PasswordHasher
}

// func NewAuthUsecase(rep AuthRepository, tokenProvider TokenProvider, passwordHasher PasswordHasher) AuthUsecase {
//  	return &authUsecase{
//  		repository:     rep,
//  		tokenProvider:  tokenProvider,
//  		passwordHasher: passwordHasher,
//  	}
//  }

func (au *authUsecase) ValidateToken(ctx context.Context, tokenString string) (*domain.User, error) {
	userID, err := au.tokenProvider.ParseToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("validateToken: parse: %w", err)
	}

	user, err := au.repository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("validateToken: get user: %w", err)
	}

	return user, nil
}

func (au *authUsecase) Register(ctx context.Context, phone, password, name string) (string, string, error) {
	hashedPassword, err := au.passwordHasher.HashPassword(password)
	if err != nil {
		return "", "", fmt.Errorf("hash password: %w", err)
	}

	user := &domain.User{
		Phone:        phone,
		PasswordHash: hashedPassword,
		Name:         name,
		Role:         "user",
	}

	if err := au.repository.CreateUser(ctx, user); err != nil {
		return "", "", fmt.Errorf("create user: %w", err)
	}

	accessToken, refreshToken, err := au.generateTokenPair(ctx, user)
	if err != nil {
		return "", "", fmt.Errorf("generate tokens: %w", err)
	}

	return accessToken, refreshToken, nil
}

func (au *authUsecase) generateTokenPair(ctx context.Context, user *domain.User) (string, string, error) {
	accessToken, err := au.tokenProvider.GenerateAccessToken(user.ID, string(user.Role))
	if err != nil {
		return "", "", fmt.Errorf("generate access token: %w", err)
	}

	refreshToken, err := au.tokenProvider.GenerateRefreshToken()
	if err != nil {
		return "", "", fmt.Errorf("generate refresh token: %w", err)
	}

	session := &repository.UserSession{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(time.Hour * 24 * 30),
	}

	if err := au.repository.SaveSession(ctx, session); err != nil {
		return "", "", fmt.Errorf("save session: %w", err)
	}

	return accessToken, refreshToken, nil
}
