# MoonFoxBox

这是一个个人学习项目：

golang版本： go1.16

recent Router

| Method | URL Pattern     | Handler         | Action                       |
| ------ | --------------- | --------------- | ---------------------------- |
| ANY    | /               | home            | Display the home page        |
| ANY    | /snippet?id=1   | showSnippet     | Display a specific snippet   |
| POST   | /snippet/create | createSnippet   | Create a new snippet         |
| ANY    | /static/        | http.FileServer | Serve a spectifc static file |



2021-9-28 16:43:41 createProject

2021-9-28 23:53:08 chapter2.4 router requests

2021-9-29 23:59:53 Chapter 2.5. Customizing HTTP Headers

