package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type CalcData struct {
	IP        string
	Result    string
	First     string
	Second    string
	Operation string
}

var tpl = `
<!DOCTYPE html>
<html>
<head>
	<title>Star Wars Calculator</title>
	<style>
		body {
			background-color: #1a1a1a;
			color: #FFE81F;
			font-family: Arial, sans-serif;
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;
			height: 100vh;
		}
		form {
			background-color: #222;
			padding: 30px;
			border-radius: 10px;
			box-shadow: 0 0 15px #FFE81F;
		}
		input, select {
			padding: 10px;
			margin: 10px;
			width: 100px;
			font-size: 1em;
			border: 1px solid #FFE81F;
			background: #000;
			color: #FFE81F;
		}
		input[type=submit] {
			cursor: pointer;
			width: auto;
		}
		h1 {
			color: #FFE81F;
		}
	</style>
</head>
<body>
	<h3>Local IP: {{.IP}}</h3>
	<h1>Star Wars Calculator</h1>
	<form method="POST">
		<input type="text" name="first" placeholder="First" value="{{.First}}">
		<select name="operation">
			<option value="add" {{if eq .Operation "add"}}selected{{end}}>+</option>
			<option value="subtract" {{if eq .Operation "subtract"}}selected{{end}}>-</option>
			<option value="multiply" {{if eq .Operation "multiply"}}selected{{end}}>*</option>
			<option value="divide" {{if eq .Operation "divide"}}selected{{end}}>/</option>
		</select>
		<input type="text" name="second" placeholder="Second" value="{{.Second}}">
		<br>
		<input type="submit" value="Calculate">
	</form>
	<h2>Result: {{.Result}}</h2>
</body>
</html>
`

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "Unknown"
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "Unknown"
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := CalcData{
		IP: getLocalIP(),
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		data.First = r.FormValue("first")
		data.Second = r.FormValue("second")
		data.Operation = r.FormValue("operation")

		a, err1 := strconv.ParseFloat(data.First, 64)
		b, err2 := strconv.ParseFloat(data.Second, 64)

		if err1 != nil || err2 != nil {
			data.Result = "Invalid input"
		} else {
			switch data.Operation {
			case "add":
				data.Result = fmt.Sprintf("%.2f", a+b)
			case "subtract":
				data.Result = fmt.Sprintf("%.2f", a-b)
			case "multiply":
				data.Result = fmt.Sprintf("%.2f", a*b)
			case "divide":
				if b == 0 {
					data.Result = "Cannot divide by zero"
				} else {
					data.Result = fmt.Sprintf("%.2f", a/b)
				}
			default:
				data.Result = "Unknown operation"
			}
		}
	}

	t := template.Must(template.New("calc").Parse(tpl))
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
