---
title: "Host your own password manager with Vaultwarden and Nginx"
description: "Vaultwarden comes with all applications and browser extensions you need, it's simple to install, so why not just using it?"
slug: "host-your-own-password-manager-with-vaultwarden-and-nginx"
date: "January 30, 2025"
id: 2
---

![vaultwarder-login-page](/images/vaultwarder-login-page.png)

## Introduction

A few years ago, I used Nextcloud a lot, combined with its [Passwords](https://apps.nextcloud.com/apps/passwords) app. I can't say it was a bad user experience, but it wasn't a good one either. Later, I dropped Nextcloud because I realized that I didn't need it that much, so I no longer had a password manager (yes, I know, it's bad, but who cares?).

Now that I'm in a new phase of trying and hosting things, I wanted to give [Vaultwarden](https://github.com/dani-garcia/vaultwarden) a try. And you know what? I loved it so much that I started writing this little blog post about it, because it deserves it, and I really want you to give it a chance.

I must admit that I'm not a big fan of containers. While I can't say I like Podman, I don't dislike it as much as Docker. So when I read Vaultwarden's installation documentation and saw "You can replace Docker with Podman if you prefer to use Podman." it really appealed to me.

Well, enough talking! Let's get into what we like...

## Requirements

Oh yes, big announcement: I'm no longer into Raspberry Pi since my last post. So here's the setup I have and will assume you have for the technical parts. The versions don't have to be exactly the same, we don't really care, and anyway, if it causes an issue, you'll find out soon enough.

Note that at the time of writing, the current version of Vaultwarden is `1.33.0`.

- Debian 12
- Nginx 1.22.1 (we'll be using a reverse proxy)
- Certbot 2.1.0 (because in HTTPS we trust — though you probably shouldn't). While I won’t explain this step, I still recommend setting it up on your own.
- Podman 4.3.1 (don't forget that you can easily replace it by Docker)

## Run Vaultwarden using Podman

We will be using the `/vw-data` directory, so make sure to create it first and ensure you have the proper user permissions on it (root or whatever works for you). Feel free to adapt the command below.

```bash
podman run --detach \
    --name vaultwarden \
    --env DOMAIN="https://sub.domain.tld" \
    --env ADMIN_TOKEN="SOME_RANDOM_STRING" \
    --volume /vw-data/:/data/ \
    --restart unless-stopped \
    --publish 127.0.0.1:8080:80 \
    docker.io/vaultwarden/server:latest
```

- Don't worry, thanks to the `--restart unless-stopped` parameter, it will always be up, even after a reboot.
- The `ADMIN_TOKEN` will allow you to access the admin backend at `https://sub.domain.tld/admin`.

## Nginx reverse proxy

> This is an HTTP configuration. If you plan to expose it to the web, you should really add an HTTPS configuration.

```
server {
    listen 80;
    listen [::]:80;
    
    server_name sub.domain.tld;

    location / {
        proxy_pass http://127.0.0.1:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_redirect off;
    }
}
```

Add this configuration in your `/etc/nginx/sites-available` directory then enable it:

```
sudo nano /etc/nginx/sites-available/sub.domain.tld
sudo ln -s /etc/nginx/sites-available/sub.domain.tld /etc/nginx/sites-enabled
sudo systemctl restart nginx
```

From here you can try to access to `http://sub.domain.tld/` — it should work.

Don’t forget to check out `http://sub.domain.tld/admin` as well, using your `ADMIN_TOKEN` from the previous step.
