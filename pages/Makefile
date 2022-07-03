PAGES = index.html \
        contacts.html \

all: $(PAGES)

index.html: index.md
	smu $< | shtpl > $@

%.html: %.md
	smu $< > $@

clean:
	rm -f *.html

.PHONY: all clean
