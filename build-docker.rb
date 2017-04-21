project = 'piriscope'
package = "#{project}_1.0-1_all.deb"

Dir.chdir(File.dirname(__FILE__))

system("docker build -t #{project} .") || fail
system("docker run -it #{project}") || fail
id = `docker ps -l -q`.strip
system("docker cp '#{id}:/src/#{package}' .") || fail
system("docker rm #{id}") || fail
