{{range .Entries}}
<article>
  <h2><a href={{.}}>{{.}}</a></h2>
  <p>{{. | html2str}}</p>
</article>
{{end}}
