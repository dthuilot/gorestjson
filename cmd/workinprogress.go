package main

import "net/http"

func workInProgress(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`<!DOCTYPE html>
	<html>
	
	<head>
		<title>goWebServer</title>
		<meta charset="utf-8">
		<meta name="MobileOptimized" content="320">
		<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1,user-scalable=no">
		<meta name="theme-color" content="#eb2328">
		<meta name="apple-mobile-web-app-capable" content="yes">
		<meta http-equiv="X-UA-Compatible" content="IE=Edge">
		<meta name="content-type" content="text/html;charset=utf-8">
		<script src="//assets.adobedtm.com/10cb5d082fb7/7363b27fd56b/launch-644b0be21518.min.js"></script>
		<link rel="stylesheet" type="text/css" href="https://i.annihil.us/u/prod/marvel/html_pages_assets/error-pages/prod/main.0e491aa7.css">
		<script src="https://i.annihil.us/u/prod/marvel/html_pages_assets/error-pages/prod/main.c6892173.js"></script>
	</head>
	
	<body class="error-status-page">
		<header class="static-header">
			<nav>
				<div class="upper-nav">
					<ol> <a id="marvel-logo" href="https://www.marvel.com"></a>
					</ol>
				</div>
			</nav>
		</header>
		<section id="error-status" class="error-status" data-state="">
			<div class="overlay"></div>
			<div class="flex-wrapper no-pad">
				<div class="flex-col">
					<div class="error-copy">
						<h1>goWebServer is a work in progress</h1>
						<h4 class="dynamic-msg"></h4>
						<p>Check that you typed the address correctly, go back to your previous page or try using our site search to find something specific.</p>
					</div>
				</div>
				<div class="flex-col toAnimate no-pad">
					<div class="error-image-animate"></div>
				</div>
			</div>
		</section>
		<script>
			window.establishDigitalData("404"), _satellite.track("errorPageLoad"), _satellite.pageBottom();
		</script>
	</body>
	
	</html>`))
}
