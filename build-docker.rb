project = 'piriscope'
package = "#{project}_1.0-1_all.deb"

Dir.chdir(File.dirname(__FILE__))

system("docker build -t #{project} .")
system("docker run -it #{project}")
id = `docker ps -l -q`.strip
system("docker cp '#{id}:/src/#{package}' .")
system("docker rm #{id}")
