To manage my website I use 3 simple tools:

* Make (or mk on [9front](/plan9.html))
* [saait](http://codemadness.org/saait.html)
* [smu](http://github.com/Gottox/smu)

## Workflow

0. Write a markdown document
0. Generate html source with Make (or mk)
0. Done

## My setup

This is the hierarchy of my website:

* Articles go in the root of the website
* I put public files (like pictures and stylesheets) in ``/pub``
* I handle everything that is not blogging-related in subdirectories

### Adding an article

I have copied and adapted page templates from the examples in saait
documentation, the only things I need to do to add a new ``page.html``
entry to the website is:

0. Add a link to it in the header template
0. Add it to the Makefile
0. Write the article (in smu markdown)
0. Add a ``.cfg`` file with some metadata:

	filename = page.html
	title = test page
	description = no one ever reads this

### Generating the html

The structure enforced by saait looks like this:

	pages/ # the actual source
	output/ # the result
	templates/ # guess this one
	config.cfg # global website config

I simply add a Makefile in the root of the website:

	all: articles subdir1 subdir2
		find pages -type f -name "*.cfg" -print0 | sort -zr | xargs -0 saait
		cp style.css output/pub/
	
	articles: pages
		cd pages ; $(MAKE) all
	
	subdir1: subdir1
		cd subdir1 ; $(MAKE) all
		cp -r subdir1/output/* output/subdir1/
	
	subdir2: subdir2
		cd subdir2 ; $(MAKE) all
		cp -r subdir2/output/* output/subdir2/

	clean:
		rm output/*.html output/pub/*.css
		cd pages ; $(MAKE) clean
		cd subdir1 ; $(MAKE) clean
		cd subdir2 ; $(MAKE) clean
	
	.PHONY: all clean

and one in the ``pages`` directory:

	PAGES = article1.html \
	        article2.html \
	        article3.html
	
	all: $(PAGES)
	
	%.html: %.md
		smu $< > $@
	
	clean:
		rm *.html
	
	.PHONY: all clean

Then I repeat this for every subdirectory.  Issuing ``make`` in the
root of the tree simply generates all html sources and moves them to
the ``output`` directory.

## Final structure

This is what the tree of my website directory looks like in the end:

	Makefile
	config.cfg
	style.css
	output/index.html
	output/article1.html
	output/article2.html
	output/subdir1/index.html
	output/subdir1/article.html
	output/subdir2/article2.html
	output/pub/style.css
	output/pub/pics/pic1.png
	pages/article1.cfg
	pages/article1.html
	pages/article1.md
	pages/article2.cfg
	pages/article2.html
	pages/article2.md
	pages/index.cfg
	pages/index.html
	pages/index.md
	pages/Makefile
	subdir1/Makefile
	subdir1/config.cfg
	subdir1/output/index.html
	subdir1/output/article.html
	subdir1/pages/article.cfg
	subdir1/pages/article.html
	subdir1/pages/article.md
	subdir1/pages/index.cfg
	subdir1/pages/index.html
	subdir1/pages/index.md
	subdir1/templates/page/footer.html
	subdir1/templates/page/header.html
	subdir1/templates/page/item.html
	subdir2/Makefile
	subdir2/config.cfg
	subdir2/output/index.html
	subdir2/output/article.html
	subdir2/pages/article.cfg
	subdir2/pages/article.html
	subdir2/pages/article.md
	subdir2/pages/index.cfg
	subdir2/pages/index.html
	subdir2/pages/index.md
	subdir2/templates/page/footer.html
	subdir2/templates/page/header.html
	subdir2/templates/page/item.html
	templates/page/footer.html
	templates/page/header.html
	templates/page/item.html

# Final remarks

I'm pleased with the result: the site is very fast (because it is
static), very easy to update (I don't need to write HTML by hand
anymore) and mantainable.
