# jwtx
用于实现jwt的常用方法，替代以前的`github.com/dgrijalva/jwt-go`依赖

### Sign(payload interface{})  
签发token ，payload使用结构体指针

###  Payload(tokenString string, payload interface{}) (err error) 
解析token指针验证其正确性，并将内容解析到payload中