#🧾 Step 1: Write a Simple Go App
#📁 Step 2: Create a Project Directory
Organize your files like this:
/my-golang-app
│
├── main.go
└── Dockerfile
#🔧 Fix: Initialize a go.mod file
Before building your Docker image, make sure you initialize a Go module in your app directory:

📍 If you're building outside Docker (recommended before containerizing):
Run this in your my-golang-app folder:
command - go mod init my-golang-app
This creates a go.mod file, which Go needs to manage dependencies and build correctly.

After that, your project should look like:
/my-golang-app
├── main.go
├── go.mod
└── Dockerfile

🔁 Now Rebuild the Docker Image

docker build -t golang-app .
Then run it:

docker run -d -p 8081:8081 golang-app
