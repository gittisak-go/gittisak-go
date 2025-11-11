<div align="center">
  <p>
    <a align="center" href="https://developers.moralis.com/" target="_blank">
      <img src="https://raw.githubusercontent.com/gittisak-go/gittisak-go/refs/heads/main/images/tgittisak-logo.svg" alt="Moralis Analytics" height=200/>
    </a>
    <h1 align="center">Model Context Protocol (MCP)</h1>
  </p>
  <p>

  </p>
  <br/>
</div>

# Gittisak Go - MCP Server Implementation

**🚀 A production-ready Model Context Protocol server written in Go**

This repository contains a complete MCP server implementation that enables AI applications like Claude Desktop, VSCode, Perplexity, and Figma to connect to your local machine and access tools and resources.

## Quick Start

```bash
# Build the server
make build

# The binary will be at: bin/mcp-server
```

For detailed setup instructions, see [README-MCP.md](README-MCP.md)

## Features

✅ **stdio Transport** - Standard MCP communication protocol  
✅ **JSON-RPC 2.0** - Fully compliant with MCP specification  
✅ **Built-in Tools** - Echo, get time, read file tools included  
✅ **Extensible** - Easy to add custom tools  
✅ **Zero Dependencies** - Pure Go implementation  

## Client Support

This server works with:
- 🤖 Claude Desktop
- 💻 VSCode (with MCP extension)
- 🔍 Perplexity
- 🎨 Figma (via MCP plugins)
- Any MCP-compatible client

---

# go

MCP is an open protocol that standardizes how applications provide context to LLMs.

Think of MCP like a USB-C port for AI applications. Just as USB-C provides a standardized way to connect your devices to various peripherals and accessories, MCP provides a standardized way to connect AI models to different data sources and tools.

## Go Build MCP products

# What is the Model Context Protocol (MCP)?

MCP (Model Context Protocol) is an open-source standard for connecting AI applications to external systems.

Using MCP, AI applications like Claude or ChatGPT can connect to data sources (e.g. local files, databases), tools (e.g. search engines, calculators) and workflows (e.g. specialized prompts)—enabling them to access key information and perform tasks.

Think of MCP like a USB-C port for AI applications. Just as USB-C provides a standardized way to connect electronic devices, MCP provides a standardized way to connect AI applications to external systems.

<Frame>
  <img src="https://mintcdn.com/mcp/bEUxYpZqie0DsluH/images/mcp-simple-diagram.png?fit=max&auto=format&n=bEUxYpZqie0DsluH&q=85&s=35268aa0ad50b8c385913810e7604550" data-og-width="3840" width="3840" data-og-height="1500" height="1500" data-path="images/mcp-simple-diagram.png" data-optimize="true" data-opv="3" srcset="https://mintcdn.com/mcp/bEUxYpZqie0DsluH/images/mcp-simple-diagram.png?w=280&fit=max&auto=format&n=bEUxYpZqie0DsluH&q=85&s=0cea440365b03c2f2a299b0104375b8b 280w, https://mintcdn.com/mcp/bEUxYpZqie0DsluH/images/mcp-simple-diagram.png?w=560&fit=max&auto=format&n=bEUxYpZqie0DsluH&q=85&s=2391513484df96fa7203739dae5e53b0 560w, https://mintcdn.com/mcp/bEUxYpZqie0DsluH/images/mcp-simple-diagram.png?w=840&fit=max&auto=format&n=bEUxYpZqie0DsluH&q=85&s=96f5e553bee1051dc882db6c832b15bc 840w, https://mintcdn.com/mcp/bEUxYpZqie0DsluH/images/mcp-simple-diagram.png?w=1100&fit=max&auto=format&n=bEUxYpZqie0DsluH&q=85&s=341b88d6308188ab06bf05748c80a494 1100w, https://mintcdn.com/mcp/bEUxYpZqie0DsluH/images/mcp-simple-diagram.png?w=1650&fit=max&auto=format&n=bEUxYpZqie0DsluH&q=85&s=a131a609c7b6a70f342f493bbad57fcb 1650w, https://mintcdn.com/mcp/bEUxYpZqie0DsluH/images/mcp-simple-diagram.png?w=2500&fit=max&auto=format&n=bEUxYpZqie0DsluH&q=85&s=dc4ab238184b6c70e06e871681c921c5 2500w" />
</Frame>

## What can MCP enable?

