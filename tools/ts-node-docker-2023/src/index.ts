// Why we cast to number? Because process.env.PORT is a string
// Anything that comes from .env is always a string
const port = Number(process.env.PORT) || 3000;


// passing the port as the second argument to app.listen
// its way to doit for docker container
// you dont want to 
app.listen(port,"0.0.0.0" ,() => {
  console.log("`Server is running on http://localhost:${port}`");
});
