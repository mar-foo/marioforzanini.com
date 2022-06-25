PAGES = index.html \
        contacts.html \

all: $(PAGES)

%.html: %.md
	smu $< > $@

clean:
	rm *.html

.PHONY: all clean
