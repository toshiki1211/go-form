<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <form action="/confirm" method="post">
        ユーザ名:<input type="text" name="username" value="{{ if eq .ValidateError true }}{{ .Username }}{{end}}">
        {{ if eq .ValidateError true }}
            {{ with .ErrorMessages.username }}{{ . }}{{ end }}
        {{ end }}
        パスワード:<input type="password" name="password" value="{{ if eq .ValidateError true }}{{ .Password }}{{end}}">
        {{ if eq .ValidateError true }}
            {{ with .ErrorMessages.password }}{{ . }}{{ end }}
        {{ end }}
        <input type="submit" value="ログイン">
    </form>
</body>
</html>