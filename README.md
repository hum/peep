# peep

A super easy to use CLI tool to find available domain TLDs with a specific domain name.

### TODO:
  - [ ] Parse responses from the WHOIS servers
  - [ ] Make the CLI output pretty
  - [ ] Better README
  - [ ] Test inputs

### Usage (WIP)
#### Installation
```bash
> go install github.com/hum/peep
```
#### Run
You can either include your own `domain_list_text_file` or the app will fetch the one present in this repository. (e.g. `domains.txt`)
```bash
> peep -n [domain_name] -f [domain_list_text_file]
```

#### Example
```bash
> peep -n test                     # Will fetch from this git repo
> peep -n test -f domain_list.txt  # Will use local list
```
