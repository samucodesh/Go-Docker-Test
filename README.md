## Go-Docker-Test

**Description:**

This project is a simple Go application that demonstrates how to build and run a Docker image. It's a great starting point for learning how to containerize Go applications.

**Prerequisites:**

* **Git:** To clone the repository and manage version control.
* **Docker:** To build and run the Docker image.

**Cloning the Repository:**

1. **Initialize a new local repository:**
   ```bash
   git init
   ```
2. **Add a remote:**
   ```bash
   git remote add origin git@github.com:samucodesh/Go-Docker-Test.git
   ```
> [!NOTE]
> Replace `git@github.com:samucodesh/Go-Docker-Test.git` with your ssh link

3. **Clone the repository:**
   ```bash
   git pull origin main
   ```

**Building the Image:**

Navigate to the root directory of the project and run the following command:

```bash
docker build -t mygoapp .
```

This will build the Docker image from the Dockerfile in the current directory and tag it as `mygoapp`.

**Running the Image:**

To run the image in a container, use the following command:

```bash
docker run -p 8080:8080 mygoapp
```

This will start a container from the `mygoapp` image and map port 8080 of the container to port 8080 of your local machine. You can access your application at `http://localhost:8080`.
