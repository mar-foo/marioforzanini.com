PAGES = index.html \
        lezione1.html \
        lezione2.html \
        lezione4.html \
        lezione5.html \
        lezione6.html \
        lezione7.html \
        lezione8.html \
        programmazione1.html \
        programmazione2.html \
        programmazione3.html \
        programmazione4.html \
        programmazione6.html \
        programmazione8.html \
        programmazione9.html \
        programmazione10.html \
        programmazione11.html \
        programmazione12.html \
        programmazione13.html \
        programmazione14.html

all: $(PAGES)

%.html: %.md
	shtpl < $< | smu > $@

clean:
	rm -f *.html

.PHONY: all clean
