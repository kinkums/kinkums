<html>
  <head>
    <title> Upload an Image </title>
  </head>
  <body>
    <form enctype = "multipart/form-data" action = "http://127.0.0.1:9090/upload" method = "post">
      {{/* 1. File input */}}
      <input type = "file" name = "uploadfile"/>
      {{/* 2. Submit button */}}
      <input type = "submit" value = "Upload an Image"/>
    </form>
  </body>
</html>
