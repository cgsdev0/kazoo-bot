## What is this?

don't worry about it

## How to use

1. Navigate to Kazoo
2. Sign in
3. Grab your `_session_id` cookie
4. Run `./setup [session-id]`

## Example Usage

```
$ ./kazoo "thanks to all these great peoples" 2 `./names-to-ids` $session

$ ./kazoo "also shoutout to these github users" 2 `./usernames-to-ids` $session
```

## Github PR watcher

Make a file called `github_api.json` that looks like this:

```json
{
    "token": "YOUR_PERSONAL_GITHUB_TOKEN_HERE"
}
```

run `./setup-pr-approval-watcher`
