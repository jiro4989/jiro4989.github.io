now := `date +%Y/%m/%dT%H:%M:%S+09:00`

deploy:
	git add .
	git commit -m "deploy ${now}"
	git push origin master

