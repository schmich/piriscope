$version = `git tag`.lines.last&.strip || '1.5.0'
if !$version
  puts 'No tagged version, see git tag.'
  exit 1
end

$commit = `git rev-parse HEAD`.strip
