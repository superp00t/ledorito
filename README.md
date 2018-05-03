# LeDorito

ElDewrito asset decoder, shamelessly ported to Go from [cra0kalo's work](https://github.com/cra0kalo/Halo-Online-ElDorado-DAT-Extractor/)

I decided to translate eldorado_dat just because I hate using C# and Visual Studio.

## Installation

```
go get -u -v github.com/superp00t/ledorito
go install -v github.com/superp00t/cmd/ledorito_dat
```

## Usage

```
$ ledorito_dat <halo online .dat file> <output directory>
```