## rce-checker

A simple and benign RCE checker written in Go. Red teamers can drop and execute it to validate remote code execution on a Windows/Linux host.    

I created this tool to assist in RCE validation during a security assessment at work. While there are many ways to verify RCE, this is built mainly for having a consistent and purpose-built binary that allows for simplicity, easy retrieval and standardisation.


### What is does (by default)
* Windows:
    * Writes an indicator file to `C:\Users\Public\exploit_success.txt` with the contents `Exploited`.
    * Launches `calc.exe` as a visible proof of execution.
    * Prints minimal status to stdout. Ignores all CLI args.
* Linux:
    * Writes an indicator file to `/tmp/exploit_success.txt` with the contents `Exploited`.
    * Launches `echo "Exploited" > /tmp/exploit_success.txt` as a visible proof of execution.
    * Prints minimal status to stdout. Ignores all CLI args.

* No network, no persistence, no privilege escalation, no lateral movement. Purely a simple local execution signal.

### Quick start

* Option 1: Use one of the prebuilt binaries in `bin/` (if you trust me).
* Option 2: Build from source.

Then run on the target (locally or via your RCE vector).

### To-do

* Add a simple dropper to this repo for remote retrieval and execution.
