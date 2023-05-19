<h1 align="center" style="border-bottom: none">
    <b>
        <a href="https://docker.nsddd.top">go-project-layout</a><br>
    </b>
</h1>
<h3 align="center" style="border-bottom: none">
      ⭐️  Template for a typical module written on Go.  ⭐️ <br>
<h3>


<p align=center>
<a href="https://goreportcard.com/report/github.com/kubecub/go-project-layout"><img src="https://goreportcard.com/badge/github.com/kubecub/go-project-layout" alt="A+"></a>
<a href="https://github.com/issues?q=org%25kubecub+is%3Aissue+label%3A%22good+first+issue%22"><img src="https://img.shields.io/github/issues/kubecub/go-project-layout/good%20first%20issue?logo=%22github%22" alt="good first"></a>
<a href="https://github.com/kubecub/go-project-layout"><img src="https://img.shields.io/github/stars/kubecub/go-project-layout.svg?style=flat&logo=github&colorB=deeppink&label=stars"></a>
<a href="https://join.slack.com/t/kubecub/shared_invite/zt-1se0k2bae-lkYzz0_T~BYh3rjkvlcUqQ"><img src="https://img.shields.io/badge/Slack-100%2B-blueviolet?logo=slack&amp;logoColor=white"></a>
<a href="https://github.com/kubecub/go-project-layout/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-Apache--2.0-green"></a>
<a href="https://golang.org/"><img src="https://img.shields.io/badge/Language-Go-blue.svg"></a>
</p>

</p>

<p align="center">
    <a href="./README.md"><b>English</b></a> •
    <a href="./README_zh-CN.md"><b>中文</b></a>
</p>

</p>

----


## 🧩 Awesome features

A system that monitors and tracks ESG (Environmental, Social and Governance) disclosures and ratings from various public data sources. It collects ESG data points from sources like Feishu spreadsheets and makes them available through an administrative interface.

## 🛫 Quick start 

> **Note**: You can get started quickly with feishu-sheet-parser.


<details>
  <summary>Work with Makefile</summary>

```bash
❯ make help    # show help
❯ make build   # build binary
```

</details>
<details>
  <summary>Work with actions</summary>

Actions provide handling of PR and issue.
We used the bot @kubbot, It can detect issues in Chinese and translate them to English, and you can interact with it using the command `/comment`.

Comment in an issue:

```bash
❯ /intive
```

</details>
<details>
  <summary>Work with Tools</summary>

```bash
❯ make tools
```

</details>
<details>
  <summary>Use feishu SDK</summary>

Init client
```go
import (
    parsetsdk "github.com/kubecub/feishu-sheet-parser"
)

cli := parsetsdk.MustNewClient("cli_a253xxxxb500b", "6YKq9xxxxpZv6D")
```

Call open interface(CURD)
> example inset
```go
func demo() {
    ctx := context.Background()

    baseToken := "bascntXKzhxxxxxVKPRe"
    tableID := "tbl8xxxxAu"
    fields := map[string]interface{}{
        "multi-line text": "test content",
    }

    record, err := cli.CreateRecord(ctx, baseToken, tableID, fields)
    if err != nil {
        fmt.Printf("createRecord failed, err=%v\n", err)
    } else {
        fmt.Printf("createRecord success, record=%+#v\n", record)
    }
}
```

</details>


## 🕋 architecture diagram
```mermaid
graph LR

subgraph External Services
feishuAPI --> sheetParser
end

subgraph Sheet Parser & Manager
sheetParser[Sheet Parser] --> versionControl
versionControl[Version Control] --> sheetDisplay
end

subgraph Display
sheetDisplay[Sheet Display] --> UI
end

subgraph Backend
versionControl --> API
end

subgraph Frontend
UI[User Interface]
end

API --> UI
UI --> feishuAPI 
```

**MVC Architecture Design:**
```mermaid
graph LR
A[View] -->|1. User interaction| B(Controller)
B -->|2. Requests data| C(Model)
C -->|3. Returns data| B
B -->|4. Updates view| A
```

## 🤖 File Directory Description


## community meeting

We welcome everyone to join us and contribute to feishu-sheet-parser, whether you are new to open source or professional. We are committed to promoting an open source culture, so we offer community members neighborhood prizes and reward money in recognition of their contributions. We believe that by working together, we can build a strong community and make valuable open source tools and resources available to more people. So if you are interested in feishu-sheet-parser, please join our community and start contributing your ideas and skills!

We take notes of each [biweekly meeting](https://github.com/kubecub/feishu-sheet-parser/issues/2) in [GitHub discussions](https://github.com/kubecub/feishu-sheet-parser/discussions/categories/meeting), and our minutes are written in [Google Docs](https://docs.google.com/document/d/1nx8MDpuG74NASx081JcCpxPgDITNTpIIos0DS6Vr9GU/edit?usp=sharing).

feishu-sheet-parser maintains a [public roadmap](https://github.com/kubecub/community/tree/main/roadmaps). It gives a a high-level view of the main priorities for the project, the maturity of different features and projects, and how to influence the project direction.


## License

Sealer is licensed under the Apache License, Version 2.0. See [LICENSE](https://github.com/kubecub/feishu-sheet-parser/tree/main/LICENSE) for the full license text.

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fsealerio%2Fsealer.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fkubecub%2Ffeishu-sheet-parser?ref=badge_large)


## Thanks to our contributors!

<a href="https://github.com/kubecub/feishu-sheet-parser/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=kubecub/feishu-sheet-parser" />
</a>