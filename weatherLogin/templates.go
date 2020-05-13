package main

//	form action="/login" Указывает обработчик, к которому обращаются данные формы при их отправке на сервер(submit)
//	method метод
//	Login: <input type="text" name="login"> Label и Edit, параметр login
//	Password: <input type="password" name="password"> тоже самое, только скрытое
//	<input type="submit" value="Login"> Button с названием Login

var (
	loginPageTmpl = string(`
	<html>
		<body>
		<form action="/login" method="post">
			Login: <input type="text" name="login">
			Password: <input type="password" name="password">
			<input type="submit" value="Login">
		</form>
		</body>
	</html>
`)

	innerPageTmpl = string(`
	<!DOCTYPE HTML>
	<html>
		<head>
			<meta charset = "utf-8">
			<title>Login</title>
		</head>
		<body>		
			Welcome, {{.Login}} <br />
			<form method="POST" action="/">
			Cities: <input type="text" name="cities" value="Moscow, Volgodonsk"> <br /><br />
			Period Begin: <select id="monthIdStart" name="monthStart">
				<option value=1>January</option>
				<option value=2>February</option>
				<option value=3>March</option>
				<option value=4>April</option>
				<option value=5>May</option>
				<option value=6>June</option>
				<option value=7>July</option>
				<option value=8>August</option>
				<option value=9>September</option>
				<option value=10>October</option>
				<option value=11 selected>November</option>
				<option value=12>December</option>
			</select><br /><br />
			Period End: <select id="monthIdEnd" name="monthEnd"><!--  -->
				<option value=1>January</option>
				<option value=2 selected>February</option>
				<option value=3>March</option>
				<option value=4>April</option>
				<option value=5>May</option>
				<option value=6>June</option>
				<option value=7>July</option>
				<option value=8>August</option>
				<option value=9>September</option>
				<option value=10>October</option>
				<option value=11>November</option>
				<option value=12>December</option>
			</select><br /><br />
			<input type="submit" value="Start">
			<input type="button" onclick="location.href='/logout'" value="Logout"> 
		</form><br />
		</body>
	</html>
	`)
)
