This is a repo that I am using to learn more about how to use Vagrant provisioning to create multiple machines and send messages between them.

In order to build, installing everything needed for Vagrant then running vagrant up should bring up both systems.

After that run the following commands to log into the server and start NSQ:
	
	vagrant ssh server
	nohup nsqlookupd &
	nohup nsqd --lookupd-tcp-address=127.0.0.1:4160 &
	nohup nsqadmin --lookupd-http-address=127.0.0.1:4161 &
