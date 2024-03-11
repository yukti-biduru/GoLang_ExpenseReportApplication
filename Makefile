export UDIR= .
export GOC = x86_64-xen-ethos-6g
export GOL = x86_64-xen-ethos-6l
export ETN2GO = etn2go
export ET2G   = et2g
export EG2GO  = eg2go

export GOARCH = amd64
export TARGET_ARCH = x86_64
export GOETHOSINCLUDE=ethos
export GOLINUXINCLUDE=linux
export BUILD=ethos

export ETHOSROOT=client/rootfs
export MINIMALTDROOT=client/minimaltdfs


.PHONY: all install clean
all: ExpenseClient ExpenseServer

ethos:
	mkdir ethos
	cp -pr /usr/lib64/go/pkg/ethos_$(GOARCH)/* ethos

myRpc.go: myRpc.t
	$(ETN2GO) . myRpc $^

myRpc.goo.ethos : myRpc.go ethos
	ethosGoPackage  myRpc ethos myRpc.go

ExpenseServer: ExpenseServer.go myRpc.goo.ethos
	ethosGo ExpenseServer.go

ExpenseClient: ExpenseClient.go myRpc.goo.ethos
	ethosGo ExpenseClient.go

# install types, service,
install: all
	sudo rm -rf client
	(ethosParams client && cd client && ethosMinimaltdBuilder)
	ethosTypeInstall myRpc
	ethosDirCreate $(ETHOSROOT)/services/myRpc   $(ETHOSROOT)/types/spec/myRpc/MyRpc all
	install -D  ExpenseClient ExpenseServer                   $(ETHOSROOT)/programs
	ethosStringEncode /programs/ExpenseServer    > $(ETHOSROOT)/etc/init/services/ExpenseServer
	ethosStringEncode /programs/ExpenseClient       > $(ETHOSROOT)/etc/init/services/ExpenseClient

# remove build artifacts
clean:
	rm -rf myRpc/ myRpcIndex/ ethos clent
	rm -f myRpc.go
	rm -f ExpenseClient
	rm -f ExpenseServer
	rm -f myRpc.goo.ethos
