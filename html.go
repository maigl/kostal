package main

var html = `
<!DOCTYPE html>
<html>

<head>
	<title>power</title>
	<style>
		body {
			padding: 0px;
			margin: 0px;
		}

		div.value {
			text-align: right;
			font-size: 21em;
			padding-right: 31%;
		}

		div.unitblock {
			overflow: visible;
			float: right;
			padding-right: 10%;
			padding-top: 18%;
			width: 20%;
		}

		div.unitsymbol {
			font-size: 5em;
		}

		div.unitname {
			font-size: 3em;
		}

		#battery {
			background-color: #434c5e;
			color: #697794;
		}

		#yield {
			background-color: #668585;
			color: #81a7a7;
		}

		#consumption {
			background-color: #5e81ac;
			color: #3d536e;
	</style>
	<script type="text/javascript">
  		setTimeout(function(){ location = '' },60000)
    </script>
</head>

<body>
	{{ range $key, $item := .}}
	<div class=topic id={{$key}}>
		<div class=unitblock>
			<div class=unitsymbol>{{$item.Unit}}</div>
			<div class=unitname>{{$item.Label}}</div>
		</div>
		<div class=value>{{$item.Value}}</div>
	</div>
	{{end}}
</body>

`
