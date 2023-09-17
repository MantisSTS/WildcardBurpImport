# WildcardBurpImport
A tool to convert a list of wildcard domains into Burp's scope

## Install

```
go install -v github.com/MantisSTS/WildcardBurpImport@latest
```

## Usage

**wildcards.txt**

```txt
yahoo.com
google.com
```

```bash
$ WildcardBurpImport -f wildcards.txt -o results.json
```

![burpimport](https://github.com/MantisSTS/WildcardBurpImport/assets/818959/f389cd14-66ca-404f-96b3-fc987b20387f)