* Agents can access your Google Calendar and Notion, acting as a more personalized AI assistant.
* Claude Code can generate an entire web app using a Figma design.
* Enterprise chatbots can connect to multiple databases across an organization, empowering users to analyze data using chat.
* AI models can create 3D designs on Blender and print them out using a 3D printer.

## Why does MCP matter?

Depending on where you sit in the ecosystem, MCP can have a range of benefits.

* **Developers**: MCP reduces development time and complexity when building, or integrating with, an AI application or agent.
* **AI applications or agents**: MCP provides access to an ecosystem of data sources, tools and apps which will enhance capabilities and improve the end-user experience.
* **End-users**: MCP results in more capable AI applications or agents which can access your data and take actions on your behalf when necessary.

## Start Building

<CardGroup cols={2}>
  <Card title="Build servers" icon="server" href="/docs/develop/build-server">
    Create MCP servers to expose your data and tools
  </Card>

  <Card title="Build clients" icon="computer" href="/docs/develop/build-client">
    Develop applications that connect to MCP servers
  </Card>
</CardGroup>

## Learn more

<CardGroup cols={2}>
  <Card title="Understand concepts" icon="book" href="/docs/learn/architecture">
    Learn the core concepts and architecture of MCP
  </Card>
</CardGroup>

<Card title="MCP Documentation" icon="book" href="https://modelcontextprotocol.io">
  Learn more about the protocol, how to build servers and clients, and discover those made by others.
</Card>

## MCP in Anthropic products

<CardGroup>
  <Card title="MCP in the Messages API" icon="cloud" href="/en/docs/agents-and-tools/mcp-connector">
    Use the MCP connector in the Messages API to connect to MCP servers.
  </Card>

  <Card title="MCP in Claude Code" icon="head-side-gear" href="/en/docs/claude-code/mcp">
    Add your MCP servers to Claude Code, or use Claude Code as a server.
  </Card>

  <Card title="MCP in Claude.ai" icon="comments" href="https://support.claude.com/en/articles/11175166-getting-started-with-custom-connectors-using-remote-mcp">
    Enable MCP connectors for your team in Claude.ai.
  </Card>

  <Card title="MCP in Claude Desktop" icon="desktop" href="https://support.claude.com/en/articles/10949351-getting-started-with-local-mcp-servers-on-claude-desktop">
    Add MCP servers to Claude Desktop.
  </Card>
</CardGroup>

---

## 🚀 Production-Ready Examples / ตัวอย่างที่พร้อมใช้งานจริง

### React MCP Server Example

🎯 **ตัวอย่าง MCP Server พร้อม React Client ที่สามารถใช้งานได้จริง**

ตัวอย่างที่สมบูรณ์แบบสำหรับการสร้าง MCP Server ด้วย Node.js และ React Client พร้อมคุณสมบัติ:

- ✅ **Task Management** - ระบบจัดการงานพร้อมระดับความสำคัญ
- ✅ **Note-taking** - ระบบบันทึกโน้ตพร้อมแท็กและการค้นหา
- ✅ **Weather Info** - ข้อมูลสภาพอากาศ (จำลอง)
- ✅ **Full Documentation** - เอกสารครบถ้วนภาษาไทยและอังกฤษ
- ✅ **Production-Ready** - พร้อมใช้งานจริง

**📂 ตำแหน่ง / Location:** [`react-mcp-example/`](react-mcp-example/)

**📖 Quick Start:**
```bash
# Clone and setup
cd react-mcp-example

# Install server dependencies
cd server && npm install

# Install client dependencies  
cd ../client && npm install

# Start server
cd ../server && npm start

# Start client (in new terminal)
cd ../client && npm run dev
```

**🔗 เอกสาร / Documentation:**
- [README - เอกสารหลัก](react-mcp-example/README.md)
- [Quick Start Guide - คู่มือเริ่มต้นอย่างรวดเร็ว](react-mcp-example/QUICKSTART.md)
- [Server Documentation - รายละเอียด Server](react-mcp-example/server/README.md)

---

<!--
**gittisak-go/gittisak-go** is a ✨ _special_ ✨ repository because its `README.md` (this file) appears on your GitHub profile.

Here are some ideas to get you started:

- 🔭 I’m currently working on ...
- 🌱 I’m currently learning ...
- 👯 I’m looking to collaborate on ...
- 🤔 I’m looking for help with ...
- 💬 Ask me about ...
- 📫 How to reach me: ...
- 😄 Pronouns: ...
- ⚡ Fun fact: ...
-->
