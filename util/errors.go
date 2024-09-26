package util

import (
	"errors"

	"github.com/albr74/commons/util"
)

var (
	UnexpectedError         = errors.New("unexpected error")
	UserNotfound            = errors.New("user not found")
	InvalidUsernamePassword = errors.New("invalid username/password")
	UserAlreadyExists       = errors.New("user already exists")
	TenantAlreadyExists     = errors.New("tenant already exists")
	CountryCodeMandatory    = errors.New("country code is mandatory")
	TaxNumberMandatory      = errors.New("tax number is mandatory")
	CompanyNameMandatory    = errors.New("company name is mandatory")
	PasswordMandatory       = errors.New("password is mandatory")
	EmailMandatory          = errors.New("email is mandatory")
	FirstNameMandatory      = errors.New("first name is mandatory")
	LastNameMandatory       = errors.New("last name is mandatory")
	InvalidEmail            = errors.New("invalid email")
	InvalidPassword         = errors.New("invalid password")
	InsecurePassword        = errors.New("insecure password")
	TokenAlreadyRevoked     = errors.New("token already revoked")
	TokenNotFound           = errors.New("token not found")
	TenantBlocked           = errors.New("tenant is blocked")
	TenantNotFound          = errors.New("tenant not found")
	TenantWithoutAdminUser  = errors.New("tenant without admin user")
	UserWithoutTenantHome   = errors.New("user without tenant home")
	UsernameMandatory       = errors.New("username is mandatory")
)

func ConvertError(err error) *util.ErrorResponse {
	if errors.Is(err, UnexpectedError) {
		return util.UnexpectedErrorResponse()
	}
	if errors.Is(err, UserNotfound) {
		return UserNotFoundResponse()
	}
	if errors.Is(err, InvalidUsernamePassword) {
		return InvalidUsernamePasswordResponse()
	}
	if errors.Is(err, UserAlreadyExists) {
		return UserAlreadyExistsResponse()
	}
	if errors.Is(err, TenantAlreadyExists) {
		return TenantAlreadyExistsResponse()
	}
	if errors.Is(err, CountryCodeMandatory) {
		return util.MandatoryFieldResponse("Country code")
	}
	if errors.Is(err, TaxNumberMandatory) {
		return util.MandatoryFieldResponse("Tax number")
	}
	if errors.Is(err, CompanyNameMandatory) {
		return util.MandatoryFieldResponse("Company name")
	}
	if errors.Is(err, PasswordMandatory) {
		return util.MandatoryFieldResponse("Password")
	}
	if errors.Is(err, EmailMandatory) {
		return util.MandatoryFieldResponse("Email")
	}
	if errors.Is(err, FirstNameMandatory) {
		return util.MandatoryFieldResponse("First name")
	}
	if errors.Is(err, LastNameMandatory) {
		return util.MandatoryFieldResponse("last name")
	}
	if errors.Is(err, InvalidEmail) {
		return InvalidEmailResponse()
	}
	if errors.Is(err, InvalidPassword) {
		return InvalidPasswordResponse()
	}
	if errors.Is(err, InsecurePassword) {
		return InsecurePasswordResponse()
	}
	if errors.Is(err, TokenAlreadyRevoked) {
		return TokenAlreadyRevokedResponse()
	}
	if errors.Is(err, TokenNotFound) {
		return TokenNotFoundResponse()
	}
	if errors.Is(err, TenantBlocked) {
		return TenantBlockedResponse()
	}
	if errors.Is(err, TenantNotFound) {
		return TenantNotFoundResponse()
	}
	if errors.Is(err, TenantBlocked) {
		return TenantBlockedResponse()
	}
	if errors.Is(err, TenantNotFound) {
		return TenantNotFoundResponse()
	}
	if errors.Is(err, TenantWithoutAdminUser) {
		return TenantWithoutAdminUserResponse()
	}
	if errors.Is(err, UserWithoutTenantHome) {
		return UserWithoutTenantHomeResponse()
	}
	if errors.Is(err, UsernameMandatory) {
		return UsernameMandatoryResponse()
	}
	return util.UnexpectedErrorResponse()
}
