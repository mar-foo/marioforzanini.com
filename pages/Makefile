PAGES = index.html \
        contacts.html \

all: $(PAGES)

index.html: index.md
	smu $< | shtpl > $@

%.html: %.md
	smu $< > $@

clean:
	rm *.html

.PHONY: all clean
