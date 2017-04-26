deb: *.deb

piriscope: piriscope.go
	ruby -I. build-go-docker.rb

%.deb: piriscope
	ruby -I. build-deb-docker.rb
