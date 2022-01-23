# makemine

Write user information into a global location for desktop linux computers. Desktop owner information is often baked into parts of the desktop configuration. It's often convenient to have a local, central source of truth for the the user's full name, local user account email adddress. makemine takes the user info from a few possible sources and writes it to JSON, YAML and source files in /etc/makefile/[makefile.json, makefile,yaml, makefile.sh].  It's intended to be as simple as possible

NOTE: because it writes to etc, it has to be run with root privileges


## Usage

If you just built your computer and you're logged in with the default root user, run makemine to be prompted for your data
```shell
sudo makemine
Full Name (ex. Firstly Lastly):
Firstly Lastly
local computer user account(ex. flastly):
flastly
Email address (ex. flastly@somedomain.com):
flastly@somedomain.com
# the useradd/passwd line is printed by makemine for your convenience
sudo useradd -m -d /home/flastly -s /bin/bash -g sudo flastly && sudo passwd flastly
# to see what it did:
❯ cat /etc/makemine/makemine.yaml
full_name: Firstly Lastly
local_user: flastly
email: flastly@somedomain.com
❯ cat /etc/makemine/makemine.json
{
 "fullName": "Firstly Lastly",
 "localUser": "flastly",
 "email": "flastly@somedomain.com"
}%
❯ cat /etc/makemine/makemine.sh
export FULLNAME="Firstly Lastly"
export LOCALUSER="flastly"
export EMAIL="flastly@somedomain.com"
```
If you do this kind of thing often, you can store your data in a local json file or  public, url-accessible json file and makemine will iterate through your arguments to find it  and use it rather than prompting you. It tries every argument as a file path, then as a URL. If none of them work, it will prompt.  I've used github gists for this in testing, but anything should work.  The debug option below shows the general flow


```shell
sudo build/current/darwin/amd64/makemine -debug ggg hhh
Password:
{"level":"debug","version":"0.0.2","time":"2021-11-28T13:19:27-05:00","message":"version: 0.0.2"}
{"level":"debug","version":"0.0.2","time":"2021-11-28T13:19:27-05:00","message":"Unable to get MyData from file: ggg"}
{"level":"debug","version":"0.0.2","time":"2021-11-28T13:19:27-05:00","message":"Unable to get MyData from url: ggg"}
{"level":"debug","version":"0.0.2","time":"2021-11-28T13:19:27-05:00","message":"Unable to get MyData from file: hhh"}
{"level":"debug","version":"0.0.2","time":"2021-11-28T13:19:27-05:00","message":"Unable to get MyData from url: hhh"}
Full Name (ex. Firstly Lastly):
rrr ttt
local computer user account(ex. flastly):
rttt
Email address (ex. flastly@somedomain.com):
rttt@gmail.com
sudo useradd -m -d /home/rttt -s /bin/bash -g sudo rttt && passwd rttt
❯ sudo build/current/darwin/amd64/makemine ggg hhh
Full Name (ex. Firstly Lastly):
ttt yyy
local computer user account(ex. flastly):
tyyy
Email address (ex. flastly@somedomain.com):
tyyy@google.com
sudo useradd -m -d /home/tyyy -s /bin/bash -g sudo tyyy && passwd tyyy
```