# Integration testing in Go
This repository is the basis for the article: How to do integration test in Go 
(found [here](https://blog.mortenvistisen.com/how-to-do-integration-tests-in-go)).

The article only focus on how to do integration tests and therefore does not
provide any in-depth explanation of the code structure or any way to expose
the functionality to the outside world.

You are free to use this code in anyway you see fit.

## Requirements
This setup is built with macOS/linux in mind - if you are on windows, for some
reason, google is your friend. However, if you've powershell installed and
setup you should be able to run the make commands by adding it through choco.

Here's what you'll need:
- Docker

When running the testing command (make run-integration-tests) you will auto-
matically build the image containing everything you need.


