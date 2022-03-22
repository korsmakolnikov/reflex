# Reflex
Reflex is just a bad copy of netcat. Its purpose is to create a remote shell on 
the host where it runs. 

# Usage 
```bash
> reflex 20080 # runs reflex and bind to port 20080
```

On another shell you can now
```bash 
> telnet host_address 20080
...
ls # this will execute ls via reflex on the host and will returns stdout here
...
```
