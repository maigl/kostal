package main

var html = `
<!DOCTYPE html>
<html>

<head>
	<title>power</title>
	<link rel="stylesheet"
          href="https://fonts.googleapis.com/css?family=Open+Sans+Condensed:300&display=swap">
    <style>
        body {
  		    font-family: 'Open Sans Condensed', regular;
			font-size: 18px;
			line-height: 1;
			padding: 0px;
			margin: 0px;
		}

		/*https://coolors.co/c41b5c-08415c-6b818c-f1bf98-eee5e9*/
		:root {
			--color1: #08415C;
			--color2: #6B818C;
			--color3: #F1BF98;
			--color4: #C41B5C;
		}

		div.value {
			text-align: right;
			font-size: 22em;
			padding-right: 31%;
		}

		div.unitblock {
			overflow: visible;
			float: right;
			padding-right: 19%;
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
			background-color: var(--color1);
			color: var(--color2);
		}

		#consumption {
			background-color: var(--color2);
			color: var(--color3);
		}

		#grid {
			background-color: var(--color3);
			color: var(--color4);
		}

		#yield {
			background-color: var(--color4);
			color: var(--color1);
		}



	</style>
	<script type="text/javascript">
  		setTimeout(function(){ location = '' },5000)
    </script>
</head>

<body>
	{{ range $key, $item := .}}
	<div class=topic id={{$key}}>
		<div class=unitblock>
			<div class=unitsymbol>{{$item.Unit}}</div>
			<div class=unitname>{{$item.Label}}</div>
		</div>
		<div  class=value>{{$item.Value}}</div>
	</div>
	{{end}}
</body>

`
