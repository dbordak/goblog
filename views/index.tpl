{{range $title, $content := .Entries}}
<article>
  <h2><a href=>{{$title}}</a></h2>
  <p>{{$content | html2str}}</p>
</article>
{{end}}
