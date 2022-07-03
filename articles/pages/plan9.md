<img src="/pub/pics/bunny_white.jpg" width=300>

Plan 9 is a Unix-like operating system, written at Bell
Labs in the late '80s. The two main ideas behind the
design of the OS are **private namespaces**: every
process has an independent view of the file tree,
and **file interfaces**: everything in plan9 can be
accessed through a file. Based on these two premises the
system appears to be very awkward at first, as in
different from anything else, but also very coherent and
easy to predict, once you get the hang of it. It was born
as a research operating system and as such it doesn't not
strive to be a product.

I'm still experimenting with it, but I find it very
interesting and amusing to use especially because it
lacks a 'modern' browser and all the distractions it
brings.

## Resources

* [Cat -v](http://cat-v.org/) (not only plan9 related stuff)
* [The fork I'm using](http://9front.org)
* [Many useful tips](https://pspodcasting.net/dan/blog/2019/plan9_desktop.html)

# Rename

Conveniently rename files.  I basically use this to remove whitespace
and other weird characters from file names

	#!/bin/rc
	# usage: rename file pattern
	if(! ~ $#* 2){
		echo 'usage: rename file pattern' >[1=2]
		exit 'error'
	}
	
	new_name=`{echo $1 | sed $2}
	mv $1 $new_name

# Balance

I keep my balance in a simple greppable™ text file
($home/lib/balance).

File format:

	date;amount;tag1 tag2 tag3

Money spent is considered negative. To compute the total balance:

	#!/bin/awk -F; -f
	BEGIN{s=0}
		{s+=($2-0)}
	END{print s}

e.g.:

	% balance/tot < lib/balance
	-10000

To compute the total of all entries matching a pattern:

	#!/bin/rc 
	fn Search{
	awk -F';' -v 's=0' 'BEGIN{} /'$1'/ {print $2, $3 ; s+=($2-0)}
		END{print s}'
	}
	Search $1

e.g. to compute how much I spent on stuff:

	% balance/s stuff < lib/balance
	-100 stuff home
	-10000 stuff unneeded
	-10100
