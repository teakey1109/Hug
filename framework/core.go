package framework

import "net/http"

// Engine --new handler for Hug
type Engine struct {
}

func NewEngine() *Engine {
	return &Engine{}
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}
