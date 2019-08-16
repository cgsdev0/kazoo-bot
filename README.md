## What is this?

it's the demo

## How to use

1. Navigate to Kazoo
2. Sign in
3. Grab your `_session_id` cookie
4. Run `./scrape-users [session-id]`
5. Run either `./find-users` or `./kazoo`

## Examples

```
$ ./kazoo "thanks to all these great peoples" 2 `./names-to-ids` $session

$ ./kazoo "also shoutout to these github users" 2 `./usernames-to-ids` $session
```
