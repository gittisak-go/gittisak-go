<div align="center">
  <p>
    <a align="center" href="#"_blank">
      <img src="https://res.cloudinary.com/dcwjok3nu/image/upload/v1764411147/Backend-sparesos_gd5myv.jpg" alt="gittisak-go" height=100%/>
    </a>
    <h1 align="center">GtsAlpha Model Context Protocol (MCP)</h1>
    <h3 align="center">(โปรโตคอลบริบทโมเดล—Model Context Protocol: MCP)</h3>
  </p>
  <p>
  </p>
  <br/>
</div>

## Gittisak Go - MCP Server Thailand 100% (เซิร์ฟเวอร์ MCP เปิดซอร์สพร้อมใช้ สำหรับประเทศไทย)

[![Watch the video](https://img.youtube.com/vi/kOhLoixrJXo/maxresdefault.jpg)](https://www.youtube.com/watch?v=kOhLoixrJXo)

**🚀 A production-ready Model Context Protocol server written in Go**  
**🚀 เซิร์ฟเวอร์ Model Context Protocol (MCP) ที่พร้อมใช้งานจริง พัฒนาด้วยภาษา Go**

คลังข้อมูลนี้ (repository) ประกอบด้วยเซิร์ฟเวอร์ MCP ที่ครบสมบูรณ์  
เหมาะสำหรับการเชื่อมต่อแอปพลิเคชัน AI ต่างๆ (เช่น Claude Desktop, VSCode, Perplexity, Figma) กับเครื่องของคุณเพื่อใช้งานเครื่องมือ (tools) และแหล่งข้อมูล (resources) ได้อย่างมีประสิทธิภาพ

## Quick Start (เริ่มต้นใช้งานอย่างรวดเร็ว)

```bash
# Build the server (สร้างเซิร์ฟเวอร์)
make build

# The binary will be at: bin/mcp-server
# ไฟล์โปรแกรมจะอยู่ที่: bin/mcp-server
```

ดูคำแนะนำการติดตั้งแบบละเอียดที่ [README-MCP.md](README-MCP.md)

## Features (คุณสมบัติเด่น)

✅ **stdio Transport** - สื่อสารผ่านมาตรฐาน MCP  
✅ **JSON-RPC 2.0** - รองรับมาตรฐาน MCP อย่างสมบูรณ์ (Fully compliant)  
✅ **Built-in Tools** - มีเครื่องมือพื้นฐาน เช่น echo, ดูเวลา, อ่านไฟล์  
✅ **Extensible** - เพิ่มเครื่องมือใหม่ได้อย่างง่ายดาย (ต่อขยายได้)  
✅ **Zero Dependencies** - เขียนด้วย Go ล้วน ไม่ต้องมีไลบรารีจากภายนอก  

## Client Support (รองรับการเชื่อมต่อกับ)

- 🤖 Claude Desktop
- 💻 VSCode (พร้อมส่วนขยาย MCP extension)
- 🔍 Perplexity
- 🎨 Figma (ผ่านปลั๊กอิน MCP)
- หรือแอปที่รองรับ MCP อื่นๆ

---

# Go

MCP คือโปรโตคอลเปิด (Open Protocol) ที่ออกแบบมาเพื่อเป็นมาตรฐานกลางในการเชื่อมโยง application กับ context สำหรับ LLM (Large Language Models)

MCP เปรียบเสมือน USB-C ของโลก AI  
เช่นเดียวกับที่พอร์ต USB-C ให้มาตรฐานกลางในการเชื่อมอุปกรณ์ electronic ต่างๆ MCP คือมาตรฐานกลางในการเชื่อมข้อมูลและเครื่องมือ ให้กับแอปพลิเคชัน AI

## Go Build MCP products  
(พัฒนาโซลูชัน MCP ด้วยภาษา Go ได้เลย)

# What is the Model Context Protocol (MCP)?  
(MCP คืออะไร?)

Model Context Protocol (MCP) คือมาตรฐานเปิด (Open Standard) สำหรับเชื่อม Application AI เข้ากับระบบภายนอก เช่น ฐานข้อมูล เครื่องมือ หรือเวิร์กโฟลว์ต่างๆ

เช่น AI อย่าง Claude หรือ ChatGPT จะสามารถเชื่อมกับข้อมูลในไฟล์ (local files), ฐานข้อมูล, เครื่องมือ (Tools) เช่น เครื่องคิดเลข หรือเครื่องมือเฉพาะทางอื่นๆ ได้—ทำให้ AI ทำงานได้กว้างขึ้น ทรงพลังขึ้น

MCP เปรียบเหมือน USB-C สำหรับ AI applications  
ดังที่ USB-C เป็นมาตรฐานกลาง สำหรับเชื่อมอุปกรณ์ต่างๆ, MCP เป็นมาตรฐานกลางสำหรับเชื่อมแอป AI กับระบบและทรัพยากรภายนอก

<Frame>
  <img src="https://mintcdn.com/mcp/bEUxYpZqie0DsluH/images/mcp-simple-diagram.png?fit=max&auto=format&n=bEUxYpZqie0DsluH&q=85&s=35268aa0ad50b8c385913810e7604550" data-og-width="3840" width="3840" data-og-height="2160" alt="MCP Simple Architecture Diagram"/>
  <p align="center"><i>แผนภาพโครงสร้างพื้นฐานของ MCP (Simple Architecture of MCP Protocol)</i></p>
</Frame>

## What can MCP enable?  
(MCP เปิดศักยภาพอะไรได้บ้าง?)

* เอเจนต์ (Agents) เข้าถึง Google Calendar และ Notion ของผู้ใช้ ปรับแต่งให้เป็นผู้ช่วย AI ส่วนตัว
* Claude Code สร้างแอปเว็บครบชุดจากดีไซน์ Figma
* แชทบอทองค์กรเข้าถึงฐานข้อมูลหลากหลาย เพื่อช่วยวิเคราะห์ข้อมูลผ่านการสนทนา
* AI สร้าง 3D Designs ใน Blender แล้วสั่งพิมพ์ 3D Printer โดยอัตโนมัติ

## Why does MCP matter?  
(ทำไม MCP ถึงสำคัญ?)

ประโยชน์ของ MCP ขึ้นอยู่กับบทบาทของคุณใน ecosystem

* **Developers** (นักพัฒนา): ลดเวลาพัฒนาและความซับซ้อน เมื่อสร้างหรือเชื่อมต่อแอป AI/Agent
* **AI applications or agents** (แอป/Agent AI): ใช้งานแหล่งข้อมูล, เครื่องมือ, แอป อื่นๆ ได้ง่ายและหลากหลาย เพิ่มประสิทธิภาพ
* **End-users** (ผู้ใช้งาน): ได้รับประสบการณ์จาก AI ที่ฉลาดขึ้น ใช้งานได้จริงและช่วยทำงาน/ตัดสินใจได้หลากหลาย

<div align="center">
  <a href="https://www.youtube.com/watch?v=kOhLoixrJXo">
    <img src="https://img.youtube.com/vi/kOhLoixrJXo/0.jpg" alt="Video Preview">
  </a>
</div>
<br/>

## Start Building (เริ่มสร้าง—Build Now!)

<CardGroup cols={2}>
  <Card title="Build servers (สร้างเซิร์ฟเวอร์)" icon="server" href="/docs/develop/build-server">
    สร้างเซิร์ฟเวอร์ MCP สำหรับเปิดเผยข้อมูลและเครื่องมือ
  </Card>
  <Card title="Build clients (พัฒนาไคลเอนต์)" icon="computer" href="/docs/develop/build-client">
    พัฒนาแอปพลิเคชันที่เชื่อมต่อกับ MCP servers
  </Card>
</CardGroup>

## Learn more (เรียนรู้เพิ่มเติม)

<CardGroup cols={2}>
  <Card title="Understand concepts (เข้าใจแนวคิด)" icon="book" href="/docs/learn/architecture">
    เรียนรู้แนวคิดหลักและโครงสร้าง MCP
  </Card>
</CardGroup>

<Card title="MCP Documentation (เอกสาร MCP)" icon="book" href="https://modelcontextprotocol.io">
  ศึกษารายละเอียดโปรโตคอล วิธีสร้างเซิร์ฟเวอร์/ไคลเอนต์ และดูตัวอย่างการใช้งาน
</Card>

## MCP in Anthropic products (MCP กับผลิตภัณฑ์ของ Anthropic)

<CardGroup>
  <Card title="MCP in the Messages API" icon="cloud" href="/en/docs/agents-and-tools/mcp-connector">
    ใช้งาน MCP Connector บน Messages API เพื่อเชื่อมต่อ MCP servers
  </Card>
  <Card title="MCP in Claude Code" icon="head-side-gear" href="/en/docs/claude-code/mcp">
    เพิ่ม MCP server สำหรับ Claude Code หรือใช้ Claude Code เป็น server
  </Card>
  <Card title="MCP in Claude.ai" icon="comments" href="https://support.claude.com/en/articles/11175166-getting-started-with-custom-connectors-using-remote-mcp">
    เปิดใช้งาน MCP connectors ให้ทีมของคุณบน Claude.ai
  </Card>
  <Card title="MCP in Claude Desktop" icon="desktop" href="https://support.claude.com/en/articles/10949351-getting-started-with-local-mcp-servers-on-claude-desktop">
    เพิ่ม MCP server สำหรับ Claude Desktop
  </Card>
</CardGroup>

---

## 🚀 Production-Ready Examples / ตัวอย่างที่พร้อมใช้งานจริง

### React MCP Server Example (ตัวอย่างเซิร์ฟเวอร์ MCP สำหรับ React)

🎯 **ตัวอย่าง MCP Server พร้อม React Client ที่สามารถใช้งานได้จริง**  
A complete example showing how to build an MCP server using Node.js and React client—with production-ready features.

ตัวอย่างนี้เหมาะสำหรับการสร้าง MCP Server ด้วย Node.js และ React Client ที่มีฟีเจอร์พร้อมใช้งานจริง:

- ✅ **Task Management** - ระบบจัดการงานพร้อมระดับความสำคัญ (Task Management with priorities)
- ✅ **Note-taking** - ระบบจดโน้ต พร้อมแท็ก & ค้นหา (Note system with search/tags)
- ✅ **Weather Info** - ดูข้อมูลสภาพอากาศ (จำลอง)
- ✅ **Full Documentation** - คู่มือภาษาไทย/อังกฤษครบถ้วน
- ✅ **Production-Ready** - พร้อมใช้งานในระบบจริง

**📂 ตำแหน่ง / Location:** [`react-mcp-example/`](react-mcp-example/)

**📖 Quick Start (เริ่มต้นอย่างรวดเร็ว):**
```bash
# Clone and setup (โคลนโปรเจกต์และตั้งค่า)
cd react-mcp-example

# Install server dependencies (ติดตั้ง dependencies ฝั่ง server)
cd server && npm install

# Install client dependencies (ติดตั้ง dependencies ฝั่ง client)
cd ../client && npm install

# Start server (เริ่ม server)
cd ../server && npm start

# Start client (เริ่ม client - เปิด terminal ใหม่)
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
