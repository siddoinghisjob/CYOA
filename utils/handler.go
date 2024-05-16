package utils

import (
	"html/template"
	"net/http"
	"os"
)

var tpl *template.Template
var err error

func init() {
	tpl, err = template.New("").Parse(tmpl)
	if err != nil {
		os.Exit(1)
	}
}

var tmpl string = `
	<html>
		<head>
			<title>{{.Title}}Choose Your Own Adventure</title>
		</head>
		<body>
			<h1>{{.Title}}</h1>
			<div>
				{{range .Desc}}
					<p>{{.}}</p>
				{{end}}
				{{range .Options}}
				<ul>
					<li>
						<a href = {{.Arc}}> {{.Text}} </a>
					</li>
				</ul>
				{{end}}
			</div>
		</body>
	</html>
`

func (h Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == " " || path == "/" {
		path = "/intro"
	}
	data, ok := h.Data[path[1:]]

	if !ok {
		data = h.Data["intro"]
	}
	tpl.Execute(w, data)
}
