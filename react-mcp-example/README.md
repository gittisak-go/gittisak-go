# React MCP Server Example / ตัวอย่าง MCP Server พร้อม React

<div align="center">
  <h1>🚀 Production-Ready MCP Server with React</h1>
  <p>ตัวอย่างการสร้าง MCP (Model Context Protocol) Server ที่พร้อมใช้งานจริง พร้อม React Client</p>
  <p>A comprehensive, production-ready example of building an MCP server with React client</p>
</div>

---

## 📋 สารบัญ / Table of Contents

- [ภาพรวม / Overview](#overview)
- [คุณสมบัติ / Features](#features)
- [สิ่งที่ต้องเตรียม / Prerequisites](#prerequisites)
- [การติดตั้ง / Installation](#installation)
- [การใช้งาน / Usage](#usage)
- [โครงสร้างโปรเจกต์ / Project Structure](#project-structure)
- [การทำงานของ MCP / How MCP Works](#how-mcp-works)
- [ตัวอย่างการใช้งาน / Examples](#examples)
- [การพัฒนาต่อ / Development](#development)
- [การแก้ปัญหา / Troubleshooting](#troubleshooting)

---

## 🎯 ภาพรวม / Overview

โปรเจกต์นี้แสดงตัวอย่างการสร้าง MCP (Model Context Protocol) Server ที่สามารถนำไปใช้งานได้จริง พร้อมกับ React Client ที่สามารถเชื่อมต่อและใช้งาน tools ต่างๆ ที่ server จัดเตรียมไว้

This project demonstrates how to build a production-ready MCP (Model Context Protocol) server with a React client that can connect and use various tools provided by the server.

### MCP คืออะไร? / What is MCP?

MCP (Model Context Protocol) เป็นโปรโตคอลเปิดที่ช่วยให้ AI models สามารถเชื่อมต่อกับแหล่งข้อมูลและเครื่องมือภายนอกได้อย่างมาตรฐาน

MCP is an open protocol that enables AI models to connect with external data sources and tools in a standardized way.

---

## ✨ คุณสมบัติ / Features

### 📋 ระบบจัดการงาน / Task Management
- ✅ สร้างงานใหม่พร้อมระดับความสำคัญ / Create tasks with priority levels
- ✅ แสดงรายการงานแยกตามสถานะ / List tasks filtered by status
- ✅ ทำเครื่องหมายงานว่าเสร็จสิ้น / Mark tasks as completed

### 📝 ระบบบันทึกโน้ต / Note-taking System
- ✅ สร้างโน้ตพร้อมแท็ก / Create notes with tags
- ✅ ค้นหาโน้ตจากหัวข้อ เนื้อหา หรือแท็ก / Search notes by title, content, or tags
- ✅ จัดการโน้ตได้ง่าย / Easy note management

### 🌤️ ข้อมูลสภาพอากาศ / Weather Information
- ✅ ดึงข้อมูลสภาพอากาศ (จำลอง) / Get weather data (simulated)
- ✅ แสดงผลในรูปแบบที่สวยงาม / Beautiful UI display

### 🔧 MCP Server Implementation
- ✅ รองรับ stdio transport / Support for stdio transport
- ✅ เครื่องมือที่พร้อมใช้งาน / Production-ready tools
- ✅ จัดการ resources / Resource management
- ✅ ตั้งค่าได้ง่าย / Easy configuration

---

## 📦 สิ่งที่ต้องเตรียม / Prerequisites

ก่อนเริ่มใช้งาน คุณต้องติดตั้งโปรแกรมเหล่านี้:

Before you begin, ensure you have the following installed:

- **Node.js** >= 18.0.0
- **npm** or **yarn**
- **Git**

---

## 🚀 การติดตั้ง / Installation

### 1. Clone Repository

```bash
git clone https://github.com/gittisak-go/gittisak-go.git
cd gittisak-go/react-mcp-example
```

### 2. ติดตั้ง Dependencies สำหรับ Server

```bash
cd server
npm install
```

### 3. ติดตั้ง Dependencies สำหรับ Client

```bash
cd ../client
npm install
```

---

## 💻 การใช้งาน / Usage

### เริ่มต้น MCP Server

```bash
cd server
npm start
```

หรือใช้โหมด development พร้อม auto-reload:

```bash
npm run dev
```

### เริ่มต้น React Client

เปิด terminal ใหม่:

```bash
cd client
npm run dev
```

React app จะทำงานที่ `http://localhost:3000`

---

## 🗂️ โครงสร้างโปรเจกต์ / Project Structure

```
react-mcp-example/
├── server/                    # MCP Server
│   ├── index.js              # Server implementation
│   ├── package.json          # Server dependencies
│   └── README.md             # Server documentation
│
├── client/                    # React Client
│   ├── src/
│   │   ├── components/       # React components
│   │   │   ├── TaskManager.jsx
│   │   │   ├── NoteManager.jsx
│   │   │   └── WeatherWidget.jsx
│   │   ├── hooks/           # Custom hooks
│   │   │   └── useMCP.js
│   │   ├── services/        # Services
│   │   │   └── mcpService.js
│   │   ├── App.jsx          # Main app component
│   │   ├── App.css          # App styles
│   │   ├── index.css        # Global styles
│   │   └── main.jsx         # Entry point
│   ├── index.html           # HTML template
│   ├── vite.config.js       # Vite configuration
│   └── package.json         # Client dependencies
│
└── README.md                 # This file
```

---

## 🔄 การทำงานของ MCP / How MCP Works

```
┌─────────────────┐        ┌─────────────────┐        ┌─────────────────┐
│   React Client  │◄──────►│   MCP Server    │◄──────►│   AI Model      │
│                 │  JSON  │                 │  JSON  │   (Claude,      │
│  - Task Manager │  RPC   │  - Tools        │  RPC   │    GPT, etc)    │
│  - Note Manager │        │  - Resources    │        │                 │
│  - Weather      │        │  - Prompts      │        │                 │
└─────────────────┘        └─────────────────┘        └─────────────────┘
```

### ขั้นตอนการทำงาน / Workflow

1. **Client** เรียกใช้งาน tool ผ่าน MCP protocol
2. **Server** ประมวลผลคำขอและทำงานตาม tool ที่เรียก
3. **Server** ส่งผลลัพธ์กลับไปยัง client
4. **Client** แสดงผลข้อมูลให้ผู้ใช้เห็น

---

## 📚 ตัวอย่างการใช้งาน / Examples

### การสร้าง Task ใหม่ / Creating a New Task

```javascript
const result = await callTool('create_task', {
  title: 'Complete project documentation',
  description: 'Write comprehensive README',
  priority: 'high'
})
```

### การค้นหา Notes / Searching Notes

```javascript
const result = await callTool('search_notes', {
  query: 'react'
})
```

### การดึงข้อมูลสภาพอากาศ / Getting Weather

```javascript
const result = await callTool('get_weather', {
  city: 'Bangkok'
})
```

---

## 🛠️ การพัฒนาต่อ / Development

### เพิ่ม Tool ใหม่ / Adding New Tools

แก้ไขไฟล์ `server/index.js`:

```javascript
server.setRequestHandler(ListToolsRequestSchema, async () => {
  return {
    tools: [
      // ... existing tools
      {
        name: 'your_new_tool',
        description: 'Description of your tool',
        inputSchema: {
          type: 'object',
          properties: {
            param1: {
              type: 'string',
              description: 'Parameter description',
            },
          },
          required: ['param1'],
        },
      },
    ],
  }
})

// Add handler for your tool
server.setRequestHandler(CallToolRequestSchema, async (request) => {
  // ... existing cases
  case 'your_new_tool': {
    // Your implementation
    return {
      content: [{ type: 'text', text: 'result' }],
    }
  }
})
```

### Debug Mode

ใช้ MCP Inspector เพื่อ debug:

```bash
cd server
npm run inspector
```

---

## 🐛 การแก้ปัญหา / Troubleshooting

### ปัญหา: Server ไม่สามารถเริ่มต้นได้

**วิธีแก้:**
- ตรวจสอบว่าติดตั้ง Node.js >= 18.0.0
- ลบ `node_modules` และติดตั้งใหม่: `rm -rf node_modules && npm install`

### ปัญหา: Client ไม่สามารถเชื่อมต่อกับ Server

**วิธีแก้:**
- ตรวจสอบว่า server ทำงานอยู่
- ตรวจสอบ port ที่ใช้งาน
- ดู console log เพื่อหาข้อผิดพลาด

### ปัญหา: การติดตั้ง dependencies ล้มเหลว

**วิธีแก้:**
```bash
# Clear npm cache
npm cache clean --force

# ติดตั้งใหม่
npm install
```

---

## 🔗 แหล่งข้อมูลเพิ่มเติม / Additional Resources

- [Model Context Protocol Documentation](https://modelcontextprotocol.io)
- [MCP SDK on GitHub](https://github.com/modelcontextprotocol/typescript-sdk)
- [React Documentation](https://react.dev)
- [Vite Documentation](https://vitejs.dev)

---

## 📄 License

MIT License - ดูรายละเอียดใน [LICENSE](../../LICENSE)

---

## 🤝 การมีส่วนร่วม / Contributing

ยินดีรับ contributions! กรุณา:

1. Fork repository
2. สร้าง feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add some AmazingFeature'`)
4. Push ไปยัง branch (`git push origin feature/AmazingFeature`)
5. เปิด Pull Request

---

## 👨‍💻 ผู้พัฒนา / Author

**gittisak-go**

---

## 🙏 ขอบคุณ / Acknowledgments

- Anthropic สำหรับ Model Context Protocol
- React team สำหรับ React framework
- ชุมชน open-source ทั้งหมด

---

<div align="center">
  <p>สร้างด้วย ❤️ โดย gittisak-go</p>
  <p>Made with ❤️ by gittisak-go</p>
</div>
