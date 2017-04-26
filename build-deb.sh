cd `dirname $0`

mkdir build && cd build
mkdir -p usr/bin etc/init.d
cp ../piriscope usr/bin/piriscope
cp ../init.d etc/init.d/piriscope

fpm --input-type dir \
  --output-type deb \
  --architecture armhf \
  --depends x264 \
  --deb-init ./etc/init.d/piriscope \
  --name piriscope \
  --version $VERSION \
  ./usr/bin/piriscope
