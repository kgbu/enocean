# Thanks for inspiration
# - https://gist.github.com/yanatan16/2951128
# - https://github.com/mackerelio/mackerel-agent/blob/master/Makefile

# Followings are ommited from 'deps' because of need of sudo 
# sudo gox -build-toolchain
# sudo gem install fpm

FPM = fpm
REPO="github.com/kgbu/enocean/cmd/enocean"
enoceanTEST_LIST := tests

TAG=0.0.1
ARTIFACTS=downloads
BUILDDIR=build
LDFLAGS=-ldflags "-X main.version `git describe --tags --always`"

ALL_LIST = $(TEST_LIST)

GOXPATH=${GOPATH}/bin/gox
SUDOPATH=${PATH}

all: build test raspi raspi2 edison armadillo

fmt:
	go fmt ./...

doc: fmt
	go fmt ./...
#	golint ./...
	go build github.com/kgbu/enocean/cmd/enocean
	godoc github.com/kgbu/enocean

build: deps doc
	go fmt ./...
#	golint ./...
	go build
	go build $(LDFLAGS) github.com/kgbu/enocean/samples/mqttworker

###### target by architecture

arm5: deps
	sudo GOARM=5 PATH=${SUDOPATH} ${GOXPATH} -build-toolchain -osarch=linux/arm
	GOARM=5 gox $(LDFLAGS) -os="linux" -arch="arm" -output=$(BUILDDIR)/arm5/enocean/enocean-gw $(REPO)
	cd $(BUILDDIR)/arm5/ && tar zcvf ../../$(ARTIFACTS)/enocean-gw_$(TAG)_arm5.tar.gz enocean
	echo 'linux arm5 build completed'

arm6: deps
	sudo GOARM=6 PATH=${SUDOPATH} ${GOXPATH} -build-toolchain -osarch=linux/arm
	GOARM=6 gox $(LDFLAGS) -os="linux" -arch="arm" -output=$(BUILDDIR)/arm6/enocean/enocean-gw $(REPO)
	cd $(BUILDDIR)/arm6/ && tar zcvf ../../$(ARTIFACTS)/enocean-gw_$(TAG)_arm6.tar.gz enocean
	echo 'linux arm6 build completed'

arm7: deps
	sudo GOARM=7 PATH=${SUDOPATH} ${GOXPATH} -build-toolchain -osarch=linux/arm
	GOARM=7 gox $(LDFLAGS) -os="linux" -arch="arm" -output=$(BUILDDIR)/arm7/enocean/enocean-gw $(REPO)
	cd $(BUILDDIR)/arm7/ && tar zcvf ../../$(ARTIFACTS)/enocean-gw_$(TAG)_arm7.tar.gz enocean
	echo 'linux arm7 build completed'

linux_386: deps
	sudo PATH=${SUDOPATH} ${GOXPATH} -build-toolchain -osarch=linux/386
	gox $(LDFLAGS) -os="linux" -arch="386" -output=$(BUILDDIR)/linux_386/enocean/enocean-gw $(REPO)
	cd $(BUILDDIR)/linux_386/ && tar zcvf ../../$(ARTIFACTS)/enocean-gw_$(TAG)_linux_386.tar.gz enocean
	echo 'linux 386 build completed'

linux_amd64: deps
	sudo PATH=${SUDOPATH} ${GOXPATH} -build-toolchain -osarch=linux/amd64
	gox $(LDFLAGS) -os="linux" -arch="amd64" -output=$(BUILDDIR)/linux_amd64/enocean/enocean-gw $(REPO)
	cd $(BUILDDIR)/linux_amd64/ && tar zcvf ../../$(ARTIFACTS)/enocean-gw_$(TAG)_linux_amd64.tar.gz enocean
	echo 'linux amd64 build completed'


###### packaging

raspi: deps arm6
	gem install fpm
	if [ -d $(BUILDDIR)/packages_raspi ] ; then rm -rf $(BUILDDIR)/packages_raspi ; fi
	mkdir -p $(BUILDDIR)/packages_raspi/usr/local/bin
	mkdir -p $(BUILDDIR)/packages_raspi/etc/enocean-gw
	mkdir -p $(BUILDDIR)/packages_raspi/etc/init.d
	cp -p $(BUILDDIR)/arm6/enocean/enocean-gw $(BUILDDIR)/packages_raspi/usr/local/bin/enocean-gw
	cp -p packages/config.simple.ini.example $(BUILDDIR)/packages_raspi/etc/enocean-gw/config.ini
	cp -p packages/enocean-gw.init $(BUILDDIR)/packages_raspi/etc/init.d/
	cd $(BUILDDIR)/packages_raspi; $(FPM) -s dir -t deb -a armhf -n enocean-gw -v $(TAG) -p ../../$(ARTIFACTS) --deb-init ./etc/init.d/enocean-gw.init .
	mv $(ARTIFACTS)/enocean-gw_$(TAG)_armhf.deb $(ARTIFACTS)/enocean-gw_$(TAG)_raspi_arm6.deb

