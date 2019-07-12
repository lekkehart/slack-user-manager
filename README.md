# slack-user-manager

Tool for manipulating users in Slack workspace

Main purpose was to make some kind of clean up of my Slack workspace:

* Remove titles and phone numbers from users.
* Activate/Deactivate users.

Anyhow, use this as an example for creating your own specific Slack manipulation tools.

## Prerequisites

As input it requires file `list_of_userids.in` listing the IDs of the users that are to be manipulated.

By means of a search query in Slack select create file `list_of_userids.in`.
An example can be found here [list_of_userids.in.example](list_of_userids.in.example).

Furthermore, you need a token for accessing your Slack workspace.

## Run

Run the program by means of:

```bash
go.exe run main.go -token <TOKEN> -v=1
```

### Run With Activation/Deactivation Functionality

I temporarily disabled the activation/deactivation functionality, see `// TODO ENABLE`.

With it enabled the command would be:

```bash
go.exe run main.go -token <TOKEN> -active <FALSE|TRUE> -v=1
```
