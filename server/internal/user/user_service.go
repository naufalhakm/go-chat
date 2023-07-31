package user

import (
	"context"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/naufalhakm/go-chat/util"
	"github.com/naufalhakm/go-chat/util/token"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

const (
	SecretKey = "secret"
)

func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	// hash Password
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	res := &CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}

	return res, nil
}

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &LoginUserRes{}, err
	}

	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &LoginUserRes{}, err
	}

	token, err := token.GenerateToken(int(u.ID))
	if err != nil {
		return nil, err
	}

	// token := jwt.NewWithClaims(jwt.SigningMethodES256, MyJWTClaims{
	// 	ID:       strconv.Itoa(int(u.ID)),
	// 	Username: u.Username,
	// 	RegisteredClaims: jwt.RegisteredClaims{
	// 		Issuer:    strconv.Itoa(int(u.ID)),
	// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	// 	},
	// })

	// ss, err := token.SignedString([]byte("secret"))
	// if err != nil {
	// 	return &LoginUserRes{}, err
	// }

	return &LoginUserRes{
		accessToken: token,
		ID:          strconv.Itoa(int(u.ID)),
		Username:    u.Username,
	}, nil
}
