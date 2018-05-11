<html>
<title>TimeSheets - addClient</title>
<body>
{{if .}}
    <h2>That worked, {{.}} has been added</h2>
{{end}}
<form action="/addClient" method="post">
  First name:<br>
  <input type="text" name="firstname">
  <br>
  Last name:<br>
  <input type="text" name="lastname">
  <br>
  Middle name:<br>
  <input type="text" name="middlename">
  <br><br>
  <input type="submit" value="Submit">
</form> 
<p>See existing clients <a href=/listClients>Here</a></p>
</body>
</html>
