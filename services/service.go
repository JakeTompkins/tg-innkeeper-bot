package services

type resultCode int

const (
	SUCCESS   resultCode = 200
	NOT_FOUND resultCode = 404
	ERROR     resultCode = 500
)

type ServiceCommand struct {
	Command string
	Args    []string
}

type ServiceResult struct {
	Code    resultCode
	Message string
	Data    interface{}
}

type IService interface {
	Execute(command ServiceCommand) (ServiceResult, error)
	HelpMessage() string
}

type Service struct {
	Name          string
	ValidCommands []string
}
