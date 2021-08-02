<p align="center"><img src="static/logo.png" width="100" height="100"/></p>

# peep
A work-in-progress CLI tool to search for available domain names in the list of supported TLDs.

## Usage
#### Installation
```bash
> go install github.com/hum/peep
```
#### Run
You can either include your own `domain_list_file` or the [default list](https://github.com/hum/peep/blob/main/domains.txt) stored in this repo will be used.
```bash
> peep -n [domain_name] -f [domain_list_file]
```

#### Example
```bash
> peep -n test                     # Will fetch from this git repo
> peep -n test -f domain_list.txt  # Will use local list
```
