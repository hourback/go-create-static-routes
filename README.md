go-create-static-routes
=======================

I was using a Windows batch file to add static routes to my Windows 7 workstation.  I decided to implement this in Go and make it more dynamic.

I initially investigated how to do this using Windows APIs instead of command line programs ("cmd", "route") but that seemed too much to bite off initially, particularly since I have no Windows programming experience.  :-)

## Configuration
The mask and gateway IPs are hard-coded in the program.  Also hard-coded are the default domains to look up and include in the routing table.  If you provide your own domains as command line arguments, these default domains are ignored.

## Usage

### Specifying your own domains
```
create-static-routes.exe my-domain.com your-domain.com her-domain.com
```

### Using default domains
If you know you are going to be creating routes for the same domains over and over, you can add them to the code and simply run the program without arguments:
```
create-static-routes.exe
```

