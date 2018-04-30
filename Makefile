now := `date +%Y/%m/%dT%H:%M:%S+09:00`

.PHONY: deploy
deploy:
	git add .
	git commit -m "deploy ${now}"
	git push origin master

.PHONY: clean
clean:
	-rm -rf post/content
	-rm -rf img

