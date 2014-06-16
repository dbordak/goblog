<h1>{{.Title}}</h1>
<p>{{.Desc}}</p>
{{with .Form}}
<form action="" method="post">
  <dl>
    {{with .Name}}
    <dt><label for="f_title">{{.}}</label></dt>
    <dd><input type="text" id="f_title" value="" name="name"></dd>
    {{end}}
    {{with .Select}}
    <dt><label for="f_select">{{.Name}}</label></dt>
    <dd>
      <select id="f_select" value="" name="sel">
        {{if .DefOpt}}
        <option value="0">---</option>
        {{end}}
        {{range $id, $name := .Items}}
        <option value={{$id}}>{{$name}}</option>
        {{end}}
      </select>
    </dd>
    {{end}}
  </dl>
  {{with .Textarea}}
  <label for="f_textarea">{{.}}</label><br/>
  <textarea rows="8" cols="40" id="f_textarea" name="ta"></textarea>
  {{end}}
  <div class="actions"><input type="submit" value="Submit"></div>
</form>
{{end}}