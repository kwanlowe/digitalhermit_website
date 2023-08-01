# privateGPT-container

This is a Podman build of PrivateGPT based on [privateGPT](https://github.com/imartinez/privateGPT).

This podman build is alpha stuff, meaning it barely works. The container builds and on running, has access to 
the GPUs. Basic tests work. Use at your own risk. Only tested on RHEL and CentOS. 

Aditional certs are installed for my odd home environment. Disable this in the Containerfile if you don't need it.

There are also dotfiles in the files/home directory. Load your own here to customize.

## Host Configuration

Follow guide here: [Enable GPUs for Podman](https://www.redhat.com/en/blog/how-use-gpus-containers-bare-metal-rhel-8)

Install make: ``` sudo dnf install make ```

## Build the container:

0. Checkout this repo. ``` git clone <repo>```
0. ```cd digitalhermit_website/linux/privateGPT-container```
0. ```make build```


This creates a really huge  container. If you get an error creating it, you may need a couple other tweaks to systemd (pending).

## Download the model:

Tested so far with the groovy model. 

```
mkdir models
cd models && wget https://gpt4all.io/models/ggml-gpt4all-j-v1.3-groovy.bin

```

## Run the container

Edit the Makefile as needed. 
```
make run
```



## Miscellaneous

Logo created with: ``` echo "PrivateGPT"|figlet -f small |lolcat -S 2  -p 1.4 -F .1 --force >privateGPT-logo ```

