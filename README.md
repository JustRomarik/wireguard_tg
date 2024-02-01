# Wireguard Telegram bot 
---
## For what?
This bot is needed to simplify control over your wireguard server.  
It is needed so that if you need to log into the server, you can add a new user and receive a qr code that you can give to your friend or someone else.

## Functions
- Turn on or off wireguard service
- Add new user and get qr code with user config
- Get your current wireguard config in its entirety

## Quick start 
Step 1: Clone this repository 
```
git clone https://github.com/romus204/wireguard_tg.git
```
Step 2: Fill out the config <kbd>config.yml</kbd> file  
Step 3: Run the <kbd>main</kbd> file or make the necessary changes to the code, compile the program and run

***Note***: The configuration file <kbd>config.yml</kbd> must be in the same directory as the executable file.  

***Note***: The bot must have access to server configuration files and the ability to restart the service, so it must be run with privileged rights. 

## Сonfig description 

<kbd>server_address</kbd> - Address of you server for connection from outside  
<kbd>telegram_token</kbd> - Token received from https://t.me/BotFather  
<kbd>telegram_id_allowed</kbd> - Your telegram ID. Messages from strangers will simply be ignored  
<kbd>wireguard_publickey</kbd> - Public key of you wireguaard server  
<kbd>wireguard_config_path</kbd> - Path to wireguard config. This optional field. Default: /etc/wireguard/wg0.conf  
<kbd>wireguard_service_name</kbd> - Name of wireguard service. This optional field. Default: wg-quick@wg0.service  
<kbd>wireguard_port</kbd> - Port of you wireguard server

## Bot Command

<kbd>/echo</kbd> - Just test echo func.  
<kbd>/serveron</kbd> - Turn ON wireguard server for all users.  
<kbd>/serveroff</kbd> - Turn OFF wireguard server for all users.  
<kbd>/getconfig</kbd> - Allows you to get the current configuration of your server to view the parameters for adding a new user.  
<kbd>/newuser {username} {allowedIPs}</kbd> - Adding a new user. Replace the parameters indicated in curly brackets with the ones you need. After adding a new user, the wireguard service will restart.  

Example:
```
/newuser testuser 10.0.0.9/32
```

## TODO

- [ ] Transition to webhooks from polling.  
- [ ] More convenient procedure for adding a user.
- [ ] Reducing the number of fields in the configuration file.
- [ ] More flexible functionality for editing users (changing name, allowedIPs, endpoints).
- [ ] Include a script for wireguard deployment in the bot.
