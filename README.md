[![Build Status](https://travis-ci.org/tibbots/discord-gaming-bot-go.svg?branch=develop)](https://travis-ci.org/tibbots/discord-gaming-bot-go)
[![tibbot/discord-gaming-bot-go](https://img.shields.io/docker/pulls/tibbot/discord-gaming-bot-go.svg)](https://hub.docker.com/r/tibbot/discord-gaming-bot-go/)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)]()

# discord-gaming-bot-go
Discord Bot to share your Gaming-Accounts with others

# Usage
## Creating a Profile
In order to share your different gaming account data for Steam, Origin and other platforms, you need to create a Profile first. You may request a new profile by typing `create profile`, either by directly talking to the bot in a channel 
```
@GamingBot create profile
```
or within a private message without mentioning it:
```
create profile
```

After the profile has been successfully created, the Bot will tell you about your next Steps.

## Deleting a Profile
If you want to delete all your data, simply say 
```
@GamingBot delete profile
```
or within a private message without a mention:
```
delete profile
```

## Add Account data
After creating a profile you will be able to add your account data of the supported platforms, that are:
* Steam 
* Origin 
* Uplay
* Battlenet
* Microsoft (xbox)
* Playstation Network (psn)

Adding an account is again done by talking to the bot directly in a channel or via private message in the form:
```
@GamingBot add account [provider] [account-data]
```

For example:
```
@GamingBot add account steam steam-username/id
@GamingBot add account uplay uplay-username/id
@GamingBot add account origin origin-username/id
@GamingBot add account xbox xbox-username/id
@GamingBot add account psn psn-username/id
```

## Changing Account data
To change your data simply run 'add account' command again with the new data. It will be overridden immediatley

## Show Profile
To view your profile simply type 
```
@GamingBot show profile
```
If you want to inspect a profile of another user, simply mention him/her:
```
@GamingBot show profile @YourFriend
```

## Show Help
```
@GamingBot help
```
```
@GamingBot rtfm
```

## Show deployment information
```
@GamingBot version
```

