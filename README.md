# WebApp-Full-duplex
GoLang Gorilla Websockets, JavaScript and HTML

<h3>Steps to host on DigitalOcean (Run following commands on Terminal/Powershell:</h3>


1. cd ~\.ssh
2. ssh-keygen -t rsa (give the name to store ssh key, ex: new_key)
3. Public key will be stored in new_key.pub and Private key will be store in new_key
4. Get your public key using command: cat .\new_key.pub
5. Create droplet (virtual machine) on DigitalOcean, and Use SSH keys a mode of authentication, and copy your Public key
6. Wait a moment for droplet to initialize, and copy ip address and run following command to ssh into the system: ssh -i ~\.ssh\new_key root@ip_address_of_droplet

7. Once ssh is sucessfull: run following commands:
	* Clone your epository using git clone 
		* (or you can use scp to copy your files: scp -i ~/.ssh/new_key file_name_to_copy root@ip_addr:)
	* sudo apt update
	* sudo apt install golang
	* go build -o app_name
8. Viola, Your app is deployed, head over to ip_address of droplet and start listening to specified port number


Refer this video for more details: https://www.youtube.com/watch?v=z5rRuEJvVlM&t=495s
