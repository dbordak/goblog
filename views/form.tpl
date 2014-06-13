<h1>{{.Title}}</h1>
<form action="" method="post">
  {{with .Form}}
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
        <option value="None">---</option>
        {{end}}
        {{.Items}}
      </select>
    </dd>
    {{end}}
    {{with .Textarea}}
    <dt><label for="f_textarea">{{.}}</label></dt>
    <dd><textarea rows="8" cols="40" id="f_textarea" name="ta"></textarea></dd>
    {{end}}
  </dl>
  {{end}}
  <div class="actions"><input type="submit" value="Submit"></div>
</form>
