Vagrant.configure("2") do |config|
  config.vm.define "server" do |server|
    server.vm.box = "ubuntu/precise64"
    server.vm.hostname = 'server'
    server.vm.box_url = "ubuntu/precise64"
    server.vm.provision :shell, path: "bootstrap.sh"

    server.vm.network :private_network, ip: "192.168.56.101"

    server.vm.provider :virtualbox do |v|
      v.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
      v.customize ["modifyvm", :id, "--memory", 512]
      v.customize ["modifyvm", :id, "--name", "server"]
    end
  end

  config.vm.define "client1" do |client1|
    client1.vm.box = "ubuntu/precise64"
    client1.vm.hostname = 'client1'
    client1.vm.box_url = "ubuntu/precise64"
    client1.vm.provision :shell, path: "bootstrap.sh"

    client1.vm.network :private_network, ip: "192.168.56.102"

    client1.vm.provider :virtualbox do |v|
      v.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
      v.customize ["modifyvm", :id, "--memory", 512]
      v.customize ["modifyvm", :id, "--name", "client1"]
    end
  end
end