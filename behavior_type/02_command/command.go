package command

import "fmt"

type Command interface {
	Execute()
}

type StartCommand struct {
	mb *MotherBoard
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{
		mb: mb,
	}
}

func (s *StartCommand) Execute() {
	s.mb.Start()
}

type RebootCommand struct {
	mb *MotherBoard
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{
		mb: mb,
	}
}

func (c *RebootCommand) Execute() {
	c.mb.Reboot()
}

type MotherBoard struct{}

func (s *MotherBoard) Start() {
	fmt.Println("system starting")
}

func (s *MotherBoard) Reboot() {
	fmt.Println("system rebooting")
}

type Box struct {
	button1 Command
	button2 Command
}

func NewBox(b1, b2 Command) *Box {
	return &Box{
		button1: b1,
		button2: b2,
	}
}

func (b *Box) PressButton1() {
	b.button1.Execute()
}

func (b *Box) PressButton2() {
	b.button2.Execute()
}
