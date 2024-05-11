package isalive

import (
    "net/http"
    log "github.com/sirupsen/logrus"
)

type IsAlive struct {
    Sites map[string]bool
}

func NewIsAlive() *IsAlive {
    return &IsAlive{
        Sites: make(map[string]bool),
    }
}

func (ia *IsAlive) AddSite(s string) {
    ia.Sites[s] = false 
}

func (ia *IsAlive) CheckHTTPs() {
    for s := range ia.Sites {
        ia.Sites[s] = CheckHTTP(s)
    }
}

func (ia *IsAlive) PublishState() {
    log.Info("-----")
    for s := range ia.Sites {
        log.Infof("%-30s: %v", s, ia.Sites[s])
    }
}

func CheckHTTP (url string) bool {
    _, err := http.Get(url) 
    if err != nil { log.Debug(err) }
    return err == nil
}
