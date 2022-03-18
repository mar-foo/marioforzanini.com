all: articles cbi
	find pages -type f -name "*.cfg" -print0 | sort -zr | xargs -0 saait
	cp style.css output/pub/

articles: pages
	cd pages ; $(MAKE) all

cbi: CBI2021
	cd CBI2021 ; $(MAKE) all
	cp -r CBI2021/output/* output/CBI2021/

clean:
	rm output/*.html output/pub/*.css
	cd pages ; $(MAKE) clean
	cd CBI2021 ; $(MAKE) clean

.PHONY: all clean
