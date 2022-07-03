PAGES = index.html \
	cell_phones.html \
	cell_phones_updates.html \
        plan9.html \
        read_more.html \
        static.html

all: $(PAGES)

index.html:
	./makeindex.sh > index.html

%.html: %.md
	smu $< > $@

clean:
	rm -f *.html

.PHONY: all clean
