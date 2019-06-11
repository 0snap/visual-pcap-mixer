Visual PCAP Mixer
=================

## What this is

- A highly experimental thing that I need for myself
- Untested code, potentially dangerous

*YOU SHOULD NOT USE THIS UNLESS YOU KNOW WHAT YOU DO*

You alone are responsible for using this tool, I do not take any responsibility for any kind of harm that it may cause.


## Functions

- A go cli wrapper around existing tools that can analyze and modify pcaps. The tools are directly invoked as `cmd`. They write stuff on your harddrive. Thats why you should not use this.
- A react frontend for the browser to visualize what will happen with the pcaps.

## Config

The backend needs a `config.json`. This thing differentiates attack samples and benign traffic. Example below.

Btw, nice dataset for attack & benign traffic here: https://www.unb.ca/cic/datasets/ids-2018.html 

```
{
    "groundtruth": [
        {
            "files": [
                "/home/you/pcaps/unbca/attacks/28-02-2018/capEC2AMAZ-O4EL3NG-172.31.69.24-part1",
                "/home/you/pcaps/unbca/attacks/28-02-2018/capEC2AMAZ-O4EL3NG-172.31.69.24-part2"
            ],
            "attacks": [
                {
                    "attackers": [
                        "13.58.225.34"
                    ],
                    "victims": [
                        "172.31.69.24"
                    ],
                    "name": "Infiltration",
                    "start": "2018-02-28T10:50:00-04:00",
                    "end": "2018-02-28T12:05:00-04:00"
                },
                {
                    "attackers": [
                        "13.58.225.34"
                    ],
                    "victims": [
                        "172.31.69.24"
                    ],
                    "name": "Infiltration",
                    "start": "2018-02-28T13:42:00-04:00",
                    "end": "2018-02-28T14:40:00-04:00"
                }
            ]
        }
    ],
    "unclassifiedTraffic": [
        "/home/you/pcaps/unbca/benign/22-02-2018",
        "/home/you/pcaps/unbca/benign/28-02-2018"
    ],
    "outPath": "/home/you/pcap/apt-scenarios"
}
```


## When you really really want to use this

- add the `backend` folder to your go path
- build your own `config.json` file like above
- check the help menu `go run main.go`

You *must* first run a deep analysis over the configured files. Export the analysis results to a state file:

    $ go run main.go export -e your_state.json

Grab a coffee in case you have several hundred gigs of traffic (as I do) ...

Now take the analysed files and host a server

    $ go run main.go server -s your_state.json

Navigate to the `frontend` folder and fire it up. you need a moderately new version of `npm` / `yarn`:

    $ npm install
    $ npm start

Go to your browser, `localhost:3000`. When you did the config right the browser content looks somewhat like this:

[[https://github.com/0ortmann/visual-pcap-mixer/blob/master/screenshots/configured-contents.png|alt=configured-contents]]

#### In browser use

- create new days of an attack scenario by hitting the big `+`
- move all the stuff per drag n drop (attacks, traffic samples, days in the timeline)
- drag benign and attack traffic to your liking
- hover stuff for more info
- double click stuff to delete it
- you can rewrite IP addresses with the form in the bottom left corner
- name the scenario you created (form in lower right corner)

When you create an attack scenario the following will happen *on your computer*:

- first timestamp is taken from first traffic sample in day 1
- all other pcaps get time-adjusted, that they apprear to have been recorded in order
- IP replacements are applied
- stuff is copied to a new folder in the `outPath` that is configured in the `config.json`

Depending on your traffic samples that may fill your harddrive. again, be careful where you run this. better dont. NEVER HOST THIS ON A PUBLIC SERVER. it gives away cmd. 

[[https://github.com/0ortmann/visual-pcap-mixer/blob/master/screenshots/configured-contents.png|alt=configured-contents]]



## TODO:

- test this shit
- clean up, take out garbage
- I remember vaguely that I built in a silly assumption about filenames in the benign traffic folders. sigh. remove that.