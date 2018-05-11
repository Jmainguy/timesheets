<html>
<title>TimeSheets - addService</title>
<body>
{{if .}}
    </h2>{{.}}</h2>
{{end}}
<form action="/addService" method="post">
  Service (for example '___Assist with ambulation', see how it has the ___ in front)<br>
  <input type="text" name="name">
  <br>
  Group:<br>
  <select name="group">
    <option value="ADL Mobility">ADL Mobility</option>
    <option value="ADL Bathing">ADL Bathing</option>
    <option value="ADL Dressing">ADL Dressing</option>
    <option value="ADL Toileting">ADL Toileting</option>
    <option value="Medical Monitoring/TX">Medical Monitoring/TX</option>
    <option value="IADL Meal Prep">IADL Meal Prep</option>
    <option value="IADL Home Management">IADL Home Management</option>
    <option value="Other Tasks">Other Tasks</option>
  <br>
  <br><br>
  <input type="submit" value="Submit">
</form> 
<p>See existing services <a href=/listServices>Here</a></p>
</body>
</html>
