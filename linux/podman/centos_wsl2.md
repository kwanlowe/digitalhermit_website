# Install CentOS on WSL2

This is based on [Import any Linux distribution to use with WSL](https://docs.microsoft.com/en-us/windows/wsl/use-custom-distro?WT.mc_id=windows-c9-niner). I modifed to use Podman.

## Pre-requisites

* Podman installation on WSL2

## Create CentOS WSL Installation

This section is almost identical to the Microsoft documentation, except for using Podman instead of Docker. 

### Export CentOS Image
* Download the Centos container image:

```shell
    podman pull centos:latest
    podman run centos bash ls /
```
* Get the container image id:
Run the following command then pull the image ID. 

```shell
   podman container ls -a
```

This will return something similar to:

```shell
[kwan@XPS-Firehawk podman]$ podman container ls -a
CONTAINER ID  IMAGE                         COMMAND     CREATED        STATUS                    PORTS       NAMES
776b47a3205b  quay.io/centos/centos:latest  /bin/bash   5 minutes ago  Exited (0) 5 minutes ago              centos
```

Copy the IMAGE ID (first column) from the output.

* Export the Image to a tar file:

```shell
    podman export 776b47a3205b >/mnt/c/Users/<Windows Username>/Downloads/centos.tar
```
### Import Centos Image

Open Powershell to do this next section.

* Create a location to store the CentOS WSL image:

   mkdir C:\Users\<Windows Username>\Documents\WSLStorage

For example:

```shell
PS C:\Users\kwanl> mkdir C:\Users\kwanl\Documents\WSLStorage 

Directory: C:\Users\kwanl\Documents
 Mode                 LastWriteTime         Length Name 
 ----                 -------------         ------ ----
 d-----         7/29/2021   4:51 PM                WSLStorage
```

* Import the tar image into the new location:

```
wsl --import CentOS <WSL Storage> <tarball>
```
Where ```<WSL Storage>``` is the path to the storage location for the import and ```<tarball>``` is the centos.tar image created in the first step.

For example:

```shell
PS C:\Users\kwanl> wsl --import CentOS C:\Users\kwanl\Documents\WSLStorage\CentOS .\Downloads\centos.tar
```

### Create Initial User

This section follows [Add WSL specific components](https://docs.microsoft.com/en-us/windows/wsl/use-custom-distro?WT.mc_id=windows-c9-niner#add-wsl-specific-components-like-a-default-user) from the Microsoft guide:

* From Powershell, launch the new environment:

```shell
    wsl -d CentOS
```

* Update the environment:

```shell
    yum update -y && yum install passwd sudo -y 
```

* Add local user and set as default. For example, to add user ```homer```.

```shell
    adduser -G wheel homer
    echo -e "[user]\ndefault=homer" >> /etc/wsl.conf
    passwd homer
```

This will create the ```homer``` user and create an entry in ```/etc/wsl.conf```. The ```wsl.conf``` file is used for several functions further on.

* Once the base user is created, exit the session:

```shell
    exit
```

* Terminate the instance and restart:

```
    wsl --terminate CentOS
    wsl -d CentOS
```
If everything was done correctly, you should be the default user created in these steps.

## Add entry to Windows Terminal

* Open ***Windows Terminal***.
* Open the Windows Terminal Settings:
* Under the Profiles section, click "Add New"
* Under the Name, enter "CentOS"
* Under the Command Line, enter ```wsl -d CentOS```
* Under the Starting Directory, enter ```\\wsl$\CentOS\home\homer```
* Optionally change the icon
* Click Save

## Post-Configuration

### Disable Windows PATH and Binaries

By default, the Windows system PATH is copied to the WSL installation. To disable:

From: [Disable Windows PATH](https://stackoverflow.com/questions/51336147/how-to-remove-the-win10s-path-from-wsl)

Add the following to the /etc/wsl.conf file:

```conf
  [interop]
  enabled=false # enable launch of Windows binaries;
  appendWindowsPath=false # append Windows path to $PATH variable;
```

### Install Useful Packages
At this point, you can launch the new CentOS image as you would any other WSL instance. The installed CentOS is fairly bare, as it was designed for a container image. To make this closer to a default CentOS, you can query the list of base packages:

```shell
    dnf groupinfo "Core" 
```

This will return a list of packages, some of which are "Mandatory". For example:

```shell
Group: Core
 Description: Smallest possible installation
 Mandatory Packages:
   NetworkManager
   audit
   basesystem
   [snip]
```

Of these, I installed a subset of packages. Here are a few recommended additional packages:

* Add packages:
  - ncurses
  - cronie
  - firewalld
  - openssh-clients
  - bash-completion
  - epel-release

These can be installed with ```sudo dnf install <package-name>```.




