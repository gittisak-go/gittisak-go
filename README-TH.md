# คู่มือ MCP Server (ภาษาไทย)

เซิร์ฟเวอร์ Model Context Protocol (MCP) ที่เขียนด้วยภาษา Go สำหรับเชื่อมต่อแอปพลิเคชัน AI เช่น Claude Desktop, VSCode, Perplexity และ Figma กับเครื่องมือและแหล่งข้อมูลในเครื่อง

## คุณสมบัติหลัก

- **stdio Transport**: สื่อสารผ่านมาตรฐาน MCP
- **เครื่องมือในตัว**:
  - `echo`: ส่งข้อความกลับ
  - `get_time`: แสดงเวลาปัจจุบันของเซิร์ฟเวอร์
  - `read_file`: อ่านเนื้อหาไฟล์จากระบบ
- **ขยายได้ง่าย**: สามารถเพิ่มเครื่องมือใหม่ได้อย่างง่ายดาย
- **ไม่มี Dependencies**: เขียนด้วย Go ล้วน

## ความต้องการของระบบ

- Go 1.20 หรือสูงกว่า
- โปรแกรมที่รองรับ MCP (Claude Desktop, VSCode พร้อม MCP extension, ฯลฯ)

## การติดตั้ง

### สร้างจาก Source Code

```bash
# Clone repository
git clone https://github.com/gittisak-go/gittisak-go.git
cd gittisak-go

# สร้างเซิร์ฟเวอร์
go build -o bin/mcp-server ./cmd/mcp-server

# หรือใช้ make
make build

# ไฟล์ binary จะอยู่ที่ bin/mcp-server
```

### เริ่มต้นอย่างรวดเร็ว

```bash
# รันเซิร์ฟเวอร์ (สำหรับทดสอบ)
./bin/mcp-server

# เซิร์ฟเวอร์จะเริ่มทำงานและรอรับข้อมูลจาก stdin/stdout
# คุณสามารถส่ง JSON-RPC messages เพื่อทดสอบได้
```

## การตั้งค่า

### Claude Desktop

