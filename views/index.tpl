{{range $ent := .Entries}}
<article>
  <h2><a href=$ent.Url()>{{$ent.Title}}</a></h2>
  <p>{{$ent.Content | html2str}}</p>
</article>
{{end}}
