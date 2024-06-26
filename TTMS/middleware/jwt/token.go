package jwt

import (
	"TTMS/conf"
	"TTMS/model"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

//TestSignTokenFunction(Id string) string：生成一个测试用的JWT签名token。该函数使用了预设的密钥（config.Secret）和用户ID来生成token，并返回生成的token字符串。
//
//SignToken(user model.User) string：生成JWT签名token。该函数接受一个用户对象作为参数，使用预设的密钥（config.Secret）和用户ID生成token，并返回生成的token字符串。
//
//TestParseToken(tokenString string) (interface{}, error)：解析并验证JWT签名token。该函数接受一个token字符串作为参数，使用预设的密钥（config.Secret）进行解析和验证，并返回解析后的token对象和可能的错误。
//
//ParseToken(tokenString string) (interface{}, error)：解析并验证JWT签名token，返回其中的用户ID。该函数接受一个token字符串作为参数，使用预设的密钥（config.Secret）进行解析和验证，并返回解析后的token中的用户ID和可能的错误。
//
//ParseTokenTime(tokenString string) (any, error)：解析并验证JWT签名token，返回其中的签发日期。该函数接受一个token字符串作为参数，使用预设的密钥（config.Secret）进行解析和验证，并返回解析后的token中的签发日期和可能的错误。

// 提取密钥
var jwtSecret = []byte(conf.Secret)

// 测试签名token
func TestSignTokenFunction(Id string) string {
	// test

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// sub (subject): 主题
		"sub": "用户Id来签发token",
		//
		"foo": "bar",
		//nbf (Not Before): 生效时间
		//"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"nbf": time.Now().Unix(),
	})
	// 提取密钥
	var jwtSecret = []byte(conf.Secret)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtSecret)
	// 如果签名失败
	if err != nil {
		return ""
	}
	fmt.Println(tokenString, err)
	return tokenString
}

// 签名token
func SignToken(user *model.Customer) string {
	// dev
	Id := user.ID
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// iss (issuer): 签发人
		"iss": "weiweiwait",
		// sub (subject): 主题
		"sub": "用户Id",
		// Id
		"id": Id,
		//nbf (Not Before): 生效时间
		"nbf": time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtSecret)
	// 如果签名失败
	if err != nil {
		return "签名失败"
	}
	fmt.Println(tokenString, err)
	return tokenString
}

// 签名token
func SignTokenManager(user *model.Manager) string {
	// dev
	Id := user.ID
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// iss (issuer): 签发人
		"iss": "weiweiwait",
		// sub (subject): 主题
		"sub": "用户Id",
		// Id
		"id": Id,
		//nbf (Not Before): 生效时间
		"nbf": time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtSecret)
	// 如果签名失败
	if err != nil {
		return "签名失败"
	}
	fmt.Println(tokenString, err)
	return tokenString
}

// 测试解析token
func TestParseToken(tokenString string) (interface{}, error) {
	// sample token string taken from the New example
	//tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
		return false, err
	}
	return true, err
}

//  解析token 返回用户Id

func ParseToken(tokenString string) (interface{}, error) {

	// sample token string taken from the New example
	//tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//fmt.Println(claims["foo"], claims["nbf"])
		return claims["id"], nil
	} else {
		fmt.Println(err)
		return nil, err
	}
	//return true, err
}

// 返回token中的签发日期

func ParseTokenTime(tokenString string) (any, error) {
	// sample token string taken from the New example
	//tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//fmt.Println(claims["foo"], claims["nbf"])
		return claims["nbf"], nil
	} else {
		fmt.Println(err)
		return nil, err
	}
}
