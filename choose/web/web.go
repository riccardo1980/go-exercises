package choose

import (
	base "choose/base"
	story "choose/story"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type web struct {
	base.CommonFlags
	port int
}

func New() base.Base {
	var w web
	w.FlagSet = flag.NewFlagSet("web", flag.ExitOnError)
	w.FlagSet.IntVar(&w.port, "port", 8080, "Port")
	base.CommonFlagsDefine(&w.CommonFlags)

	return &w
}

func (w *web) Read() story.Story {
	st, _ := story.FromFile(w.InputFile)
	return st
}

func (w *web) Parse(args []string) {
	w.FlagSet.Parse(args)
}

func (w *web) Run(st story.Story) {
	mux := http.NewServeMux()

	basepath := "/story/"
	parser := func(r *http.Request) string {
		return defaultPathFn(r, basepath)
	}

	h := newHandler(st,
		WithTemplate(template.Must(template.New("").Parse(storyTmpl))),
		WithParser(parser),
	)
	mux.Handle(basepath, h)

	log.Printf("Starting server on port %d\n", w.port)
	http.ListenAndServe(fmt.Sprintf(":%d", w.port), mux)
}

func WithParser(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func newHandler(st story.Story, opts ...HandlerOption) http.Handler {
	var h handler
	h.st = st
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type HandlerOption func(h *handler)

type handler struct {
	st     story.Story
	t      *template.Template
	pathFn func(r *http.Request) string
}

func defaultPathFn(r *http.Request, prefix string) string {
	path := strings.TrimSpace(r.URL.Path)
	log.Printf("To parse: %s", path)
	path = strings.TrimPrefix(path, prefix)
	if path == "" {
		path = "intro"
	}
	log.Printf("Parsed: %s", path)
	return path
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := h.pathFn(r)
	log.Printf("Arc: %s", path)
	if arc, ok := h.st[path]; ok {
		err := h.t.Execute(w, arc)
		if err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
			log.Println(err)
			log.Fatal("Error")
		}
		return
	}
	http.Error(w, "Arc not found", http.StatusNotFound)

}

var storyTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Story}}
        <p>{{.}}</p>
      {{end}}
      <ul>
      {{range .Options}}
        <li><a href="/story/{{.Arc}}">{{.Text}}</a></li>
      {{end}}
      </ul>
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FCF6FC;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #797;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: underline;
        color: #555;
      }
      a:active,
      a:hover {
        color: #222;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`
