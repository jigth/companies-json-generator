# COMPANIES JSON Generator

Generates arbitrary amount of fake companies JSON information, useful for tests where big datasets are required (e.x: 1GB+ fake data).

_It generates a Million of random fake companies data by default in less than 10 seconds!_

This number was chosen to avoid having RAM problems (getting out of RAM) in a 16GB RAM machine and get the result really fast

## Known limitations:

Because this program is not optimized for low memory footprint using advanced techniques like batch processing it can use a lot of RAM depending on the amount of generated inputs.

Nonetheless it is perfectly viable to generate up to 10 millions of entries (1.6GB) in about 10 seconds using a 32GB RAM machine. Time may vary depending on the used SSD, processor and other computer specs.

## Example usage: 

### Running the program locally

#### Requirements

* Git
* Golang 1.22 installed and available in the path

#### Instructions

1. Get the source code using the following command

```shell
git clone
```

2. Execute the following command from a shell (replace the output path with an appropriate path)

```shell
go run src/main.go /output/path/goes/here/companies.json
```

After that you'll be prompted for the num of companies to generate, enter that number and you'll have the companies generated in the output path you specified before

### Running the program using a container (Docker)

#### Requirements

* Git (if you're going to build the container from source code and run it)
* Docker installed and available in the path
* **(For Docker Desktop Users):** Verify that Docker Desktop is running if you're using it

#### Instructions for running it easily

For convenience a DockerHub image has been published so you don't need to build it, just run it like so

##### Linux and MacOS

```shell
docker run jigth/companies-generator -v ./:/app companies.json 30 # This will create 30 JSON elements in a file called companies-sample.json in your current folder
```

##### Windows 

The only difference is the backslash because Windows uses them for creating folders hierarchies (paths)

```shell
docker run jigth/companies-generator -v .\:/app companies.json 30 # This will create 30 JSON elements in a file called companies-sample.json in your current folder
```

#### Instructions for building the container and running it.

Because this program is used to generate data and then its lifecycle ends a volume will be needed to be able to get the output.

1. First build the image

```shell
docker build -t companies-generator .
```

2. **Run it with the same instructions as the section above *(Instructions for running it easily)***, just change _"jigth/companies-generator"_ to _"companies-generator"_.


## Author

Daniel Ochoa Montes
