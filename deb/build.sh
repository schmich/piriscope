cd `dirname $0`
(cd .. && git archive --format tar.gz master) > piriscope_1.0.orig.tar.gz
tar xvf piriscope_1.0.orig.tar.gz -C piriscope
(cd piriscope && debuild -uc -us)
