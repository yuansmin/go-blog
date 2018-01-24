<!DOCTYPE html>
<html>
<head>
    <title>Blog List</title>
</head>
<body>
    {{range .blogs}}
        <h2>{{.Title}}</h2>
        <p>{{.Content}}</p>
        <a href="/blogs/{{.Id}}/delete">删除</a>
    {{end}}
</body>
</html>