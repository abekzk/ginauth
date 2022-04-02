package ginauth

import "strings"

// 参考: https://github.com/go-kit/kit/blob/a073a093d1ee02b920ab78db0fb5600cef24a10e/auth/jwt/transport.go#L78-L85
func extractTokenFromAuthHeader(val string) (token string, ok bool) {
	authHeaderParts := strings.Split(val, " ")
	if len(authHeaderParts) != 2 || !strings.EqualFold(authHeaderParts[0], "bearer") {
		return "", false
	}

	return authHeaderParts[1], true
}