1. หาไฟล์ config ของ Claude Desktop:
   - **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
   - **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`
   - **Linux**: `~/.config/Claude/claude_desktop_config.json`

2. เพิ่มการตั้งค่าเซิร์ฟเวอร์:

```json
{
  "mcpServers": {
    "gittisak-go": {
      "command": "/path/ที่อยู่แบบเต็ม/ไปยัง/gittisak-go/bin/mcp-server",
      "args": [],
      "env": {}
    }
  }
}
```

**สำคัญ**: ต้องใช้ path แบบเต็ม (absolute path) เท่านั้น

3. รีสตาร์ท Claude Desktop

### VSCode

หากคุณใช้ VSCode พร้อม MCP extension:

1. เปิด VSCode settings (JSON)
2. เพิ่มการตั้งค่า MCP server:

```json
{
  "mcp.servers": {
    "gittisak-go": {
      "command": "/path/ที่อยู่แบบเต็ม/ไปยัง/gittisak-go/bin/mcp-server",
      "args": [],
      "env": {}
    }
  }
}
```

### โปรแกรมอื่นๆ ที่รองรับ MCP

สำหรับโปรแกรมอื่นๆ (Perplexity, Figma plugins, ฯลฯ) ดูเอกสารเฉพาะของแต่ละโปรแกรม โดยทั่วไปคุณจะต้องระบุ:

- **Command**: Path ไปยัง binary `mcp-server`
- **Args**: array ว่าง `[]`
- **Transport**: stdio (ค่าเริ่มต้นสำหรับโปรแกรมส่วนใหญ่)

## เครื่องมือที่มี

### echo

ส่งข้อความกลับตามที่ป้อนเข้ามา

**พารามิเตอร์:**
- `message` (string, จำเป็น): ข้อความที่ต้องการส่งกลับ

**ตัวอย่าง:**
```json
{
  "name": "echo",
  "arguments": {
    "message": "สวัสดี MCP!"
  }
}
```

### get_time

แสดงเวลาปัจจุบันของเซิร์ฟเวอร์ในรูปแบบ RFC3339

**พารามิเตอร์:** ไม่มี

**ตัวอย่าง:**
```json
{
  "name": "get_time",
  "arguments": {}
}
```

### read_file

อ่านและแสดงเนื้อหาของไฟล์จากระบบไฟล์ในเครื่อง

**พารามิเตอร์:**
- `path` (string, จำเป็น): path ของไฟล์ (แบบเต็มหรือแบบสัมพัทธ์)

**ตัวอย่าง:**
```json
{
  "name": "read_file",
  "arguments": {
    "path": "/path/to/file.txt"
  }
}
```

## การพัฒนา

### โครงสร้างโปรเจกต์

```
gittisak-go/
├── cmd/
│   └── mcp-server/       # จุดเริ่มต้นของแอปพลิเคชัน
│       └── main.go
├── pkg/
│   ├── mcp/              # การทำงานของโปรโตคอล MCP
│   │   ├── types.go      # โครงสร้างของโปรโตคอล
│   │   ├── server.go     # การทำงานของเซิร์ฟเวอร์
│   │   └── server_test.go # Unit tests สำหรับเซิร์ฟเวอร์
│   └── tools/            # การทำงานของเครื่องมือ
│       ├── tools.go
│       └── tools_test.go # Unit tests สำหรับเครื่องมือ
├── examples/             # ตัวอย่างการตั้งค่า
│   ├── claude-desktop-config.json
│   └── vscode-config.json
├── bin/                  # Binary ที่สร้างจาก build
├── go.mod
└── README-TH.md          # คู่มือฉบับภาษาไทย
```

### การเพิ่มเครื่องมือใหม่

1. สร้าง handler function ใน `pkg/tools/tools.go`:

```go
func เครื่องมือของคุณ(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
    // โค้ดการทำงานของเครื่องมือ
    return &mcp.CallToolResult{
        Content: []mcp.Content{
            {
                Type: "text",
                Text: "ผลลัพธ์จากเครื่องมือของคุณ",
            },
        },
    }, nil
}
```

2. ลงทะเบียนเครื่องมือใน `cmd/mcp-server/main.go`:

```go
server.RegisterTool(
    "my_custom_tool",
    "คำอธิบายเครื่องมือของคุณ",
    mcp.InputSchema{
        Type: "object",
        Properties: map[string]interface{}{
            "param1": map[string]interface{}{
                "type":        "string",
                "description": "คำอธิบายของพารามิเตอร์ที่ 1",
            },
        },
        Required: []string{"param1"},
    },
    tools.เครื่องมือของคุณ,
)
```

3. สร้างเซิร์ฟเวอร์ใหม่:

```bash
go build -o bin/mcp-server ./cmd/mcp-server
# หรือ
make build
```

### การทดสอบ

#### ทดสอบด้วย Command Line

คุณสามารถทดสอบเซิร์ฟเวอร์โดยตรงจาก command line:

```bash
# Initialize การเชื่อมต่อ
echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"1.0"}}}' | ./bin/mcp-server

# แสดงรายการเครื่องมือที่มี
echo '{"jsonrpc":"2.0","id":2,"method":"tools/list"}' | ./bin/mcp-server

# เรียกใช้ echo tool
echo '{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"echo","arguments":{"message":"สวัสดี!"}}}' | ./bin/mcp-server
```

#### รัน Test Suite

```bash
# รัน unit tests
go test -v ./...

# หรือใช้ make
make test

# รัน integration tests
./test.sh
```

ผลลัพธ์ที่คาดหวัง:
```
======================================
MCP Server Test Suite
======================================

Running tests...

Test 1: Initialize
✓ Initialize: PASSED

Test 2: List Tools
✓ List Tools: PASSED (found echo, get_time, read_file)

Test 3: Call Echo Tool
✓ Echo Tool: PASSED

Test 4: Call Get Time Tool
✓ Get Time Tool: PASSED

Test 5: Call Read File Tool
✓ Read File Tool: PASSED

======================================
Test Suite Complete
======================================
```

## รายละเอียดของโปรโตคอล

เซิร์ฟเวอร์นี้ใช้ Model Context Protocol (MCP) เวอร์ชัน 2024-11-05 ซึ่งอิงจาก JSON-RPC 2.0

### Methods ที่รองรับ

- `initialize`: เริ่มต้นการเชื่อมต่อ MCP
- `tools/list`: แสดงรายการเครื่องมือที่มี
- `tools/call`: เรียกใช้เครื่องมือ
- `notifications/initialized`: การแจ้งเตือนว่าไคลเอนต์เริ่มต้นเรียบร้อยแล้ว

### รูปแบบของ Message

ข้อความทั้งหมดเป็นไปตามรูปแบบ JSON-RPC 2.0:

**Request:**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "method_name",
  "params": {}
}
```

