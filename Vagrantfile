# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/bionic64"

  config.vm.synced_folder "../", "/home/vagrant/go/src/"

  config.vm.network "forwarded_port", guest: 15672, host: 15672

  config.vm.provision "shell", inline: <<-SHELL
    sudo apt-get update
    sudo apt-get -y upgrade
    sudo apt-get install -y build-essential make golang-go 
    
    sudo chmod -R 777 /home/vagrant/go

    echo "export PATH=$(go env GOPATH)/bin:$(go env GOROOT)/bin:$PATH" >> ~/.bashrc

    echo "deb http://www.rabbitmq.com/debian/ testing main"  | sudo tee  /etc/apt/sources.list.d/rabbitmq.list > /dev/null
    wget https://www.rabbitmq.com/rabbitmq-signing-key-public.asc
    sudo apt-key add rabbitmq-signing-key-public.asc
    sudo apt-get update
    sudo apt-get install rabbitmq-server -y
    sudo service rabbitmq-server start
    sudo rabbitmq-plugins enable rabbitmq_management
    sudo service rabbitmq-server restart

    rabbitmqctl add_user test test
    rabbitmqctl set_user_tags test administrator
    rabbitmqctl set_permissions -p / test ".*" ".*" ".*"
  SHELL
end
