src := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
package := piriscope_1.0-1_all.deb

$(package):
	ruby build-docker.rb

deb: $(package)

inspect: deb
	docker run -it --rm -v $(src)$(package):/root/$(package) --workdir /root debian bash
