<nav>
  <a href="">A link</a><br/>
  <a href="">Another link</a>
  <br/><br/>
  <label for="sidebar_category">Category: </label><br/>
  <select name="sidebar_category" onchange="toCategoryPage()" id="side_dd">
    <option value="">---</option>
    {{range .Categories}}
    <option value={{.}}>{{.Name}}</option>
    {{end}}
  </select>
  <hr class="section" id="navsep_b"/>
  <a href="/about">About</a>
  {{if .IsAdmin}}
  <br/><a href="/admin/">Admin Page</a>
  {{end}}
</nav>
