deb: *.deb

piriscope: piriscope.go
	ruby -I. build-go-docker.rb

%.deb: piriscope
	ruby -I. build-deb-docker.rb

smoke: piriscope.go
	go build -o piriscope.tmp && rm piriscope.tmp

upload: *.deb
	ssh pi@pi "rm ~/piriscope*.deb" && \
	scp piriscope*.deb pi@pi:~/ && \
	ssh pi@pi "sudo dpkg -P piriscope && sudo dpkg -i ~/piriscope*.deb"
