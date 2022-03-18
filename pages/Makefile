PAGES = index.html \
        cell_phones.html \
        contacts.html \
        plan9.html \
        read_more.html \
        static.html

all: $(PAGES)

%.html: %.md
	smu $< > $@

clean:
	rm *.html

.PHONY: all clean
