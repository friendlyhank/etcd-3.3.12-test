package etcdtest

import (
	"hank.com/etcd-3.3.12-annotated/etcdmain"
	"os"
	"testing"
)

//启动测试服务1
func StartInfraOServer(){
	os.Args = []string{"etcd-3.3.12-test","--name","infraO",
		"--initial-advertise-peer-urls","http://127.0.0.1:2380",
		"--listen-peer-urls","http://127.0.0.1:2380",
		"--listen-client-urls","http://127.0.0.1:2379",
		"--advertise-client-urls","http://127.0.0.1:2379",
		"--initial-cluster-token","etcd-cluster-1",
		"--initial-cluster","infraO=http://127.0.0.1:2380,infral=http://127.0.0.1:2382,infra2=http://127.0.0.1:2384",
		"--initial-cluster-state","new"}
	etcdmain.Main()//服务端主入口
}

func StartInfra1Server(){
	os.Args = []string{"etcd-3.3.12-test","--name","infra1",
		"--initial-advertise-peer-urls","http://127.0.0.1:2382",
		"--listen-peer-urls","http://127.0.0.1:2382",
		"--listen-client-urls","http://127.0.0.1:2381",
		"--advertise-client-urls","http://127.0.0.1:2381",
		"--initial-cluster-token","etcd-cluster-1",
		"--initial-cluster","infraO=http://127.0.0.1:2380,infra1=http://127.0.0.1:2382,infra2=http://127.0.0.1:2384",
		"--initial-cluster-state","new"}
	etcdmain.Main()//服务端主入口
}

func StartInfra2Server(){
	os.Args = []string{"etcd-3.3.12-test","--name","infra2",
		"--initial-advertise-peer-urls","http://127.0.0.1:2384",
		"--listen-peer-urls","http://127.0.0.1:2384",
		"--listen-client-urls","http://127.0.0.1:2383",
		"--advertise-client-urls","http://127.0.0.1:2383",
		"--initial-cluster-token","etcd-cluster-1",
		"--initial-cluster","infraO=http://127.0.0.1:2380,infral=http://127.0.0.1:2382,infra2=http://127.0.0.1:2384",
		"--initial-cluster-state","new"}
	etcdmain.Main()//服务端主入口
}

func TestInfraOEtcdMain(t *testing.T){
	StartInfraOServer()
}

func TestInfra1EtcdMain(t *testing.T){
	StartInfra1Server()
}

func TestInfra2EtcdMain(t *testing.T){
	StartInfra2Server()
}
