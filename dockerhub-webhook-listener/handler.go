package listener

import "log"

type Handler interface {
	Call(HubMessage)
}

type Logger struct{}

func (l *Logger) Call(msg HubMessage) {
	log.Print(msg)
}

func execShellHandler(msg HubMessage) {
  log.Println("receiving docker hub redeploy trigger")
  out, err := exec.Command("../redeploy.sh").Output()
  if err != nil {
    log.Println("Error running redeploy.sh:")
    log.Println(err)
    return
  }
  log.Println("successful ran redeploy.sh:", string(out))
}

type Registry struct {
	entries []func(HubMessage)
}

func (r *Registry) Add(h func(msg HubMessage)) {
	r.entries = append(r.entries, h)
	return
}

func (r *Registry) Call(msg HubMessage) {
	for _, h := range r.entries {
		go h(msg)
	}
}

func MsgHandlers() Registry {
	var handlers Registry

	handlers.Add(execShellHandler)

	return handlers
}
