package errno

import (
	"errors"
	"fmt"
)

const (
	WebSocketSuccessCode       = 1000
	WebSocketLogoutSuccessCode = iota + 1000
	WebSocketTargetOnlineSuccessCode
	WebSocketErrCode
	WebSocketConnectErrCode
	WebSocketTargetOfflineErrCode

	SuccessCode    = 10000
	ServiceErrCode = iota + 10000 //未知服务错误
	ParamErrCode                  //参数错误
	CharacterBeyondLimitErrCode

	ExistUserErrCode
	NotExistUserErrCode
	AuthFailedErrCode //Authorization错误
	ReadFileErrCode
	UploadFileErrCode
	Enable2FAErrCode
	Unable2FAErrCode
	Verify2FAErrCode
	MemberNotExistErrCode
	MemberStatusDuplicateErrCode
	NotFounderErrCode
	FounderErrCode
	CommentIsNotExistErrCode
)

const (
	SuccessMsg                  = "Success"
	ServerErrMsg                = "Service is unable to start successfully"
	ParamErrMsg                 = "Wrong Parameter has been given"
	CharacterBeyondLimitErrMsg  = "the number of character beyond the limit"
	UserAlreadyExistErrMsg      = "User existed"
	UserIsNotExistErrMsg        = "User is not exist"
	PasswordIsNotVerifiedMsg    = "Username or password not verified"
	AuthErrMsg                  = "It is not your account"
	ReadFileErrMsg              = "Error when read file"
	UploadFileErrMsg            = "Upload file error"
	Enable2FAErrMsg             = "2fa verification have opened"
	Unable2FAErrMsg             = "2fa verification have closed"
	Verify2FAErrMsg             = "incorrect otp password"
	MemberNotExistErrMsg        = "member not apply or join this party"
	MemberStatusDuplicateErrMsg = "member's status is not change"
	NotFounderErrMsg            = "this user have no access to this service"
	FounderErrMsg               = "you can't join your party"
	CommentIsNotExistErrMsg     = "Comment is not exist"

	WebSocketSuccessMsg             = "Connect to server success"
	WebSocketLogoutSuccessMsg       = "logout success"
	WebSocketTargetOnlineSuccessMsg = "target user is online"
	WebSocketConnectErrMsg          = "Connect or upgrade error"
	WebSocketTargetOfflineErrMsg    = "Target user is offline"
	WebSocketErrMsg                 = "Websocket error"
)

type ErrNo struct {
	ErrorCode int64
	ErrorMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("%s", e.ErrorMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrorCode: code,
		ErrorMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo { //出现不被定义的错误时
	e.ErrorMsg = msg
	return e
}

var (
	Success                   = NewErrNo(SuccessCode, SuccessMsg)
	ServiceError              = NewErrNo(ServiceErrCode, ServerErrMsg)
	ParamError                = NewErrNo(ParamErrCode, ParamErrMsg)
	CharacterBeyondLimitError = NewErrNo(CharacterBeyondLimitErrCode, CharacterBeyondLimitErrMsg)

	ExistUserError             = NewErrNo(ExistUserErrCode, UserAlreadyExistErrMsg)
	NotExistUserError          = NewErrNo(NotExistUserErrCode, UserIsNotExistErrMsg)
	PwdError                   = NewErrNo(AuthFailedErrCode, PasswordIsNotVerifiedMsg)
	AuthorizationError         = NewErrNo(AuthFailedErrCode, AuthErrMsg)
	UploadFileError            = NewErrNo(UploadFileErrCode, UploadFileErrMsg)
	ReadFileError              = NewErrNo(ReadFileErrCode, ReadFileErrMsg)
	Enable2FAError             = NewErrNo(Enable2FAErrCode, Enable2FAErrMsg)
	Unable2FAError             = NewErrNo(Unable2FAErrCode, Unable2FAErrMsg)
	Verify2FAError             = NewErrNo(Verify2FAErrCode, Verify2FAErrMsg)
	MemberNotExistError        = NewErrNo(MemberNotExistErrCode, MemberNotExistErrMsg)
	MemberStatusDuplicateError = NewErrNo(MemberStatusDuplicateErrCode, MemberStatusDuplicateErrMsg)
	NotFounderError            = NewErrNo(NotFounderErrCode, NotFounderErrMsg)
	FounderError               = NewErrNo(FounderErrCode, FounderErrMsg)
	CommentIsNotExistError     = NewErrNo(CommentIsNotExistErrCode, CommentIsNotExistErrMsg)

	WebSocketSuccess             = NewErrNo(WebSocketSuccessCode, WebSocketSuccessMsg)
	WebSocketLogoutSuccess       = NewErrNo(WebSocketLogoutSuccessCode, WebSocketLogoutSuccessMsg)
	WebSocketTargetOnlineSuccess = NewErrNo(WebSocketTargetOnlineSuccessCode, WebSocketTargetOnlineSuccessMsg)
	WebSocketConnectError        = NewErrNo(WebSocketConnectErrCode, WebSocketConnectErrMsg)
	WebSocketTargetOfflineError  = NewErrNo(WebSocketTargetOfflineErrCode, WebSocketTargetOfflineErrMsg)
	WebSocketError               = NewErrNo(WebSocketErrCode, WebSocketErrMsg)
)

// ConvertErr convert error to ErrNo
// in Default user ServiceErrCode
func ConvertErr(err error) ErrNo {
	errno := ErrNo{}
	if errors.As(err, &errno) {
		return errno
	}

	s := ServiceError
	s.ErrorMsg = err.Error()
	return s
}
