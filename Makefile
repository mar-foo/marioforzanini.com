all: cbi art
	find pages -type f -name "*.cfg" -print0 | sort -zr | xargs -0 saait
	cp style.css output/pub/

art: pages articles
	cd pages ; $(MAKE) all
	cd articles ; $(MAKE) all
	cp -r articles/output/* output/articles/

cbi:
	cd CBI2021 ; $(MAKE) all
	cp -r CBI2021/output/* output/CBI2021/
	cp CBI2021/pages/*.go output/CBI2021/

clean:
	rm output/*.html output/pub/*.css
	cd pages ; $(MAKE) clean
	cd articles ; $(MAKE) clean
	cd CBI2021 ; $(MAKE) clean

.PHONY: all clean
