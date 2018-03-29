require 'rake'

desc "Compile and move binary to release path"
task :compile do
  sh 'docker run --rm -v "$(pwd)":/usr/src/github.com/cloudgear-images/go-app -w /usr/src/github.com/cloudgear-images/go-app golang:1.10-alpine /bin/sh -c "apk add --no-cache git && go get github.com/go-martini/martini && go get github.com/martini-contrib/render && go build -v"'
end

desc "Create Docker image"
task :build, [:version] do |t, args|
  version = args[:version]
  if version.nil?
    puts "ERROR: not version defined (e.g. rake build[0.3])"
    exit 1
  end

  sh "docker build -t cloudgear-images/go-app:#{version} ."
end
