# Run on Commit 🏃

### What is this? 🤔
Run-on Commit is a fast, compiled and minimal server-side program that runs a script of your choice when there's a new commit to a GitHub public repository.

### Lore and Purpose 🤠 
What motivated me to build this repository is the deployment and self-update of my website. I use this program with crontab to self-update my docker containers that run instances of my web [portfolio](https://github.com/joaoofreitas/vue-portfolio).

But this program can have other purposes such has run your scripts every time there's a new commit in a repository.

Also, this is my first project in [Go](https://golang.org/).

### How was it build && how does it work? 🔨

  This program is built using [Go](https://golang.org/). It uses the [GitHub API](https://api.github.com) to check the latest commits of a user repository noticing changes. Afterwards, it automatically runs a script of your choice.

### Instalation 📝
 1. Open your terminal
 2. Clone the repository

 `git clone https://github.com/joaoofreitas/run-on-commit.git`

 3. Enter the directory

 `cd run-on-commit`

 4. Give permissions to the executable file.

 `chmod +x main`

### Usage 🖥️ 
  For __help__ you can run: 

  `./main -h`

  This program uses flags to define, arguments such as the location of the script you want to run, the GitHub username and the repository.

  __Example:__

  `./main -mem <shafile> -file <nameOfTheScript> -user <githubUsername> -repo <nameOfRepository>`

  __Raw Example__
  
  `./main -mem sha -file script.sh -user joaoofreitas -repo markdown-webserver`

### Maintenance and Future Developement ℹ️

  If you have any concerns, errors, or if you want more features feel free to open an issue, or even contact me on [Twitter](https://twitter.com/joaoofreitas_).

### License 📜 

MIT License

Copyright (c) 2020  João Freitas

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
