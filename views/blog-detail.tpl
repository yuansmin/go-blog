<!DOCTYPE html>
<html>
<head>
    <title>{{.blog.Title}}</title>
</head>
<body>
    {{if not .message }}
        <h2>{{.blog.Title}}</h2>
        <p>{{.blog.Content}}</p>
    {{else}}
        <p>{{.message}}</p>
    {{end}}
</body>
</html>