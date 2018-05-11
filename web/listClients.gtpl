<html>
<title>TimeSheets - listClients</title>
<body>
<p>ClientID, FirstName, MiddleName, LastName</p>
{{ range . }}
<p>{{.ID}}, {{.FirstName}}, {{.MiddleName}}, {{.LastName}}</p>
{{ end }}
</body>
</html>
