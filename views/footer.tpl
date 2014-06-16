{{with .PrevButton}}
<p style="float:left">
  <a href="/?page={{.}}">Newer</a>
</p>
{{end}}
{{with .NextButton}}
<p style="float:right">
  <a href="/?page={{.}}">Older</a>
</p>
{{end}}
