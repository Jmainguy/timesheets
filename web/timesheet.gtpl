<!DOCTYPE html>
<html>
<head>
<style>
table {
    font-family: arial, sans-serif;
    border-collapse: collapse;
    display: inline-block;
    width: 14%;
    #font-size: 15px;
}

td, th {
    border: 2px solid black;
    text-align: left;
    padding: 1px;
}

</style>
</head>
<body>

<form enctype="multipart/form-data" action="/timesheet" method="post">
	<table>
  		<tr bgcolor='#dddddd'>
    		<th><center>Monday<br>05/07/2018</center></th>
  		</tr>
  		<tr bgcolor='#dddddd'>
    		<th><center>Aide Name</center></th>
  		</tr>
        {{ $T := . }}
        {{ range $T.ServiceGroups }}
            {{ $SGN := .GroupName }}
            <tr bgcolor='#dddddd'>
                <td><b>{{ $SGN }}</b></td>
            </tr>
            {{ range $T.Services }}
                {{ if eq .GroupName $SGN }}
                    <tr>
                        <td>{{ .Name }}<select name="{{ .Name }}">
                            <option value="BLANK"></option>
                            <option value="YES">Yes</option>
                        </td>
                    </tr>
                {{ end }}
            {{ end }}
        {{ end }}
	</table>