**Response:**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {}
}
```

## การแก้ไขปัญหา

### เซิร์ฟเวอร์ไม่ปรากฏใน Claude Desktop

1. ตรวจสอบว่า path ไปยัง binary เป็น absolute path และถูกต้อง
2. ตรวจสอบว่า binary มีสิทธิ์ execute: `chmod +x bin/mcp-server`
3. ดู logs ของ Claude Desktop (โดยปกติอยู่ในไดเรกทอรี application data)

### ข้อผิดพลาด "Permission Denied"

ให้สิทธิ์ execute กับ binary:
```bash
chmod +x bin/mcp-server
```

### ข้อผิดพลาดในการเรียกใช้เครื่องมือ

ตรวจสอบว่า:
- พารามิเตอร์ที่จำเป็นถูกส่งครบถ้วน
- path ของไฟล์ (สำหรับ read_file) สามารถเข้าถึงได้
- เซิร์ฟเวอร์มีสิทธิ์ที่เหมาะสม

## การมีส่วนร่วม

ยินดีรับ Pull Requests! กรุณาอย่าลังเลที่จะส่ง PR เข้ามา

## สิทธิ์การใช้งาน

MIT License - ดูไฟล์ LICENSE สำหรับรายละเอียด

## แหล่งข้อมูลเพิ่มเติม

- [Model Context Protocol Specification](https://modelcontextprotocol.io)
- [เอกสาร MCP](https://modelcontextprotocol.io/docs)
- [คู่มือ Claude Desktop MCP](https://support.claude.com/en/articles/10949351-getting-started-with-local-mcp-servers-on-claude-desktop)

## การสนับสนุน

สำหรับปัญหาและคำถาม:
- GitHub Issues: https://github.com/gittisak-go/gittisak-go/issues
- MCP Community: https://modelcontextprotocol.io/community

---

## ตัวอย่างการใช้งาน

### ตัวอย่างที่ 1: ใช้กับ Claude Desktop

หลังจากติดตั้งและตั้งค่าเรียบร้อยแล้ว คุณสามารถถาม Claude:
- "มีเครื่องมืออะไรบ้างที่สามารถใช้ได้?"
- "ส่งข้อความ 'สวัสดี MCP!' กลับมา"
- "เวลาบนเซิร์ฟเวอร์ตอนนี้เท่าไหร่?"
- "อ่านไฟล์ที่ /path/to/myfile.txt"

### ตัวอย่างที่ 2: สร้างเครื่องมือคำนวณ

เพิ่มเครื่องมือใหม่ใน `pkg/tools/tools.go`:

```go
// CalculateTool ทำการคำนวณทางคณิตศาสตร์พื้นฐาน
func CalculateTool(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
    operation, _ := arguments["operation"].(string)
    a, _ := arguments["a"].(float64)
    b, _ := arguments["b"].(float64)
    
    var result float64
    switch operation {
    case "add":
        result = a + b
    case "subtract":
        result = a - b
    case "multiply":
        result = a * b
    case "divide":
        if b == 0 {
            return &mcp.CallToolResult{
                Content: []mcp.Content{{Type: "text", Text: "ข้อผิดพลาด: หารด้วยศูนย์"}},
                IsError: true,
            }, nil
        }
        result = a / b
    default:
        return &mcp.CallToolResult{
            Content: []mcp.Content{{Type: "text", Text: "ข้อผิดพลาด: การดำเนินการไม่ถูกต้อง"}},
            IsError: true,
        }, nil
    }
    
    return &mcp.CallToolResult{
        Content: []mcp.Content{
            {Type: "text", Text: fmt.Sprintf("ผลลัพธ์: %.2f", result)},
        },
    }, nil
}
```

ลงทะเบียนใน `cmd/mcp-server/main.go`:

```go
server.RegisterTool(
    "calculate",
    "ทำการคำนวณทางคณิตศาสตร์พื้นฐาน",
    mcp.InputSchema{
        Type: "object",
        Properties: map[string]interface{}{
            "operation": map[string]interface{}{
                "type":        "string",
                "description": "การดำเนินการ: add, subtract, multiply, divide",
            },
            "a": map[string]interface{}{
                "type":        "number",
                "description": "ตัวเลขแรก",
            },
            "b": map[string]interface{}{
                "type":        "number",
                "description": "ตัวเลขที่สอง",
            },
        },
        Required: []string{"operation", "a", "b"},
    },
    tools.CalculateTool,
)
```

สร้างและทดสอบ:
```bash
make build
echo '{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"calculate","arguments":{"operation":"add","a":5,"b":3}}}' | ./bin/mcp-server
```

---

**สร้างโดย**: gittisak-go  
**เวอร์ชัน**: 1.0.0  
**อัปเดตล่าสุด**: ธันวาคม 2024
