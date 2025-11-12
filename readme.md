## How to use
For local, run go run ./cmd/server. Can then run UI/UI.html as a preview on live server and interact with the API locally if you change apibase on line 83 to the appropriate value (should be localhost:8080/api/v1, but might change depending on system). Alternatively, if you don't change the apibase value, it interacts with my API running on Google Cloud Run. 

The API can be interacted with as follows: The API has 2 endpoints, api/v1/calc and api/v1/history. Sending a POST to api/v1/calc with the following json format {Username, Password, Type, Expression} (all string values) will cause the corresponding calculator (depending on type) to evaluate expression, given the username and password combo is correct. Type can take 2 values: "basic" and "rpn". Basic can only take expressions of the form "x op y", whereas RPN can take any expression, as long as it only uses +, -, *, /, % or ^ operations, as well as parentheses. Both calculators can also handle some degree of text substitution, the full list is found in expressionMap.go. Sending a GET to api/v1/history will return the history of both calculators concatenated one after the other. I ran into an issue with how to sync them up properly, so this was the 2nd best solution I had to the issue of multiple histories.

Alternatively, you can visit https://fir-testing-323bd.web.app/, which is a firebase hosted webpage of the simple UI found in UI.html, which will then send HTTP requests to the api hosted on Google Cloud Run.

I will send valid authentication pairs in the mail I send to you. In the repo you will find valid usernames, and hashed valid passwords. Ideally, I wanted a way to store the credentials outside of the repo, but was drawing a blank on how to. My thought was setting up a small database to store username hashed password pairs, but felt that with the time I have left, that would be a bit overkill to begin on doing. 

Tests are found in calculator_test.go

## Issues / notes
I use a regex to help check the validity of the expressions, this is not a fully comprehensive regex. It only checks if the symbols are valid. In an ideal world, it would also check structure of the expression, but I would need to read up a bit more on that (which I will after), but I decided to keep it as is for now, even if it's not perfect, to show I know of regular expressions. 

For the issue of the histories, I was thinking of making use of an observer pattern, but I think I might have confused myself on how to do it, because I couldn't really figure out a good way to do it.

I was thinking of a Strategy pattern to implement in basic.go, to handle the cases rather than a switch statement, but I can't really see a way to make it work that isn't just setting the flag in the current switch statement and then handling the computation afterwards, which kind of defeats the point.