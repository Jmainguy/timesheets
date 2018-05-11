<html>
<title>TimeSheets - listServices</title>
<body>
<p>Name, GroupName</p>
{{ range . }}
<p>{{.Name}}, {{.GroupName}}</p>
{{ end }}
</body>
</html>
