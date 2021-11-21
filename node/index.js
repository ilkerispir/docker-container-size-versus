const http = require('http');

const port = 8080;

const server = http.createServer((req, res) => {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'text/plain');
  res.end('Hello world!');
});

server.listen(port, "0.0.0.0", () => {
  console.log(`Server running at :${port}`);
});