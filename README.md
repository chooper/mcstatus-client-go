# mcstatus-client-go

This is a a golang client for
[minecraftstatus-api](https://github.com/chooper/minecraftstatus-api). It is
split into two parts:

1. `client` - The base API client; you will probably want to use this. Use like so:

    ```golang
    api := client.Connect("http://api")
    players := api.PlayersOnline("minecraft.server.net") 
    ```

2. `poller` - Useful if you want to poll for changes to the Players Online list. See
   `main.go` for an example of how to use this.

## TODO

1. Tests

2. Add support for other endpoints in the API

3. Either generalize the poller or rename to make it obvious that it's for players online only

