package api

import (
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/wangjun861205/nblogger"
	"github.com/wangjun861205/nborm"
	context "golang.org/x/net/context"
)

var UserManager *nborm.Manager

type Server struct {
	logger nblogger.Logger
}

func InitManager(mysqlAddr string) error {
	m, err := nborm.NewManager(mysqlAddr, &User{})
	if err != nil {
		return err
	}
	UserManager = m
	return nil
}

func NewServer(logPath string) (*Server, error) {
	logger, err := nblogger.NewLogger(logPath, "bkauth", log.Llongfile|log.Ltime, 1*time.Minute, 200*nblogger.M, 5)
	if err != nil {
		return nil, err
	}
	return &Server{logger}, nil
}

func (s *Server) SignUp(ctx context.Context, req *SignUpRequest) (*SignUpResponse, error) {
	user, err := UserManager.NewModel(
		nborm.ArgMap{
			"Username": req.Username,
			"Password": req.Password,
			"Phone":    req.Phone,
			"Email":    req.Email,
			"Status":   0})

	if err != nil {
		s.logger.Log(err)
		return nil, err
	}
	err = UserManager.InsertOne(user)
	if err != nil {
		s.logger.Log(err)
		return nil, err
	}
	return &SignUpResponse{Status: 0, Error: ""}, nil
}

func (s *Server) SignIn(ctx context.Context, req *SignInRequest) (*SignInResponse, error) {
	where := fmt.Sprintf("(@Username=%q or @Phone=%q or @Email=%q) and @Password=%q",
		req.Ident, req.Ident, req.Ident, fmt.Sprintf("%x", md5.Sum([]byte(req.Password))))
	u, err := UserManager.QueryOne(where)
	if err != nil {
		s.logger.Log(err)
		return nil, err
	}
	user := u.(*User)
	sessionID := fmt.Sprintf("%x", sha256.Sum256([]byte(user.Username.String()+user.Password.String())))
	expireTime := time.Now().Add(time.Hour * 24 * 30)
	err = user.SessionID.Set(sessionID)
	if err != nil {
		s.logger.Log(err)
		return nil, err
	}
	err = user.ExpireTime.Set(expireTime)
	if err != nil {
		s.logger.Log(err)
		return nil, err
	}
	err = UserManager.UpdateOne(user)
	if err != nil {
		s.logger.Log(err)
		return nil, err
	}
	return &SignInResponse{Status: 0, Error: ""}, nil
}

func (s *Server) AuthSessionID(ctx context.Context, req *AuthSessionIDRequest) (*AuthSessionIDResponse, error) {
	u, err := UserManager.QueryOne(fmt.Sprintf("@SessionID=%q", req.SessionID))
	if err != nil {
		s.logger.Log(err)
		return nil, err
	}
	user := u.(*User)
	expireTime, ok := user.ExpireTime.ToTime()
	if !ok {
		err = errors.New("not valid expired time")
		s.logger.Log(err)
		return nil, err
	}
	if expireTime.Before(time.Now()) {
		return &AuthSessionIDResponse{Status: 1, Error: "expired"}, nil
	}
	return &AuthSessionIDResponse{Status: 0, Error: ""}, nil
}
