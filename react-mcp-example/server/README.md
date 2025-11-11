# MCP Server Implementation / การทำงานของ MCP Server

## ภาพรวม / Overview

MCP Server นี้ถูกสร้างขึ้นโดยใช้ `@modelcontextprotocol/sdk` และให้บริการ tools ต่างๆ ที่ AI models สามารถเรียกใช้งานได้

This MCP server is built using `@modelcontextprotocol/sdk` and provides various tools that AI models can call.

## Tools ที่มีให้บริการ / Available Tools

### 1. create_task
สร้างงานใหม่ในระบบ / Create a new task in the system

**Parameters:**
- `title` (required): ชื่องาน / Task title
- `description` (optional): รายละเอียดงาน / Task description
- `priority` (optional): ระดับความสำคัญ (low, medium, high) / Priority level

**Example:**
```json
{
  "title": "Complete documentation",
  "description": "Write README and API docs",
  "priority": "high"
}
```

### 2. list_tasks
แสดงรายการงานทั้งหมด / List all tasks

**Parameters:**
- `status` (optional): กรองตามสถานะ (pending, completed, all) / Filter by status

**Example:**
```json
{
  "status": "pending"
}
```

### 3. complete_task
ทำเครื่องหมายงานว่าเสร็จสิ้น / Mark a task as completed

**Parameters:**
- `taskId` (required): รหัสงาน / Task ID

**Example:**
```json
{
  "taskId": "task_1234567890"
}
```

### 4. create_note
สร้างโน้ตใหม่ / Create a new note

**Parameters:**
- `title` (required): หัวข้อโน้ต / Note title
- `content` (required): เนื้อหาโน้ต / Note content
- `tags` (optional): แท็ก / Tags (array)

**Example:**
```json
{
  "title": "Meeting Notes",
  "content": "Discussion about project timeline",
  "tags": ["meeting", "project", "important"]
}
```

### 5. search_notes
ค้นหาโน้ต / Search notes

**Parameters:**
- `query` (required): คำค้นหา / Search query

**Example:**
```json
{
  "query": "project"
}
```

### 6. get_weather
รับข้อมูลสภาพอากาศ / Get weather information

**Parameters:**
- `city` (required): ชื่อเมือง / City name

**Example:**
```json
{
  "city": "Bangkok"
}
```

## การใช้งานกับ AI Clients

### Claude Desktop Configuration

เพิ่ม configuration ใน `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "react-mcp-example": {
      "command": "node",
      "args": ["/path/to/react-mcp-example/server/index.js"],
      "env": {}
    }
  }
}
```

### ตำแหน่งไฟล์ Config / Config File Locations

- **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`
- **Linux**: `~/.config/Claude/claude_desktop_config.json`

## Transport Modes

Server รองรับ transport modes ต่างๆ:

### 1. stdio (default)
เหมาะสำหรับการใช้งานกับ AI clients ที่รัน server เป็น subprocess

```bash
node index.js
```

### 2. HTTP (ถ้าต้องการพัฒนาเพิ่มเติม)
สามารถขยายให้รองรับ HTTP transport สำหรับการเรียกใช้ผ่าน web

## การพัฒนาและ Debug

### ใช้ MCP Inspector

```bash
npm run inspector
```

Inspector จะเปิดหน้าเว็บที่สามารถ:
- ดู tools ทั้งหมด
- ทดสอบเรียก tools
- ดู request/response
- Debug การทำงาน

### Logging

Server จะ log ข้อมูลไปยัง stderr (เนื่องจากใช้ stdio สำหรับ MCP communication)

```javascript
console.error('Log message')  // ใช้ stderr สำหรับ log
```

## การจัดเก็บข้อมูล / Data Storage

ปัจจุบันใช้ in-memory storage (Map) สำหรับ demo
ในการใช้งานจริง ควรใช้:
- Database (PostgreSQL, MongoDB, etc.)
- File system
- Cloud storage

### ตัวอย่างการเชื่อมต่อ Database

```javascript
import { MongoClient } from 'mongodb'

const client = new MongoClient(process.env.MONGODB_URI)
await client.connect()
const db = client.db('mcp-example')
const tasks = db.collection('tasks')

// ใช้ในการจัดเก็บ tasks แทน Map
```

## ความปลอดภัย / Security

### Best Practices:
1. ✅ ตรวจสอบ input ทุกครั้ง / Always validate input
2. ✅ จำกัดขนาดของ input / Limit input size
3. ✅ ใช้ environment variables สำหรับ sensitive data
4. ✅ เพิ่ม rate limiting ในการใช้งานจริง
5. ✅ Sanitize output ก่อนส่งกลับ

## การขยายความสามารถ / Extending Capabilities

### เพิ่ม Tool ใหม่

1. เพิ่ม tool definition ใน `ListToolsRequestSchema` handler
2. เพิ่ม tool handler ใน `CallToolRequestSchema` handler
3. อัพเดท documentation
4. เพิ่ม tests

### เพิ่ม Resources

MCP รองรับ resources สำหรับให้ context กับ AI:

```javascript
server.setRequestHandler(ListResourcesRequestSchema, async () => {
  return {
    resources: [
      {
        uri: 'file:///path/to/resource',
        mimeType: 'text/plain',
        name: 'Resource Name',
      },
    ],
  }
})
```

## Performance Optimization

### ในการใช้งานจริง Production:

1. **Caching**: ใช้ Redis หรือ Memcached
2. **Database Connection Pooling**: จัดการ connections อย่างมีประสิทธิภาพ
3. **Async Operations**: ใช้ async/await อย่างถูกต้อง
4. **Error Handling**: จัดการ errors อย่างครอบคลุม
5. **Monitoring**: เพิ่ม logging และ monitoring

## Testing

### Unit Tests

```bash
npm test
```

### Integration Tests

```bash
npm run test:integration
```

### ใช้ MCP Inspector สำหรับ Manual Testing

```bash
npm run inspector
```

## Troubleshooting

### ปัญหาทั่วไป / Common Issues:

1. **Server ไม่ทำงาน**: ตรวจสอบ Node.js version >= 18
2. **Tools ไม่ถูกเรียก**: ตรวจสอบ tool names และ parameters
3. **Memory leak**: ในการใช้งานจริง ใช้ proper database แทน Map

---

สำหรับคำถามเพิ่มเติม ดูได้ที่ [MCP Documentation](https://modelcontextprotocol.io)
