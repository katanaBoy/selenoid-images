# Selenoid Docker Images
This repository contains [Docker](http://docker.com/) build files to be used for [Selenoid](http://github.com/aerokube/selenoid) project. You can find prebuilt images [here](https://hub.docker.com/u/selenoid/).

## Download Statistics

### Firefox: [![Firefox Docker Pulls](https://img.shields.io/docker/pulls/selenoid/firefox.svg)](https://hub.docker.com/r/selenoid/firefox)

### Chrome: [![Chrome Docker Pulls](https://img.shields.io/docker/pulls/selenoid/chrome.svg)](https://hub.docker.com/r/selenoid/chrome)

### Opera: [![Opera Docker Pulls](https://img.shields.io/docker/pulls/selenoid/opera.svg)](https://hub.docker.com/r/selenoid/opera)

### Android: [![Android Docker Pulls](https://img.shields.io/docker/pulls/selenoid/android.svg)](https://hub.docker.com/r/selenoid/android)

## How images are built

![layers](layers.png)

Each image consists of 3 or 4 layers:
1) **Base layer** - contains stuff needed in every image: Xvfb, fonts, cursor blinking fix, timezone definition and so on. This layer is always built manually.
2) **Optional Java layer** - contains latest Java Runtime Environment. Only needed for old Firefox versions incompatible with Geckodriver. This layer is always built manually.
3) **Browser layer** - contains browser binary. We create two versions: with APT cache and without it. The latter is then used to add driver layer.
4) **Driver layer** - contains either respective web driver binary or corresponding Selenium server version.

## How to build images yourself

Building procedure is automated with shell scripts ```selenium/automate_chrome.sh```, ```selenium/automate_firefox.sh``` and so on.

* Before building images you can optionally clone tests repository:
```
$ git clone https://github.com/aerokube/selenoid-container-tests.git
```
These tests require Java and Maven 3 to be installed. Tests directory should be cloned to this repository parent directory:
```
selenoid-images/ # <== this repo
selenoid-container-tests/ # <== optional tests repo
```
* To build a Firefox image use the following command:
```
$ ./automate_firefox.sh 70.0.1+build1-0ubuntu0.18.04.1 1.9.3 70.0 0.26.0
```
Here `70.0.1+build1-0ubuntu0.18.04.1` is `firefox` package version for Ubuntu 18.04, `1.9.3` is [Selenoid](https://github.com/aerokube/selenoid/releases) version to use inside image (just use latest release version here), `70.0` is Docker tag to be applied, `0.26.0` is [Geckodriver](http://github.com/mozilla/geckodriver/releases) version to use.

If you wish to automatically use the latest Selenoid and Geckodriver versions - just replace them with **latest**:
```
$ ./automate_firefox.sh 70.0.1+build1-0ubuntu0.18.04.1 latest 70.0 latest
```

If you wish to pack a local Debian package instead of APT - just replace package version with full path to **deb** file:
```
$ ./automate_firefox.sh /path/to/firefox_70.0.1+build1-0ubuntu0.18.04.1_i386.deb 1.9.3 70.0 0.26.0
``` 
It is important to use package files with full version specified name because automation scripts determine browser version by parsing package file name!

* To build a Chrome image use the following command:
```
$ ./automate_chrome.sh 78.0.3904.97-1 78.0.3904.70 78.0
```
Here `78.0.3904.97-1` is `google-chrome-stable` package version for Ubuntu 18.04, `78.0.3904.70` is [Chromedriver](https://chromedriver.storage.googleapis.com/index.html) version, `78.0` is Docker tag to be applied.  

If you wish to automatically use the latest [compatible](https://chromedriver.chromium.org/downloads/version-selection) Chromedriver version - just replace it with **latest**:
```
$ ./automate_chrome.sh 78.0.3904.97-1 latest 78.0
```
* To build an Opera image use the following command:
```
$ ./automate_opera.sh 64.0.3417.92 77.0.3865.120 64.0
```
Here `64.0.3417.92` is `opera-stable` package version for Ubuntu 18.04, `77.0.3865.120` is [Operadriver](https://github.com/operasoftware/operachromiumdriver/releases) version, `64.0` is Docker tag to be applied.  

* To build a Yandex image use the following command:
```
$ ./automate_yandex.sh 19.9.3.358-1 19.9.3 19.9
```
Here `19.9.3.358-1` is `yandex-browser-beta` package version for Ubuntu 18.04, `19.9.3` is [Yandexdriver](https://github.com/yandex/YandexDriver/tree/master/linux) version, `19.9` is Docker tag to be applied.

* To build an Android image use the following command:
```
$ ./automate_android.sh
```
This command is interactive - just answer the questions and it will build an image for you. In order to bundle custom APK to image - put it to `selenium/android` directory before running the script.

## How to build images for Beta and Dev channels

Apart from the default Stable release channel, the following ones are also supported:

| Browser | Channel | Package |
| :--- | :--- | :--- |
| Firefox | Beta | firefox [(PPA)](https://launchpad.net/~mozillateam/+archive/ubuntu/firefox-next/+packages) |
| Firefox | Dev | firefox-trunk [(PPA)](https://launchpad.net/~ubuntu-mozilla-daily/+archive/ubuntu/ppa/+packages) |
| Chrome | Beta | google-chrome-beta |
| Chrome | Dev | google-chrome-unstable |
| Opera | Beta | opera-beta | |
| Opera | Dev | opera-developer | |

* To build an image for one of the channels above use the optional argument `{beta|dev}` at the end of the corresponding command.
```
$ ./automate_firefox.sh 72.0~a1~hg20191114r501767-0ubuntu0.18.04.1~umd1 1.9.3 72.0a1 0.26.0 dev
```

## Image information
Moved to: http://aerokube.com/selenoid/latest/#_browser_image_information
