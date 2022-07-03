all: articles
	find pages -type f -name "*.cfg" -print0 | sort -zr | xargs -0 saait

articles: pages
	cd pages ; $(MAKE) all

clean:
	rm -f output/*.html
	cd pages ; $(MAKE) clean

.PHONY: all clean
