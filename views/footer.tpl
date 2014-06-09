{{with .PrevButton}}
<p style="float:left">
  <a href="/?cursor={{.}}">Newer</a>
</p>
{{end}}
{{with .NextButton}}
<p style="float:right">
  <a href="/?cursor={{.}}">Older</a>
</p>
{{end}}
