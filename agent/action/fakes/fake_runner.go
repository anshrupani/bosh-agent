package fakes

import (
	boshaction "github.com/cloudfoundry/bosh-agent/v2/agent/action"
)

type FakeRunner struct {
	RunAction          boshaction.Action
	RunPayload         []byte
	RunProtocolVersion boshaction.ProtocolVersion
	RunValue           interface{}
	RunErr             error

	ResumeAction  boshaction.Action
	ResumePayload []byte
	ResumeValue   interface{}
	ResumeErr     error
}

func (runner *FakeRunner) Run(action boshaction.Action, payload []byte, version boshaction.ProtocolVersion) (interface{}, error) {
	runner.RunAction = action
	runner.RunPayload = payload
	runner.RunProtocolVersion = version
	return runner.RunValue, runner.RunErr
}

func (runner *FakeRunner) Resume(action boshaction.Action, payload []byte) (interface{}, error) {
	runner.ResumeAction = action
	runner.ResumePayload = payload
	return runner.ResumeValue, runner.ResumeErr
}
