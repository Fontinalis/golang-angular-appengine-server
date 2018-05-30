# Sample golang server for App Engine to serve Angular apps

### Setup v1
If you want to have the server in the root of your angular app, just copy the `gae` folder next to your `.angular-cli.json`.

In your `.angular-cli.json`, just change the `outDir` to `./gae/dist`, then you can run your `gcloud app deploy` command.

### Setup v2
In case you want to have the Google App Engine server as a separate repo, just copy your `dist` folder's content into `gae/dist`. 
