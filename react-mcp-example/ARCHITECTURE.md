# Architecture Overview / ภาพรวมสถาปัตยกรรม

## System Architecture

```
┌─────────────────────────────────────────────────────────────────────┐
│                          React MCP Example                          │
└─────────────────────────────────────────────────────────────────────┘

┌─────────────────────┐           ┌─────────────────────┐
│   React Frontend    │           │   MCP Server        │
│   (Port 3000)       │◄─────────►│   (Node.js)         │
│                     │   HTTP    │                     │
│  ┌───────────────┐  │           │  ┌───────────────┐  │
│  │ TaskManager   │  │           │  │ create_task   │  │
│  │ Component     │  │           │  │ list_tasks    │  │
│  └───────────────┘  │           │  │ complete_task │  │
│                     │           │  └───────────────┘  │
│  ┌───────────────┐  │           │                     │
│  │ NoteManager   │  │           │  ┌───────────────┐  │
│  │ Component     │  │           │  │ create_note   │  │
│  └───────────────┘  │           │  │ search_notes  │  │
│                     │           │  └───────────────┘  │
│  ┌───────────────┐  │           │                     │
│  │ WeatherWidget │  │           │  ┌───────────────┐  │
│  │ Component     │  │           │  │ get_weather   │  │
│  └───────────────┘  │           │  └───────────────┘  │
│                     │           │                     │
│  ┌───────────────┐  │           │  ┌───────────────┐  │
│  │ useMCP Hook   │  │           │  │ Resources     │  │
│  └───────────────┘  │           │  │ Management    │  │
│                     │           │  └───────────────┘  │
└─────────────────────┘           └─────────────────────┘
         ▲                                  ▲
         │                                  │
         │                                  │ stdio/HTTP
         │                                  │
         ▼                                  ▼
┌─────────────────────────────────────────────────────────────────────┐
│                         AI Models (Optional)                        │
│                  Claude, GPT, or other MCP-compatible               │
└─────────────────────────────────────────────────────────────────────┘
```

## Data Flow / การไหลของข้อมูล

### 1. User Interaction Flow

```
User Input
    ↓
React Component
    ↓
useMCP Hook
    ↓
mcpService
    ↓
MCP Server (Tool Handler)
    ↓
In-Memory Storage / Database
    ↓
Response back to Component
    ↓
UI Update
```

### 2. AI Model Integration Flow

```
AI Model (e.g., Claude)
    ↓
MCP Client (Built-in)
    ↓
stdio Transport
    ↓
MCP Server (Your Server)
    ↓
Tool Execution
    ↓
Response to AI Model
    ↓
AI generates response to user
```

## Component Structure / โครงสร้างคอมโพเนนต์

```
react-mcp-example/
│
├── client/                         # React Frontend
│   ├── src/
│   │   ├── components/            # UI Components
│   │   │   ├── TaskManager.jsx   # Task management UI
│   │   │   ├── NoteManager.jsx   # Note-taking UI
│   │   │   └── WeatherWidget.jsx # Weather display UI
│   │   │
│   │   ├── hooks/                # Custom React Hooks
│   │   │   └── useMCP.js        # Hook for MCP operations
│   │   │
│   │   ├── services/            # Business Logic
│   │   │   └── mcpService.js   # MCP communication service
│   │   │
│   │   ├── App.jsx              # Main application
│   │   └── main.jsx             # Entry point
│   │
│   └── package.json
│
└── server/                        # MCP Server
    ├── index.js                  # Server implementation
    │   ├── Server Setup          # MCP SDK initialization
    │   ├── Tool Definitions      # Available tools
    │   ├── Tool Handlers         # Tool execution logic
    │   ├── Resource Management   # Resource handlers
    │   └── Transport Layer       # stdio/HTTP transport
    │
    └── package.json
```

## Tool Categories / หมวดหมู่เครื่องมือ

### Task Management Tools
```
┌─────────────────┐
│  create_task    │ → Creates new task with metadata
├─────────────────┤
│  list_tasks     │ → Lists tasks with filtering
├─────────────────┤
│  complete_task  │ → Marks task as completed
└─────────────────┘
```

### Note Management Tools
```
┌─────────────────┐
│  create_note    │ → Creates note with tags
├─────────────────┤
│  search_notes   │ → Full-text search in notes
└─────────────────┘
```

### Information Tools
```
┌─────────────────┐
│  get_weather    │ → Retrieves weather data
└─────────────────┘
```

## Technology Stack / เทคโนโลยีที่ใช้

### Frontend
- **React 18** - UI framework
- **Vite** - Build tool & dev server
- **Modern CSS** - Styling with dark/light mode

### Backend
- **Node.js 18+** - Runtime
- **@modelcontextprotocol/sdk** - MCP implementation
- **Express** (optional) - HTTP server

### Development
- **MCP Inspector** - Debugging tool
- **ESLint** - Code linting
- **Git** - Version control

## Security Considerations / ข้อควรระวังด้านความปลอดภัย

```
┌─────────────────────────────────────────┐
│          Security Layers                │
├─────────────────────────────────────────┤
│  1. Input Validation                    │
│     ├─ Type checking                    │
│     ├─ Size limits                      │
│     └─ Sanitization                     │
├─────────────────────────────────────────┤
│  2. Error Handling                      │
│     ├─ Try-catch blocks                 │
│     ├─ Error messages                   │
│     └─ Logging                          │
├─────────────────────────────────────────┤
│  3. Transport Security                  │
│     ├─ stdio (local only)               │
│     └─ HTTPS (for remote)               │
├─────────────────────────────────────────┤
│  4. Data Storage                        │
│     ├─ No sensitive data in memory      │
│     └─ Database encryption (prod)       │
└─────────────────────────────────────────┘
```

## Scaling for Production / ขยายสู่ Production

### Current (Demo)
```
In-Memory Storage
    ↓
Single Process
    ↓
Local Development
```

### Production Ready
```
Database (PostgreSQL/MongoDB)
    ↓
Multiple Processes / Containers
    ↓
Load Balancer
    ↓
Cloud Deployment (AWS/GCP/Azure)
    ↓
Monitoring & Logging
```

## Integration Patterns / รูปแบบการรวมระบบ

### Pattern 1: Direct Client Use
```
User → React App → MCP Service → Display
```

### Pattern 2: AI Model Integration
```
User → AI Model → MCP Server → Tool Execution → AI Response
```

### Pattern 3: Hybrid
```
User → React App (Direct) + AI Model (Natural Language) → MCP Server
```

## Performance Optimization / การเพิ่มประสิทธิภาพ

### Frontend
1. **React.memo** - Component memoization
2. **useMemo** - Value memoization
3. **Lazy Loading** - Code splitting
4. **Asset Optimization** - Image & bundle size

### Backend
1. **Caching** - Redis/Memcached
2. **Database Indexing** - Query optimization
3. **Connection Pooling** - Efficient DB connections
4. **Async Operations** - Non-blocking I/O

---

For more details, see:
- [Main README](README.md)
- [Server Documentation](server/README.md)
- [Quick Start Guide](QUICKSTART.md)
