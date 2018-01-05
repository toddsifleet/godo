package client

import (
	"errors"
	"net/rpc"

	ui "github.com/gizak/termui"
	"github.com/toddsifleet/godo/models"
	"github.com/toddsifleet/godo/terminal"
)

type Client struct {
	rpc              *rpc.Client
	currentDirectory string
	searchType       string
	serverMethod     string
}

func New(host, currentDirectory, searchType string) (*Client, error) {
	c, err := rpc.DialHTTP("tcp", host)
	if err != nil {
		return nil, err
	}
	serverMethod := map[string]string{
		"command":   "Server.FindCommands",
		"directory": "Server.FindDirectories",
	}[searchType]

	if serverMethod == "" {
		return nil, errors.New("invalid search type")
	}
	return &Client{
		rpc:              c,
		currentDirectory: currentDirectory,
		searchType:       searchType,
		serverMethod:     serverMethod,
	}, nil
}

func colorOptions(options []string, i int) []string {
	result := make([]string, len(options))
	if len(options) == 0 {
		return result
	}
	copy(result, options)
	result[i] = "[" + options[i] + "](fg-green)"
	return result
}

func (c *Client) search(searchTerm string, limit int) ([]string, error) {
	response := &models.Response{}
	request := models.Request{
		CurrentDirectory: c.currentDirectory,
		Limit:            limit,
		SearchTerm:       searchTerm,
	}
	err := c.rpc.Call(c.serverMethod, request, response)
	return response.Options, err
}

func (c *Client) render(searchTerm string) (resultValue, resultAction string, err error) {
	currentSelection := 0
	var options []string
	if err = ui.Init(); err != nil {
		return
	}
	defer ui.Close()

	searchPar := ui.NewPar(searchTerm)
	searchPar.Height = 2
	searchPar.Width = ui.TermWidth()
	searchPar.BorderBottom = true
	searchPar.BorderLeft = false
	searchPar.BorderTop = false
	searchPar.BorderRight = false

	optionList := ui.NewList()
	optionList.Width = ui.TermWidth()
	optionList.Height = ui.TermHeight() - 4
	optionList.Y = 2
	optionList.Border = false
	optionList.Overflow = "wrap"
	optionList.ItemFgColor = ui.ColorWhite

	reRender := func() {
		if err != nil {
			panic(err)
		}
		optionList.Items = colorOptions(options, currentSelection)
		searchPar.Text = searchTerm
		ui.Render(searchPar, optionList)
	}

	updateSearchTerm := func(s string) {
		defer reRender()
		searchTerm = s
		options, err = c.search(searchTerm, ui.TermHeight()-4)
	}

	updateSelection := func(i int) func(ui.Event) {
		return func(ui.Event) {
			defer reRender()
			switch i {
			case -1:
				if currentSelection == 0 {
					currentSelection = len(options) - 1
					return
				}
				currentSelection--
			case 1:
				if currentSelection == len(options)-1 {
					currentSelection = 0
					return
				}
				currentSelection++
			}
		}
	}

	updateSearchTerm(searchTerm)
	ui.Handle("/sys/kbd/<down>", updateSelection(1))
	ui.Handle("/sys/kbd/C-j", updateSelection(1))
	ui.Handle("/sys/kbd/<up>", updateSelection(-1))
	ui.Handle("/sys/kbd/C-k", updateSelection(-1))
	ui.Handle("/sys/kbd/<escape>", func(ui.Event) {
		resultValue = options[currentSelection]
		resultAction = "ECHO_VALUE"
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd/<enter>", func(ui.Event) {
		resultValue = options[currentSelection]
		resultAction = "RUN_VALUE"
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd/C-8", func(e ui.Event) {
		if len(searchTerm) > 0 {
			updateSearchTerm(searchTerm[:len(searchTerm)-1])
		}
	})
	ui.Handle("/sys", func(e ui.Event) {
		if k, ok := e.Data.(ui.EvtKbd); ok && len(k.KeyStr) == 1 {
			updateSearchTerm(searchTerm + k.KeyStr)
		}
	})
	ui.Loop()
	return
}

func (c *Client) commandSearch(s string) error {
	cmd, action, err := c.render(s)
	if err != nil || cmd == "" {
		return err
	}

	if action == "ECHO_VALUE" {
		return terminal.Echo(cmd)
	}
	return terminal.Run(cmd)
}

func (c *Client) directorySearch(s string) error {
	cmd, action, err := c.render(s)
	if err != nil || cmd == "" {
		return err
	}
	if cmd == "" {
		return nil
	}

	cmd = "cd " + cmd
	if action == "ECHO_VALUE" {
		return terminal.Echo(cmd)
	}
	return terminal.Run(cmd)
}

func (c *Client) Run(s string) error {
	switch c.searchType {
	case "directory":
		return c.directorySearch(s)
	case "command":
		return c.commandSearch(s)
	}
	return nil
}
