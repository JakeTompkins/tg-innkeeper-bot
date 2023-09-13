package services

import (
	"fmt"
)

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
	Execute(command *ServiceCommand) (*ServiceResult, error)
	CommandIsValid(command *ServiceCommand) bool
	HelpMessage() string
}

type Service struct {
	Name          string
	ValidCommands []string
}

func (s *Service) CommandIsValid(command *ServiceCommand) bool {
	for _, validCommand := range s.ValidCommands {
		if validCommand == command.Command {
			return true
		}
	}

	return false
}

var services = [1]IService{NewD20Service()}

func ExecuteCommand(command *ServiceCommand) (*ServiceResult, error) {
	var result *ServiceResult
	var err error

	for _, service := range services {
		if service.CommandIsValid(command) {
			result, err = service.Execute(command)

			if err != nil {
				fmt.Println(err)
			}

			break
		}
	}

	return result, err
}
