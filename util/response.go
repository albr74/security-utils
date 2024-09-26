package util

import (
	"net/http"

	"github.com/albr74/commons/util"
)

func InvalidEmailResponse() *util.ErrorResponse {
	return &util.ErrorResponse{Code: "SEC-001",
		Message:  "Invalid email",
		HttpCode: http.StatusBadRequest,
	}
}

func InvalidPasswordResponse() *util.ErrorResponse {
	return &util.ErrorResponse{Code: "SEC-002",
		Message:  "Invalid password",
		HttpCode: http.StatusBadRequest,
	}
}

func TenantBlockedResponse() *util.ErrorResponse {
	return &util.ErrorResponse{Code: "SEC-003",
		Message:  "Tenant is blocked",
		HttpCode: http.StatusPreconditionFailed,
	}
}

func TenantWithoutAdminUserResponse() *util.ErrorResponse {
	return &util.ErrorResponse{Code: "SEC-004",
		Message:  "Tenant without admin user",
		HttpCode: http.StatusPreconditionFailed,
	}
}

func UserWithoutPermissionResponse() *util.ErrorResponse {
	return &util.ErrorResponse{Code: "SEC-005",
		Message:  "User without permission",
		HttpCode: http.StatusPreconditionFailed,
	}
}

func ForbiddenResponse() *util.ErrorResponse {
	return &util.ErrorResponse{Code: "SEC-006",
		Message:  "Forbidden",
		HttpCode: http.StatusForbidden,
	}
}

func UserNotFoundResponse() *util.ErrorResponse {
	return &util.ErrorResponse{Code: "SEC-0007",
		Message:  "User not found",
		HttpCode: http.StatusNotFound,
	}
}

func UserWithoutTenantHomeResponse() *util.ErrorResponse {
	return &util.ErrorResponse{Code: "SEC-008",
		Message:  "User without tenant home",
		HttpCode: http.StatusPreconditionFailed,
	}
}

func TokenExpiredResponse() *util.ErrorResponse {
	return &util.ErrorResponse{Code: "SEC-009",
		Message:  "Token expired",
		HttpCode: http.StatusUnauthorized,
	}
}

func TokenAlreadyRevokedResponse() *util.ErrorResponse {
	return &util.ErrorResponse{Code: "SEC-010",
		Message:  "Token already revogated",
		HttpCode: http.StatusPreconditionFailed,
	}
}

func TokenRevokedResponse() *util.ErrorResponse {
	return &util.ErrorResponse{Code: "SEC-011",
		Message:  "Token is revoked",
		HttpCode: http.StatusPreconditionFailed,
	}
}

func InsecurePasswordResponse() *util.ErrorResponse {
	return &util.ErrorResponse{Code: "SEC-012",
		Message:  "Insecure password",
		HttpCode: http.StatusPreconditionFailed,
	}
}

func InvalidUsernamePasswordResponse() *util.ErrorResponse {
	return &util.ErrorResponse{Code: "SEC-013",
		Message:  "Invalid username/password",
		HttpCode: http.StatusPreconditionFailed,
	}
}

func UserAlreadyExistsResponse() *util.ErrorResponse {
	return &util.ErrorResponse{Code: "SEC-014",
		Message:  "User already exists",
		HttpCode: http.StatusPreconditionFailed,
	}
}

func TenantAlreadyExistsResponse() *util.ErrorResponse {
	return &util.ErrorResponse{Code: "SEC-015",
		Message:  "Tenant already exists",
		HttpCode: http.StatusPreconditionFailed,
	}
}

func TokenNotFoundResponse() *util.ErrorResponse {
	return &util.ErrorResponse{Code: "SEC-016",
		Message:  "Token not found",
		HttpCode: http.StatusPreconditionFailed,
	}
}

func TenantNotFoundResponse() *util.ErrorResponse {
	return &util.ErrorResponse{Code: "SEC-017",
		Message:  "Tenant not found",
		HttpCode: http.StatusPreconditionFailed,
	}
}

func UsernameMandatoryResponse() *util.ErrorResponse {
	return &util.ErrorResponse{Code: "SEC-018",
		Message:  "Username is mandatory",
		HttpCode: http.StatusBadRequest,
	}
}
