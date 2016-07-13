Vagrant.configure("2") do |config|
  config.vm.define "server" do |server|
    server.vm.box = "ubuntu/precise64"
    server.vm.hostname = 'server'
    server.vm.box_url = "ubuntu/precise64"
    server.vm.provision :shell, path: "bootstrap-all.sh", env: {GOPATH: "/home/vagrant/godev"}

    server.vm.network :private_network, ip: "192.168.56.101"

    server.vm.provider :virtualbox do |v|
      v.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
      v.customize ["modifyvm", :id, "--memory", 512]
      v.customize ["modifyvm", :id, "--name", "server"]
    end
  end

  config.vm.define "client" do |client|
    client.vm.box = "ubuntu/precise64"
    client.vm.hostname = 'client'
    client.vm.box_url = "ubuntu/precise64"
    client.vm.provision :shell, path: "bootstrap-all.sh", env: {GOPATH: "/home/vagrant/godev"}

    client.vm.network :private_network, ip: "192.168.56.102"

    client.vm.provider :virtualbox do |v|
      v.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
      v.customize ["modifyvm", :id, "--memory", 512]
      v.customize ["modifyvm", :id, "--name", "client"]
    end
  end
end