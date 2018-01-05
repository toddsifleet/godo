package tail

import (
	"fmt"

	realTail "github.com/hpcloud/tail"
	"github.com/toddsifleet/godo/index"
	"github.com/toddsifleet/godo/models"
)

type Tail struct {
	tail  *realTail.Tail
	index index.Index
}

func (t *Tail) Run() {
	defer t.tail.Done()
	for line := range t.tail.Lines {
		if line.Err != nil {
			panic(line.Err)
		}
		cmd, err := models.CommandFromLogLine(line.Text)
		if err != nil {
			fmt.Println("Error Parsing Line:", string(line.Text))
			continue
		}
		t.index.AddCommand(cmd)
	}
}

func New(idx index.Index, logPath string) (*Tail, error) {
	t, err := realTail.TailFile(logPath, realTail.Config{Follow: true})
	if err != nil {
		return nil, err
	}

	return &Tail{
		index: idx,
		tail:  t,
	}, nil
}
