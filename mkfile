all:V: articles cbi
	for(i in `{walk pages | grep cfg | sort -r})
		saait $i
	cp style.css output/pub/

articles:V: pages
	cd pages; mk all
	cd ..

cbi:V: CBI2021
	cd CBI2021 ; mk all
	cd ..
	dircp CBI2021/output/ output/CBI2021/

clean:Q:
	rm output/*.html output/pub/*.css
	cd pages ; mk clean
	cd ..
	cd CBI2021 ; mk clean
	cd ..

