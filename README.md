# Golang GAE endpoints for Facebook Messenger Bot  

This is a simple example of server that echoes messages.

Follow these steps for configuring your local environment:

1. Install [Go](https://golang.org/dl/)
2. Install [Google App Engine SDK](https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go)
3. Clone repository
4. Install libraries

    ```sh
    $ goapp get google.golang.org/appengine
    $ goapp get github.com/gorilla/mux
    ```

5. Rename file `config.go.template` into `config.go` and fill it with your configuration.
6. Rename file `app.yaml.template` into `app.yaml` and fill it with your GAE project id.

6. Run Application. This is useful for checking eventual errors before deploying, but Facebook Messenger bots sadly work just in production (no localhost, unless you use solutions like [ngrok](https://ngrok.com/)).

    ```sh
    goapp serve
    ```

7. Deploy Application

    ```sh
    goapp deploy
    ```
