# COMPANIES JSON Generator

Generates arbitrary amount of fake companies JSON information, useful for tests where big datasets are required (e.x: 1GB+ fake data).

_It generates a Million of random fake companies data by default in less than 10 seconds!_

This number was chosen to avoid having RAM problems (getting out of RAM) in a 16GB RAM machine and get the result really fast

## Example usage: 

First execute the following command from a shell (replace the output path with an appropriate path)

```
▶ go run src/main.go /output/path/goes/here/companies.json
```

After that you'll be prompted for the num of companies to generate, enter that number and you'll have the companies generated in the output path you specified before
