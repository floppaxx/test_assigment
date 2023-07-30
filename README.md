# DevOps test assigment
<p>In this test assigment I have made an go application that sends three basic API requests and shows the metrics of all endpoints on the localhost. I have configured Github Actions, that on every push commit the code will be tested, linted and published as a container to DockerHub. On every pull requesrt the "Deployed" will be shown</p>
<br>
<h2>How to run this application</h2>
<p>You can either run it on the the go application itself or run the container</p>
<h3>Running the application itseld</h3>
<p>1. Clone the github repository</p>

```git clone https://github.com/floppaxx/test_assigment.git```

<p>2. Go to the directory and run the application (make sure that go is installed on your machine)</p>

```cd test_assigment```
<br>
```go run test_assigment.go```

<p>3. Open your browser and enter (make sure that the port is open)</p>

 ```http://localhost:8080``` 

 <h3>Running this application as container</h3>
 <p>1. Open terminal and pull the image</p>

 ```docker pull floppax/test_assigment:latest```

 <p>2. Run the docker container (docker image is updated with each push request)</p>

 ```docker run -p 8080:8080 floppax/test_assigment```

 
