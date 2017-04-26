require 'version'

project = 'piriscope'

Dir.chdir(File.dirname(__FILE__))

system("docker build -f Dockerfile-go -t #{project}-go .") || fail
system("docker run -it -e VERSION=#{$version} -e COMMIT=#{$commit} #{project}-go") || fail
id = `docker ps -l -q`.strip
system("docker cp '#{id}:/go/src/github.com/schmich/#{project}/#{project}' .") || fail
system("docker rm #{id}") || fail
