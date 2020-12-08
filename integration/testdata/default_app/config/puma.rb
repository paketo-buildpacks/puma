workers 2

threads 1,2

app do |env|
  body = 'Hello World, Puma!'

  [
    200,
    {
      'Content-Type' => 'text/plain',
      'Content-Length' => body.length.to_s
    },
    [body]
  ]
end
