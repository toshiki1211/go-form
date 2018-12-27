<!DOCTYPE html>
<html lang="ja">
<head>
    <title></title>
</head>
<body>
    <div>
        <p>ユーザ名:{{ .Username }}</p>
        <p>パスワード:{{ .Password }}</p>
    </div>
    <form action="/" method="post">
        <input type="hidden" name="username" value="{{ .Username }}">
        <input type="hidden" name="password" value="{{ .Password }}">
        <input type="submit" value="修正">
    </form>
    <form action="/complate" method="post">
        <input type="hidden" name="username" value="{{ .Username }}">
        <input type="hidden" name="password" value="{{ .Password }}">
        <input type="submit" value="完了">
    </form>
</body>
</html>