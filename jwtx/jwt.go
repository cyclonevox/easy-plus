package jwtx

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// todo 可以考虑用sonic加速json的序列化和范序列化

// JWTConfig JWT配置
type JWTConfig struct {
	// 签名算法
	Method string `default:"HS256" yaml:"method" validate:"required,oneof=HS256 HS384 HS512"`
	// 密钥
	Key string `yaml:"key" validate:"required"`
	// 统一前缀
	Scheme string `yaml:"scheme"`
	// 有效期，单位分组
	Expiration int `default:"720" yaml:"expiration" validate:"required"`
}

// 额外的Token验证方法
// 例如只允许一个Token有效
type extraKeyFunc func(token *jwt.Token) bool

type JwtTool struct {
	// 签名算法
	method string
	// 签名密钥
	key []byte
	// Token前缀
	scheme string
	// 有效期，单位分钟
	expiration int
	// Token校验方法
	keyFunc jwt.Keyfunc
}

type Claims struct {
	Payload []byte
	*jwt.RegisteredClaims
}

func NewJWT(config JWTConfig, extra ...extraKeyFunc) *JwtTool {
	var validator = jwt.NewValidator()

	return &JwtTool{
		method:     config.Method,
		key:        []byte(config.Key),
		scheme:     config.Scheme,
		expiration: config.Expiration,
		keyFunc: func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != config.Method {
				return nil, jwt.ErrInvalidKey
			}

			var (
				claims *Claims
				ok     bool
			)
			if claims, ok = token.Claims.(*Claims); !ok {
				return nil, jwt.ErrInvalidKey
			}

			if err := validator.Validate(claims); err != nil {
				return nil, err
			}

			if len(extra) != 0 {
				if !extra[0](token) {
					return nil, jwt.ErrInvalidKey
				}
			}

			return []byte(config.Key), nil
		},
	}
}

// Sign 签署token，payload为具体的机构题信息
func (j *JwtTool) Sign(payload interface{}) (string, error) {
	var (
		data []byte
		err  error
	)

	if data, err = json.Marshal(payload); err != nil {
		return "", err
	}

	now := time.Now()

	regClaims := &Claims{
		data,
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(j.expiration) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(j.method), regClaims)

	var tokenString string
	if tokenString, err = token.SignedString(j.key); err != nil {
		return "", err
	}

	if j.scheme != "" {
		tokenString = j.scheme + " " + tokenString
	}

	return tokenString, nil
}

func (j *JwtTool) Parse(tokenString string) (*Claims, map[string]interface{}, error) {
	var (
		err   error
		token *jwt.Token
	)

	if j.scheme != "" && strings.HasPrefix(tokenString, j.scheme) {
		tokenString = tokenString[len(j.scheme)+1:]
	}

	claims := new(Claims)
	if token, err = jwt.ParseWithClaims(tokenString, claims, j.keyFunc); err != nil {
		return nil, nil, err
	}

	return token.Claims.(*Claims), token.Header, nil
}

// Payload 直接解析到结构体指针中。故 payload 请使用结构体的指针
func (j *JwtTool) Payload(tokenString string, payload interface{}) (err error) {
	if j.scheme != "" && strings.HasPrefix(tokenString, j.scheme) {
		tokenString = tokenString[len(j.scheme)+1:]
	}

	claims := new(Claims)
	if _, err = jwt.ParseWithClaims(tokenString, claims, j.keyFunc); err != nil {
		return
	}

	if err = json.Unmarshal(claims.Payload, payload); err != nil {
		return
	}

	return
}
