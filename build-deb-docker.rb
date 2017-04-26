require 'version'

project = 'piriscope'
package = "#{project}_#{$version}_armhf.deb"

Dir.chdir(File.dirname(__FILE__))

system("docker build -f Dockerfile-deb -t #{project}-deb .") || fail
system("docker run -it -e VERSION=#{$version} #{project}-deb") || fail
id = `docker ps -l -q`.strip
system("docker cp '#{id}:/src/build/#{package}' .") || fail
system("docker rm #{id}") || fail
