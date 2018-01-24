<!DOCTYPE html>
<html>
<head>
    <title>Edit Blog</title>
</head>
<body>
    <form action="" method="post">
        标题：<input type="text" name="Title" value="{{.blog.Title}}"><br \>
        内容：<textarea name="Content" colspan="3" rowspan="10">{{.blog.Content}}</textarea><br \>
        <input type="submit" value="修改">
    </form>
</body>
</html>