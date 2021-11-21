require "sinatra"

set :bind, "0.0.0.0"
port = "8080"
set :port, port

get "/" do
  "Hello world!"
end