raspi2: deps arm7
	gem install fpm
	if [ -d $(BUILDDIR)/packages_raspi2 ] ; then rm -rf $(BUILDDIR)/packages_raspi2 ; fi
	mkdir -p $(BUILDDIR)/packages_raspi2/usr/local/bin
	mkdir -p $(BUILDDIR)/packages_raspi2/etc/enocean-gw
	cp -p $(BUILDDIR)/arm7/enocean/enocean-gw $(BUILDDIR)/packages_raspi2/usr/local/bin/enocean-gw
	cp -p packages/config.simple.ini.example $(BUILDDIR)/packages_raspi2/etc/enocean-gw/config.ini
	cd $(BUILDDIR)/packages_raspi2; $(FPM) -s dir -t deb -a armhf -n enocean-gw -v $(TAG) -p ../../$(ARTIFACTS) .
	mv $(ARTIFACTS)/enocean-gw_$(TAG)_armhf.deb $(ARTIFACTS)/enocean-gw_$(TAG)_raspi2_arm7.deb

armadillo: deps arm5
	gem install fpm
	if [ -d $(BUILDDIR)/packages_armadillo ] ; then rm -rf $(BUILDDIR)/packages_armadillo ; fi
	mkdir -p $(BUILDDIR)/packages_armadillo/usr/local/bin
	mkdir -p $(BUILDDIR)/packages_armadillo/etc/enocean-gw
	cp -p $(BUILDDIR)/arm5/enocean/enocean-gw $(BUILDDIR)/packages_armadillo/usr/local/bin/enocean-gw
	cp -p packages/config.simple.ini.example $(BUILDDIR)/packages_armadillo/etc/enocean-gw/config.ini
	cd $(BUILDDIR)/packages_armadillo; $(FPM) -s dir -t deb -a armle -n enocean-gw -v $(TAG) -p .. .

edison: linux_386
	if [ ! -d opkg-utils ] ; then git clone http://git.yoctoproject.org/git/opkg-utils ; fi
	if [ -d $(BUILDDIR)/packages_edison ] ; then rm -rf $(BUILDDIR)/packages_edison ; fi
	mkdir -p $(BUILDDIR)/packages_edison/usr/local/bin
	mkdir -p $(BUILDDIR)/packages_edison/etc/enocean-gw

	cp -p $(BUILDDIR)/linux_386/enocean/enocean-gw $(BUILDDIR)/packages_edison/usr/local/bin/enocean-gw
	cp -p packages/config.simple.ini.example $(BUILDDIR)/packages_edison/etc/enocean-gw/config.ini

	mkdir -p $(BUILDDIR)/packages_edison/CONTROL
	sed -i -e 's/FUJI_GIT_TAG/$(TAG)/' packages/opkg_files/control
	cp packages/opkg_files/control $(BUILDDIR)/packages_edison/CONTROL
	cd packages/opkg_files && tar czf control.tar.gz control
	cd $(BUILDDIR)/packages_edison/ && sudo ../../opkg-utils/opkg-build -o root -g root . /tmp
	mv /tmp/enocean-gw_$(TAG)_edison.ipk $(ARTIFACTS)/enocean-gw_$(TAG)_edison_386.ipk

linux_amd64_deb: linux_amd64
	gem install fpm
	if [ -d $(BUILDDIR)/packages_linux_amd64 ] ; then rm -rf $(BUILDDIR)/packages_linux_amd64 ; fi
	mkdir -p $(BUILDDIR)/packages_linux_amd64/usr/local/bin
	mkdir -p $(BUILDDIR)/packages_linux_amd64/etc/enocean-gw
	cp -p $(BUILDDIR)/linux_amd64/enocean/enocean-gw $(BUILDDIR)/packages_linux_amd64/usr/local/bin/enocean-gw
	cp -p packages/config.simple.ini.example $(BUILDDIR)/packages_linux_amd64/etc/enocean-gw/config.ini
	cd $(BUILDDIR)/packages_linux_amd64; $(FPM) -s dir -t deb -a x86_64 -n enocean-gw -v $(TAG) -p ../../$(ARTIFACTS) .


test: $(ALL_LIST)
	go get golang.org/x/tools/cmd/cover
	go test -coverpkg github.com/kgbu/enocean ./...

deps:
	if [ -d $(ARTIFACTS) ] ; then rm -rf $(ARTIFACTS) ; fi
	if [ -d $(BUILDDIR) ] ; then rm -rf $(BUILDDIR) ; fi
	mkdir -p $(ARTIFACTS)
	mkdir -p $(BUILDDIR)

	go get -d -v -t ./...
	go get github.com/mitchellh/gox
	go get github.com/kr/pty
#	go get -u github.com/golang/lint/golint
	godep restore
