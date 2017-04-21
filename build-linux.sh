cd `dirname $0`
git archive --format tar.gz master > piriscope_1.0.orig.tar.gz
tar xvf piriscope_1.0.orig.tar.gz -C deb
(cd deb && debuild -uc -us)
