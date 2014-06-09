{{range .Entries}}
<article>
  <h2><a href={{.Url}}>{{.Title}}</a></h2>
  <p>{{.Content | html2str}}</p>
</article>
{{end}}
