# Bug Out Monitor

The Bug Out Monitor is a Python script that monitors news articles for mentions of specific keywords related to international conflicts. When the script detects relevant keywords in the news, it sends notifications and plays a voice alert on your system.

As you all know, Wagner et al are eager to make a move, with opposing leaders throwing their arms up saying "We don't know how we can stop them". So I decided to create a bug out monitor, so that it can alert me of any invasions to bug out, if needed.

If you'd like to contribute, you are free to do so. There are no rules, besides me reviewing PRs. But most likely I will approve it. 

However there is 1 requirement, please refrain from using Natural Language processors. As they are resource heavy, and I would like this to be lightweight monitor.

## Prerequisites

Before running the Bug Out Monitor, ensure you have Go installed on your system.

You will also need the following additional packages for the Text-to-Speech (TTS) feature:

- notify-send command-line utility (for Linux desktop notifications)
- espeak package (for TTS support on Linux)

You can install the required Python libraries using the following command:

On Linux, you can install the espeak package using the package manager specific to your distribution. For example:

Arch:

```bash
sudo pacman -S espeak-ng
```

Ubuntu/Debian:

```bash
sudo apt-get install espeak
```

Fedora:

```bash
sudo dnf install espeak
```

In some cases, one will need to make a symlink for espeak-ng:

```bash
sudo ln -s /usr/lib/libespeak-ng.so.1 /usr/lib/libespeak.so.1
```

## Usage

If you've gotten this far, you should be able to compile and run it yourself. You're a big boy dev now

## Running in Cron

To run the Bug Out Monitor periodically using cron, add the following line to your crontab:

```bash
*/30 * * * * /path/to/your/bom-ng
```

Replace /path/to/your/bom-ng with the full path to the bom-ng binary you have compiled on your system.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
