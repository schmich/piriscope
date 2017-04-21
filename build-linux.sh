cd `dirname $0`

# Remove project Makefile to avoid conflict with Debian package Makefile.
rm Makefile

# Package source and extract next to Debian package.
tar --exclude='./deb' --exclude='.git' -zcvf piriscope_1.0.orig.tar.gz . 
tar -xvf piriscope_1.0.orig.tar.gz -C deb

# Build Debian package.
(cd deb && debuild -uc -us)
